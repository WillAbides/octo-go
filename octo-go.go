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
	"time"
)

//go:generate go run ./generator -schema "api.github.com.json" -pkgpath "github.com/willabides/octo-go" -pkg octo

var relLinkExp = regexp.MustCompile(`<(.+?)>\s*;\s*rel="([^"]*)"`)

//ResponseRelLink returns a rel link from resp's Link header
func ResponseRelLink(resp *http.Response, linkName string) string {
	if resp == nil {
		return ""
	}
	for _, link := range resp.Header.Values("Link") {
		for _, match := range relLinkExp.FindAllStringSubmatch(link, -1) {
			if match[2] == linkName {
				return match[1]
			}
		}
	}
	return ""
}

//ResponseNextPageURL returns the next page url or ""
func ResponseNextPageURL(resp *http.Response) string {
	return ResponseRelLink(resp, "next")
}

//ResponsePrevPageURL returns the previous page url or ""
func ResponsePrevPageURL(resp *http.Response) string {
	return ResponseRelLink(resp, "prev")
}

//ResponseFirstPageURL returns the first page url or ""
func ResponseFirstPageURL(resp *http.Response) string {
	return ResponseRelLink(resp, "first")
}

//ResponseLastPageURL returns the last page url or ""
func ResponseLastPageURL(resp *http.Response) string {
	return ResponseRelLink(resp, "last")
}

//UnmarshalResponseBody unmarshals resp's body to target and closes the body
func UnmarshalResponseBody(resp *http.Response, target interface{}) (err error) {
	defer func() {
		err = resp.Body.Close()
	}()
	return json.NewDecoder(resp.Body).Decode(target)
}

//ResponseNextPageReq returns an http request to get the next page or nil if there is no next page
func ResponseNextPageReq(ctx context.Context, req *http.Request, resp *http.Response) (*http.Request, error) {
	return responseRelReq(ctx, req, resp, "next")
}

func responseRelReq(ctx context.Context, req *http.Request, resp *http.Response, link string) (*http.Request, error) {
	np := ResponseRelLink(resp, link)
	if np == "" {
		return nil, nil
	}
	u, err := url.Parse(np)
	if err != nil {
		return nil, fmt.Errorf("could not parse url")
	}
	result := req.Clone(ctx)
	result.URL = u
	return result, nil
}

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

func httpRequest(ctx context.Context, urlPath, method string, urlQuery url.Values, header http.Header, body interface{}, opts []RequestOption) (*http.Request, error) {
	options := buildRequestOptions(opts)
	var bodyReader io.ReadCloser
	if body != nil {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(&body)
		if err != nil {
			return nil, err
		}
		bodyReader = ioutil.NopCloser(&buf)
	}
	u := options.BaseURL
	u.Path = path.Join(u.Path, urlPath)
	u.RawQuery = urlQuery.Encode()
	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header = header
	req.Header.Set("User-Agent", options.UserAgent)
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

//Int returns a pointer to i
func Int(i int) *int {
	return &i
}
