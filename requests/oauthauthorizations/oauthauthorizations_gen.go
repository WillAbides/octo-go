// Code generated by octo-go; DO NOT EDIT.

package oauthauthorizations

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	requests "github.com/willabides/octo-go/requests"
	"net/http"
	"net/url"
	"strconv"
)

/*
CreateAuthorization performs requests for "oauth-authorizations/create-authorization"

Create a new authorization.

  POST /authorizations

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
*/
func CreateAuthorization(ctx context.Context, req *CreateAuthorizationReq, opt ...requests.Option) (*CreateAuthorizationResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(CreateAuthorizationReq)
	}
	resp := &CreateAuthorizationResponse{}

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
CreateAuthorization performs requests for "oauth-authorizations/create-authorization"

Create a new authorization.

  POST /authorizations

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) CreateAuthorization(ctx context.Context, req *CreateAuthorizationReq, opt ...requests.Option) (*CreateAuthorizationResponse, error) {
	return CreateAuthorization(ctx, req, append(c, opt...)...)
}

/*
CreateAuthorizationReq is request data for Client.CreateAuthorization

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type CreateAuthorizationReq struct {
	_url        string
	RequestBody CreateAuthorizationReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *CreateAuthorizationReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       strPtr("application/json"),
			"content-type": strPtr("application/json"),
		},
		Method:  "POST",
		Options: opt,
		URLPath: fmt.Sprintf("/authorizations"),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CreateAuthorizationReq) Rel(link string, resp *CreateAuthorizationResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
CreateAuthorizationReqBody is a request body for oauth-authorizations/create-authorization

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
*/
type CreateAuthorizationReqBody struct {

	// The OAuth app client key for which to create the token.
	ClientId *string `json:"client_id,omitempty"`

	// The OAuth app client secret for which to create the token.
	ClientSecret *string `json:"client_secret,omitempty"`

	// A unique string to distinguish an authorization from others created for the same client ID and user.
	Fingerprint *string `json:"fingerprint,omitempty"`

	// A note to remind you what the OAuth token is for.
	Note *string `json:"note,omitempty"`

	// A URL to remind you what app the OAuth token is for.
	NoteUrl *string `json:"note_url,omitempty"`

	// A list of scopes that this authorization is in.
	Scopes []string `json:"scopes,omitempty"`
}

/*
CreateAuthorizationResponse is a response for CreateAuthorization

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
*/
type CreateAuthorizationResponse struct {
	httpResponse *http.Response
	Data         components.Authorization
}

// HTTPResponse returns the *http.Response
func (r *CreateAuthorizationResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *CreateAuthorizationResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{201, 304})
	if err != nil {
		return err
	}
	if intInSlice(resp.StatusCode, []int{201}) {
		err = unmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
DeleteAuthorization performs requests for "oauth-authorizations/delete-authorization"

Delete an authorization.

  DELETE /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization
*/
func DeleteAuthorization(ctx context.Context, req *DeleteAuthorizationReq, opt ...requests.Option) (*DeleteAuthorizationResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(DeleteAuthorizationReq)
	}
	resp := &DeleteAuthorizationResponse{}

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
DeleteAuthorization performs requests for "oauth-authorizations/delete-authorization"

Delete an authorization.

  DELETE /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) DeleteAuthorization(ctx context.Context, req *DeleteAuthorizationReq, opt ...requests.Option) (*DeleteAuthorizationResponse, error) {
	return DeleteAuthorization(ctx, req, append(c, opt...)...)
}

/*
DeleteAuthorizationReq is request data for Client.DeleteAuthorization

https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type DeleteAuthorizationReq struct {
	_url string

	// authorization_id parameter
	AuthorizationId int64
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *DeleteAuthorizationReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		Method:      "DELETE",
		Options:     opt,
		URLPath:     fmt.Sprintf("/authorizations/%v", r.AuthorizationId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *DeleteAuthorizationReq) Rel(link string, resp *DeleteAuthorizationResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
DeleteAuthorizationResponse is a response for DeleteAuthorization

https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization
*/
type DeleteAuthorizationResponse struct {
	httpResponse *http.Response
}

