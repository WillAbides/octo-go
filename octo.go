package octo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:generate go run ./generator -schema "api.github.com.json" -pkgpath "github.com/willabides/octo-go" -pkg octo

// RelLink is the name for a relative link in a Link header. Used for paging.
type RelName string

// Common RelLink values
const (
	RelNext  RelName = "next"
	RelPrev  RelName = "prev"
	RelFirst RelName = "first"
	RelLast  RelName = "last"
)

type response struct {
	opts          requestOpts
	httpResponse  *http.Response
	httpRequester httpRequester
}

func hasEndpointAttribute(builder requestBuilder, attribute endpointAttribute) bool {
	for _, attr := range builder.endpointAttributes() {
		if attr == attribute {
			return true
		}
	}
	return false
}

func (r *response) decodeBody(target interface{}) error {
	if hasEndpointAttribute(r.httpRequester, attrRedirectOnly) {
		return nil
	}
	origBody := r.httpResponse.Body
	var bodyReader io.Reader = origBody
	if r.opts.preserveResponseBody {
		var buf bytes.Buffer
		bodyReader = io.TeeReader(r.httpResponse.Body, &buf)
		r.httpResponse.Body = ioutil.NopCloser(&buf)
	}
	//nolint:errcheck // If there's an error draining the response body, there was probably already an error reported.
	defer func() {
		_, _ = ioutil.ReadAll(bodyReader)
		_ = origBody.Close()
	}()
	if !r.statusCodeInList(r.httpRequester.dataStatuses()) {
		return nil
	}
	if target == nil {
		return nil
	}
	return json.NewDecoder(bodyReader).Decode(target)
}

func (r *response) statusCodeInList(codes []int) bool {
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

// boolResult maps a 204 status code to true and 404 to false
//  returns an error if the response is any other value
func (r *response) setBoolResult(ptr *bool) error {
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
func (r *response) HTTPResponse() *http.Response {
	return r.httpResponse
}

// HasRelLink returns true if lnk exists in the response's Link header
func (r *response) HasRelLink(lnk RelName) bool {
	return r.RelLink(lnk) != ""
}

// RelLink returns the content of lnk from the response's Link header or "" if it does not exist
func (r *response) RelLink(lnk RelName) string {
	if r == nil {
		return ""
	}
	for _, link := range r.httpResponse.Header.Values("Link") {
		for _, match := range relLinkExp.FindAllStringSubmatch(link, -1) {
			if match[2] == string(lnk) {
				return match[1]
			}
		}
	}
	return ""
}

func (r *response) intHeaderOrNegOne(headerName string) int {
	hdr := r.httpResponse.Header.Get(headerName)
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
func (r *response) RateLimit() int {
	return r.intHeaderOrNegOne("X-RateLimit-Limit")
}

// RateLimitRemaining - The number of requests remaining in the current rate limit window.
//  returns -1 if no X-RateLimit-Remaining value exists in the response header
func (r *response) RateLimitRemaining() int {
	return r.intHeaderOrNegOne("X-RateLimit-Remaining")
}

// RateLimitReset - X-RateLimit-Remaining
//  returns time.Zero if no X-RateLimit-Reset value exists in the response header
func (r *response) RateLimitReset() time.Time {
	i := r.intHeaderOrNegOne("X-RateLimit-Reset")
	if i == -1 {
		return time.Time{}
	}
	return time.Unix(int64(i), 0)
}

type httpRequester interface {
	HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error)
	requestBuilder
}

var relLinkExp = regexp.MustCompile(`<(.+?)>\s*;\s*rel="([^"]*)"`)

func requestHeaders(headers map[string]*string, previews map[string]bool) http.Header {
	header := make(http.Header, len(headers)+len(previews)+1)
	header.Set("Accept", "application/vnd.github.v3+json")
	for k, v := range headers {
		if v == nil {
			continue
		}
		header.Add(k, *v)
	}
	for previewName, ok := range previews {
		if !ok {
			continue
		}
		header.Add("Accept", fmt.Sprintf("application/vnd.github.%s-preview+json", previewName))
	}
	return header
}

type requestBuilder interface {
	url() string
	urlPath() string
	method() string
	urlQuery() url.Values
	header(requiredPreviews, allPreviews bool) http.Header
	body() interface{}
	validStatuses() []int
	dataStatuses() []int
	endpointAttributes() []endpointAttribute
}

func httpRequestURL(builder requestBuilder, options requestOpts) (string, error) {
	expURL := builder.url()
	if expURL != "" {
		if !hasEndpointAttribute(builder, attrExplicitURL) {
			return expURL, nil
		}
		// get rid of any {?templates}
		expURL = strings.SplitN(expURL, "{?", 2)[0]

		u, err := url.Parse(expURL)
		if err != nil {
			return "", err
		}
		expQuery := u.Query()
		for key, vals := range builder.urlQuery() {
			expQuery.Del(key)
			for _, val := range vals {
				expQuery.Add(key, val)
			}
		}
		u.RawQuery = expQuery.Encode()
		return u.String(), nil
	}
	if hasEndpointAttribute(builder, attrExplicitURL) {
		return "", fmt.Errorf("URL must be set")
	}

	u := options.baseURL
	u.Path = path.Join(u.Path, builder.urlPath())
	urlQuery := builder.urlQuery()
	if urlQuery != nil {
		u.RawQuery = urlQuery.Encode()
	}
	return u.String(), nil
}

func buildHTTPRequest(ctx context.Context, builder requestBuilder, opts []RequestOption) (*http.Request, error) {
	options, err := buildRequestOptions(opts)
	if err != nil {
		return nil, err
	}
	var bodyReader io.Reader
	body := builder.body()
	switch {
	case body == nil:
	case hasEndpointAttribute(builder, attrJSONRequestBody):
		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(&body)
		if err != nil {
			return nil, err
		}
		bodyReader = &buf
	case hasEndpointAttribute(builder, attrBodyUploader):
		bodyReader = body.(io.Reader)
	}
	urlString, err := httpRequestURL(builder, options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, builder.method(), urlString, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header = builder.header(options.requiredPreviews, options.allPreviews)
	req.Header.Set("User-Agent", options.userAgent)
	return req, nil
}

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

func doRequest(ctx context.Context, requester httpRequester, opt ...RequestOption) (*response, error) {
	req, err := requester.HTTPRequest(ctx, opt...)
	if err != nil {
		return nil, err
	}
	ro, err := buildRequestOptions(opt)
	if err != nil {
		return nil, err
	}
	if ro.authProvider != nil {
		var authHeader string
		authHeader, err = ro.authProvider.AuthorizationHeader(ctx)
		if err != nil {
			return nil, fmt.Errorf("error setting authorization header: %v", err)
		}
		req.Header.Set("Authorization", authHeader)
	}
	httpClient := ro.httpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	httpResponse, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	resp := &response{
		httpResponse:  httpResponse,
		httpRequester: requester,
		opts:          ro,
	}

	err = errorCheck(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Client is a set of options to apply to requests
type Client []RequestOption

// NewClient returns a new Client
func NewClient(opt ...RequestOption) Client {
	return Client(opt)
}
