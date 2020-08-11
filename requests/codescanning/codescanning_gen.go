// Code generated by octo-go; DO NOT EDIT.

package codescanning

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
GetAlert performs requests for "code-scanning/get-alert"

Get a code scanning alert.

  GET /repos/{owner}/{repo}/code-scanning/alerts/{alert_id}

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
func GetAlert(ctx context.Context, req *GetAlertReq, opt ...requests.Option) (*GetAlertResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(GetAlertReq)
	}
	resp := &GetAlertResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.CodeScanningAlert{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
GetAlert performs requests for "code-scanning/get-alert"

Get a code scanning alert.

  GET /repos/{owner}/{repo}/code-scanning/alerts/{alert_id}

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
func (c Client) GetAlert(ctx context.Context, req *GetAlertReq, opt ...requests.Option) (*GetAlertResponse, error) {
	return GetAlert(ctx, req, append(c, opt...)...)
}

/*
GetAlertReq is request data for Client.GetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type GetAlertReq struct {
	_url  string
	Owner string
	Repo  string

	// alert_id parameter
	AlertId int64
}

// HTTPRequest builds an *http.Request
func (r *GetAlertReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *GetAlertReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": internal.String("application/json")},
		Method:           "GET",
		OperationID:      "code-scanning/get-alert",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/repos/%v/%v/code-scanning/alerts/%v", r.Owner, r.Repo, r.AlertId),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetAlertReq) Rel(link string, resp *GetAlertResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetAlertResponse is a response for GetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type GetAlertResponse struct {
	requests.Response
	request *GetAlertReq
	Data    components.CodeScanningAlert
}

/*
ListAlertsForRepo performs requests for "code-scanning/list-alerts-for-repo"

List code scanning alerts for a repository.

  GET /repos/{owner}/{repo}/code-scanning/alerts

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
func ListAlertsForRepo(ctx context.Context, req *ListAlertsForRepoReq, opt ...requests.Option) (*ListAlertsForRepoResponse, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(ListAlertsForRepoReq)
	}
	resp := &ListAlertsForRepoResponse{request: req}
	builder := req.requestBuilder()
	r, err := internal.DoRequest(ctx, builder, opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = []components.CodeScanningAlert{}
	err = internal.DecodeResponseBody(r, builder, opts, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
ListAlertsForRepo performs requests for "code-scanning/list-alerts-for-repo"

List code scanning alerts for a repository.

  GET /repos/{owner}/{repo}/code-scanning/alerts

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
func (c Client) ListAlertsForRepo(ctx context.Context, req *ListAlertsForRepoReq, opt ...requests.Option) (*ListAlertsForRepoResponse, error) {
	return ListAlertsForRepo(ctx, req, append(c, opt...)...)
}

/*
ListAlertsForRepoReq is request data for Client.ListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type ListAlertsForRepoReq struct {
	_url  string
	Owner string
	Repo  string

	// Set to `closed` to list only closed code scanning alerts.
	State *string

	/*
	Returns a list of code scanning alerts for a specific brach reference. The `ref`
	must be formatted as `heads/<branch name>`.
	*/
	Ref *string
}

// HTTPRequest builds an *http.Request
func (r *ListAlertsForRepoReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	opts, err := requests.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *ListAlertsForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}
	if r.State != nil {
		query.Set("state", *r.State)
	}
	if r.Ref != nil {
		query.Set("ref", *r.Ref)
	}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": internal.String("application/json")},
		Method:           "GET",
		OperationID:      "code-scanning/list-alerts-for-repo",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/repos/%v/%v/code-scanning/alerts", r.Owner, r.Repo),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *ListAlertsForRepoReq) Rel(link string, resp *ListAlertsForRepoResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
ListAlertsForRepoResponse is a response for ListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type ListAlertsForRepoResponse struct {
	requests.Response
	request *ListAlertsForRepoReq
	Data    []components.CodeScanningAlert
}
