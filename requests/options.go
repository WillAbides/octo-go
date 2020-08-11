package requests

import (
	"context"
	"net/http"
	"net/url"
)

// AuthProvider sets the Authorization header authenticate you with the API
type AuthProvider interface {
	AuthorizationHeader(ctx context.Context) (string, error)
}

// Option is a request option
type Option func(opts *Options) error

var defaultOptions = Options{
	baseURL: url.URL{
		Host:   "api.github.com",
		Scheme: "https",
	},
	userAgent:  "octo-go",
	httpClient: http.DefaultClient,
}

// BuildOptions turns a list of opt into *Options
func BuildOptions(opt ...Option) (*Options, error) {
	result := defaultOptions
	for _, o := range opt {
		if o == nil {
			continue
		}
		err := o(&result)
		if err != nil {
			return nil, nil
		}
	}
	return &result, nil
}

// Options is options
type Options struct {
	baseURL              url.URL
	userAgent            string
	requiredPreviews     bool
	allPreviews          bool
	preserveResponseBody bool
	authProvider         AuthProvider
	httpClient           *http.Client
}

// HttpClient return httpClient
func (o *Options) HttpClient() *http.Client {
	return o.httpClient
}

// SetHttpClient sets httpClient
func (o *Options) SetHttpClient(httpClient *http.Client) {
	o.httpClient = httpClient
}

// AuthProvider returns AuthProvider
func (o *Options) AuthProvider() AuthProvider {
	return o.authProvider
}

// SetAuthProvider sets AuthProvider
func (o *Options) SetAuthProvider(authProvider AuthProvider) {
	o.authProvider = authProvider
}

// PreserveResponseBody returns bool
func (o *Options) PreserveResponseBody() bool {
	return o.preserveResponseBody
}

// SetPreserveResponseBody sets bool
func (o *Options) SetPreserveResponseBody(preserveResponseBody bool) {
	o.preserveResponseBody = preserveResponseBody
}

// AllPreviews returns bool
func (o *Options) AllPreviews() bool {
	return o.allPreviews
}

// SetAllPreviews sets bool
func (o *Options) SetAllPreviews(allPreviews bool) {
	o.allPreviews = allPreviews
}

// RequiredPreviews returns bool
func (o *Options) RequiredPreviews() bool {
	return o.requiredPreviews
}

// SetRequiredPreviews sets bool
func (o *Options) SetRequiredPreviews(requiredPreviews bool) {
	o.requiredPreviews = requiredPreviews
}

// UserAgent returns userAgent
func (o *Options) UserAgent() string {
	return o.userAgent
}

// SetUserAgent sets userAgent
func (o *Options) SetUserAgent(userAgent string) {
	o.userAgent = userAgent
}

// BaseURL returns baseURL
func (o *Options) BaseURL() url.URL {
	return o.baseURL
}

// SetBaseURL sets baseURL
func (o *Options) SetBaseURL(baseURL url.URL) {
	o.baseURL = baseURL
}
