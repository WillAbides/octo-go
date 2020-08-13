// Code generated by octo-go; DO NOT EDIT.

package scim

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	internal "github.com/willabides/octo-go/internal"
	requests "github.com/willabides/octo-go/requests"
	"net/http"
	"net/url"
	"strconv"
)

// Client is a set of options to apply to requests
type Client []requests.Option

// NewClient returns a new Client
func NewClient(opt ...requests.Option) Client {
	return opt
}

/*
DeleteUserFromOrg performs requests for "scim/delete-user-from-org"

Delete a SCIM user from an organization.

  DELETE /scim/v2/organizations/{org}/Users/{scim_user_id}

https://developer.github.com/v3/scim/#delete-a-scim-user-from-an-organization
*/
func DeleteUserFromOrg(ctx context.Context, req *DeleteUserFromOrgReq, opt ...requests.Option) (*DeleteUserFromOrgResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(DeleteUserFromOrgReq)
	}
	resp := &DeleteUserFromOrgResponse{}

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
DeleteUserFromOrg performs requests for "scim/delete-user-from-org"

Delete a SCIM user from an organization.

  DELETE /scim/v2/organizations/{org}/Users/{scim_user_id}

https://developer.github.com/v3/scim/#delete-a-scim-user-from-an-organization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) DeleteUserFromOrg(ctx context.Context, req *DeleteUserFromOrgReq, opt ...requests.Option) (*DeleteUserFromOrgResponse, error) {
	return DeleteUserFromOrg(ctx, req, append(c, opt...)...)
}

/*
DeleteUserFromOrgReq is request data for Client.DeleteUserFromOrg

https://developer.github.com/v3/scim/#delete-a-scim-user-from-an-organization

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type DeleteUserFromOrgReq struct {
	_url string
	Org  string

	// scim_user_id parameter
	ScimUserId string
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *DeleteUserFromOrgReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		Body:        nil,
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{},
		Method:      "DELETE",
		Options:     opt,
		URLPath:     fmt.Sprintf("/scim/v2/organizations/%v/Users/%v", r.Org, r.ScimUserId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *DeleteUserFromOrgReq) Rel(link string, resp *DeleteUserFromOrgResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
DeleteUserFromOrgResponse is a response for DeleteUserFromOrg

https://developer.github.com/v3/scim/#delete-a-scim-user-from-an-organization
*/
type DeleteUserFromOrgResponse struct {
	httpResponse *http.Response
}

// HTTPResponse returns the *http.Response
func (r *DeleteUserFromOrgResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *DeleteUserFromOrgResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{204, 304})
	if err != nil {
		return err
	}
	return nil
}

/*
GetProvisioningInformationForUser performs requests for "scim/get-provisioning-information-for-user"

Get SCIM provisioning information for a user.

  GET /scim/v2/organizations/{org}/Users/{scim_user_id}

https://developer.github.com/v3/scim/#get-scim-provisioning-information-for-a-user
*/
func GetProvisioningInformationForUser(ctx context.Context, req *GetProvisioningInformationForUserReq, opt ...requests.Option) (*GetProvisioningInformationForUserResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetProvisioningInformationForUserReq)
	}
	resp := &GetProvisioningInformationForUserResponse{}

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
GetProvisioningInformationForUser performs requests for "scim/get-provisioning-information-for-user"

Get SCIM provisioning information for a user.

  GET /scim/v2/organizations/{org}/Users/{scim_user_id}

https://developer.github.com/v3/scim/#get-scim-provisioning-information-for-a-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) GetProvisioningInformationForUser(ctx context.Context, req *GetProvisioningInformationForUserReq, opt ...requests.Option) (*GetProvisioningInformationForUserResponse, error) {
	return GetProvisioningInformationForUser(ctx, req, append(c, opt...)...)
}

/*
GetProvisioningInformationForUserReq is request data for Client.GetProvisioningInformationForUser

https://developer.github.com/v3/scim/#get-scim-provisioning-information-for-a-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetProvisioningInformationForUserReq struct {
	_url string
	Org  string

	// scim_user_id parameter
	ScimUserId string
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetProvisioningInformationForUserReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		Body:        nil,
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": internal.String("application/scim+json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/scim/v2/organizations/%v/Users/%v", r.Org, r.ScimUserId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetProvisioningInformationForUserReq) Rel(link string, resp *GetProvisioningInformationForUserResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetProvisioningInformationForUserResponse is a response for GetProvisioningInformationForUser

https://developer.github.com/v3/scim/#get-scim-provisioning-information-for-a-user
*/
type GetProvisioningInformationForUserResponse struct {
	httpResponse *http.Response
	Data         components.ScimUser
}

