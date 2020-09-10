// Code generated by octo-go; DO NOT EDIT.

package markdown

import (
	"context"
	"fmt"
	requests "github.com/willabides/octo-go/requests"
	"io"
	"net/http"
)

/*
Render performs requests for "markdown/render"

Render a Markdown document.

  POST /markdown

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document
*/
func Render(ctx context.Context, req *RenderReq, opt ...requests.Option) (*RenderResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(RenderReq)
	}
	resp := &RenderResponse{}

	httpReq, err := req.HTTPRequest(ctx, opt...)
	if err != nil {
		return nil, err
	}

	r, err := opts.HttpClient().Do(httpReq)
	if err != nil {
		return nil, err
	}

	err = resp.ReadResponse(r)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
Render performs requests for "markdown/render"

Render a Markdown document.

  POST /markdown

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) Render(ctx context.Context, req *RenderReq, opt ...requests.Option) (*RenderResponse, error) {
	return Render(ctx, req, append(c, opt...)...)
}

/*
RenderReq is request data for Client.Render

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type RenderReq struct {
	_url        string
	RequestBody RenderReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *RenderReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"content-type": strPtr("application/json")},
		Method:      "POST",
		Options:     opt,
		URLPath:     fmt.Sprintf("/markdown"),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RenderReq) Rel(link string, resp *RenderResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
RenderReqBody is a request body for markdown/render

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document
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

https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document
*/
type RenderResponse struct {
	httpResponse *http.Response
}

// HTTPResponse returns the *http.Response
func (r *RenderResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *RenderResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{200, 304})
	if err != nil {
		return err
	}
	return nil
}

/*
RenderRaw performs requests for "markdown/render-raw"

Render a Markdown document in raw mode.

  POST /markdown/raw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode
*/
func RenderRaw(ctx context.Context, req *RenderRawReq, opt ...requests.Option) (*RenderRawResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(RenderRawReq)
	}
	resp := &RenderRawResponse{}

	httpReq, err := req.HTTPRequest(ctx, opt...)
	if err != nil {
		return nil, err
	}

	r, err := opts.HttpClient().Do(httpReq)
	if err != nil {
		return nil, err
	}

	err = resp.ReadResponse(r)
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

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) RenderRaw(ctx context.Context, req *RenderRawReq, opt ...requests.Option) (*RenderRawResponse, error) {
	return RenderRaw(ctx, req, append(c, opt...)...)
}

/*
RenderRawReq is request data for Client.RenderRaw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type RenderRawReq struct {
	_url string

	// http request's body
	RequestBody io.Reader
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *RenderRawReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"content-type": strPtr("text/x-markdown")},
		Method:      "POST",
		Options:     opt,
		URLPath:     fmt.Sprintf("/markdown/raw"),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RenderRawReq) Rel(link string, resp *RenderRawResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
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
	httpResponse *http.Response
}

// HTTPResponse returns the *http.Response
func (r *RenderRawResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *RenderRawResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{200, 304})
	if err != nil {
		return err
	}
	return nil
}
