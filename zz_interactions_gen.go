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
)

/*
InteractionsGetRestrictionsForOrg performs requests for "interactions/get-restrictions-for-org"

Get interaction restrictions for an organization.

  GET /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
func InteractionsGetRestrictionsForOrg(ctx context.Context, req *InteractionsGetRestrictionsForOrgReq, opt ...options.Option) (*InteractionsGetRestrictionsForOrgResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(InteractionsGetRestrictionsForOrgReq)
	}
	resp := &InteractionsGetRestrictionsForOrgResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.InteractionLimit{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
InteractionsGetRestrictionsForOrg performs requests for "interactions/get-restrictions-for-org"

Get interaction restrictions for an organization.

  GET /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
func (c Client) InteractionsGetRestrictionsForOrg(ctx context.Context, req *InteractionsGetRestrictionsForOrgReq, opt ...options.Option) (*InteractionsGetRestrictionsForOrgResponse, error) {
	return InteractionsGetRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
InteractionsGetRestrictionsForOrgReq is request data for Client.InteractionsGetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
type InteractionsGetRestrictionsForOrgReq struct {
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
func (r *InteractionsGetRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *InteractionsGetRestrictionsForOrgReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{"sombra"},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "interactions/get-restrictions-for-org",
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *InteractionsGetRestrictionsForOrgReq) Rel(link string, resp *InteractionsGetRestrictionsForOrgResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
InteractionsGetRestrictionsForOrgResponse is a response for InteractionsGetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
type InteractionsGetRestrictionsForOrgResponse struct {
	common.Response
	request *InteractionsGetRestrictionsForOrgReq
	Data    components.InteractionLimit
}

/*
InteractionsGetRestrictionsForRepo performs requests for "interactions/get-restrictions-for-repo"

Get interaction restrictions for a repository.

  GET /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
func InteractionsGetRestrictionsForRepo(ctx context.Context, req *InteractionsGetRestrictionsForRepoReq, opt ...options.Option) (*InteractionsGetRestrictionsForRepoResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(InteractionsGetRestrictionsForRepoReq)
	}
	resp := &InteractionsGetRestrictionsForRepoResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.InteractionLimit{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
InteractionsGetRestrictionsForRepo performs requests for "interactions/get-restrictions-for-repo"

Get interaction restrictions for a repository.

  GET /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
func (c Client) InteractionsGetRestrictionsForRepo(ctx context.Context, req *InteractionsGetRestrictionsForRepoReq, opt ...options.Option) (*InteractionsGetRestrictionsForRepoResponse, error) {
	return InteractionsGetRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
InteractionsGetRestrictionsForRepoReq is request data for Client.InteractionsGetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
type InteractionsGetRestrictionsForRepoReq struct {
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
func (r *InteractionsGetRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *InteractionsGetRestrictionsForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{"sombra"},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "interactions/get-restrictions-for-repo",
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *InteractionsGetRestrictionsForRepoReq) Rel(link string, resp *InteractionsGetRestrictionsForRepoResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
InteractionsGetRestrictionsForRepoResponse is a response for InteractionsGetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
type InteractionsGetRestrictionsForRepoResponse struct {
	common.Response
	request *InteractionsGetRestrictionsForRepoReq
	Data    components.InteractionLimit
}

/*
InteractionsRemoveRestrictionsForOrg performs requests for "interactions/remove-restrictions-for-org"

Remove interaction restrictions for an organization.

  DELETE /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
func InteractionsRemoveRestrictionsForOrg(ctx context.Context, req *InteractionsRemoveRestrictionsForOrgReq, opt ...options.Option) (*InteractionsRemoveRestrictionsForOrgResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(InteractionsRemoveRestrictionsForOrgReq)
	}
	resp := &InteractionsRemoveRestrictionsForOrgResponse{request: req}
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
InteractionsRemoveRestrictionsForOrg performs requests for "interactions/remove-restrictions-for-org"

Remove interaction restrictions for an organization.

  DELETE /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
func (c Client) InteractionsRemoveRestrictionsForOrg(ctx context.Context, req *InteractionsRemoveRestrictionsForOrgReq, opt ...options.Option) (*InteractionsRemoveRestrictionsForOrgResponse, error) {
	return InteractionsRemoveRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
InteractionsRemoveRestrictionsForOrgReq is request data for Client.InteractionsRemoveRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
type InteractionsRemoveRestrictionsForOrgReq struct {
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
func (r *InteractionsRemoveRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *InteractionsRemoveRestrictionsForOrgReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{"sombra"},
		Body:             nil,
		DataStatuses:     []int{},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{},
		Method:           "DELETE",
		OperationID:      "interactions/remove-restrictions-for-org",
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
		URLQuery:         query,
		ValidStatuses:    []int{204},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *InteractionsRemoveRestrictionsForOrgReq) Rel(link string, resp *InteractionsRemoveRestrictionsForOrgResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
InteractionsRemoveRestrictionsForOrgResponse is a response for InteractionsRemoveRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
type InteractionsRemoveRestrictionsForOrgResponse struct {
	common.Response
	request *InteractionsRemoveRestrictionsForOrgReq
}

/*
InteractionsRemoveRestrictionsForRepo performs requests for "interactions/remove-restrictions-for-repo"

Remove interaction restrictions for a repository.

  DELETE /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
func InteractionsRemoveRestrictionsForRepo(ctx context.Context, req *InteractionsRemoveRestrictionsForRepoReq, opt ...options.Option) (*InteractionsRemoveRestrictionsForRepoResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(InteractionsRemoveRestrictionsForRepoReq)
	}
	resp := &InteractionsRemoveRestrictionsForRepoResponse{request: req}
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
InteractionsRemoveRestrictionsForRepo performs requests for "interactions/remove-restrictions-for-repo"

Remove interaction restrictions for a repository.

  DELETE /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
func (c Client) InteractionsRemoveRestrictionsForRepo(ctx context.Context, req *InteractionsRemoveRestrictionsForRepoReq, opt ...options.Option) (*InteractionsRemoveRestrictionsForRepoResponse, error) {
	return InteractionsRemoveRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
InteractionsRemoveRestrictionsForRepoReq is request data for Client.InteractionsRemoveRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
type InteractionsRemoveRestrictionsForRepoReq struct {
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
func (r *InteractionsRemoveRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *InteractionsRemoveRestrictionsForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{"sombra"},
		Body:             nil,
		DataStatuses:     []int{},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{},
		Method:           "DELETE",
		OperationID:      "interactions/remove-restrictions-for-repo",
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
		URLQuery:         query,
		ValidStatuses:    []int{204},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *InteractionsRemoveRestrictionsForRepoReq) Rel(link string, resp *InteractionsRemoveRestrictionsForRepoResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
InteractionsRemoveRestrictionsForRepoResponse is a response for InteractionsRemoveRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
type InteractionsRemoveRestrictionsForRepoResponse struct {
	common.Response
	request *InteractionsRemoveRestrictionsForRepoReq
}

/*
InteractionsSetRestrictionsForOrg performs requests for "interactions/set-restrictions-for-org"

Set interaction restrictions for an organization.

  PUT /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
func InteractionsSetRestrictionsForOrg(ctx context.Context, req *InteractionsSetRestrictionsForOrgReq, opt ...options.Option) (*InteractionsSetRestrictionsForOrgResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(InteractionsSetRestrictionsForOrgReq)
	}
	resp := &InteractionsSetRestrictionsForOrgResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.InteractionLimit{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
InteractionsSetRestrictionsForOrg performs requests for "interactions/set-restrictions-for-org"

Set interaction restrictions for an organization.

  PUT /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
func (c Client) InteractionsSetRestrictionsForOrg(ctx context.Context, req *InteractionsSetRestrictionsForOrgReq, opt ...options.Option) (*InteractionsSetRestrictionsForOrgResponse, error) {
	return InteractionsSetRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
InteractionsSetRestrictionsForOrgReq is request data for Client.InteractionsSetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
type InteractionsSetRestrictionsForOrgReq struct {
	_url        string
	Org         string
	RequestBody InteractionsSetRestrictionsForOrgReqBody

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

// HTTPRequest builds an *http.Request
func (r *InteractionsSetRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *InteractionsSetRestrictionsForOrgReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:  []string{"sombra"},
		Body:         r.RequestBody,
		DataStatuses: []int{200},
		ExplicitURL:  r._url,
		HeaderVals: map[string]*string{
			"accept":       String("application/json"),
			"content-type": String("application/json"),
		},
		Method:           "PUT",
		OperationID:      "interactions/set-restrictions-for-org",
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *InteractionsSetRestrictionsForOrgReq) Rel(link string, resp *InteractionsSetRestrictionsForOrgResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
InteractionsSetRestrictionsForOrgReqBody is a request body for interactions/set-restrictions-for-org

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
type InteractionsSetRestrictionsForOrgReqBody struct {

	/*
	Specifies the group of GitHub users who can comment, open issues, or create pull
	requests in public repositories for the given organization. Must be one of:
	`existing_users`, `contributors_only`, or `collaborators_only`.
	*/
	Limit *string `json:"limit"`
}

