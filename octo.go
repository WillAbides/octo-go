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
	if hasEndpointAttribute(r.httpRequester, attrRedirect) {
		return nil
	}
	origBody := r.httpResponse.Body
	var bodyReader io.Reader = origBody
	if r.opts.preserveResponseBody {
		var buf bytes.Buffer
		bodyReader = io.TeeReader(r.httpResponse.Body, &buf)
		r.httpResponse.Body = ioutil.NopCloser(&buf)
	}
	defer func() {
		_, _ = ioutil.ReadAll(bodyReader) //nolint:errcheck
		_ = origBody.Close()              //nolint:errcheck
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

// Client is a client for the GitHub API
type Client struct {
	requestOpts requestOpts
	HttpClient  *http.Client
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, opt ...RequestOption) *Client {
	client := &Client{
		HttpClient:  httpClient,
		requestOpts: buildRequestOptions(opt),
	}
	return client
}

type httpRequester interface {
	httpRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error)
	requestBuilder
}

func (c *Client) doRequest(ctx context.Context, requester httpRequester, opt ...RequestOption) (*response, error) {
	if c.HttpClient == nil {
		c.HttpClient = http.DefaultClient
	}
	opts := make([]RequestOption, 0, len(opt)+1)
	opts = append(opts, resetOptions(c.requestOpts))
	opts = append(opts, opt...)
	req, err := requester.httpRequest(ctx, opts...)
	if err != nil {
		return nil, err
	}
	httpResponse, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	resp := &response{
		httpResponse:  httpResponse,
		httpRequester: requester,
		opts:          buildRequestOptions(opts),
	}

	err = errorCheck(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetHttpClient sets a client's underlying *http.Client
func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.HttpClient = httpClient
}

// SetRequestOptions sets options that will be used on all requests this client makes
func (c *Client) SetRequestOptions(opt ...RequestOption) {
	for _, o := range opt {
		o(&c.requestOpts)
	}
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
	pagingURL() string
	urlPath() string
	method() string
	urlQuery() url.Values
	header(requiredPreviews, allPreviews bool) http.Header
	body() interface{}
	validStatuses() []int
	dataStatuses() []int
	endpointAttributes() []endpointAttribute
}

func httpRequestURL(builder requestBuilder, options requestOpts) string {
	if builder.pagingURL() != "" {
		return builder.pagingURL()
	}
	u := options.baseURL
	u.Path = path.Join(u.Path, builder.urlPath())
	urlQuery := builder.urlQuery()
	if urlQuery != nil {
		u.RawQuery = urlQuery.Encode()
	}
	return u.String()
}

func buildHTTPRequest(ctx context.Context, builder requestBuilder, opts []RequestOption) (*http.Request, error) {
	options := buildRequestOptions(opts)
	var bodyReader io.ReadCloser
	body := builder.body()
	if body != nil {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(&body)
		if err != nil {
			return nil, err
		}
		bodyReader = ioutil.NopCloser(&buf)
	}
	urlString := httpRequestURL(builder, options)
	req, err := http.NewRequestWithContext(ctx, builder.method(), urlString, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header = builder.header(options.requiredPreviews, options.allPreviews)
	req.Header.Set("User-Agent", options.userAgent)
	return req, nil
}

//ISOTimeString returns a pointer to tm formated as an iso8601/rfc3339 string
func ISOTimeString(tm time.Time) *string {
	return String(tm.Format(time.RFC3339))
}

//String returns a pointer to s
func String(s string) *string {
	return &s
}

//Bool returns a pointer to b
func Bool(b bool) *bool {
	return &b
}

//Int64 returns a pointer to i
func Int64(i int64) *int64 {
	return &i
}
