// Code generated by octo-go; DO NOT EDIT.

package interactions

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
GetRestrictionsForOrg performs requests for "interactions/get-restrictions-for-org"

Get interaction restrictions for an organization.

  GET /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
func GetRestrictionsForOrg(ctx context.Context, req *GetRestrictionsForOrgReq, opt ...requests.Option) (*GetRestrictionsForOrgResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(GetRestrictionsForOrgReq)
	}
	resp := &GetRestrictionsForOrgResponse{}
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

	return NewGetRestrictionsForOrgResponse(r, opts.PreserveResponseBody())
}

// NewGetRestrictionsForOrgResponse builds a new *GetRestrictionsForOrgResponse from an *http.Response
func NewGetRestrictionsForOrgResponse(resp *http.Response, preserveBody bool) (*GetRestrictionsForOrgResponse, error) {
	var result GetRestrictionsForOrgResponse
	result.httpResponse = resp
	err := internal.ErrorCheck(resp, []int{200})
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
GetRestrictionsForOrg performs requests for "interactions/get-restrictions-for-org"

Get interaction restrictions for an organization.

  GET /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
func (c Client) GetRestrictionsForOrg(ctx context.Context, req *GetRestrictionsForOrgReq, opt ...requests.Option) (*GetRestrictionsForOrgResponse, error) {
	return GetRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
GetRestrictionsForOrgReq is request data for Client.GetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
type GetRestrictionsForOrgReq struct {
	_url string
	Org  string

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

// HTTPRequest builds an *http.Request
func (r *GetRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *GetRestrictionsForOrgReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{"sombra"},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"accept": internal.String("application/json")},
		Method:             "GET",
		OperationID:        "interactions/get-restrictions-for-org",
		Previews:           map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews:   []string{"sombra"},
		URLPath:            fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
		URLQuery:           query,
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetRestrictionsForOrgReq) Rel(link string, resp *GetRestrictionsForOrgResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetRestrictionsForOrgResponse is a response for GetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
type GetRestrictionsForOrgResponse struct {
	httpResponse *http.Response
	Data         components.InteractionLimit
}

func (r *GetRestrictionsForOrgResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

/*
GetRestrictionsForRepo performs requests for "interactions/get-restrictions-for-repo"

Get interaction restrictions for a repository.

  GET /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
func GetRestrictionsForRepo(ctx context.Context, req *GetRestrictionsForRepoReq, opt ...requests.Option) (*GetRestrictionsForRepoResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(GetRestrictionsForRepoReq)
	}
	resp := &GetRestrictionsForRepoResponse{}
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

	return NewGetRestrictionsForRepoResponse(r, opts.PreserveResponseBody())
}

// NewGetRestrictionsForRepoResponse builds a new *GetRestrictionsForRepoResponse from an *http.Response
func NewGetRestrictionsForRepoResponse(resp *http.Response, preserveBody bool) (*GetRestrictionsForRepoResponse, error) {
	var result GetRestrictionsForRepoResponse
	result.httpResponse = resp
	err := internal.ErrorCheck(resp, []int{200})
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
GetRestrictionsForRepo performs requests for "interactions/get-restrictions-for-repo"

Get interaction restrictions for a repository.

  GET /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
func (c Client) GetRestrictionsForRepo(ctx context.Context, req *GetRestrictionsForRepoReq, opt ...requests.Option) (*GetRestrictionsForRepoResponse, error) {
	return GetRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
GetRestrictionsForRepoReq is request data for Client.GetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
type GetRestrictionsForRepoReq struct {
	_url  string
	Owner string
	Repo  string

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

// HTTPRequest builds an *http.Request
func (r *GetRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *GetRestrictionsForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{"sombra"},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{"accept": internal.String("application/json")},
		Method:             "GET",
		OperationID:        "interactions/get-restrictions-for-repo",
		Previews:           map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews:   []string{"sombra"},
		URLPath:            fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
		URLQuery:           query,
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetRestrictionsForRepoReq) Rel(link string, resp *GetRestrictionsForRepoResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetRestrictionsForRepoResponse is a response for GetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
type GetRestrictionsForRepoResponse struct {
	httpResponse *http.Response
	Data         components.InteractionLimit
}

func (r *GetRestrictionsForRepoResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

/*
RemoveRestrictionsForOrg performs requests for "interactions/remove-restrictions-for-org"

Remove interaction restrictions for an organization.

  DELETE /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
func RemoveRestrictionsForOrg(ctx context.Context, req *RemoveRestrictionsForOrgReq, opt ...requests.Option) (*RemoveRestrictionsForOrgResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(RemoveRestrictionsForOrgReq)
	}
	resp := &RemoveRestrictionsForOrgResponse{}
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

	return NewRemoveRestrictionsForOrgResponse(r, opts.PreserveResponseBody())
}

// NewRemoveRestrictionsForOrgResponse builds a new *RemoveRestrictionsForOrgResponse from an *http.Response
func NewRemoveRestrictionsForOrgResponse(resp *http.Response, preserveBody bool) (*RemoveRestrictionsForOrgResponse, error) {
	var result RemoveRestrictionsForOrgResponse
	result.httpResponse = resp
	err := internal.ErrorCheck(resp, []int{204})
	if err != nil {
		return &result, err
	}
	return &result, nil
}

/*
RemoveRestrictionsForOrg performs requests for "interactions/remove-restrictions-for-org"

Remove interaction restrictions for an organization.

  DELETE /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
func (c Client) RemoveRestrictionsForOrg(ctx context.Context, req *RemoveRestrictionsForOrgReq, opt ...requests.Option) (*RemoveRestrictionsForOrgResponse, error) {
	return RemoveRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
RemoveRestrictionsForOrgReq is request data for Client.RemoveRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
type RemoveRestrictionsForOrgReq struct {
	_url string
	Org  string

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

// HTTPRequest builds an *http.Request
func (r *RemoveRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *RemoveRestrictionsForOrgReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{"sombra"},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{},
		Method:             "DELETE",
		OperationID:        "interactions/remove-restrictions-for-org",
		Previews:           map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews:   []string{"sombra"},
		URLPath:            fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
		URLQuery:           query,
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RemoveRestrictionsForOrgReq) Rel(link string, resp *RemoveRestrictionsForOrgResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
RemoveRestrictionsForOrgResponse is a response for RemoveRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
type RemoveRestrictionsForOrgResponse struct {
	httpResponse *http.Response
}

func (r *RemoveRestrictionsForOrgResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

/*
RemoveRestrictionsForRepo performs requests for "interactions/remove-restrictions-for-repo"

Remove interaction restrictions for a repository.

  DELETE /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
func RemoveRestrictionsForRepo(ctx context.Context, req *RemoveRestrictionsForRepoReq, opt ...requests.Option) (*RemoveRestrictionsForRepoResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(RemoveRestrictionsForRepoReq)
	}
	resp := &RemoveRestrictionsForRepoResponse{}
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

	return NewRemoveRestrictionsForRepoResponse(r, opts.PreserveResponseBody())
}

// NewRemoveRestrictionsForRepoResponse builds a new *RemoveRestrictionsForRepoResponse from an *http.Response
func NewRemoveRestrictionsForRepoResponse(resp *http.Response, preserveBody bool) (*RemoveRestrictionsForRepoResponse, error) {
	var result RemoveRestrictionsForRepoResponse
	result.httpResponse = resp
	err := internal.ErrorCheck(resp, []int{204})
	if err != nil {
		return &result, err
	}
	return &result, nil
}

/*
RemoveRestrictionsForRepo performs requests for "interactions/remove-restrictions-for-repo"

Remove interaction restrictions for a repository.

  DELETE /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
func (c Client) RemoveRestrictionsForRepo(ctx context.Context, req *RemoveRestrictionsForRepoReq, opt ...requests.Option) (*RemoveRestrictionsForRepoResponse, error) {
	return RemoveRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
RemoveRestrictionsForRepoReq is request data for Client.RemoveRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
type RemoveRestrictionsForRepoReq struct {
	_url  string
	Owner string
	Repo  string

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

// HTTPRequest builds an *http.Request
func (r *RemoveRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *RemoveRestrictionsForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{"sombra"},
		Body:               nil,
		EndpointAttributes: []internal.EndpointAttribute{},
		ExplicitURL:        r._url,
		HeaderVals:         map[string]*string{},
		Method:             "DELETE",
		OperationID:        "interactions/remove-restrictions-for-repo",
		Previews:           map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews:   []string{"sombra"},
		URLPath:            fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
		URLQuery:           query,
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *RemoveRestrictionsForRepoReq) Rel(link string, resp *RemoveRestrictionsForRepoResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
RemoveRestrictionsForRepoResponse is a response for RemoveRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
type RemoveRestrictionsForRepoResponse struct {
	httpResponse *http.Response
}

func (r *RemoveRestrictionsForRepoResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

/*
SetRestrictionsForOrg performs requests for "interactions/set-restrictions-for-org"

Set interaction restrictions for an organization.

  PUT /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
func SetRestrictionsForOrg(ctx context.Context, req *SetRestrictionsForOrgReq, opt ...requests.Option) (*SetRestrictionsForOrgResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(SetRestrictionsForOrgReq)
	}
	resp := &SetRestrictionsForOrgResponse{}
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

	return NewSetRestrictionsForOrgResponse(r, opts.PreserveResponseBody())
}

// NewSetRestrictionsForOrgResponse builds a new *SetRestrictionsForOrgResponse from an *http.Response
func NewSetRestrictionsForOrgResponse(resp *http.Response, preserveBody bool) (*SetRestrictionsForOrgResponse, error) {
	var result SetRestrictionsForOrgResponse
	result.httpResponse = resp
	err := internal.ErrorCheck(resp, []int{200})
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
SetRestrictionsForOrg performs requests for "interactions/set-restrictions-for-org"

Set interaction restrictions for an organization.

  PUT /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
func (c Client) SetRestrictionsForOrg(ctx context.Context, req *SetRestrictionsForOrgReq, opt ...requests.Option) (*SetRestrictionsForOrgResponse, error) {
	return SetRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
SetRestrictionsForOrgReq is request data for Client.SetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
type SetRestrictionsForOrgReq struct {
	_url        string
	Org         string
	RequestBody SetRestrictionsForOrgReqBody

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

// HTTPRequest builds an *http.Request
func (r *SetRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *SetRestrictionsForOrgReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{"sombra"},
		Body:               r.RequestBody,
		EndpointAttributes: []internal.EndpointAttribute{internal.AttrJSONRequestBody},
		ExplicitURL:        r._url,
		HeaderVals: map[string]*string{
			"accept":       internal.String("application/json"),
			"content-type": internal.String("application/json"),
		},
		Method:           "PUT",
		OperationID:      "interactions/set-restrictions-for-org",
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
		URLQuery:         query,
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SetRestrictionsForOrgReq) Rel(link string, resp *SetRestrictionsForOrgResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SetRestrictionsForOrgReqBody is a request body for interactions/set-restrictions-for-org

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
type SetRestrictionsForOrgReqBody struct {

	/*
	Specifies the group of GitHub users who can comment, open issues, or create pull
	requests in public repositories for the given organization. Must be one of:
	`existing_users`, `contributors_only`, or `collaborators_only`.
	*/
	Limit *string `json:"limit"`
}

/*
SetRestrictionsForOrgResponse is a response for SetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
type SetRestrictionsForOrgResponse struct {
	httpResponse *http.Response
	Data         components.InteractionLimit
}

func (r *SetRestrictionsForOrgResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

/*
SetRestrictionsForRepo performs requests for "interactions/set-restrictions-for-repo"

Set interaction restrictions for a repository.

  PUT /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
func SetRestrictionsForRepo(ctx context.Context, req *SetRestrictionsForRepoReq, opt ...requests.Option) (*SetRestrictionsForRepoResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(SetRestrictionsForRepoReq)
	}
	resp := &SetRestrictionsForRepoResponse{}
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

	return NewSetRestrictionsForRepoResponse(r, opts.PreserveResponseBody())
}

// NewSetRestrictionsForRepoResponse builds a new *SetRestrictionsForRepoResponse from an *http.Response
func NewSetRestrictionsForRepoResponse(resp *http.Response, preserveBody bool) (*SetRestrictionsForRepoResponse, error) {
	var result SetRestrictionsForRepoResponse
	result.httpResponse = resp
	err := internal.ErrorCheck(resp, []int{200})
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
SetRestrictionsForRepo performs requests for "interactions/set-restrictions-for-repo"

Set interaction restrictions for a repository.

  PUT /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
func (c Client) SetRestrictionsForRepo(ctx context.Context, req *SetRestrictionsForRepoReq, opt ...requests.Option) (*SetRestrictionsForRepoResponse, error) {
	return SetRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
SetRestrictionsForRepoReq is request data for Client.SetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
type SetRestrictionsForRepoReq struct {
	_url        string
	Owner       string
	Repo        string
	RequestBody SetRestrictionsForRepoReqBody

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

// HTTPRequest builds an *http.Request
func (r *SetRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *SetRestrictionsForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:        []string{"sombra"},
		Body:               r.RequestBody,
		EndpointAttributes: []internal.EndpointAttribute{internal.AttrJSONRequestBody},
		ExplicitURL:        r._url,
		HeaderVals: map[string]*string{
			"accept":       internal.String("application/json"),
			"content-type": internal.String("application/json"),
		},
		Method:           "PUT",
		OperationID:      "interactions/set-restrictions-for-repo",
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
		URLQuery:         query,
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SetRestrictionsForRepoReq) Rel(link string, resp *SetRestrictionsForRepoResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SetRestrictionsForRepoReqBody is a request body for interactions/set-restrictions-for-repo

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
type SetRestrictionsForRepoReqBody struct {

	/*
	Specifies the group of GitHub users who can comment, open issues, or create pull
	requests for the given repository. Must be one of: `existing_users`,
	`contributors_only`, or `collaborators_only`.
	*/
	Limit *string `json:"limit"`
}

/*
SetRestrictionsForRepoResponse is a response for SetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
type SetRestrictionsForRepoResponse struct {
	httpResponse *http.Response
	Data         components.InteractionLimit
}

func (r *SetRestrictionsForRepoResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}