/*
InteractionsSetRestrictionsForOrgResponse is a response for InteractionsSetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
type InteractionsSetRestrictionsForOrgResponse struct {
	common.Response
	request *InteractionsSetRestrictionsForOrgReq
	Data    components.InteractionLimit
}

/*
InteractionsSetRestrictionsForRepo performs requests for "interactions/set-restrictions-for-repo"

Set interaction restrictions for a repository.

  PUT /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
func InteractionsSetRestrictionsForRepo(ctx context.Context, req *InteractionsSetRestrictionsForRepoReq, opt ...options.Option) (*InteractionsSetRestrictionsForRepoResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(InteractionsSetRestrictionsForRepoReq)
	}
	resp := &InteractionsSetRestrictionsForRepoResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.InteractionLimit{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
InteractionsSetRestrictionsForRepo performs requests for "interactions/set-restrictions-for-repo"

Set interaction restrictions for a repository.

  PUT /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
func (c Client) InteractionsSetRestrictionsForRepo(ctx context.Context, req *InteractionsSetRestrictionsForRepoReq, opt ...options.Option) (*InteractionsSetRestrictionsForRepoResponse, error) {
	return InteractionsSetRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
InteractionsSetRestrictionsForRepoReq is request data for Client.InteractionsSetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
type InteractionsSetRestrictionsForRepoReq struct {
	_url        string
	Owner       string
	Repo        string
	RequestBody InteractionsSetRestrictionsForRepoReqBody

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

// HTTPRequest builds an *http.Request
func (r *InteractionsSetRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *InteractionsSetRestrictionsForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:  []string{"sombra"},
		Body:         r.RequestBody,
		DataStatuses: []int{200},
		ExplicitURL:  r._url,
		HeaderVals: map[string]*string{
			"accept":       String("application/json"),
			"content-type": String("application/json"),
		},
		Method:           "PUT",
		OperationID:      "interactions/set-restrictions-for-repo",
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *InteractionsSetRestrictionsForRepoReq) Rel(link string, resp *InteractionsSetRestrictionsForRepoResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
InteractionsSetRestrictionsForRepoReqBody is a request body for interactions/set-restrictions-for-repo

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
type InteractionsSetRestrictionsForRepoReqBody struct {

	/*
	Specifies the group of GitHub users who can comment, open issues, or create pull
	requests for the given repository. Must be one of: `existing_users`,
	`contributors_only`, or `collaborators_only`.
	*/
	Limit *string `json:"limit"`
}

/*
InteractionsSetRestrictionsForRepoResponse is a response for InteractionsSetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
type InteractionsSetRestrictionsForRepoResponse struct {
	common.Response
	request *InteractionsSetRestrictionsForRepoReq
	Data    components.InteractionLimit
}
