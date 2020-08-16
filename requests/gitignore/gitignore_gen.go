// Code generated by octo-go; DO NOT EDIT.

package gitignore

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	requests "github.com/willabides/octo-go/requests"
	"net/http"
)

/*
GetAllTemplates performs requests for "gitignore/get-all-templates"

Get all gitignore templates.

  GET /gitignore/templates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
func GetAllTemplates(ctx context.Context, req *GetAllTemplatesReq, opt ...requests.Option) (*GetAllTemplatesResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetAllTemplatesReq)
	}
	resp := &GetAllTemplatesResponse{}

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
GetAllTemplates performs requests for "gitignore/get-all-templates"

Get all gitignore templates.

  GET /gitignore/templates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetAllTemplates(ctx context.Context, req *GetAllTemplatesReq, opt ...requests.Option) (*GetAllTemplatesResponse, error) {
	return GetAllTemplates(ctx, req, append(c, opt...)...)
}

/*
GetAllTemplatesReq is request data for Client.GetAllTemplates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetAllTemplatesReq struct {
	_url string
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetAllTemplatesReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": strPtr("application/json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/gitignore/templates"),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetAllTemplatesReq) Rel(link string, resp *GetAllTemplatesResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
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

// HTTPResponse returns the *http.Response
func (r *GetAllTemplatesResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetAllTemplatesResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{200, 304})
	if err != nil {
		return err
	}
	if intInSlice(resp.StatusCode, []int{200}) {
		err = unmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
GetTemplate performs requests for "gitignore/get-template"

Get a gitignore template.

  GET /gitignore/templates/{name}

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
func GetTemplate(ctx context.Context, req *GetTemplateReq, opt ...requests.Option) (*GetTemplateResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetTemplateReq)
	}
	resp := &GetTemplateResponse{}

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
GetTemplate performs requests for "gitignore/get-template"

Get a gitignore template.

  GET /gitignore/templates/{name}

https://developer.github.com/v3/gitignore/#get-a-gitignore-template

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetTemplate(ctx context.Context, req *GetTemplateReq, opt ...requests.Option) (*GetTemplateResponse, error) {
	return GetTemplate(ctx, req, append(c, opt...)...)
}

/*
GetTemplateReq is request data for Client.GetTemplate

https://developer.github.com/v3/gitignore/#get-a-gitignore-template

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetTemplateReq struct {
	_url string

	// name parameter
	Name string
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetTemplateReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": strPtr("application/json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/gitignore/templates/%v", r.Name),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetTemplateReq) Rel(link string, resp *GetTemplateResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
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

// HTTPResponse returns the *http.Response
func (r *GetTemplateResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetTemplateResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{200, 304})
	if err != nil {
		return err
	}
	if intInSlice(resp.StatusCode, []int{200}) {
		err = unmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
