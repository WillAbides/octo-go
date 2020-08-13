// Code generated by octo-go; DO NOT EDIT.

package codesofconduct

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	internal "github.com/willabides/octo-go/internal"
	requests "github.com/willabides/octo-go/requests"
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
GetAllCodesOfConduct performs requests for "codes-of-conduct/get-all-codes-of-conduct"

Get all codes of conduct.

  GET /codes_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct
*/
func GetAllCodesOfConduct(ctx context.Context, req *GetAllCodesOfConductReq, opt ...requests.Option) (*GetAllCodesOfConductResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetAllCodesOfConductReq)
	}
	resp := &GetAllCodesOfConductResponse{}

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
GetAllCodesOfConduct performs requests for "codes-of-conduct/get-all-codes-of-conduct"

Get all codes of conduct.

  GET /codes_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetAllCodesOfConduct(ctx context.Context, req *GetAllCodesOfConductReq, opt ...requests.Option) (*GetAllCodesOfConductResponse, error) {
	return GetAllCodesOfConduct(ctx, req, append(c, opt...)...)
}

/*
GetAllCodesOfConductReq is request data for Client.GetAllCodesOfConduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetAllCodesOfConductReq struct {
	_url string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetAllCodesOfConductReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	query := url.Values{}

	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:        []string{"scarlet-witch"},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"accept": internal.String("application/json")},
		Method:             "GET",
		Options:            opt,
		Previews:           map[string]bool{"scarlet-witch": r.ScarletWitchPreview},
		RequiredPreviews:   []string{"scarlet-witch"},
		URLPath:            fmt.Sprintf("/codes_of_conduct"),
		URLQuery:           query,
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetAllCodesOfConductReq) Rel(link string, resp *GetAllCodesOfConductResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetAllCodesOfConductResponse is a response for GetAllCodesOfConduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct
*/
type GetAllCodesOfConductResponse struct {
	httpResponse *http.Response
	Data         []components.CodeOfConduct
}

// HTTPResponse returns the *http.Response
func (r *GetAllCodesOfConductResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetAllCodesOfConductResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{200, 304})
	if err != nil {
		return err
	}
	if internal.IntInSlice(resp.StatusCode, []int{200}) {
		err = internal.UnmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
GetConductCode performs requests for "codes-of-conduct/get-conduct-code"

Get a code of conduct.

  GET /codes_of_conduct/{key}

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct
*/
func GetConductCode(ctx context.Context, req *GetConductCodeReq, opt ...requests.Option) (*GetConductCodeResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetConductCodeReq)
	}
	resp := &GetConductCodeResponse{}

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
GetConductCode performs requests for "codes-of-conduct/get-conduct-code"

Get a code of conduct.

  GET /codes_of_conduct/{key}

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetConductCode(ctx context.Context, req *GetConductCodeReq, opt ...requests.Option) (*GetConductCodeResponse, error) {
	return GetConductCode(ctx, req, append(c, opt...)...)
}

/*
GetConductCodeReq is request data for Client.GetConductCode

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetConductCodeReq struct {
	_url string

	// key parameter
	Key string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetConductCodeReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	query := url.Values{}

	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:        []string{"scarlet-witch"},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"accept": internal.String("application/json")},
		Method:             "GET",
		Options:            opt,
		Previews:           map[string]bool{"scarlet-witch": r.ScarletWitchPreview},
		RequiredPreviews:   []string{"scarlet-witch"},
		URLPath:            fmt.Sprintf("/codes_of_conduct/%v", r.Key),
		URLQuery:           query,
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetConductCodeReq) Rel(link string, resp *GetConductCodeResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetConductCodeResponse is a response for GetConductCode

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct
*/
type GetConductCodeResponse struct {
	httpResponse *http.Response
	Data         components.CodeOfConduct
}

// HTTPResponse returns the *http.Response
func (r *GetConductCodeResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetConductCodeResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{200, 304})
	if err != nil {
		return err
	}
	if internal.IntInSlice(resp.StatusCode, []int{200}) {
		err = internal.UnmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
GetForRepo performs requests for "codes-of-conduct/get-for-repo"

Get the code of conduct for a repository.

  GET /repos/{owner}/{repo}/community/code_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository
*/
func GetForRepo(ctx context.Context, req *GetForRepoReq, opt ...requests.Option) (*GetForRepoResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetForRepoReq)
	}
	resp := &GetForRepoResponse{}

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
GetForRepo performs requests for "codes-of-conduct/get-for-repo"

Get the code of conduct for a repository.

  GET /repos/{owner}/{repo}/community/code_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetForRepo(ctx context.Context, req *GetForRepoReq, opt ...requests.Option) (*GetForRepoResponse, error) {
	return GetForRepo(ctx, req, append(c, opt...)...)
}

/*
GetForRepoReq is request data for Client.GetForRepo

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetForRepoReq struct {
	_url  string
	Owner string
	Repo  string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetForRepoReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	query := url.Values{}

	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:        []string{"scarlet-witch"},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"accept": internal.String("application/json")},
		Method:             "GET",
		Options:            opt,
		Previews:           map[string]bool{"scarlet-witch": r.ScarletWitchPreview},
		RequiredPreviews:   []string{"scarlet-witch"},
		URLPath:            fmt.Sprintf("/repos/%v/%v/community/code_of_conduct", r.Owner, r.Repo),
		URLQuery:           query,
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetForRepoReq) Rel(link string, resp *GetForRepoResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetForRepoResponse is a response for GetForRepo

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository
*/
type GetForRepoResponse struct {
	httpResponse *http.Response
	Data         components.CodeOfConduct
}

// HTTPResponse returns the *http.Response
func (r *GetForRepoResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetForRepoResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{200})
	if err != nil {
		return err
	}
	if internal.IntInSlice(resp.StatusCode, []int{200}) {
		err = internal.UnmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
