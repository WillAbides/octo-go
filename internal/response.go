package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/willabides/octo-go/options"
)

// Response is a query response
type Response struct {
	opts         *options.Options
	httpResponse *http.Response
	reqBuilder   *RequestBuilder
}

// DecodeResponseBody unmarshals a response body onto target
func DecodeResponseBody(r *Response, target interface{}) error {
	if r.reqBuilder.HasAttribute(AttrRedirectOnly) {
		return nil
	}
	origBody := r.httpResponse.Body
	var bodyReader io.Reader = origBody
	if r.opts.PreserveResponseBody() {
		var buf bytes.Buffer
		bodyReader = io.TeeReader(r.httpResponse.Body, &buf)
		r.httpResponse.Body = ioutil.NopCloser(&buf)
	}
	//nolint:errcheck // If there's an error draining the response body, there was probably already an error reported.
	defer func() {
		_, _ = ioutil.ReadAll(bodyReader)
		_ = origBody.Close()
	}()
	if !r.statusCodeInList(r.reqBuilder.DataStatuses) {
		return nil
	}
	if target == nil {
		return nil
	}
	return json.NewDecoder(bodyReader).Decode(target)
}

func (r *Response) statusCodeInList(codes []int) bool {
	if r.httpResponse == nil {
		return false
	}
	for _, code := range codes {
		if r.httpResponse.StatusCode == code {
			return true
		}
	}
	return false
}

// SetBoolResult sets the value of ptr to true if r has a 204 status code to true or false if the status code is 404
//  returns an error if the response is any other value
func SetBoolResult(r *Response, ptr *bool) error {
	switch r.httpResponse.StatusCode {
	case 204:
		*ptr = true
	case 404:
		*ptr = false
	default:
		return fmt.Errorf("non-boolean response status")
	}
	return nil
}

// HTTPResponse returns a response's underlying *http.Response
func (r *Response) HTTPResponse() *http.Response {
	return r.httpResponse
}

var relLinkExp = regexp.MustCompile(`<(.+?)>\s*;\s*rel="([^"]*)"`)

// RelLink returns the content of lnk from the response's Link header or "" if it does not exist
func (r *Response) RelLink(lnk string) string {
	if r == nil {
		return ""
	}
	for _, link := range r.httpResponse.Header.Values("Link") {
		for _, match := range relLinkExp.FindAllStringSubmatch(link, -1) {
			if match[2] == lnk {
				return match[1]
			}
		}
	}
	return ""
}

// HasRelLink returns true if lnk exists in the response's Link header
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
//  returns -1 if no X-RateLimit-Limit value exists in the response header
func (r *Response) RateLimit() int {
	return intResponseHeaderOrNegOne(r.httpResponse, "X-RateLimit-Limit")
}

// RateLimitRemaining - The number of requests remaining in the current rate limit window.
//  returns -1 if no X-RateLimit-Remaining value exists in the response header
func (r *Response) RateLimitRemaining() int {
	return intResponseHeaderOrNegOne(r.httpResponse, "X-RateLimit-Remaining")
}

// RateLimitReset - X-RateLimit-Remaining
//  returns time.Zero if no X-RateLimit-Reset value exists in the response header
func (r *Response) RateLimitReset() time.Time {
	i := intResponseHeaderOrNegOne(r.httpResponse, "X-RateLimit-Reset")
	if i == -1 {
		return time.Time{}
	}
	return time.Unix(int64(i), 0)
}
