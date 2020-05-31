package octo_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"golang.org/x/oauth2"
)

func vcrClient(t *testing.T, cassette string) *octo.Client {
	t.Helper()
	ctx := context.Background()
	oauthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	))
	cassette = strings.ReplaceAll(cassette, "/", "_")
	cassette = filepath.Join(filepath.FromSlash("testdata/vcr/"), cassette)
	r, err := recorder.NewAsMode(cassette, recorder.ModeReplaying, oauthClient.Transport)

	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, r.Stop())
	})

	return octo.NewClient(&http.Client{
		Transport: r,
	})
}

func TestResposUploadReleaseAsset(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name())
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
		require.Equal(t, "LICENSE", uploadResp.Data.Name)
	})
}

func TestIssuesCheckAssignee(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name())

	t.Run("true", func(t *testing.T) {
		result, err := client.IssuesCheckAssignee(ctx, &octo.IssuesCheckAssigneeReq{
			Owner:    "WillAbides",
			Repo:     "octo-go",
			Assignee: "WillAbides",
		})
		require.NoError(t, err)
		require.True(t, result.Data)
	})

	t.Run("false", func(t *testing.T) {
		result, err := client.IssuesCheckAssignee(ctx, &octo.IssuesCheckAssigneeReq{
			Owner:    "WillAbides",
			Repo:     "octo-go",
			Assignee: "defunkt",
		})
		require.NoError(t, err)
		require.False(t, result.Data)
	})
}

func TestReposGetArchiveLink(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name())

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

func TestCreateGist(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name())

	createResp, err := client.GistsCreate(ctx, &octo.GistsCreateReq{
		RequestBody: octo.GistsCreateReqBody{
			Description: octo.String("test gist, pls delete"),
			Public:      octo.Bool(false),
			Files: map[string]*octo.GistsCreateReqBodyFiles{
				"foo.md": {
					Content: octo.String(`not much here`),
				},
			},
		},
	})
	require.NoError(t, err)
	fooFile := createResp.Data.Files["foo.md"]
	require.Equal(t, `not much here`, fooFile.Content)
}

func TestIssuesListComments(t *testing.T) {
	t.Run("paging", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name())
		var commentIDs []int64
		req := &octo.IssuesListCommentsReq{
			Owner:       "golang",
			Repo:        "go",
			IssueNumber: 1,
			PerPage:     octo.Int64(4),
		}
		ok := true
		for ok {
			resp, err := client.IssuesListComments(ctx, req)
			require.NoError(t, err)
			if resp.Data != nil {
				for _, r := range *resp.Data {
					commentIDs = append(commentIDs, r.Id)
				}
			}
			ok = req.Rel(octo.RelNext, resp)
		}
		require.Len(t, commentIDs, 12)
	})
}

func TestMarkdownRender(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name())
	response, err := client.MarkdownRender(ctx, &octo.MarkdownRenderReq{
		RequestBody: octo.MarkdownRenderReqBody{
			Text: octo.String("this is my body"),
		},
	}, octo.RequestPreserveResponseBody())
	require.NoError(t, err)
	rendered, err := ioutil.ReadAll(response.HTTPResponse().Body)
	require.NoError(t, err)
	require.Equal(t, "<p>this is my body</p>\n", string(rendered))
}

func TestMarkdownRenderRaw(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name())
	response, err := client.MarkdownRenderRaw(ctx, &octo.MarkdownRenderRawReq{
		RequestBody:       strings.NewReader("this is my body"),
		ContentTypeHeader: octo.String("text/plain"),
	}, octo.RequestPreserveResponseBody())
	require.NoError(t, err)
	rendered, err := ioutil.ReadAll(response.HTTPResponse().Body)
	require.NoError(t, err)
	require.Equal(t, "<p>this is my body</p>\n", string(rendered))
}
