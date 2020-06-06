// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

/*
EmojisGet performs requests for "emojis/get"

Get.

  GET /emojis

https://developer.github.com/v3/emojis/#emojis
*/
func EmojisGet(ctx context.Context, req *EmojisGetReq, opt ...RequestOption) (*EmojisGetResponse, error) {
	if req == nil {
		req = new(EmojisGetReq)
	}
	resp := &EmojisGetResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	err = r.decodeBody(nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
EmojisGet performs requests for "emojis/get"

Get.

  GET /emojis

https://developer.github.com/v3/emojis/#emojis
*/
func (c *Client) EmojisGet(ctx context.Context, req *EmojisGetReq, opt ...RequestOption) (*EmojisGetResponse, error) {
	return EmojisGet(ctx, req, append(c.opts, opt...)...)
}

/*
EmojisGetReq is request data for Client.EmojisGet

https://developer.github.com/v3/emojis/#emojis
*/
type EmojisGetReq struct {
	_url string
}

func (r *EmojisGetReq) url() string {
	return r._url
}

func (r *EmojisGetReq) urlPath() string {
	return fmt.Sprintf("/emojis")
}

func (r *EmojisGetReq) method() string {
	return "GET"
}

func (r *EmojisGetReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *EmojisGetReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *EmojisGetReq) body() interface{} {
	return nil
}

func (r *EmojisGetReq) dataStatuses() []int {
	return []int{}
}

func (r *EmojisGetReq) validStatuses() []int {
	return []int{200}
}

func (r *EmojisGetReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *EmojisGetReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *EmojisGetReq) Rel(link RelName, resp *EmojisGetResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
EmojisGetResponse is a response for EmojisGet

https://developer.github.com/v3/emojis/#emojis
*/
type EmojisGetResponse struct {
	response
	request *EmojisGetReq
}
