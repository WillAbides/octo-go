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
MetaGet performs requests for "meta/get"

Get.

  GET /meta

https://developer.github.com/v3/meta/#meta
*/
func (c *Client) MetaGet(ctx context.Context, req *MetaGetReq, opt ...RequestOption) (*MetaGetResponse, error) {
	resp := &MetaGetResponse{request: req}
	r, err := c.doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = new(MetaGetResponseBody)
	err = r.decodeBody(resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
MetaGetReq is request data for Client.MetaGet

https://developer.github.com/v3/meta/#meta
*/
type MetaGetReq struct {
	_url string
}

func (r *MetaGetReq) url() string {
	return r._url
}

func (r *MetaGetReq) urlPath() string {
	return fmt.Sprintf("/meta")
}

func (r *MetaGetReq) method() string {
	return "GET"
}

func (r *MetaGetReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *MetaGetReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *MetaGetReq) body() interface{} {
	return nil
}

func (r *MetaGetReq) dataStatuses() []int {
	return []int{200}
}

func (r *MetaGetReq) validStatuses() []int {
	return []int{200}
}

func (r *MetaGetReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// httpRequest creates an http request
func (r *MetaGetReq) httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *MetaGetReq) Rel(link RelName, resp *MetaGetResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
MetaGetResponseBody is a response body for MetaGet

https://developer.github.com/v3/meta/#meta
*/
type MetaGetResponseBody struct {
	components.ApiOverview
}

/*
MetaGetResponse is a response for MetaGet

https://developer.github.com/v3/meta/#meta
*/
type MetaGetResponse struct {
	response
	request *MetaGetReq
	Data    *MetaGetResponseBody
}
