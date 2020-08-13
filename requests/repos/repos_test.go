package repos_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/components"
	"github.com/willabides/octo-go/internal/testutil"
	"github.com/willabides/octo-go/requests"
	"github.com/willabides/octo-go/requests/repos"
)

func vcrClient(t *testing.T, cas string, opts ...requests.Option) repos.Client {
	return testutil.VCRClient(t, cas, opts...).Repos()
}

func TestGetContent(t *testing.T) {
	t.Run("file", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), testutil.PATAuth())

		response, err := client.GetContent(ctx, &repos.GetContentReq{
			Owner: "WillAbides",
			Repo:  "octo-go",
			Path:  "generator/main.go",
		})
		require.NoError(t, err)
		gotVal, ok := response.Data.Value().(components.ContentFile)
		require.True(t, ok)
		require.Equal(t, "main.go", gotVal.Name)
	})

	t.Run("directory", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), testutil.PATAuth())

		response, err := client.GetContent(ctx, &repos.GetContentReq{
			Owner: "WillAbides",
			Repo:  "octo-go",
			Path:  "generator",
		})
		require.NoError(t, err)
		fmt.Printf("%T\n", response.Data.Value())
		gotVal, ok := response.Data.Value().(components.ContentDirectory)
		require.True(t, ok)
		require.Greater(t, len(gotVal), 1)
	})
}

func TestDownloadTarballArchive(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), testutil.PATAuth())

	resp, err := client.DownloadTarballArchive(ctx, &repos.DownloadTarballArchiveReq{
		Owner: "octocat",
		Repo:  "Hello-World",
		Ref:   "master",
	})
	require.NoError(t, err)
	g, err := ioutil.ReadAll(resp.HTTPResponse().Body)
	require.NoError(t, err)
	require.True(t, len(g) > 100)
}

func TestCompareCommits(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), testutil.PATAuth())
	got, err := client.CompareCommits(ctx, &repos.CompareCommitsReq{
		Owner: "WillAbides",
		Repo:  "octo-go",
		Base:  "2c80673f15b275cbb22fc167cb1b9aa43ba7a27a",
		Head:  "ceac7c6d9a134326a0871174423bea19acbb122a",
	})
	require.NoError(t, err)
	require.Equal(t, "ahead", got.Data.Status)
}

func TestUploadReleaseAsset(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), testutil.PATAuth())
		releaseResp, err := client.GetReleaseByTag(ctx, &repos.GetReleaseByTagReq{
			Owner: "octo-cli-testorg",
			Repo:  "scratch",
			Tag:   "v0.0.2",
		})
		require.NoError(t, err)
		licenseFile := filepath.Join(testutil.ProjectRoot(t), "LICENSE")
		licenseBytes, err := ioutil.ReadFile(licenseFile)
		require.NoError(t, err)
		uploadResp, err := client.UploadReleaseAsset(ctx, &repos.UploadReleaseAssetReq{
			URL:               releaseResp.Data.UploadUrl,
			Name:              octo.String("LICENSE"),
			RequestBody:       bytes.NewBuffer(licenseBytes),
			ContentTypeHeader: octo.String("text/plain"),
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, err := client.DeleteReleaseAsset(ctx, &repos.DeleteReleaseAssetReq{
				Owner:   "octo-cli-testorg",
				Repo:    "scratch",
				AssetId: uploadResp.Data.Id,
			})
			require.NoError(t, err)
		})

		require.Equal(t, "LICENSE", uploadResp.Data.Name)
	})
}
