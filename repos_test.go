package octo_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
)

func TestReposGetArchiveLink(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), patAuth())

	resp, err := client.ReposGetArchiveLink(ctx, &octo.ReposGetArchiveLinkReq{
		Owner:         "octocat",
		Repo:          "Hello-World",
		ArchiveFormat: "tarball",
		Ref:           "master",
	})
	require.NoError(t, err)
	g, err := ioutil.ReadAll(resp.HTTPResponse().Body)
	require.NoError(t, err)
	require.True(t, len(g) > 100)
}

func TestResposUploadReleaseAsset(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), patAuth())
		releaseResp, err := client.ReposGetReleaseByTag(ctx, &octo.ReposGetReleaseByTagReq{
			Owner: "octo-cli-testorg",
			Repo:  "scratch",
			Tag:   "v0.0.2",
		})
		require.NoError(t, err)
		licenseBytes, err := ioutil.ReadFile("LICENSE")
		require.NoError(t, err)
		uploadResp, err := client.ReposUploadReleaseAsset(ctx, &octo.ReposUploadReleaseAssetReq{
			URL:               releaseResp.Data.UploadUrl,
			Name:              octo.String("LICENSE"),
			RequestBody:       bytes.NewBuffer(licenseBytes),
			ContentTypeHeader: octo.String("text/plain"),
		})
		require.NoError(t, err)

		t.Cleanup(func() {
			_, err := client.ReposDeleteReleaseAsset(ctx, &octo.ReposDeleteReleaseAssetReq{
				Owner:   "octo-cli-testorg",
				Repo:    "scratch",
				AssetId: uploadResp.Data.Id,
			})
			require.NoError(t, err)
		})

		require.Equal(t, "LICENSE", uploadResp.Data.Name)
	})
}