// HTTPResponse returns the *http.Response
func (r *DeleteAuthorizationResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *DeleteAuthorizationResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{204, 304})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteGrant performs requests for "oauth-authorizations/delete-grant"

Delete a grant.

  DELETE /applications/grants/{grant_id}

https://developer.github.com/v3/oauth_authorizations/#delete-a-grant
*/
func DeleteGrant(ctx context.Context, req *DeleteGrantReq, opt ...requests.Option) (*DeleteGrantResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(DeleteGrantReq)
	}
	resp := &DeleteGrantResponse{}

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
DeleteGrant performs requests for "oauth-authorizations/delete-grant"

Delete a grant.

  DELETE /applications/grants/{grant_id}

https://developer.github.com/v3/oauth_authorizations/#delete-a-grant

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) DeleteGrant(ctx context.Context, req *DeleteGrantReq, opt ...requests.Option) (*DeleteGrantResponse, error) {
	return DeleteGrant(ctx, req, append(c, opt...)...)
}

/*
DeleteGrantReq is request data for Client.DeleteGrant

https://developer.github.com/v3/oauth_authorizations/#delete-a-grant

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type DeleteGrantReq struct {
	_url string

	// grant_id parameter
	GrantId int64
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *DeleteGrantReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		Method:      "DELETE",
		Options:     opt,
		URLPath:     fmt.Sprintf("/applications/grants/%v", r.GrantId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *DeleteGrantReq) Rel(link string, resp *DeleteGrantResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
DeleteGrantResponse is a response for DeleteGrant

https://developer.github.com/v3/oauth_authorizations/#delete-a-grant
*/
type DeleteGrantResponse struct {
	httpResponse *http.Response
}

// HTTPResponse returns the *http.Response
func (r *DeleteGrantResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *DeleteGrantResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{204, 304})
	if err != nil {
		return err
	}
	return nil
}

/*
GetAuthorization performs requests for "oauth-authorizations/get-authorization"

Get a single authorization.

  GET /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization
*/
func GetAuthorization(ctx context.Context, req *GetAuthorizationReq, opt ...requests.Option) (*GetAuthorizationResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetAuthorizationReq)
	}
	resp := &GetAuthorizationResponse{}

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
GetAuthorization performs requests for "oauth-authorizations/get-authorization"

Get a single authorization.

  GET /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetAuthorization(ctx context.Context, req *GetAuthorizationReq, opt ...requests.Option) (*GetAuthorizationResponse, error) {
	return GetAuthorization(ctx, req, append(c, opt...)...)
}

/*
GetAuthorizationReq is request data for Client.GetAuthorization

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetAuthorizationReq struct {
	_url string

	// authorization_id parameter
	AuthorizationId int64
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetAuthorizationReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": strPtr("application/json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/authorizations/%v", r.AuthorizationId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetAuthorizationReq) Rel(link string, resp *GetAuthorizationResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetAuthorizationResponse is a response for GetAuthorization

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization
*/
type GetAuthorizationResponse struct {
	httpResponse *http.Response
	Data         components.Authorization
}

// HTTPResponse returns the *http.Response
func (r *GetAuthorizationResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetAuthorizationResponse) ReadResponse(resp *http.Response) error {
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
GetGrant performs requests for "oauth-authorizations/get-grant"

Get a single grant.

  GET /applications/grants/{grant_id}

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant
*/
func GetGrant(ctx context.Context, req *GetGrantReq, opt ...requests.Option) (*GetGrantResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetGrantReq)
	}
	resp := &GetGrantResponse{}

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
GetGrant performs requests for "oauth-authorizations/get-grant"

Get a single grant.

  GET /applications/grants/{grant_id}

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetGrant(ctx context.Context, req *GetGrantReq, opt ...requests.Option) (*GetGrantResponse, error) {
	return GetGrant(ctx, req, append(c, opt...)...)
}

