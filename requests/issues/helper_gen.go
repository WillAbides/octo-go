// Code generated by octo-go; DO NOT EDIT.

package issues

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
	"strings"

	"github.com/willabides/octo-go/components"
	"github.com/willabides/octo-go/requests"
)

// Client is a set of options to apply to requests
type Client []requests.Option

// NewClient returns a new Client
func NewClient(opt ...requests.Option) Client {
	return opt
}

func strPtr(s string) *string { return &s }

// buildHTTPRequestOptions builds http requests
type buildHTTPRequestOptions struct {
	ExplicitURL        string
	Method             string
	RequiredPreviews   []string
	AllPreviews        []string
	HeaderVals         map[string]*string
	Previews           map[string]bool
	Body               interface{}
	URLQuery           url.Values
	URLPath            string
	Options            []requests.Option
	RequireExplicitURL bool
}

func requestHeaders(b buildHTTPRequestOptions) http.Header {
	opts := requests.BuildOptions(b.Options...)
	previews := b.Previews
	headers := b.HeaderVals
	if opts.RequiredPreviews() {
		for _, preview := range b.RequiredPreviews {
			previews[preview] = true
		}
	}
	if opts.AllPreviews() {
		for _, preview := range b.AllPreviews {
			previews[preview] = true
		}
	}
	header := make(http.Header, len(headers)+len(previews)+1)
	for k, v := range headers {
		if v == nil {
			continue
		}
		header.Add(k, *v)
	}
	if header.Get("accept") == "" {
		header.Set("Accept", "application/vnd.github.v3+json")
	}
	for previewName, ok := range previews {
		if !ok {
			continue
		}
		header.Add("Accept", fmt.Sprintf("application/vnd.github.%s-preview+json", previewName))
	}
	return header
}

// updateURLQuery updates u's query with vals
func updateURLQuery(u *url.URL, vals url.Values) {
	if len(vals) == 0 {
		return
	}
	q := u.Query()
	if len(q) == 0 {
		u.RawQuery = vals.Encode()
		return
	}
	for key, vals := range vals {
		q.Del(key)
		for _, val := range vals {
			q.Add(key, val)
		}
	}
	u.RawQuery = q.Encode()
}

func requestURL(b buildHTTPRequestOptions) (string, error) {
	expURL := b.ExplicitURL
	if expURL != "" {
		if !b.RequireExplicitURL {
			return expURL, nil
		}

		// get rid of any {?templates}
		expURL = strings.SplitN(expURL, "{?", 2)[0]

		u, err := url.Parse(expURL)
		if err != nil {
			return "", err
		}
		updateURLQuery(u, b.URLQuery)
		return u.String(), nil
	}
	if b.RequireExplicitURL {
		return "", fmt.Errorf("ExplicitURL must be set")
	}
	opts := requests.BuildOptions(b.Options...)
	u := new(url.URL)
	*u = opts.BaseURL()
	u.Path = path.Join(u.Path, b.URLPath)
	updateURLQuery(u, b.URLQuery)
	return u.String(), nil
}

func makeBodyReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	if rdr, ok := body.(io.Reader); ok {
		return rdr, nil
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&body)
	if err != nil {
		return nil, newRequestError("error marshaling json body")
	}
	return &buf, nil
}

