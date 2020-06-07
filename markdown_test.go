package octo_test

import (
	"context"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
)

func TestMarkdownRender(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name())
	response, err := client.MarkdownRender(ctx, &octo.MarkdownRenderReq{
		RequestBody: octo.MarkdownRenderReqBody{
			Text: octo.String("this is my body"),
		},
	}, octo.PreserveResponseBody())
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
	}, octo.PreserveResponseBody())
	require.NoError(t, err)
	rendered, err := ioutil.ReadAll(response.HTTPResponse().Body)
	require.NoError(t, err)
	require.Equal(t, "<p>this is my body</p>\n", string(rendered))
}
