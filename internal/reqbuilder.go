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

// BuildHTTPRequestOptions builds http requests
type BuildHTTPRequestOptions struct {
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
	Options            []requests.Option
}

func requestHeaders(b BuildHTTPRequestOptions) http.Header {
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

func requestURL(b BuildHTTPRequestOptions) (string, error) {
	expURL := b.ExplicitURL
	if expURL != "" {
		if !hasAttribute(AttrExplicitURL, b.EndpointAttributes) {
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
	if hasAttribute(AttrExplicitURL, b.EndpointAttributes) {
		return "", fmt.Errorf("ExplicitURL must be set")
	}
	opts := requests.BuildOptions(b.Options...)
	u := new(url.URL)
	*u = opts.BaseURL()
	u.Path = path.Join(u.Path, b.URLPath)
	updateURLQuery(u, b.URLQuery)
	return u.String(), nil
}

// BuildHTTPRequest builds an *http.Request. All errors are *errors.RequestError.
func BuildHTTPRequest(ctx context.Context, b BuildHTTPRequestOptions) (*http.Request, error) {
	opts := requests.BuildOptions(b.Options...)
	var bodyReader io.Reader
	var err error
	switch {
	case b.Body == nil:
	case hasAttribute(AttrJSONRequestBody, b.EndpointAttributes):
		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(&b.Body)
		if err != nil {
			return nil, newRequestError("error marshaling json body")
		}
		bodyReader = &buf
	case hasAttribute(AttrBodyUploader, b.EndpointAttributes):
		bodyReader = b.Body.(io.Reader)
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

func hasAttribute(want EndpointAttribute, attributes []EndpointAttribute) bool {
	for _, attr := range attributes {
		if want == attr {
			return true
		}
	}
	return false
}
