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
func (c *Client) CodeScanningGetAlert(ctx context.Context, req *CodeScanningGetAlertReq, opt ...RequestOption) (*CodeScanningGetAlertResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &CodeScanningGetAlertResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(CodeScanningGetAlertResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodeScanningGetAlertReq is request data for Client.CodeScanningGetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type CodeScanningGetAlertReq struct {
	pgURL   string
	Owner   string
	Repo    string
	AlertId int64
}

func (r *CodeScanningGetAlertReq) pagingURL() string {
	return r.pgURL
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
	headerVals := map[string]*string{}
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

func (r *CodeScanningGetAlertReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *CodeScanningGetAlertReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
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
	r.pgURL = u
	return true
}

/*
CodeScanningGetAlertResponseBody is a response body for CodeScanningGetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type CodeScanningGetAlertResponseBody struct {
	components.CodeScanningAlert
}

/*
CodeScanningGetAlertResponse is a response for CodeScanningGetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type CodeScanningGetAlertResponse struct {
	response
	request *CodeScanningGetAlertReq
	Data    *CodeScanningGetAlertResponseBody
}

/*
CodeScanningListAlertsForRepo performs requests for "code-scanning/list-alerts-for-repo"

List code scanning alerts for a repository.

  GET /repos/{owner}/{repo}/code-scanning/alerts

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
func (c *Client) CodeScanningListAlertsForRepo(ctx context.Context, req *CodeScanningListAlertsForRepoReq, opt ...RequestOption) (*CodeScanningListAlertsForRepoResponse, error) {
	r, err := c.doRequest(ctx, req, opt...)
	if err != nil {
		return nil, err
	}
	resp := &CodeScanningListAlertsForRepoResponse{
		request:  req,
		response: *r,
	}
	resp.Data = new(CodeScanningListAlertsForRepoResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodeScanningListAlertsForRepoReq is request data for Client.CodeScanningListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type CodeScanningListAlertsForRepoReq struct {
	pgURL string
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

func (r *CodeScanningListAlertsForRepoReq) pagingURL() string {
	return r.pgURL
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
	headerVals := map[string]*string{}
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

func (r *CodeScanningListAlertsForRepoReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *CodeScanningListAlertsForRepoReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
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
	r.pgURL = u
	return true
}

/*
CodeScanningListAlertsForRepoResponseBody is a response body for CodeScanningListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type CodeScanningListAlertsForRepoResponseBody []struct {
	components.CodeScanningAlert
}

/*
CodeScanningListAlertsForRepoResponse is a response for CodeScanningListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type CodeScanningListAlertsForRepoResponse struct {
	response
	request *CodeScanningListAlertsForRepoReq
	Data    *CodeScanningListAlertsForRepoResponseBody
}
