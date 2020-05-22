package octo

import (
	"net/url"
)

type RequestOption func(opts *requestOpts)

func RequestHTTPScheme(scheme string) RequestOption {
	return func(opts *requestOpts) {
		opts.BaseURL.Scheme = scheme
	}
}

func RequestBaseURL(baseURL url.URL) RequestOption {
	return func(opts *requestOpts) {
		opts.BaseURL = baseURL
	}
}

//RequestOptions is a convenience function for when you want to send the same set of options to multiple requests
func RequestOptions(option ...RequestOption) RequestOption {
	return func(opts *requestOpts) {
		for _, requestOption := range option {
			requestOption(opts)
		}
	}
}

type requestOpts struct {
	BaseURL   url.URL
	UserAgent string
}

var defaultRequestOpts = requestOpts{
	BaseURL: url.URL{
		Host:   "api.github.com",
		Scheme: "https",
	},
	UserAgent: "octo-go",
}

func buildRequestOptions(opts []RequestOption) requestOpts {
	result := defaultRequestOpts
	for _, opt := range opts {
		opt(&result)
	}
	return result
}
