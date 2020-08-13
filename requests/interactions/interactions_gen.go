// Code generated by octo-go; DO NOT EDIT.

package interactions

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	internal "github.com/willabides/octo-go/internal"
	requests "github.com/willabides/octo-go/requests"
	"net/http"
)

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
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetRestrictionsForOrgReq)
	}
	resp := &GetRestrictionsForOrgResponse{}

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
GetRestrictionsForOrg performs requests for "interactions/get-restrictions-for-org"

Get interaction restrictions for an organization.

  GET /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetRestrictionsForOrg(ctx context.Context, req *GetRestrictionsForOrgReq, opt ...requests.Option) (*GetRestrictionsForOrgResponse, error) {
	return GetRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
GetRestrictionsForOrgReq is request data for Client.GetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
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

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:      []string{"sombra"},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": internal.String("application/json")},
		Method:           "GET",
		Options:          opt,
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
	})
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

// HTTPResponse returns the *http.Response
func (r *GetRestrictionsForOrgResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetRestrictionsForOrgResponse) ReadResponse(resp *http.Response) error {
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

/*
GetRestrictionsForRepo performs requests for "interactions/get-restrictions-for-repo"

Get interaction restrictions for a repository.

  GET /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
func GetRestrictionsForRepo(ctx context.Context, req *GetRestrictionsForRepoReq, opt ...requests.Option) (*GetRestrictionsForRepoResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetRestrictionsForRepoReq)
	}
	resp := &GetRestrictionsForRepoResponse{}

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
GetRestrictionsForRepo performs requests for "interactions/get-restrictions-for-repo"

Get interaction restrictions for a repository.

  GET /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetRestrictionsForRepo(ctx context.Context, req *GetRestrictionsForRepoReq, opt ...requests.Option) (*GetRestrictionsForRepoResponse, error) {
	return GetRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
GetRestrictionsForRepoReq is request data for Client.GetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
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

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:      []string{"sombra"},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": internal.String("application/json")},
		Method:           "GET",
		Options:          opt,
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
	})
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

// HTTPResponse returns the *http.Response
func (r *GetRestrictionsForRepoResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetRestrictionsForRepoResponse) ReadResponse(resp *http.Response) error {
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

/*
RemoveRestrictionsForOrg performs requests for "interactions/remove-restrictions-for-org"

Remove interaction restrictions for an organization.

  DELETE /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
func RemoveRestrictionsForOrg(ctx context.Context, req *RemoveRestrictionsForOrgReq, opt ...requests.Option) (*RemoveRestrictionsForOrgResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(RemoveRestrictionsForOrgReq)
	}
	resp := &RemoveRestrictionsForOrgResponse{}

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
RemoveRestrictionsForOrg performs requests for "interactions/remove-restrictions-for-org"

Remove interaction restrictions for an organization.

  DELETE /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) RemoveRestrictionsForOrg(ctx context.Context, req *RemoveRestrictionsForOrgReq, opt ...requests.Option) (*RemoveRestrictionsForOrgResponse, error) {
	return RemoveRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
RemoveRestrictionsForOrgReq is request data for Client.RemoveRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
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

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *RemoveRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:      []string{"sombra"},
		ExplicitURL:      r._url,
		Method:           "DELETE",
		Options:          opt,
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
	})
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

// HTTPResponse returns the *http.Response
func (r *RemoveRestrictionsForOrgResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *RemoveRestrictionsForOrgResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{204})
	if err != nil {
		return err
	}
	return nil
}

/*
RemoveRestrictionsForRepo performs requests for "interactions/remove-restrictions-for-repo"

Remove interaction restrictions for a repository.

  DELETE /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
func RemoveRestrictionsForRepo(ctx context.Context, req *RemoveRestrictionsForRepoReq, opt ...requests.Option) (*RemoveRestrictionsForRepoResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(RemoveRestrictionsForRepoReq)
	}
	resp := &RemoveRestrictionsForRepoResponse{}

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
RemoveRestrictionsForRepo performs requests for "interactions/remove-restrictions-for-repo"

Remove interaction restrictions for a repository.

  DELETE /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) RemoveRestrictionsForRepo(ctx context.Context, req *RemoveRestrictionsForRepoReq, opt ...requests.Option) (*RemoveRestrictionsForRepoResponse, error) {
	return RemoveRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
RemoveRestrictionsForRepoReq is request data for Client.RemoveRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
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

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *RemoveRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews:      []string{"sombra"},
		ExplicitURL:      r._url,
		Method:           "DELETE",
		Options:          opt,
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
	})
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

// HTTPResponse returns the *http.Response
func (r *RemoveRestrictionsForRepoResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *RemoveRestrictionsForRepoResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{204})
	if err != nil {
		return err
	}
	return nil
}

/*
SetRestrictionsForOrg performs requests for "interactions/set-restrictions-for-org"

Set interaction restrictions for an organization.

  PUT /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization
*/
func SetRestrictionsForOrg(ctx context.Context, req *SetRestrictionsForOrgReq, opt ...requests.Option) (*SetRestrictionsForOrgResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(SetRestrictionsForOrgReq)
	}
	resp := &SetRestrictionsForOrgResponse{}

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
SetRestrictionsForOrg performs requests for "interactions/set-restrictions-for-org"

Set interaction restrictions for an organization.

  PUT /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) SetRestrictionsForOrg(ctx context.Context, req *SetRestrictionsForOrgReq, opt ...requests.Option) (*SetRestrictionsForOrgResponse, error) {
	return SetRestrictionsForOrg(ctx, req, append(c, opt...)...)
}

/*
SetRestrictionsForOrgReq is request data for Client.SetRestrictionsForOrg

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
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

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *SetRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews: []string{"sombra"},
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       internal.String("application/json"),
			"content-type": internal.String("application/json"),
		},
		Method:           "PUT",
		Options:          opt,
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/orgs/%v/interaction-limits", r.Org),
	})
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

// HTTPResponse returns the *http.Response
func (r *SetRestrictionsForOrgResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *SetRestrictionsForOrgResponse) ReadResponse(resp *http.Response) error {
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

/*
SetRestrictionsForRepo performs requests for "interactions/set-restrictions-for-repo"

Set interaction restrictions for a repository.

  PUT /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository
*/
func SetRestrictionsForRepo(ctx context.Context, req *SetRestrictionsForRepoReq, opt ...requests.Option) (*SetRestrictionsForRepoResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(SetRestrictionsForRepoReq)
	}
	resp := &SetRestrictionsForRepoResponse{}

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
SetRestrictionsForRepo performs requests for "interactions/set-restrictions-for-repo"

Set interaction restrictions for a repository.

  PUT /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) SetRestrictionsForRepo(ctx context.Context, req *SetRestrictionsForRepoReq, opt ...requests.Option) (*SetRestrictionsForRepoResponse, error) {
	return SetRestrictionsForRepo(ctx, req, append(c, opt...)...)
}

/*
SetRestrictionsForRepoReq is request data for Client.SetRestrictionsForRepo

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
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

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *SetRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		AllPreviews: []string{"sombra"},
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       internal.String("application/json"),
			"content-type": internal.String("application/json"),
		},
		Method:           "PUT",
		Options:          opt,
		Previews:         map[string]bool{"sombra": r.SombraPreview},
		RequiredPreviews: []string{"sombra"},
		URLPath:          fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo),
	})
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

// HTTPResponse returns the *http.Response
func (r *SetRestrictionsForRepoResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *SetRestrictionsForRepoResponse) ReadResponse(resp *http.Response) error {
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
