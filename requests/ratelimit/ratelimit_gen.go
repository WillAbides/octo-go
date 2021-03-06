// Code generated by octo-go; DO NOT EDIT.

package ratelimit

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	requests "github.com/willabides/octo-go/requests"
	"net/http"
)

/*
Get performs requests for "rate-limit/get"

Get rate limit status for the authenticated user.

  GET /rate_limit

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user
*/
func Get(ctx context.Context, req *GetReq, opt ...requests.Option) (*GetResponse, error) {
	opts := requests.BuildOptions(opt...)
	if req == nil {
		req = new(GetReq)
	}
	resp := &GetResponse{}

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
Get performs requests for "rate-limit/get"

Get rate limit status for the authenticated user.

  GET /rate_limit

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
func (c Client) Get(ctx context.Context, req *GetReq, opt ...requests.Option) (*GetResponse, error) {
	return Get(ctx, req, append(c, opt...)...)
}

/*
GetReq is request data for Client.Get

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.
*/
type GetReq struct {
	_url string
}

// HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.
func (r *GetReq) HTTPRequest(ctx context.Context, opt ...requests.Option) (*http.Request, error) {
	return buildHTTPRequest(ctx, buildHTTPRequestOptions{
		ExplicitURL: r._url,
		HeaderVals:  map[string]*string{"accept": strPtr("application/json")},
		Method:      "GET",
		Options:     opt,
		URLPath:     fmt.Sprintf("/rate_limit"),
	})
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GetReq) Rel(link string, resp *GetResponse) bool {
	u := getRelLink(resp.HTTPResponse(), link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GetResponse is a response for Get

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user
*/
type GetResponse struct {
	httpResponse *http.Response
	Data         components.RateLimitOverview
}

// HTTPResponse returns the *http.Response
func (r *GetResponse) HTTPResponse() *http.Response {
	return r.httpResponse
}

// ReadResponse reads an *http.Response. Non-nil errors will have the type octo.ResponseError.
func (r *GetResponse) ReadResponse(resp *http.Response) error {
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