// HTTPResponse returns the *http.Response
func (r *GetProvisioningInformationForUserResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetProvisioningInformationForUserResponse) ReadResponse(resp *http.Response) error {
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
ListProvisionedIdentities performs requests for "scim/list-provisioned-identities"

List SCIM provisioned identities.

  GET /scim/v2/organizations/{org}/Users

https://developer.github.com/v3/scim/#list-scim-provisioned-identities
*/
func ListProvisionedIdentities(ctx context.Context, req *ListProvisionedIdentitiesReq, opt ...requests.Option) (*ListProvisionedIdentitiesResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(ListProvisionedIdentitiesReq)
	}
	resp := &ListProvisionedIdentitiesResponse{}

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
ListProvisionedIdentities performs requests for "scim/list-provisioned-identities"

List SCIM provisioned identities.

  GET /scim/v2/organizations/{org}/Users

https://developer.github.com/v3/scim/#list-scim-provisioned-identities

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) ListProvisionedIdentities(ctx context.Context, req *ListProvisionedIdentitiesReq, opt ...requests.Option) (*ListProvisionedIdentitiesResponse, error) {
	return ListProvisionedIdentities(ctx, req, append(c, opt...)...)
}

/*
ListProvisionedIdentitiesReq is request data for Client.ListProvisionedIdentities

https://developer.github.com/v3/scim/#list-scim-provisioned-identities

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type ListProvisionedIdentitiesReq struct {
	_url string
	Org  string

	// Used for pagination: the index of the first result to return.
	StartIndex *int64

	// Used for pagination: the number of results to return.
	Count *int64

	/*
	Filters results using the equals query parameter operator (`eq`). You can filter
	results that are equal to `id`, `userName`, `emails`, and `external_id`. For
	example, to search for an identity with the `userName` Octocat, you would use
	this query:

	`?filter=userName%20eq%20\"Octocat\"`.

	To filter results for for the identity with the email `octocat@github.com`, you
	would use this query:

	`?filter=emails%20eq%20\"octocat@github.com\"`.
	*/
	Filter *string
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *ListProvisionedIdentitiesReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	query := url.Values{}
	if r.StartIndex != nil {
		query.Set("startIndex", strconv.FormatInt(*r.StartIndex, 10))
	}
	if r.Count != nil {
		query.Set("count", strconv.FormatInt(*r.Count, 10))
	}
	if r.Filter != nil {
		query.Set("filter", *r.Filter)
	}

	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		Body:        nil,
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": internal.String("application/scim+json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/scim/v2/organizations/%v/Users", r.Org),
		URLQuery:    query,
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *ListProvisionedIdentitiesReq) Rel(link string, resp *ListProvisionedIdentitiesResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
ListProvisionedIdentitiesResponse is a response for ListProvisionedIdentities

https://developer.github.com/v3/scim/#list-scim-provisioned-identities
*/
type ListProvisionedIdentitiesResponse struct {
	httpResponse *http.Response
	Data         components.ScimUserList
}

// HTTPResponse returns the *http.Response
func (r *ListProvisionedIdentitiesResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *ListProvisionedIdentitiesResponse) ReadResponse(resp *http.Response) error {
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
ProvisionAndInviteUser performs requests for "scim/provision-and-invite-user"

Provision and invite a SCIM user.

  POST /scim/v2/organizations/{org}/Users

https://developer.github.com/v3/scim/#provision-and-invite-a-scim-user
*/
func ProvisionAndInviteUser(ctx context.Context, req *ProvisionAndInviteUserReq, opt ...requests.Option) (*ProvisionAndInviteUserResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(ProvisionAndInviteUserReq)
	}
	resp := &ProvisionAndInviteUserResponse{}

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
ProvisionAndInviteUser performs requests for "scim/provision-and-invite-user"

Provision and invite a SCIM user.

  POST /scim/v2/organizations/{org}/Users

https://developer.github.com/v3/scim/#provision-and-invite-a-scim-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) ProvisionAndInviteUser(ctx context.Context, req *ProvisionAndInviteUserReq, opt ...requests.Option) (*ProvisionAndInviteUserResponse, error) {
	return ProvisionAndInviteUser(ctx, req, append(c, opt...)...)
}

/*
ProvisionAndInviteUserReq is request data for Client.ProvisionAndInviteUser

https://developer.github.com/v3/scim/#provision-and-invite-a-scim-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type ProvisionAndInviteUserReq struct {
	_url        string
	Org         string
	RequestBody ProvisionAndInviteUserReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *ProvisionAndInviteUserReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       internal.String("application/scim+json"),
			"content-type": internal.String("application/json"),
		},
		JSONRequestBody: true,
		Method:          "POST",
		Options:         opt,
		URLPath:         fmt.Sprintf("/scim/v2/organizations/%v/Users", r.Org),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *ProvisionAndInviteUserReq) Rel(link string, resp *ProvisionAndInviteUserResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

// ProvisionAndInviteUserReqBodyEmails is a value for ProvisionAndInviteUserReqBody's Emails field
type ProvisionAndInviteUserReqBodyEmails struct {
	Primary *bool   `json:"primary,omitempty"`
	Type    *string `json:"type,omitempty"`
	Value   *string `json:"value"`
}

// ProvisionAndInviteUserReqBodyName is a value for ProvisionAndInviteUserReqBody's Name field
type ProvisionAndInviteUserReqBodyName struct {
	FamilyName *string `json:"familyName"`
	GivenName  *string `json:"givenName"`
}

/*
ProvisionAndInviteUserReqBody is a request body for scim/provision-and-invite-user

https://developer.github.com/v3/scim/#provision-and-invite-a-scim-user
*/
type ProvisionAndInviteUserReqBody struct {
	Active      *bool   `json:"active,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`

	// user emails
	Emails     []ProvisionAndInviteUserReqBodyEmails `json:"emails"`
	ExternalId *string                               `json:"externalId,omitempty"`
	Groups     []string                              `json:"groups,omitempty"`
	Name       *ProvisionAndInviteUserReqBodyName    `json:"name"`
	Schemas    []string                              `json:"schemas,omitempty"`

	// Configured by the admin. Could be an email, login, or username
	UserName *string `json:"userName"`
}

/*
ProvisionAndInviteUserResponse is a response for ProvisionAndInviteUser

https://developer.github.com/v3/scim/#provision-and-invite-a-scim-user
*/
type ProvisionAndInviteUserResponse struct {
	httpResponse *http.Response
	Data         components.ScimUser
}

// HTTPResponse returns the *http.Response
func (r *ProvisionAndInviteUserResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *ProvisionAndInviteUserResponse) ReadResponse(resp *http.Response) error {
	r.httpResponse = resp
	err := internal.ResponseErrorCheck(resp, []int{201, 304})
	if err != nil {
		return err
	}
	if internal.IntInSlice(resp.StatusCode, []int{201}) {
		err = internal.UnmarshalResponseBody(resp, &r.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
SetInformationForProvisionedUser performs requests for "scim/set-information-for-provisioned-user"

Update a provisioned organization membership.

  PUT /scim/v2/organizations/{org}/Users/{scim_user_id}

https://developer.github.com/v3/scim/#set-scim-information-for-a-provisioned-user
*/
func SetInformationForProvisionedUser(ctx context.Context, req *SetInformationForProvisionedUserReq, opt ...requests.Option) (*SetInformationForProvisionedUserResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(SetInformationForProvisionedUserReq)
	}
	resp := &SetInformationForProvisionedUserResponse{}

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
SetInformationForProvisionedUser performs requests for "scim/set-information-for-provisioned-user"

Update a provisioned organization membership.

  PUT /scim/v2/organizations/{org}/Users/{scim_user_id}

https://developer.github.com/v3/scim/#set-scim-information-for-a-provisioned-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) SetInformationForProvisionedUser(ctx context.Context, req *SetInformationForProvisionedUserReq, opt ...requests.Option) (*SetInformationForProvisionedUserResponse, error) {
	return SetInformationForProvisionedUser(ctx, req, append(c, opt...)...)
}

/*
SetInformationForProvisionedUserReq is request data for Client.SetInformationForProvisionedUser

https://developer.github.com/v3/scim/#set-scim-information-for-a-provisioned-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type SetInformationForProvisionedUserReq struct {
	_url string
	Org  string

	// scim_user_id parameter
	ScimUserId  string
	RequestBody SetInformationForProvisionedUserReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *SetInformationForProvisionedUserReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       internal.String("application/scim+json"),
			"content-type": internal.String("application/json"),
		},
		JSONRequestBody: true,
		Method:          "PUT",
		Options:         opt,
		URLPath:         fmt.Sprintf("/scim/v2/organizations/%v/Users/%v", r.Org, r.ScimUserId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SetInformationForProvisionedUserReq) Rel(link string, resp *SetInformationForProvisionedUserResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

// SetInformationForProvisionedUserReqBodyEmails is a value for SetInformationForProvisionedUserReqBody's Emails field
type SetInformationForProvisionedUserReqBodyEmails struct {
	Primary *bool   `json:"primary,omitempty"`
	Type    *string `json:"type,omitempty"`
	Value   *string `json:"value"`
}

// SetInformationForProvisionedUserReqBodyName is a value for SetInformationForProvisionedUserReqBody's Name field
type SetInformationForProvisionedUserReqBodyName struct {
	FamilyName *string `json:"familyName"`
	GivenName  *string `json:"givenName"`
}

/*
SetInformationForProvisionedUserReqBody is a request body for scim/set-information-for-provisioned-user

https://developer.github.com/v3/scim/#set-scim-information-for-a-provisioned-user
*/
type SetInformationForProvisionedUserReqBody struct {
	Active      *bool   `json:"active,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`

	// user emails
	Emails     []SetInformationForProvisionedUserReqBodyEmails `json:"emails"`
	ExternalId *string                                         `json:"externalId,omitempty"`
	Groups     []string                                        `json:"groups,omitempty"`
	Name       *SetInformationForProvisionedUserReqBodyName    `json:"name"`
	Schemas    []string                                        `json:"schemas,omitempty"`

	// Configured by the admin. Could be an email, login, or username
	UserName *string `json:"userName"`
}

/*
SetInformationForProvisionedUserResponse is a response for SetInformationForProvisionedUser

https://developer.github.com/v3/scim/#set-scim-information-for-a-provisioned-user
*/
type SetInformationForProvisionedUserResponse struct {
	httpResponse *http.Response
	Data         components.ScimUser
}

// HTTPResponse returns the *http.Response
func (r *SetInformationForProvisionedUserResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *SetInformationForProvisionedUserResponse) ReadResponse(resp *http.Response) error {
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
UpdateAttributeForUser performs requests for "scim/update-attribute-for-user"

Update an attribute for a SCIM user.

  PATCH /scim/v2/organizations/{org}/Users/{scim_user_id}

https://developer.github.com/v3/scim/#update-an-attribute-for-a-scim-user
*/
func UpdateAttributeForUser(ctx context.Context, req *UpdateAttributeForUserReq, opt ...requests.Option) (*UpdateAttributeForUserResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(UpdateAttributeForUserReq)
	}
	resp := &UpdateAttributeForUserResponse{}

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
UpdateAttributeForUser performs requests for "scim/update-attribute-for-user"

Update an attribute for a SCIM user.

  PATCH /scim/v2/organizations/{org}/Users/{scim_user_id}

https://developer.github.com/v3/scim/#update-an-attribute-for-a-scim-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) UpdateAttributeForUser(ctx context.Context, req *UpdateAttributeForUserReq, opt ...requests.Option) (*UpdateAttributeForUserResponse, error) {
	return UpdateAttributeForUser(ctx, req, append(c, opt...)...)
}

/*
UpdateAttributeForUserReq is request data for Client.UpdateAttributeForUser

https://developer.github.com/v3/scim/#update-an-attribute-for-a-scim-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type UpdateAttributeForUserReq struct {
	_url string
	Org  string

	// scim_user_id parameter
	ScimUserId  string
	RequestBody UpdateAttributeForUserReqBody
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *UpdateAttributeForUserReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return internal.BuildHTTPRequest(ctx, internal.BuildHTTPRequestOptions{
		Body:        r.RequestBody,
		ExplicitURL: r._url,
		HeaderVals: map[string]*string{
			"accept":       internal.String("application/scim+json"),
			"content-type": internal.String("application/json"),
		},
		JSONRequestBody: true,
		Method:          "PATCH",
		Options:         opt,
		URLPath:         fmt.Sprintf("/scim/v2/organizations/%v/Users/%v", r.Org, r.ScimUserId),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *UpdateAttributeForUserReq) Rel(link string, resp *UpdateAttributeForUserResponse) bool {
	u := internal.RelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

// UpdateAttributeForUserReqBodyOperations is a value for UpdateAttributeForUserReqBody's Operations field
type UpdateAttributeForUserReqBodyOperations struct {
	Op    *string                                       `json:"op"`
	Path  *string                                       `json:"path,omitempty"`
	Value *UpdateAttributeForUserReqBodyOperationsValue `json:"value,omitempty"`
}

// UpdateAttributeForUserReqBodyOperationsValue is a value for UpdateAttributeForUserReqBodyOperations's Value field
type UpdateAttributeForUserReqBodyOperationsValue struct {
	Active     *bool   `json:"active,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	FamilyName *string `json:"familyName,omitempty"`
	GivenName  *string `json:"givenName,omitempty"`
	UserName   *string `json:"userName,omitempty"`
}

/*
UpdateAttributeForUserReqBody is a request body for scim/update-attribute-for-user

https://developer.github.com/v3/scim/#update-an-attribute-for-a-scim-user
*/
type UpdateAttributeForUserReqBody struct {

	// Set of operations to be performed
	Operations []UpdateAttributeForUserReqBodyOperations `json:"Operations"`
	Schemas    []string                                  `json:"schemas,omitempty"`
}

/*
UpdateAttributeForUserResponse is a response for UpdateAttributeForUser

https://developer.github.com/v3/scim/#update-an-attribute-for-a-scim-user
*/
type UpdateAttributeForUserResponse struct {
	httpResponse *http.Response
	Data         components.ScimUser
}

// HTTPResponse returns the *http.Response
func (r *UpdateAttributeForUserResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *UpdateAttributeForUserResponse) ReadResponse(resp *http.Response) error {
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
