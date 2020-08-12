package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/willabides/octo-go/requests"
)

// RequestBuilder builds http requests
type RequestBuilder struct {
	OperationID        string
	ExplicitURL        string
	Method             string
	RequiredPreviews   []string
	AllPreviews        []string
	HeaderVals         map[string]*string
	Previews           map[string]bool
	Body               interface{}
	URLQuery           url.Values
	URLPath            string
	EndpointAttributes []EndpointAttribute
}

func (b *RequestBuilder) requestHeaders(opts requests.Options) http.Header {
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

// HasAttribute return true if the endpoint has the attribute
func (b *RequestBuilder) HasAttribute(attribute EndpointAttribute) bool {
	for _, endpointAttribute := range b.EndpointAttributes {
		if endpointAttribute == attribute {
			return true
		}
	}
	return false
}

func (b *RequestBuilder) setURLQuery(u *url.URL) {
	if b.URLQuery == nil {
		return
	}
	q := u.Query()
	if len(q) == 0 {
		u.RawQuery = b.URLQuery.Encode()
		return
	}
	for key, vals := range b.URLQuery {
		q.Del(key)
		for _, val := range vals {
			q.Add(key, val)
		}
	}
	u.RawQuery = q.Encode()
}

func (b *RequestBuilder) requestURL(opts requests.Options) (string, error) {
	expURL := b.ExplicitURL
	if expURL != "" {
		if !b.HasAttribute(AttrExplicitURL) {
			return expURL, nil
		}

		// get rid of any {?templates}
		expURL = strings.SplitN(expURL, "{?", 2)[0]

		u, err := url.Parse(expURL)
		if err != nil {
			return "", err
		}
		b.setURLQuery(u)
		return u.String(), nil
	}
	if b.HasAttribute(AttrExplicitURL) {
		return "", fmt.Errorf("ExplicitURL must be set")
	}
	u := new(url.URL)
	*u = opts.BaseURL()
	u.Path = path.Join(u.Path, b.URLPath)
	b.setURLQuery(u)
	return u.String(), nil
}

// HTTPRequest returns an http request
func (b *RequestBuilder) HTTPRequest(ctx context.Context, opts *requests.Options) (*http.Request, error) {
	var bodyReader io.Reader
	var err error
	switch {
	case b.Body == nil:
	case b.HasAttribute(AttrJSONRequestBody):
		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(&b.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = &buf
	case b.HasAttribute(AttrBodyUploader):
		bodyReader = b.Body.(io.Reader)
	}
	urlString, err := b.requestURL(*opts)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, b.Method, urlString, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header = b.requestHeaders(*opts)
	req.Header.Set("User-Agent", opts.UserAgent())

	authProvider := opts.AuthProvider()
	if authProvider != nil {
		var authHeader string
		authHeader, err = authProvider.AuthorizationHeader(ctx)
		if err != nil {
			return nil, fmt.Errorf("error setting authorization header: %v", err)
		}
		req.Header.Set("Authorization", authHeader)
	}

	return req, nil
}
