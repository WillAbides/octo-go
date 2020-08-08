// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	"net/http"
	"net/url"
)

/*
CodesOfConductGetAllCodesOfConduct performs requests for "codes-of-conduct/get-all-codes-of-conduct"

Get all codes of conduct.

  GET /codes_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct
*/
func CodesOfConductGetAllCodesOfConduct(ctx context.Context, req *CodesOfConductGetAllCodesOfConductReq, opt ...RequestOption) (*CodesOfConductGetAllCodesOfConductResponse, error) {
	if req == nil {
		req = new(CodesOfConductGetAllCodesOfConductReq)
	}
	resp := &CodesOfConductGetAllCodesOfConductResponse{request: req}
	r, err := doRequest(ctx, req, "codes-of-conduct/get-all-codes-of-conduct", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = []components.CodeOfConduct{}
	err = r.decodeBody(&resp.Data, "codes-of-conduct/get-all-codes-of-conduct")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodesOfConductGetAllCodesOfConduct performs requests for "codes-of-conduct/get-all-codes-of-conduct"

Get all codes of conduct.

  GET /codes_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct
*/
func (c Client) CodesOfConductGetAllCodesOfConduct(ctx context.Context, req *CodesOfConductGetAllCodesOfConductReq, opt ...RequestOption) (*CodesOfConductGetAllCodesOfConductResponse, error) {
	return CodesOfConductGetAllCodesOfConduct(ctx, req, append(c, opt...)...)
}

/*
CodesOfConductGetAllCodesOfConductReq is request data for Client.CodesOfConductGetAllCodesOfConduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct
*/
type CodesOfConductGetAllCodesOfConductReq struct {
	_url string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

func (r *CodesOfConductGetAllCodesOfConductReq) url() string {
	return r._url
}

func (r *CodesOfConductGetAllCodesOfConductReq) urlPath() string {
	return fmt.Sprintf("/codes_of_conduct")
}

func (r *CodesOfConductGetAllCodesOfConductReq) method() string {
	return "GET"
}

func (r *CodesOfConductGetAllCodesOfConductReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *CodesOfConductGetAllCodesOfConductReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{"scarlet-witch": r.ScarletWitchPreview}
	if requiredPreviews {
		previewVals["scarlet-witch"] = true
	}
	if allPreviews {
		previewVals["scarlet-witch"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *CodesOfConductGetAllCodesOfConductReq) body() interface{} {
	return nil
}

func (r *CodesOfConductGetAllCodesOfConductReq) dataStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetAllCodesOfConductReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *CodesOfConductGetAllCodesOfConductReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "codes-of-conduct/get-all-codes-of-conduct", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodesOfConductGetAllCodesOfConductReq) Rel(link RelName, resp *CodesOfConductGetAllCodesOfConductResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
CodesOfConductGetAllCodesOfConductResponse is a response for CodesOfConductGetAllCodesOfConduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct
*/
type CodesOfConductGetAllCodesOfConductResponse struct {
	response
	request *CodesOfConductGetAllCodesOfConductReq
	Data    []components.CodeOfConduct
}

/*
CodesOfConductGetConductCode performs requests for "codes-of-conduct/get-conduct-code"

Get a code of conduct.

  GET /codes_of_conduct/{key}

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct
*/
func CodesOfConductGetConductCode(ctx context.Context, req *CodesOfConductGetConductCodeReq, opt ...RequestOption) (*CodesOfConductGetConductCodeResponse, error) {
	if req == nil {
		req = new(CodesOfConductGetConductCodeReq)
	}
	resp := &CodesOfConductGetConductCodeResponse{request: req}
	r, err := doRequest(ctx, req, "codes-of-conduct/get-conduct-code", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.CodeOfConduct{}
	err = r.decodeBody(&resp.Data, "codes-of-conduct/get-conduct-code")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodesOfConductGetConductCode performs requests for "codes-of-conduct/get-conduct-code"

Get a code of conduct.

  GET /codes_of_conduct/{key}

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct
*/
func (c Client) CodesOfConductGetConductCode(ctx context.Context, req *CodesOfConductGetConductCodeReq, opt ...RequestOption) (*CodesOfConductGetConductCodeResponse, error) {
	return CodesOfConductGetConductCode(ctx, req, append(c, opt...)...)
}

/*
CodesOfConductGetConductCodeReq is request data for Client.CodesOfConductGetConductCode

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct
*/
type CodesOfConductGetConductCodeReq struct {
	_url string

	// key parameter
	Key string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

func (r *CodesOfConductGetConductCodeReq) url() string {
	return r._url
}

func (r *CodesOfConductGetConductCodeReq) urlPath() string {
	return fmt.Sprintf("/codes_of_conduct/%v", r.Key)
}

func (r *CodesOfConductGetConductCodeReq) method() string {
	return "GET"
}

func (r *CodesOfConductGetConductCodeReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *CodesOfConductGetConductCodeReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{"scarlet-witch": r.ScarletWitchPreview}
	if requiredPreviews {
		previewVals["scarlet-witch"] = true
	}
	if allPreviews {
		previewVals["scarlet-witch"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *CodesOfConductGetConductCodeReq) body() interface{} {
	return nil
}

func (r *CodesOfConductGetConductCodeReq) dataStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetConductCodeReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *CodesOfConductGetConductCodeReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "codes-of-conduct/get-conduct-code", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodesOfConductGetConductCodeReq) Rel(link RelName, resp *CodesOfConductGetConductCodeResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
CodesOfConductGetConductCodeResponse is a response for CodesOfConductGetConductCode

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct
*/
type CodesOfConductGetConductCodeResponse struct {
	response
	request *CodesOfConductGetConductCodeReq
	Data    components.CodeOfConduct
}

/*
CodesOfConductGetForRepo performs requests for "codes-of-conduct/get-for-repo"

Get the code of conduct for a repository.

  GET /repos/{owner}/{repo}/community/code_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository
*/
func CodesOfConductGetForRepo(ctx context.Context, req *CodesOfConductGetForRepoReq, opt ...RequestOption) (*CodesOfConductGetForRepoResponse, error) {
	if req == nil {
		req = new(CodesOfConductGetForRepoReq)
	}
	resp := &CodesOfConductGetForRepoResponse{request: req}
	r, err := doRequest(ctx, req, "codes-of-conduct/get-for-repo", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.CodeOfConduct{}
	err = r.decodeBody(&resp.Data, "codes-of-conduct/get-for-repo")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodesOfConductGetForRepo performs requests for "codes-of-conduct/get-for-repo"

Get the code of conduct for a repository.

  GET /repos/{owner}/{repo}/community/code_of_conduct

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository
*/
func (c Client) CodesOfConductGetForRepo(ctx context.Context, req *CodesOfConductGetForRepoReq, opt ...RequestOption) (*CodesOfConductGetForRepoResponse, error) {
	return CodesOfConductGetForRepo(ctx, req, append(c, opt...)...)
}

/*
CodesOfConductGetForRepoReq is request data for Client.CodesOfConductGetForRepo

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository
*/
type CodesOfConductGetForRepoReq struct {
	_url  string
	Owner string
	Repo  string

	/*
	The Codes of Conduct API is currently available for developers to preview.

	To access the API during the preview period, you must set this to true.
	*/
	ScarletWitchPreview bool
}

func (r *CodesOfConductGetForRepoReq) url() string {
	return r._url
}

func (r *CodesOfConductGetForRepoReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/community/code_of_conduct", r.Owner, r.Repo)
}

func (r *CodesOfConductGetForRepoReq) method() string {
	return "GET"
}

func (r *CodesOfConductGetForRepoReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *CodesOfConductGetForRepoReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{"scarlet-witch": r.ScarletWitchPreview}
	if requiredPreviews {
		previewVals["scarlet-witch"] = true
	}
	if allPreviews {
		previewVals["scarlet-witch"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *CodesOfConductGetForRepoReq) body() interface{} {
	return nil
}

func (r *CodesOfConductGetForRepoReq) dataStatuses() []int {
	return []int{200}
}

func (r *CodesOfConductGetForRepoReq) validStatuses() []int {
	return []int{200}
}

// HTTPRequest builds an *http.Request
func (r *CodesOfConductGetForRepoReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "codes-of-conduct/get-for-repo", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodesOfConductGetForRepoReq) Rel(link RelName, resp *CodesOfConductGetForRepoResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
CodesOfConductGetForRepoResponse is a response for CodesOfConductGetForRepo

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository
*/
type CodesOfConductGetForRepoResponse struct {
	response
	request *CodesOfConductGetForRepoReq
	Data    components.CodeOfConduct
}