/*
GetGrantReq is request data for Client.GetGrant

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetGrantReq struct {
	_url string

	// grant_id parameter
	GrantId int64
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetGrantReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": strPtr("application/json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/applications/grants/%v", r.GrantId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetGrantReq) Rel(link string, resp *GetGrantResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetGrantResponse is a response for GetGrant

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant
*/
type GetGrantResponse struct {
	httpResponse *http.Response
	Data         components.ApplicationGrant
}

// HTTPResponse returns the *http.Response
func (r *GetGrantResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetGrantResponse) ReadResponse(resp *http.Response) error {
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
GetOrCreateAuthorizationForApp performs requests for "oauth-authorizations/get-or-create-authorization-for-app"

Get-or-create an authorization for a specific app.

  PUT /authorizations/clients/{client_id}

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app
*/
func GetOrCreateAuthorizationForApp(ctx context.Context, req *GetOrCreateAuthorizationForAppReq, opt ...requests.Option) (*GetOrCreateAuthorizationForAppResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetOrCreateAuthorizationForAppReq)
	}
	resp := &GetOrCreateAuthorizationForAppResponse{}

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
GetOrCreateAuthorizationForApp performs requests for "oauth-authorizations/get-or-create-authorization-for-app"

Get-or-create an authorization for a specific app.

  PUT /authorizations/clients/{client_id}

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetOrCreateAuthorizationForApp(ctx context.Context, req *GetOrCreateAuthorizationForAppReq, opt ...requests.Option) (*GetOrCreateAuthorizationForAppResponse, error) {
	return GetOrCreateAuthorizationForApp(ctx, req, append(c, opt...)...)
}

/*
GetOrCreateAuthorizationForAppReq is request data for Client.GetOrCreateAuthorizationForApp

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetOrCreateAuthorizationForAppReq struct {
	_url        string
	ClientId    string
	RequestBody GetOrCreateAuthorizationForAppReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetOrCreateAuthorizationForAppReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       strPtr("application/json"),
			"content-type": strPtr("application/json"),
		},
		Method:  "PUT",
		Options: opt,
		URLPath: fmt.Sprintf("/authorizations/clients/%v", r.ClientId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetOrCreateAuthorizationForAppReq) Rel(link string, resp *GetOrCreateAuthorizationForAppResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetOrCreateAuthorizationForAppReqBody is a request body for oauth-authorizations/get-or-create-authorization-for-app

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app
*/
type GetOrCreateAuthorizationForAppReqBody struct {

	// The OAuth app client secret for which to create the token.
	ClientSecret *string `json:"client_secret"`

	// A unique string to distinguish an authorization from others created for the same client ID and user.
	Fingerprint *string `json:"fingerprint,omitempty"`

	// A note to remind you what the OAuth token is for.
	Note *string `json:"note,omitempty"`

	// A URL to remind you what app the OAuth token is for.
	NoteUrl *string `json:"note_url,omitempty"`

	// A list of scopes that this authorization is in.
	Scopes []string `json:"scopes,omitempty"`
}

/*
GetOrCreateAuthorizationForAppResponse is a response for GetOrCreateAuthorizationForApp

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app
*/
type GetOrCreateAuthorizationForAppResponse struct {
	httpResponse *http.Response
	Data         components.Authorization
}

// HTTPResponse returns the *http.Response
func (r *GetOrCreateAuthorizationForAppResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetOrCreateAuthorizationForAppResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{200, 201, 304})
	if err != nil {
		return err
	}
	if intInSlice(resp.StatusCode, []int{200, 201}) {
		err = unmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
GetOrCreateAuthorizationForAppAndFingerprint performs requests for "oauth-authorizations/get-or-create-authorization-for-app-and-fingerprint"

Get-or-create an authorization for a specific app and fingerprint.

  PUT /authorizations/clients/{client_id}/{fingerprint}

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint
*/
func GetOrCreateAuthorizationForAppAndFingerprint(ctx context.Context, req *GetOrCreateAuthorizationForAppAndFingerprintReq, opt ...requests.Option) (*GetOrCreateAuthorizationForAppAndFingerprintResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetOrCreateAuthorizationForAppAndFingerprintReq)
	}
	resp := &GetOrCreateAuthorizationForAppAndFingerprintResponse{}

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
GetOrCreateAuthorizationForAppAndFingerprint performs requests for "oauth-authorizations/get-or-create-authorization-for-app-and-fingerprint"

