// Code generated by octo-go; DO NOT EDIT.

package markdown

import (
	"context"
	"fmt"
	internal "github.com/willabides/octo-go/internal"
	requests "github.com/willabides/octo-go/requests"
	"io"
	"net/http"
	"net/url"
)

func strPtr(s string) *string { return &s }

// Client is a set of options to apply to requests
type Client []requests.Option

// NewClient returns a new Client
func NewClient(opt ...requests.Option) Client {
	return opt
}

/*
Render performs requests for "markdown/render"

Render a Markdown document.

  POST /markdown

https://developer.github.com/v3/markdown/#render-a-markdown-document
*/
func Render(ctx context.Context, req *RenderReq, opt ...requests.Option) (*RenderResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(RenderReq)
	}
	resp := &RenderResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	err = internal.DecodeResponseBody(r, builder, opts, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
Render performs requests for "markdown/render"

Render a Markdown document.

  POST /markdown

https://developer.github.com/v3/markdown/#render-a-markdown-document
*/
func (c Client) Render(ctx context.Context, req *RenderReq, opt ...requests.Option) (*RenderResponse, error) {
	return Render(ctx, req, append(c, opt...)...)
}

/*
RenderReq is request data for Client.Render

https://developer.github.com/v3/markdown/#render-a-markdown-document
*/
type RenderReq struct {
	_url        string
	RequestBody RenderReqBody
}

// HTTPRequest builds an *http.Request
func (r *RenderReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *RenderReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{},
		Body:               r.RequestBody,
		DataStatuses:       []int{},
		EndpointAttributes: []internal.EndpointAttribute{internal.AttrJSONRequestBody},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"content-type": internal.String("application/json")},
		Method:             "POST",
		OperationID:        "markdown/render",
		Previews:           map[string]bool{},
		RequiredPreviews:   []string{},
		URLPath:            fmt.Sprintf("/markdown"),
		URLQuery:           query,
		ValidStatuses:      []int{200, 304},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RenderReq) Rel(link string, resp *RenderResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
RenderReqBody is a request body for markdown/render

https://developer.github.com/v3/markdown/#render-a-markdown-document
*/
type RenderReqBody struct {

	// The repository context to use when creating references in `gfm` mode.
	Context *string `json:"context,omitempty"`

	// The rendering mode.
	Mode *string `json:"mode,omitempty"`

	// The Markdown text to render in HTML.
	Text *string `json:"text"`
}

/*
RenderResponse is a response for Render

https://developer.github.com/v3/markdown/#render-a-markdown-document
*/
type RenderResponse struct {
	requests.Response
	request *RenderReq
}

/*
RenderRaw performs requests for "markdown/render-raw"

Render a Markdown document in raw mode.

  POST /markdown/raw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
func RenderRaw(ctx context.Context, req *RenderRawReq, opt ...requests.Option) (*RenderRawResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(RenderRawReq)
	}
	resp := &RenderRawResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	err = internal.DecodeResponseBody(r, builder, opts, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
RenderRaw performs requests for "markdown/render-raw"

Render a Markdown document in raw mode.

  POST /markdown/raw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
func (c Client) RenderRaw(ctx context.Context, req *RenderRawReq, opt ...requests.Option) (*RenderRawResponse, error) {
	return RenderRaw(ctx, req, append(c, opt...)...)
}

/*
RenderRawReq is request data for Client.RenderRaw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
type RenderRawReq struct {
	_url string

	// http request's body
	RequestBody io.Reader
}

// HTTPRequest builds an *http.Request
func (r *RenderRawReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *RenderRawReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{},
		Body:               r.RequestBody,
		DataStatuses:       []int{},
		EndpointAttributes: []internal.EndpointAttribute{internal.AttrBodyUploader},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"content-type": internal.String("text/x-markdown")},
		Method:             "POST",
		OperationID:        "markdown/render-raw",
		Previews:           map[string]bool{},
		RequiredPreviews:   []string{},
		URLPath:            fmt.Sprintf("/markdown/raw"),
		URLQuery:           query,
		ValidStatuses:      []int{200, 304},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RenderRawReq) Rel(link string, resp *RenderRawResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
RenderRawResponse is a response for RenderRaw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
type RenderRawResponse struct {
	requests.Response
	request *RenderRawReq
}
