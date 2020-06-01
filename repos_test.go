package octo_test

import (
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
