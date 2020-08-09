package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	// got a request that isn't matched by any matcher
	StatusUnexpectedRequest = 600
	// couldn't clone the request body
	StatusErrorCloningRequestBody = 601
)

// RequestMatcher matches requests with responses
type RequestMatcher interface {
	Match(req *http.Request, requestCount int) bool
}

type matcherFunc func(req *http.Request, requestCount int) bool

func (fn matcherFunc) Match(req *http.Request, requestCount int) bool {
	return fn(req, requestCount)
}

// MatchBody matches when the request's body is equal to wantBody
func MatchBody(wantBody string) RequestMatcher {
	return matcherFunc(func(req *http.Request, requestCount int) bool {
		clone, err := cloneRequestWithBody(req)
		if err != nil {
			return false
		}
		var body string
		if clone.Body != nil {
			b, err := ioutil.ReadAll(clone.Body)
			if err != nil {
				return false
			}
			body = string(b)
		}
		return body == wantBody
	})
}

// MatchRequestPath matches when the request's URL path equals wantPath
func MatchRequestPath(wantPath string) RequestMatcher {
	return matcherFunc(func(req *http.Request, requestCount int) bool {
		return wantPath == req.URL.Path
	})
}

// MatchRequestQuery matches when the request's URL query matches wantQuery
func MatchRequestQuery(wantQuery string) RequestMatcher {
	return matcherFunc(func(req *http.Request, requestCount int) bool {
		return wantQuery == req.URL.RawQuery
	})
}

// MatchAll matches when all matchers match the request
func MatchAll(matcher ...RequestMatcher) RequestMatcher {
	return matcherFunc(func(req *http.Request, requestCount int) bool {
		for _, rm := range matcher {
			clone, err := cloneRequestWithBody(req)
			if err != nil {
				return false
			}
			if !rm.Match(clone, requestCount) {
				return false
			}
		}
		return true
	})
}

// ExpectedRequest is a request and common for your test
type ExpectedRequest struct {
	Matcher RequestMatcher
	Handler http.Handler
}

// RequestHandler is a handler that will serve responses for expected requests
type RequestHandler struct {
	expectedRequests []*ExpectedRequest
	count            int
	mu               sync.Mutex
}

// Expect adds an expected request to the handler
func (r *RequestHandler) Expect(expectedRequest *ExpectedRequest) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.expectedRequests = append(r.expectedRequests, expectedRequest)
}

// Count returns the number of requests the handler has served
func (r *RequestHandler) Count() int {
	r.mu.Lock()
	count := r.count
	r.mu.Unlock()
	return count
}

func cloneRequestWithBody(req *http.Request) (*http.Request, error) {
	clone := req.Clone(req.Context())
	if req.Body == nil {
		return clone, nil
	}
	ogBody := req.Body
	defer func() {
		_ = ogBody.Close() //nolint:errcheck // no problem if this fails
	}()
	buf, err := ioutil.ReadAll(ogBody)
	if err != nil {
		return nil, err
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	clone.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return clone, nil
}

func (r *RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mu.Lock()
	defer func() {
		r.count++
		r.mu.Unlock()
	}()
	for _, request := range r.expectedRequests {
		clone, err := cloneRequestWithBody(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("error from cloneRequestWithBody: %v", err), StatusErrorCloningRequestBody)
			return
		}
		if request.Matcher.Match(clone, r.count) {
			request.Handler.ServeHTTP(w, req)
			return
		}
	}
	http.Error(w, "no ExpectedRequest matches req", StatusUnexpectedRequest)
}
