package octo

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/willabides/octo-go/components"
	"github.com/willabides/octo-go/requests"
)

//go:generate go run ./generator -schema "api.github.com.json" -pkgpath "github.com/willabides/octo-go" -pkg octo

// Common values for rel links
const (
	RelNext  = "next"
	RelPrev  = "prev"
	RelFirst = "first"
	RelLast  = "last"
)

// ISOTimeString returns a pointer to tm formated as an iso8601/rfc3339 string
func ISOTimeString(tm time.Time) *string {
	return String(tm.Format(time.RFC3339))
}

// String returns a pointer to s
func String(s string) *string {
	return &s
}

// Bool returns a pointer to b
func Bool(b bool) *bool {
	return &b
}

// Int64 returns a pointer to i
func Int64(i int64) *int64 {
	return &i
}

// Client is a set of options to apply to requests
type Client []requests.Option

// NewClient returns a new Client
func NewClient(opt ...requests.Option) Client {
	return opt
}

var relLinkExp = regexp.MustCompile(`<(.+?)>\s*;\s*rel="([^"]*)"`)

// RelLink returns the content of lnk from the response's Link header or "" if it does not exist
func RelLink(resp *http.Response, lnk string) string {
	for _, link := range resp.Header.Values("Link") {
		for _, match := range relLinkExp.FindAllStringSubmatch(link, -1) {
			if match[2] == lnk {
				return match[1]
			}
		}
	}
	return ""
}

// RateLimit - The maximum number of requests you're permitted to make per hour.
//  returns -1 if no X-RateLimit-Limit value exists in the response header
func RateLimit(r *http.Response) int {
	return intResponseHeaderOrNegOne(r, "X-RateLimit-Limit")
}

// RateLimitRemaining - The number of requests remaining in the current rate limit window.
//  returns -1 if no X-RateLimit-Remaining value exists in the response header
func RateLimitRemaining(r *http.Response) int {
	return intResponseHeaderOrNegOne(r, "X-RateLimit-Remaining")
}

// RateLimitReset - X-RateLimit-Reset
//  returns time.Zero if no X-RateLimit-Reset value exists in the response header
func RateLimitReset(r *http.Response) time.Time {
	resetTS := intResponseHeaderOrNegOne(r, "X-RateLimit-Reset")
	if resetTS == -1 {
		return time.Time{}
	}
	return time.Unix(int64(resetTS), 0)
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

// ResponseError is an error from an *http.Response.
type ResponseError interface {
	HttpResponse() *http.Response
	Error() string
	Data() *components.ResponseErrorData // data from the error body if it can be unmarshalled
	IsClientError() bool                 // true if the http status is in the 4xx range
	IsServerError() bool                 // true if the http status is in the 5xx range
}
