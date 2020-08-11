package markdown_test

import (
	"context"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/internal/testutil"
	"github.com/willabides/octo-go/requests"
	"github.com/willabides/octo-go/requests/markdown"
)

func vcrClient(t *testing.T, cas string, opts ...requests.Option) markdown.Client {
	return markdown.NewClient(testutil.VCRClient(t, cas, opts...))
}

func TestRender(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), octo.PreserveResponseBody())
	response, err := client.Render(ctx, &markdown.RenderReq{
		RequestBody: markdown.RenderReqBody{
			Text: octo.String("this is my body"),
		},
	})
	require.NoError(t, err)
	rendered, err := ioutil.ReadAll(response.HTTPResponse().Body)
	require.NoError(t, err)
	require.Equal(t, "<p>this is my body</p>\n", string(rendered))
}

func TestRenderRaw(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), octo.PreserveResponseBody())
	response, err := client.RenderRaw(ctx, &markdown.RenderRawReq{
		RequestBody: strings.NewReader("this is my body"),
	})
	require.NoError(t, err)
	rendered, err := ioutil.ReadAll(response.HTTPResponse().Body)
	require.NoError(t, err)
	require.Equal(t, "<p>this is my body</p>\n", string(rendered))
}