// buildHTTPRequest builds an *http.Request. All errors are *errors.RequestError.
func buildHTTPRequest(ctx context.Context, b buildHTTPRequestOptions) (*http.Request, error) {
	opts := requests.BuildOptions(b.Options...)
	bodyReader, err := makeBodyReader(b.Body)
	if err != nil {
		return nil, err
	}
	urlString, err := requestURL(b)
	if err != nil {
		return nil, newRequestError(err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, b.Method, urlString, bodyReader)
	if err != nil {
		return nil, newRequestError(err.Error())
	}
	req.Header = requestHeaders(b)
	req.Header.Set("User-Agent", opts.UserAgent())

	authProvider := opts.AuthProvider()
	if authProvider != nil {
		var authHeader string
		authHeader, err = authProvider.AuthorizationHeader(ctx)
		if err != nil {
			return nil, newRequestError("error setting Authorization header")
		}
		req.Header.Set("Authorization", authHeader)
	}

	return req, nil
}

// newResponseError returns a new *responseError
func newResponseError(msg string, resp *http.Response) *responseError {
	data, err := unmarshalErrorData(resp)
	if err != nil {
		data = nil
	}
	return &responseError{
		resp: resp,
		msg:  msg,
		data: data,
	}
}

// responseError implements errors.responseError
type responseError struct {
	resp *http.Response
	msg  string
	data *components.ResponseErrorData
}

// HttpResponse implements errors.responseError
func (r *responseError) HttpResponse() *http.Response {
	return r.resp
}

func (r *responseError) Error() string {
	msg := r.msg
	if r.data != nil && r.data.Message != "" {
		msg += ": " + r.data.Message
	}
	return msg
}

// Data implements errors.responseError
func (r *responseError) Data() *components.ResponseErrorData {
	return r.data
}

// IsClientError implements errors.responseError
func (r *responseError) IsClientError() bool {
	return r.resp != nil && r.resp.StatusCode >= 400 && r.resp.StatusCode < 500
}

// IsServerError implements errors.responseError
func (r *responseError) IsServerError() bool {
	return r.resp != nil && r.resp.StatusCode >= 500 && r.resp.StatusCode < 600
}

// NewRequestError returns a new RequestError
func newRequestError(msg string) error {
	return &requests.RequestError{
		Message: msg,
	}
}

// responseErrorCheck checks for error responses
func responseErrorCheck(resp *http.Response, validStatuses []int) error {
	code := resp.StatusCode
	for _, wantStatus := range validStatuses {
		if code == wantStatus {
			return nil
		}
	}
	switch {
	case code >= 400 && code < 600:
		return newResponseError(fmt.Sprintf("client error %d", resp.StatusCode), resp)
	case code >= 500 && code < 600:
		return newResponseError(fmt.Sprintf("server error %d", resp.StatusCode), resp)
	}
	if isRedirectOnly(validStatuses) && code < 300 {
		return nil
	}
	msg := fmt.Sprintf("received unexpected http status code %d, expected codes are %v", code, validStatuses)
	return newResponseError(msg, resp)
}

func isRedirectOnly(validStatuses []int) bool {
	if len(validStatuses) == 0 {
		return false
	}
	for _, vs := range validStatuses {
		if vs < 300 || vs > 399 {
			return false
		}
	}
	return true
}

func unmarshalErrorData(resp *http.Response) (*components.ResponseErrorData, error) {
	if resp.Body == nil {
		return nil, fmt.Errorf("no body")
	}
	var nextBody bytes.Buffer
	bodyReader := io.TeeReader(resp.Body, &nextBody)
	//nolint:errcheck // If there's an error draining the response body, there was probably already an error reported.
	defer func() {
		_, _ = ioutil.ReadAll(bodyReader)
		_ = resp.Body.Close()
		resp.Body = ioutil.NopCloser(&nextBody)
	}()
	var errorData components.ResponseErrorData
	err := json.NewDecoder(bodyReader).Decode(&errorData)
	if err != nil {
		return nil, err
	}
	return &errorData, nil
}

var relLinkExp = regexp.MustCompile(`<(.+?)>\s*;\s*rel="([^"]*)"`)

// getRelLink returns the content of lnk from the response's Link header or "" if it does not exist
func getRelLink(resp *http.Response, lnk string) string {
	for _, link := range resp.Header.Values("Link") {
		for _, match := range relLinkExp.FindAllStringSubmatch(link, -1) {
			if match[2] == lnk {
				return match[1]
			}
		}
	}
	return ""
}

// unmarshalResponseBody unmarshalls a response body onto target. Non-nil errors will have the type *errors.ResponseError.
func unmarshalResponseBody(r *http.Response, target interface{}) error {
	body := r.Body
	bb, err := ioutil.ReadAll(body)
	if err != nil {
		return newResponseError("could not read response body", r)
	}
	err = body.Close()
	if err != nil {
		return newResponseError("could not close response body", r)
	}
	err = json.Unmarshal(bb, &target)
	if err != nil {
		return newResponseError("could not unmarshal json from response body", r)
	}
	return nil
}

// intInSlice returns true if i is in want
func intInSlice(i int, want []int) bool {
	for _, code := range want {
		if i == code {
			return true
		}
	}
	return false
}

// setBoolResult sets the value of ptr to true if r has a 204 status code to true or false if the status code is 404
//  returns an error if the response is any other value
func setBoolResult(r *http.Response, ptr *bool) error {
	switch r.StatusCode {
	case 204:
		*ptr = true
	case 404:
		*ptr = false
	default:
		return newResponseError("non-boolean response status", r)
	}
	return nil
}
