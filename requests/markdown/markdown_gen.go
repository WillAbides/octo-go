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

	err = resp.Load(r)
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

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) Render(ctx context.Context, req *RenderReq, opt ...requests.Option) (*RenderResponse, error) {
	return Render(ctx, req, append(c, opt...)...)
}

/*
RenderReq is request data for Client.Render

https://developer.github.com/v3/markdown/#render-a-markdown-document

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type RenderReq struct {
	_url        string
	RequestBody RenderReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *RenderReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	query := url.Values{}

	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:        []string{},
		Body:               r.RequestBody,
		EndpointAttributes: []internal.EndpointAttribute{internal.AttrJSONRequestBody},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"content-type": internal.String("application/json")},
		Method:             "POST",
		Options:            opt,
		Previews:           map[string]bool{},
		RequiredPreviews:   []string{},
		URLPath:            fmt.Sprintf("/markdown"),
		URLQuery:           query,
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RenderReq) Rel(link string, resp *RenderResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
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
	httpResponse *http.Response
}

// HTTPResponse returns the *http.Response
func (r *RenderResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// Load loads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *RenderResponse) Load(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{200, 304})
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

	err = resp.Load(r)
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
	query := url.Values{}

	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:        []string{},
		Body:               r.RequestBody,
		EndpointAttributes: []internal.EndpointAttribute{internal.AttrBodyUploader},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"content-type": internal.String("text/x-markdown")},
		Method:             "POST",
		Options:            opt,
		Previews:           map[string]bool{},
		RequiredPreviews:   []string{},
		URLPath:            fmt.Sprintf("/markdown/raw"),
		URLQuery:           query,
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RenderRawReq) Rel(link string, resp *RenderRawResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
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

// Load loads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *RenderRawResponse) Load(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{200, 304})
	if err != nil {
		return err
	}
	return nil
}
