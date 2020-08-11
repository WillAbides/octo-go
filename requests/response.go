package requests

import (
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// NewResponse returns a Response
func NewResponse(httpResponse *http.Response) *Response {
	return &Response{
		httpResponse: httpResponse,
	}
}

// Response is a query common
type Response struct {
	httpResponse *http.Response
}

// HTTPResponse returns a common's underlying *http.Response
func (r *Response) HTTPResponse() *http.Response {
	return r.httpResponse
}

var relLinkExp = regexp.MustCompile(`<(.+?)>\s*;\s*rel="([^"]*)"`)

// RelLink returns the content of lnk from the common's Link header or "" if it does not exist
func (r *Response) RelLink(lnk string) string {
	if r == nil {
		return ""
	}
	for _, link := range r.HTTPResponse().Header.Values("Link") {
		for _, match := range relLinkExp.FindAllStringSubmatch(link, -1) {
			if match[2] == lnk {
				return match[1]
			}
		}
	}
	return ""
}

// HasRelLink returns true if lnk exists in the common's Link header
func (r *Response) HasRelLink(lnk string) bool {
	return r.RelLink(lnk) != ""
}

func intResponseHeaderOrNegOne(resp *http.Response, headerName string) int {
	hdr := resp.Header.Get(headerName)
	if hdr == "" {
		return -1
	}
	i, err := strconv.Atoi(hdr)
	if err != nil {
		return -1
	}
	return i
}

// RateLimit - The maximum number of requests you're permitted to make per hour.
//  returns -1 if no X-RateLimit-Limit value exists in the common header
func (r *Response) RateLimit() int {
	return intResponseHeaderOrNegOne(r.httpResponse, "X-RateLimit-Limit")
}

// RateLimitRemaining - The number of requests remaining in the current rate limit window.
//  returns -1 if no X-RateLimit-Remaining value exists in the common header
func (r *Response) RateLimitRemaining() int {
	return intResponseHeaderOrNegOne(r.httpResponse, "X-RateLimit-Remaining")
}

// RateLimitReset - X-RateLimit-Remaining
//  returns time.Zero if no X-RateLimit-Reset value exists in the common header
func (r *Response) RateLimitReset() time.Time {
	i := intResponseHeaderOrNegOne(r.httpResponse, "X-RateLimit-Reset")
	if i == -1 {
		return time.Time{}
	}
	return time.Unix(int64(i), 0)
}
