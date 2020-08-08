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
CodeScanningGetAlert performs requests for "code-scanning/get-alert"

Get a code scanning alert.

  GET /repos/{owner}/{repo}/code-scanning/alerts/{alert_id}

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
func CodeScanningGetAlert(ctx context.Context, req *CodeScanningGetAlertReq, opt ...RequestOption) (*CodeScanningGetAlertResponse, error) {
	if req == nil {
		req = new(CodeScanningGetAlertReq)
	}
	resp := &CodeScanningGetAlertResponse{request: req}
	r, err := doRequest(ctx, req, "code-scanning/get-alert", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.CodeScanningAlert{}
	err = r.decodeBody(&resp.Data, "code-scanning/get-alert")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodeScanningGetAlert performs requests for "code-scanning/get-alert"

Get a code scanning alert.

  GET /repos/{owner}/{repo}/code-scanning/alerts/{alert_id}

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
func (c Client) CodeScanningGetAlert(ctx context.Context, req *CodeScanningGetAlertReq, opt ...RequestOption) (*CodeScanningGetAlertResponse, error) {
	return CodeScanningGetAlert(ctx, req, append(c, opt...)...)
}

/*
CodeScanningGetAlertReq is request data for Client.CodeScanningGetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type CodeScanningGetAlertReq struct {
	_url  string
	Owner string
	Repo  string

	// alert_id parameter
	AlertId int64
}

func (r *CodeScanningGetAlertReq) url() string {
	return r._url
}

func (r *CodeScanningGetAlertReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/code-scanning/alerts/%v", r.Owner, r.Repo, r.AlertId)
}

func (r *CodeScanningGetAlertReq) method() string {
	return "GET"
}

func (r *CodeScanningGetAlertReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *CodeScanningGetAlertReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *CodeScanningGetAlertReq) body() interface{} {
	return nil
}

func (r *CodeScanningGetAlertReq) dataStatuses() []int {
	return []int{200}
}

func (r *CodeScanningGetAlertReq) validStatuses() []int {
	return []int{200}
}

// HTTPRequest builds an *http.Request
func (r *CodeScanningGetAlertReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "code-scanning/get-alert", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodeScanningGetAlertReq) Rel(link RelName, resp *CodeScanningGetAlertResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
CodeScanningGetAlertResponse is a response for CodeScanningGetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type CodeScanningGetAlertResponse struct {
	response
	request *CodeScanningGetAlertReq
	Data    components.CodeScanningAlert
}

/*
CodeScanningListAlertsForRepo performs requests for "code-scanning/list-alerts-for-repo"

List code scanning alerts for a repository.

  GET /repos/{owner}/{repo}/code-scanning/alerts

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
func CodeScanningListAlertsForRepo(ctx context.Context, req *CodeScanningListAlertsForRepoReq, opt ...RequestOption) (*CodeScanningListAlertsForRepoResponse, error) {
	if req == nil {
		req = new(CodeScanningListAlertsForRepoReq)
	}
	resp := &CodeScanningListAlertsForRepoResponse{request: req}
	r, err := doRequest(ctx, req, "code-scanning/list-alerts-for-repo", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = []components.CodeScanningAlert{}
	err = r.decodeBody(&resp.Data, "code-scanning/list-alerts-for-repo")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodeScanningListAlertsForRepo performs requests for "code-scanning/list-alerts-for-repo"

List code scanning alerts for a repository.

  GET /repos/{owner}/{repo}/code-scanning/alerts

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
func (c Client) CodeScanningListAlertsForRepo(ctx context.Context, req *CodeScanningListAlertsForRepoReq, opt ...RequestOption) (*CodeScanningListAlertsForRepoResponse, error) {
	return CodeScanningListAlertsForRepo(ctx, req, append(c, opt...)...)
}

/*
CodeScanningListAlertsForRepoReq is request data for Client.CodeScanningListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type CodeScanningListAlertsForRepoReq struct {
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

func (r *CodeScanningListAlertsForRepoReq) url() string {
	return r._url
}

func (r *CodeScanningListAlertsForRepoReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/code-scanning/alerts", r.Owner, r.Repo)
}

func (r *CodeScanningListAlertsForRepoReq) method() string {
	return "GET"
}

func (r *CodeScanningListAlertsForRepoReq) urlQuery() url.Values {
	query := url.Values{}
	if r.State != nil {
		query.Set("state", *r.State)
	}
	if r.Ref != nil {
		query.Set("ref", *r.Ref)
	}
	return query
}

func (r *CodeScanningListAlertsForRepoReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *CodeScanningListAlertsForRepoReq) body() interface{} {
	return nil
}

func (r *CodeScanningListAlertsForRepoReq) dataStatuses() []int {
	return []int{200}
}

func (r *CodeScanningListAlertsForRepoReq) validStatuses() []int {
	return []int{200}
}

// HTTPRequest builds an *http.Request
func (r *CodeScanningListAlertsForRepoReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "code-scanning/list-alerts-for-repo", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodeScanningListAlertsForRepoReq) Rel(link RelName, resp *CodeScanningListAlertsForRepoResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
CodeScanningListAlertsForRepoResponse is a response for CodeScanningListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type CodeScanningListAlertsForRepoResponse struct {
	response
	request *CodeScanningListAlertsForRepoReq
	Data    []components.CodeScanningAlert
}
