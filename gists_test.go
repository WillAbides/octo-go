package octo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
)

func TestCreateGist(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), patAuth())

	createResp, err := client.GistsCreate(ctx, &octo.GistsCreateReq{
		RequestBody: octo.GistsCreateReqBody{
			Description: octo.String("test gist, pls delete"),
			Public:      octo.Bool(false),
			Files: map[string]octo.GistsCreateReqBodyFiles{
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
