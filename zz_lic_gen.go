// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	common "github.com/willabides/octo-go/common"
	components "github.com/willabides/octo-go/components"
	internal "github.com/willabides/octo-go/internal"
	options "github.com/willabides/octo-go/options"
	"net/http"
	"net/url"
	"strconv"
)

/*
LicensesGet performs requests for "licenses/get"

Get a license.

  GET /licenses/{license}

https://developer.github.com/v3/licenses/#get-a-license
*/
func LicensesGet(ctx context.Context, req *LicensesGetReq, opt ...options.Option) (*LicensesGetResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(LicensesGetReq)
	}
	resp := &LicensesGetResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.License{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesGet performs requests for "licenses/get"

Get a license.

  GET /licenses/{license}

https://developer.github.com/v3/licenses/#get-a-license
*/
func (c Client) LicensesGet(ctx context.Context, req *LicensesGetReq, opt ...options.Option) (*LicensesGetResponse, error) {
	return LicensesGet(ctx, req, append(c, opt...)...)
}

/*
LicensesGetReq is request data for Client.LicensesGet

https://developer.github.com/v3/licenses/#get-a-license
*/
type LicensesGetReq struct {
	_url string

	// license parameter
	License string
}

// HTTPRequest builds an *http.Request
func (r *LicensesGetReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *LicensesGetReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "licenses/get",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/licenses/%v", r.License),
		URLQuery:         query,
		ValidStatuses:    []int{200, 304},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *LicensesGetReq) Rel(link string, resp *LicensesGetResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
LicensesGetResponse is a response for LicensesGet

https://developer.github.com/v3/licenses/#get-a-license
*/
type LicensesGetResponse struct {
	common.Response
	request *LicensesGetReq
	Data    components.License
}

/*
LicensesGetAllCommonlyUsed performs requests for "licenses/get-all-commonly-used"

Get all commonly used licenses.

  GET /licenses

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
func LicensesGetAllCommonlyUsed(ctx context.Context, req *LicensesGetAllCommonlyUsedReq, opt ...options.Option) (*LicensesGetAllCommonlyUsedResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(LicensesGetAllCommonlyUsedReq)
	}
	resp := &LicensesGetAllCommonlyUsedResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = []components.LicenseSimple{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesGetAllCommonlyUsed performs requests for "licenses/get-all-commonly-used"

Get all commonly used licenses.

  GET /licenses

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
func (c Client) LicensesGetAllCommonlyUsed(ctx context.Context, req *LicensesGetAllCommonlyUsedReq, opt ...options.Option) (*LicensesGetAllCommonlyUsedResponse, error) {
	return LicensesGetAllCommonlyUsed(ctx, req, append(c, opt...)...)
}

/*
LicensesGetAllCommonlyUsedReq is request data for Client.LicensesGetAllCommonlyUsed

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
type LicensesGetAllCommonlyUsedReq struct {
	_url     string
	Featured *bool

	// Results per page (max 100)
	PerPage *int64
}

// HTTPRequest builds an *http.Request
func (r *LicensesGetAllCommonlyUsedReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *LicensesGetAllCommonlyUsedReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}
	if r.Featured != nil {
		query.Set("featured", strconv.FormatBool(*r.Featured))
	}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "licenses/get-all-commonly-used",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/licenses"),
		URLQuery:         query,
		ValidStatuses:    []int{200, 304},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *LicensesGetAllCommonlyUsedReq) Rel(link string, resp *LicensesGetAllCommonlyUsedResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
LicensesGetAllCommonlyUsedResponse is a response for LicensesGetAllCommonlyUsed

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses
*/
type LicensesGetAllCommonlyUsedResponse struct {
	common.Response
	request *LicensesGetAllCommonlyUsedReq
	Data    []components.LicenseSimple
}

/*
LicensesGetForRepo performs requests for "licenses/get-for-repo"

Get the license for a repository.

  GET /repos/{owner}/{repo}/license

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository
*/
func LicensesGetForRepo(ctx context.Context, req *LicensesGetForRepoReq, opt ...options.Option) (*LicensesGetForRepoResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(LicensesGetForRepoReq)
	}
	resp := &LicensesGetForRepoResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.LicenseContent{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
LicensesGetForRepo performs requests for "licenses/get-for-repo"

Get the license for a repository.

  GET /repos/{owner}/{repo}/license

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository
*/
func (c Client) LicensesGetForRepo(ctx context.Context, req *LicensesGetForRepoReq, opt ...options.Option) (*LicensesGetForRepoResponse, error) {
	return LicensesGetForRepo(ctx, req, append(c, opt...)...)
}

/*
LicensesGetForRepoReq is request data for Client.LicensesGetForRepo

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository
*/
type LicensesGetForRepoReq struct {
	_url  string
	Owner string
	Repo  string
}

// HTTPRequest builds an *http.Request
func (r *LicensesGetForRepoReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *LicensesGetForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "licenses/get-for-repo",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/repos/%v/%v/license", r.Owner, r.Repo),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *LicensesGetForRepoReq) Rel(link string, resp *LicensesGetForRepoResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
LicensesGetForRepoResponse is a response for LicensesGetForRepo

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository
*/
type LicensesGetForRepoResponse struct {
	common.Response
	request *LicensesGetForRepoReq
	Data    components.LicenseContent
}
