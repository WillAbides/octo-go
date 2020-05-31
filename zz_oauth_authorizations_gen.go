// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	"net/http"
	"net/url"
	"strconv"
)

/*
OauthAuthorizationsCreateAuthorization performs requests for "oauth-authorizations/create-authorization"

Create a new authorization.

  POST /authorizations

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
*/
func (c *Client) OauthAuthorizationsCreateAuthorization(ctx context.Context, req *OauthAuthorizationsCreateAuthorizationReq, opt ...RequestOption) (*OauthAuthorizationsCreateAuthorizationResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsCreateAuthorizationResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(OauthAuthorizationsCreateAuthorizationResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsCreateAuthorizationReq is request data for Client.OauthAuthorizationsCreateAuthorization

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
*/
type OauthAuthorizationsCreateAuthorizationReq struct {
	pgURL       string
	RequestBody OauthAuthorizationsCreateAuthorizationReqBody
}

func (r *OauthAuthorizationsCreateAuthorizationReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsCreateAuthorizationReq) urlPath() string {
	return fmt.Sprintf("/authorizations")
}

func (r *OauthAuthorizationsCreateAuthorizationReq) method() string {
	return "POST"
}

func (r *OauthAuthorizationsCreateAuthorizationReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *OauthAuthorizationsCreateAuthorizationReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsCreateAuthorizationReq) body() interface{} {
	return r.RequestBody
}

func (r *OauthAuthorizationsCreateAuthorizationReq) dataStatuses() []int {
	return []int{201}
}

func (r *OauthAuthorizationsCreateAuthorizationReq) validStatuses() []int {
	return []int{201}
}

func (r *OauthAuthorizationsCreateAuthorizationReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsCreateAuthorizationReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsCreateAuthorizationReq) Rel(link RelName, resp *OauthAuthorizationsCreateAuthorizationResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsCreateAuthorizationReqBody is a request body for oauth-authorizations/create-authorization

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
*/
type OauthAuthorizationsCreateAuthorizationReqBody struct {

	// The 20 character OAuth app client key for which to create the token.
	ClientId *string `json:"client_id,omitempty"`

	// The 40 character OAuth app client secret for which to create the token.
	ClientSecret *string `json:"client_secret,omitempty"`

	/*
	   A unique string to distinguish an authorization from others created for the same
	   client ID and user.
	*/
	Fingerprint *string `json:"fingerprint,omitempty"`

	/*
	   A note to remind you what the OAuth token is for. Tokens not associated with a
	   specific OAuth application (i.e. personal access tokens) must have a unique
	   note.
	*/
	Note *string `json:"note"`

	// A URL to remind you what app the OAuth token is for.
	NoteUrl *string `json:"note_url,omitempty"`

	// A list of scopes that this authorization is in.
	Scopes []string `json:"scopes,omitempty"`
}

/*
OauthAuthorizationsCreateAuthorizationResponseBody is a response body for OauthAuthorizationsCreateAuthorization

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
*/
type OauthAuthorizationsCreateAuthorizationResponseBody struct {
	components.Authorization
}

/*
OauthAuthorizationsCreateAuthorizationResponse is a response for OauthAuthorizationsCreateAuthorization

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
*/
type OauthAuthorizationsCreateAuthorizationResponse struct {
	response
	request *OauthAuthorizationsCreateAuthorizationReq
	Data    *OauthAuthorizationsCreateAuthorizationResponseBody
}

/*
OauthAuthorizationsDeleteAuthorization performs requests for "oauth-authorizations/delete-authorization"

Delete an authorization.

  DELETE /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization
*/
func (c *Client) OauthAuthorizationsDeleteAuthorization(ctx context.Context, req *OauthAuthorizationsDeleteAuthorizationReq, opt ...RequestOption) (*OauthAuthorizationsDeleteAuthorizationResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsDeleteAuthorizationResponse{
		request:  req,
		response: *r,
	}
	err = r.decodeBody(nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsDeleteAuthorizationReq is request data for Client.OauthAuthorizationsDeleteAuthorization

https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization
*/
type OauthAuthorizationsDeleteAuthorizationReq struct {
	pgURL           string
	AuthorizationId int64
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) urlPath() string {
	return fmt.Sprintf("/authorizations/%v", r.AuthorizationId)
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) method() string {
	return "DELETE"
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) body() interface{} {
	return nil
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) dataStatuses() []int {
	return []int{}
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) validStatuses() []int {
	return []int{204}
}

func (r *OauthAuthorizationsDeleteAuthorizationReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsDeleteAuthorizationReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsDeleteAuthorizationReq) Rel(link RelName, resp *OauthAuthorizationsDeleteAuthorizationResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsDeleteAuthorizationResponse is a response for OauthAuthorizationsDeleteAuthorization

https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization
*/
type OauthAuthorizationsDeleteAuthorizationResponse struct {
	response
	request *OauthAuthorizationsDeleteAuthorizationReq
}

/*
OauthAuthorizationsDeleteGrant performs requests for "oauth-authorizations/delete-grant"

Delete a grant.

  DELETE /applications/grants/{grant_id}

https://developer.github.com/v3/oauth_authorizations/#delete-a-grant
*/
func (c *Client) OauthAuthorizationsDeleteGrant(ctx context.Context, req *OauthAuthorizationsDeleteGrantReq, opt ...RequestOption) (*OauthAuthorizationsDeleteGrantResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsDeleteGrantResponse{
		request:  req,
		response: *r,
	}
	err = r.decodeBody(nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsDeleteGrantReq is request data for Client.OauthAuthorizationsDeleteGrant

https://developer.github.com/v3/oauth_authorizations/#delete-a-grant
*/
type OauthAuthorizationsDeleteGrantReq struct {
	pgURL   string
	GrantId int64
}

func (r *OauthAuthorizationsDeleteGrantReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsDeleteGrantReq) urlPath() string {
	return fmt.Sprintf("/applications/grants/%v", r.GrantId)
}

func (r *OauthAuthorizationsDeleteGrantReq) method() string {
	return "DELETE"
}

func (r *OauthAuthorizationsDeleteGrantReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *OauthAuthorizationsDeleteGrantReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsDeleteGrantReq) body() interface{} {
	return nil
}

func (r *OauthAuthorizationsDeleteGrantReq) dataStatuses() []int {
	return []int{}
}

func (r *OauthAuthorizationsDeleteGrantReq) validStatuses() []int {
	return []int{204}
}

func (r *OauthAuthorizationsDeleteGrantReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsDeleteGrantReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsDeleteGrantReq) Rel(link RelName, resp *OauthAuthorizationsDeleteGrantResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsDeleteGrantResponse is a response for OauthAuthorizationsDeleteGrant

https://developer.github.com/v3/oauth_authorizations/#delete-a-grant
*/
type OauthAuthorizationsDeleteGrantResponse struct {
	response
	request *OauthAuthorizationsDeleteGrantReq
}

/*
OauthAuthorizationsGetAuthorization performs requests for "oauth-authorizations/get-authorization"

Get a single authorization.

  GET /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization
*/
func (c *Client) OauthAuthorizationsGetAuthorization(ctx context.Context, req *OauthAuthorizationsGetAuthorizationReq, opt ...RequestOption) (*OauthAuthorizationsGetAuthorizationResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsGetAuthorizationResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(OauthAuthorizationsGetAuthorizationResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsGetAuthorizationReq is request data for Client.OauthAuthorizationsGetAuthorization

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization
*/
type OauthAuthorizationsGetAuthorizationReq struct {
	pgURL           string
	AuthorizationId int64
}

func (r *OauthAuthorizationsGetAuthorizationReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsGetAuthorizationReq) urlPath() string {
	return fmt.Sprintf("/authorizations/%v", r.AuthorizationId)
}

func (r *OauthAuthorizationsGetAuthorizationReq) method() string {
	return "GET"
}

func (r *OauthAuthorizationsGetAuthorizationReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *OauthAuthorizationsGetAuthorizationReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsGetAuthorizationReq) body() interface{} {
	return nil
}

func (r *OauthAuthorizationsGetAuthorizationReq) dataStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsGetAuthorizationReq) validStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsGetAuthorizationReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsGetAuthorizationReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsGetAuthorizationReq) Rel(link RelName, resp *OauthAuthorizationsGetAuthorizationResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsGetAuthorizationResponseBody is a response body for OauthAuthorizationsGetAuthorization

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization
*/
type OauthAuthorizationsGetAuthorizationResponseBody struct {
	components.Authorization
}

/*
OauthAuthorizationsGetAuthorizationResponse is a response for OauthAuthorizationsGetAuthorization

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization
*/
type OauthAuthorizationsGetAuthorizationResponse struct {
	response
	request *OauthAuthorizationsGetAuthorizationReq
	Data    *OauthAuthorizationsGetAuthorizationResponseBody
}

/*
OauthAuthorizationsGetGrant performs requests for "oauth-authorizations/get-grant"

Get a single grant.

  GET /applications/grants/{grant_id}

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant
*/
func (c *Client) OauthAuthorizationsGetGrant(ctx context.Context, req *OauthAuthorizationsGetGrantReq, opt ...RequestOption) (*OauthAuthorizationsGetGrantResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsGetGrantResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(OauthAuthorizationsGetGrantResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsGetGrantReq is request data for Client.OauthAuthorizationsGetGrant

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant
*/
type OauthAuthorizationsGetGrantReq struct {
	pgURL   string
	GrantId int64
}

func (r *OauthAuthorizationsGetGrantReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsGetGrantReq) urlPath() string {
	return fmt.Sprintf("/applications/grants/%v", r.GrantId)
}

func (r *OauthAuthorizationsGetGrantReq) method() string {
	return "GET"
}

func (r *OauthAuthorizationsGetGrantReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *OauthAuthorizationsGetGrantReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsGetGrantReq) body() interface{} {
	return nil
}

func (r *OauthAuthorizationsGetGrantReq) dataStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsGetGrantReq) validStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsGetGrantReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsGetGrantReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsGetGrantReq) Rel(link RelName, resp *OauthAuthorizationsGetGrantResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsGetGrantResponseBody is a response body for OauthAuthorizationsGetGrant

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant
*/
type OauthAuthorizationsGetGrantResponseBody struct {
	components.ApplicationGrant
}

/*
OauthAuthorizationsGetGrantResponse is a response for OauthAuthorizationsGetGrant

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant
*/
type OauthAuthorizationsGetGrantResponse struct {
	response
	request *OauthAuthorizationsGetGrantReq
	Data    *OauthAuthorizationsGetGrantResponseBody
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForApp performs requests for "oauth-authorizations/get-or-create-authorization-for-app"

Get-or-create an authorization for a specific app.

  PUT /authorizations/clients/{client_id}

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app
*/
func (c *Client) OauthAuthorizationsGetOrCreateAuthorizationForApp(ctx context.Context, req *OauthAuthorizationsGetOrCreateAuthorizationForAppReq, opt ...RequestOption) (*OauthAuthorizationsGetOrCreateAuthorizationForAppResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsGetOrCreateAuthorizationForAppResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(OauthAuthorizationsGetOrCreateAuthorizationForAppResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppReq is request data for Client.OauthAuthorizationsGetOrCreateAuthorizationForApp

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app
*/
type OauthAuthorizationsGetOrCreateAuthorizationForAppReq struct {
	pgURL       string
	ClientId    string
	RequestBody OauthAuthorizationsGetOrCreateAuthorizationForAppReqBody
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) urlPath() string {
	return fmt.Sprintf("/authorizations/clients/%v", r.ClientId)
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) method() string {
	return "PUT"
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) body() interface{} {
	return r.RequestBody
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) dataStatuses() []int {
	return []int{200, 201}
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) validStatuses() []int {
	return []int{200, 201}
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppReq) Rel(link RelName, resp *OauthAuthorizationsGetOrCreateAuthorizationForAppResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppReqBody is a request body for oauth-authorizations/get-or-create-authorization-for-app

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app
*/
type OauthAuthorizationsGetOrCreateAuthorizationForAppReqBody struct {

	/*
	   The 40 character OAuth app client secret associated with the client ID specified
	   in the URL.
	*/
	ClientSecret *string `json:"client_secret"`

	/*
	   A unique string to distinguish an authorization from others created for the same
	   client and user. If provided, this API is functionally equivalent to
	   [Get-or-create an authorization for a specific app and
	   fingerprint](https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint).
	*/
	Fingerprint *string `json:"fingerprint,omitempty"`

	// A note to remind you what the OAuth token is for.
	Note *string `json:"note,omitempty"`

	// A URL to remind you what app the OAuth token is for.
	NoteUrl *string `json:"note_url,omitempty"`

	// A list of scopes that this authorization is in.
	Scopes []string `json:"scopes,omitempty"`
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppResponseBody is a response body for OauthAuthorizationsGetOrCreateAuthorizationForApp

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app
*/
type OauthAuthorizationsGetOrCreateAuthorizationForAppResponseBody struct {
	components.Authorization
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppResponse is a response for OauthAuthorizationsGetOrCreateAuthorizationForApp

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app
*/
type OauthAuthorizationsGetOrCreateAuthorizationForAppResponse struct {
	response
	request *OauthAuthorizationsGetOrCreateAuthorizationForAppReq
	Data    *OauthAuthorizationsGetOrCreateAuthorizationForAppResponseBody
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprint performs requests for "oauth-authorizations/get-or-create-authorization-for-app-and-fingerprint"

Get-or-create an authorization for a specific app and fingerprint.

  PUT /authorizations/clients/{client_id}/{fingerprint}

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint
*/
func (c *Client) OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprint(ctx context.Context, req *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq, opt ...RequestOption) (*OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq is request data for Client.OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprint

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint
*/
type OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq struct {
	pgURL       string
	ClientId    string
	Fingerprint string
	RequestBody OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReqBody
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) urlPath() string {
	return fmt.Sprintf("/authorizations/clients/%v/%v", r.ClientId, r.Fingerprint)
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) method() string {
	return "PUT"
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) body() interface{} {
	return r.RequestBody
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) dataStatuses() []int {
	return []int{200, 201}
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) validStatuses() []int {
	return []int{200, 201}
}

func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq) Rel(link RelName, resp *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReqBody is a request body for oauth-authorizations/get-or-create-authorization-for-app-and-fingerprint

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint
*/
type OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReqBody struct {

	/*
	   The 40 character OAuth app client secret associated with the client ID specified
	   in the URL.
	*/
	ClientSecret *string `json:"client_secret"`

	// A note to remind you what the OAuth token is for.
	Note *string `json:"note,omitempty"`

	// A URL to remind you what app the OAuth token is for.
	NoteUrl *string `json:"note_url,omitempty"`

	// A list of scopes that this authorization is in.
	Scopes []string `json:"scopes,omitempty"`
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponseBody is a response body for OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprint

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint
*/
type OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponseBody struct {
	components.Authorization
}

/*
OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponse is a response for OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprint

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint
*/
type OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponse struct {
	response
	request *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq
	Data    *OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintResponseBody
}

/*
OauthAuthorizationsListAuthorizations performs requests for "oauth-authorizations/list-authorizations"

List your authorizations.

  GET /authorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations
*/
func (c *Client) OauthAuthorizationsListAuthorizations(ctx context.Context, req *OauthAuthorizationsListAuthorizationsReq, opt ...RequestOption) (*OauthAuthorizationsListAuthorizationsResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsListAuthorizationsResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(OauthAuthorizationsListAuthorizationsResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsListAuthorizationsReq is request data for Client.OauthAuthorizationsListAuthorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations
*/
type OauthAuthorizationsListAuthorizationsReq struct {
	pgURL string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64
}

func (r *OauthAuthorizationsListAuthorizationsReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsListAuthorizationsReq) urlPath() string {
	return fmt.Sprintf("/authorizations")
}

func (r *OauthAuthorizationsListAuthorizationsReq) method() string {
	return "GET"
}

func (r *OauthAuthorizationsListAuthorizationsReq) urlQuery() url.Values {
	query := url.Values{}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}
	return query
}

func (r *OauthAuthorizationsListAuthorizationsReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsListAuthorizationsReq) body() interface{} {
	return nil
}

func (r *OauthAuthorizationsListAuthorizationsReq) dataStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsListAuthorizationsReq) validStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsListAuthorizationsReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsListAuthorizationsReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsListAuthorizationsReq) Rel(link RelName, resp *OauthAuthorizationsListAuthorizationsResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsListAuthorizationsResponseBody is a response body for OauthAuthorizationsListAuthorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations
*/
type OauthAuthorizationsListAuthorizationsResponseBody []struct {
	components.Authorization
}

/*
OauthAuthorizationsListAuthorizationsResponse is a response for OauthAuthorizationsListAuthorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations
*/
type OauthAuthorizationsListAuthorizationsResponse struct {
	response
	request *OauthAuthorizationsListAuthorizationsReq
	Data    *OauthAuthorizationsListAuthorizationsResponseBody
}

/*
OauthAuthorizationsListGrants performs requests for "oauth-authorizations/list-grants"

List your grants.

  GET /applications/grants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants
*/
func (c *Client) OauthAuthorizationsListGrants(ctx context.Context, req *OauthAuthorizationsListGrantsReq, opt ...RequestOption) (*OauthAuthorizationsListGrantsResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsListGrantsResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(OauthAuthorizationsListGrantsResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsListGrantsReq is request data for Client.OauthAuthorizationsListGrants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants
*/
type OauthAuthorizationsListGrantsReq struct {
	pgURL string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64
}

func (r *OauthAuthorizationsListGrantsReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsListGrantsReq) urlPath() string {
	return fmt.Sprintf("/applications/grants")
}

func (r *OauthAuthorizationsListGrantsReq) method() string {
	return "GET"
}

func (r *OauthAuthorizationsListGrantsReq) urlQuery() url.Values {
	query := url.Values{}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}
	return query
}

func (r *OauthAuthorizationsListGrantsReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsListGrantsReq) body() interface{} {
	return nil
}

func (r *OauthAuthorizationsListGrantsReq) dataStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsListGrantsReq) validStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsListGrantsReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsListGrantsReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsListGrantsReq) Rel(link RelName, resp *OauthAuthorizationsListGrantsResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsListGrantsResponseBody is a response body for OauthAuthorizationsListGrants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants
*/
type OauthAuthorizationsListGrantsResponseBody []struct {
	components.ApplicationGrant
}

/*
OauthAuthorizationsListGrantsResponse is a response for OauthAuthorizationsListGrants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants
*/
type OauthAuthorizationsListGrantsResponse struct {
	response
	request *OauthAuthorizationsListGrantsReq
	Data    *OauthAuthorizationsListGrantsResponseBody
}

/*
OauthAuthorizationsUpdateAuthorization performs requests for "oauth-authorizations/update-authorization"

Update an existing authorization.

  PATCH /authorizations/{authorization_id}

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization
*/
func (c *Client) OauthAuthorizationsUpdateAuthorization(ctx context.Context, req *OauthAuthorizationsUpdateAuthorizationReq, opt ...RequestOption) (*OauthAuthorizationsUpdateAuthorizationResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &OauthAuthorizationsUpdateAuthorizationResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(OauthAuthorizationsUpdateAuthorizationResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
OauthAuthorizationsUpdateAuthorizationReq is request data for Client.OauthAuthorizationsUpdateAuthorization

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization
*/
type OauthAuthorizationsUpdateAuthorizationReq struct {
	pgURL           string
	AuthorizationId int64
	RequestBody     OauthAuthorizationsUpdateAuthorizationReqBody
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) pagingURL() string {
	return r.pgURL
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) urlPath() string {
	return fmt.Sprintf("/authorizations/%v", r.AuthorizationId)
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) method() string {
	return "PATCH"
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) body() interface{} {
	return r.RequestBody
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) dataStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) validStatuses() []int {
	return []int{200}
}

func (r *OauthAuthorizationsUpdateAuthorizationReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *OauthAuthorizationsUpdateAuthorizationReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *OauthAuthorizationsUpdateAuthorizationReq) Rel(link RelName, resp *OauthAuthorizationsUpdateAuthorizationResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r.pgURL = u
	return true
}

/*
OauthAuthorizationsUpdateAuthorizationReqBody is a request body for oauth-authorizations/update-authorization

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization
*/
type OauthAuthorizationsUpdateAuthorizationReqBody struct {

	// A list of scopes to add to this authorization.
	AddScopes []string `json:"add_scopes,omitempty"`

	/*
	   A unique string to distinguish an authorization from others created for the same
	   client ID and user.
	*/
	Fingerprint *string `json:"fingerprint,omitempty"`

	/*
	   A note to remind you what the OAuth token is for. Tokens not associated with a
	   specific OAuth application (i.e. personal access tokens) must have a unique
	   note.
	*/
	Note *string `json:"note,omitempty"`

	// A URL to remind you what app the OAuth token is for.
	NoteUrl *string `json:"note_url,omitempty"`

	// A list of scopes to remove from this authorization.
	RemoveScopes []string `json:"remove_scopes,omitempty"`

	// Replaces the authorization scopes with these.
	Scopes []string `json:"scopes,omitempty"`
}

/*
OauthAuthorizationsUpdateAuthorizationResponseBody is a response body for OauthAuthorizationsUpdateAuthorization

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization
*/
type OauthAuthorizationsUpdateAuthorizationResponseBody struct {
	components.Authorization
}

/*
OauthAuthorizationsUpdateAuthorizationResponse is a response for OauthAuthorizationsUpdateAuthorization

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization
*/
type OauthAuthorizationsUpdateAuthorizationResponse struct {
	response
	request *OauthAuthorizationsUpdateAuthorizationReq
	Data    *OauthAuthorizationsUpdateAuthorizationResponseBody
}
