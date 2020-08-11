package gists_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/internal/testutil"
	"github.com/willabides/octo-go/requests"
	"github.com/willabides/octo-go/requests/gists"
)

func vcrClient(t *testing.T, cas string, opts ...requests.Option) gists.Client {
	return gists.NewClient(testutil.VCRClient(t, cas, opts...))
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), testutil.PATAuth())

	createResp, err := client.Create(ctx, &gists.CreateReq{
		RequestBody: gists.CreateReqBody{
			Description: octo.String("test gist, pls delete"),
			Public:      octo.Bool(false),
			Files: map[string]gists.CreateReqBodyFiles{
				"foo.md": {
					Content: octo.String(`not much here`),
				},
			},
		},
	})
	require.NoError(t, err)
	fooFile := createResp.Data.Files["foo.md"]
	require.Equal(t, `not much here`, fooFile.Content)

	_, err = client.Delete(ctx, &gists.DeleteReq{
		GistId: createResp.Data.Id,
	})
	require.NoError(t, err)
}