Get-or-create an authorization for a specific app and fingerprint.

  PUT /authorizations/clients/{client_id}/{fingerprint}

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetOrCreateAuthorizationForAppAndFingerprint(ctx context.Context, req *GetOrCreateAuthorizationForAppAndFingerprintReq, opt ...requests.Option) (*GetOrCreateAuthorizationForAppAndFingerprintResponse, error) {
	return GetOrCreateAuthorizationForAppAndFingerprint(ctx, req, append(c, opt...)...)
}

/*
GetOrCreateAuthorizationForAppAndFingerprintReq is request data for Client.GetOrCreateAuthorizationForAppAndFingerprint

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetOrCreateAuthorizationForAppAndFingerprintReq struct {
	_url     string
	ClientId string

	// fingerprint parameter
	Fingerprint string
	RequestBody GetOrCreateAuthorizationForAppAndFingerprintReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetOrCreateAuthorizationForAppAndFingerprintReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       strPtr("application/json"),
			"content-type": strPtr("application/json"),
		},
		Method:  "PUT",
		Options: opt,
		URLPath: fmt.Sprintf("/authorizations/clients/%v/%v", r.ClientId, r.Fingerprint),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetOrCreateAuthorizationForAppAndFingerprintReq) Rel(link string, resp *GetOrCreateAuthorizationForAppAndFingerprintResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetOrCreateAuthorizationForAppAndFingerprintReqBody is a request body for oauth-authorizations/get-or-create-authorization-for-app-and-fingerprint

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint
*/
type GetOrCreateAuthorizationForAppAndFingerprintReqBody struct {

	// The OAuth app client secret for which to create the token.
	ClientSecret *string `json:"client_secret"`

	// A note to remind you what the OAuth token is for.
	Note *string `json:"note,omitempty"`

	// A URL to remind you what app the OAuth token is for.
	NoteUrl *string `json:"note_url,omitempty"`

	// A list of scopes that this authorization is in.
	Scopes []string `json:"scopes,omitempty"`
}

/*
GetOrCreateAuthorizationForAppAndFingerprintResponse is a response for GetOrCreateAuthorizationForAppAndFingerprint

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint
*/
type GetOrCreateAuthorizationForAppAndFingerprintResponse struct {
	httpResponse *http.Response
	Data         components.Authorization
}

// HTTPResponse returns the *http.Response
func (r *GetOrCreateAuthorizationForAppAndFingerprintResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetOrCreateAuthorizationForAppAndFingerprintResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{200, 201})
	if err != nil {
		return err
	}
	if intInSlice(resp.StatusCode, []int{200, 201}) {
		err = unmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
ListAuthorizations performs requests for "oauth-authorizations/list-authorizations"

List your authorizations.

  GET /authorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations
*/
func ListAuthorizations(ctx context.Context, req *ListAuthorizationsReq, opt ...requests.Option) (*ListAuthorizationsResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(ListAuthorizationsReq)
	}
	resp := &ListAuthorizationsResponse{}

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
ListAuthorizations performs requests for "oauth-authorizations/list-authorizations"

List your authorizations.

  GET /authorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) ListAuthorizations(ctx context.Context, req *ListAuthorizationsReq, opt ...requests.Option) (*ListAuthorizationsResponse, error) {
	return ListAuthorizations(ctx, req, append(c, opt...)...)
}

/*
ListAuthorizationsReq is request data for Client.ListAuthorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type ListAuthorizationsReq struct {
	_url string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *ListAuthorizationsReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	query := url.Values{}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}

	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": strPtr("application/json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/authorizations"),
		URLQuery:    query,
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *ListAuthorizationsReq) Rel(link string, resp *ListAuthorizationsResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
ListAuthorizationsResponse is a response for ListAuthorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations
*/
type ListAuthorizationsResponse struct {
	httpResponse *http.Response
	Data         []components.Authorization
}

