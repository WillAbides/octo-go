// Code generated by octo-go; DO NOT EDIT.

package gitignore

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	internal "github.com/willabides/octo-go/internal"
	requests "github.com/willabides/octo-go/requests"
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
GetAllTemplates performs requests for "gitignore/get-all-templates"

Get all gitignore templates.

  GET /gitignore/templates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
func GetAllTemplates(ctx context.Context, req *GetAllTemplatesReq, opt ...requests.Option) (*GetAllTemplatesResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(GetAllTemplatesReq)
	}
	resp := &GetAllTemplatesResponse{}
	builder := req.requestBuilder()

	httpReq, err := builder.HTTPRequest(ctx, opts)
	if err != nil {
		return resp, err
	}

	r, err := opts.HttpClient().Do(httpReq)
	if err != nil {
		return resp, err
	}
	resp.httpResponse = r

	return NewGetAllTemplatesResponse(r, opts.PreserveResponseBody())
}

// NewGetAllTemplatesResponse builds a new *GetAllTemplatesResponse from an *http.Response
func NewGetAllTemplatesResponse(resp *http.Response, preserveBody bool) (*GetAllTemplatesResponse, error) {
	var result GetAllTemplatesResponse
	result.httpResponse = resp
	err := internal.ErrorCheck(resp, []int{200, 304})
	if err != nil {
		return &result, err
	}
	if internal.IntInSlice(resp.StatusCode, []int{200}) {
		err = internal.DecodeResponseBody(resp, &result.Data, preserveBody)
		if err != nil {
			return &result, err
		}
	}
	return &result, nil
}

/*
GetAllTemplates performs requests for "gitignore/get-all-templates"

Get all gitignore templates.

  GET /gitignore/templates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
func (c Client) GetAllTemplates(ctx context.Context, req *GetAllTemplatesReq, opt ...requests.Option) (*GetAllTemplatesResponse, error) {
	return GetAllTemplates(ctx, req, append(c, opt...)...)
}

/*
GetAllTemplatesReq is request data for Client.GetAllTemplates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
type GetAllTemplatesReq struct {
	_url string
}

// HTTPRequest builds an *http.Request
func (r *GetAllTemplatesReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *GetAllTemplatesReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"accept": internal.String("application/json")},
		Method:             "GET",
		OperationID:        "gitignore/get-all-templates",
		Previews:           map[string]bool{},
		RequiredPreviews:   []string{},
		URLPath:            fmt.Sprintf("/gitignore/templates"),
		URLQuery:           query,
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetAllTemplatesReq) Rel(link string, resp *GetAllTemplatesResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetAllTemplatesResponseBody is a response body for GetAllTemplates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
type GetAllTemplatesResponseBody []string

/*
GetAllTemplatesResponse is a response for GetAllTemplates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
type GetAllTemplatesResponse struct {
	httpResponse *http.Response
	Data         GetAllTemplatesResponseBody
}

func (r *GetAllTemplatesResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

/*
GetTemplate performs requests for "gitignore/get-template"

Get a gitignore template.

  GET /gitignore/templates/{name}

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
func GetTemplate(ctx context.Context, req *GetTemplateReq, opt ...requests.Option) (*GetTemplateResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(GetTemplateReq)
	}
	resp := &GetTemplateResponse{}
	builder := req.requestBuilder()

	httpReq, err := builder.HTTPRequest(ctx, opts)
	if err != nil {
		return resp, err
	}

	r, err := opts.HttpClient().Do(httpReq)
	if err != nil {
		return resp, err
	}
	resp.httpResponse = r

	return NewGetTemplateResponse(r, opts.PreserveResponseBody())
}

// NewGetTemplateResponse builds a new *GetTemplateResponse from an *http.Response
func NewGetTemplateResponse(resp *http.Response, preserveBody bool) (*GetTemplateResponse, error) {
	var result GetTemplateResponse
	result.httpResponse = resp
	err := internal.ErrorCheck(resp, []int{200, 304})
	if err != nil {
		return &result, err
	}
	if internal.IntInSlice(resp.StatusCode, []int{200}) {
		err = internal.DecodeResponseBody(resp, &result.Data, preserveBody)
		if err != nil {
			return &result, err
		}
	}
	return &result, nil
}

/*
GetTemplate performs requests for "gitignore/get-template"

Get a gitignore template.

  GET /gitignore/templates/{name}

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
func (c Client) GetTemplate(ctx context.Context, req *GetTemplateReq, opt ...requests.Option) (*GetTemplateResponse, error) {
	return GetTemplate(ctx, req, append(c, opt...)...)
}

/*
GetTemplateReq is request data for Client.GetTemplate

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
type GetTemplateReq struct {
	_url string

	// name parameter
	Name string
}

// HTTPRequest builds an *http.Request
func (r *GetTemplateReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *GetTemplateReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"accept": internal.String("application/json")},
		Method:             "GET",
		OperationID:        "gitignore/get-template",
		Previews:           map[string]bool{},
		RequiredPreviews:   []string{},
		URLPath:            fmt.Sprintf("/gitignore/templates/%v", r.Name),
		URLQuery:           query,
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetTemplateReq) Rel(link string, resp *GetTemplateResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetTemplateResponse is a response for GetTemplate

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
type GetTemplateResponse struct {
	httpResponse *http.Response
	Data         components.GitignoreTemplate
}

func (r *GetTemplateResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}