// HTTPResponse returns the *http.Response
func (r *ListAuthorizationsResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *ListAuthorizationsResponse) ReadResponse(resp *http.Response) error {
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
ListGrants performs requests for "oauth-authorizations/list-grants"

List your grants.

  GET /applications/grants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants
*/
func ListGrants(ctx context.Context, req *ListGrantsReq, opt ...requests.Option) (*ListGrantsResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(ListGrantsReq)
	}
	resp := &ListGrantsResponse{}

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
ListGrants performs requests for "oauth-authorizations/list-grants"

List your grants.

  GET /applications/grants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) ListGrants(ctx context.Context, req *ListGrantsReq, opt ...requests.Option) (*ListGrantsResponse, error) {
	return ListGrants(ctx, req, append(c, opt...)...)
}

/*
ListGrantsReq is request data for Client.ListGrants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type ListGrantsReq struct {
	_url string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *ListGrantsReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	query := url.Values{}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}

	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": strPtr("application/json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/applications/grants"),
		URLQuery:    query,
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *ListGrantsReq) Rel(link string, resp *ListGrantsResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
ListGrantsResponse is a response for ListGrants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants
*/
type ListGrantsResponse struct {
	httpResponse *http.Response
	Data         []components.ApplicationGrant
}

// HTTPResponse returns the *http.Response
func (r *ListGrantsResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *ListGrantsResponse) ReadResponse(resp *http.Response) error {
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
UpdateAuthorization performs requests for "oauth-authorizations/update-authorization"

Update an existing authorization.

  PATCH /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization
*/
func UpdateAuthorization(ctx context.Context, req *UpdateAuthorizationReq, opt ...requests.Option) (*UpdateAuthorizationResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(UpdateAuthorizationReq)
	}
	resp := &UpdateAuthorizationResponse{}

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
UpdateAuthorization performs requests for "oauth-authorizations/update-authorization"

Update an existing authorization.

  PATCH /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) UpdateAuthorization(ctx context.Context, req *UpdateAuthorizationReq, opt ...requests.Option) (*UpdateAuthorizationResponse, error) {
	return UpdateAuthorization(ctx, req, append(c, opt...)...)
}

/*
UpdateAuthorizationReq is request data for Client.UpdateAuthorization

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type UpdateAuthorizationReq struct {
	_url string

	// authorization_id parameter
	AuthorizationId int64
	RequestBody     UpdateAuthorizationReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *UpdateAuthorizationReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       strPtr("application/json"),
			"content-type": strPtr("application/json"),
		},
		Method:  "PATCH",
		Options: opt,
		URLPath: fmt.Sprintf("/authorizations/%v", r.AuthorizationId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *UpdateAuthorizationReq) Rel(link string, resp *UpdateAuthorizationResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
UpdateAuthorizationReqBody is a request body for oauth-authorizations/update-authorization

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization
*/
type UpdateAuthorizationReqBody struct {

	// A list of scopes to add to this authorization.
	AddScopes []string `json:"add_scopes,omitempty"`

	// A unique string to distinguish an authorization from others created for the same client ID and user.
	Fingerprint *string `json:"fingerprint,omitempty"`

	// A note to remind you what the OAuth token is for.
	Note *string `json:"note,omitempty"`

	// A URL to remind you what app the OAuth token is for.
	NoteUrl *string `json:"note_url,omitempty"`

	// A list of scopes to remove from this authorization.
	RemoveScopes []string `json:"remove_scopes,omitempty"`

	// A list of scopes that this authorization is in.
	Scopes []string `json:"scopes,omitempty"`
}

/*
UpdateAuthorizationResponse is a response for UpdateAuthorization

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization
*/
type UpdateAuthorizationResponse struct {
	httpResponse *http.Response
	Data         components.Authorization
}

// HTTPResponse returns the *http.Response
func (r *UpdateAuthorizationResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *UpdateAuthorizationResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := responseErrorCheck(resp, []int{200})
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
