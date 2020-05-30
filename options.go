package octo

import (
	"net/url"
)

//RequestOption is an option for building an http request
type RequestOption func(opts *requestOpts)

func resetOptions(newOpts requestOpts) RequestOption {
	return func(opts *requestOpts) {
		*opts = newOpts
	}
}

//RequestBaseURL set the baseURL to use. Default is https://api.github.com
func RequestBaseURL(baseURL url.URL) RequestOption {
	return func(opts *requestOpts) {
		opts.baseURL = baseURL
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

//RequestEnableRequirePreviews enables any previews that are required for your request
func RequestEnableRequirePreviews() RequestOption {
	return func(opts *requestOpts) {
		opts.requiredPreviews = true
	}
}

//RequestEnableAllPreviews enables all previews that are available for your request
func RequestEnableAllPreviews() RequestOption {
	return func(opts *requestOpts) {
		opts.allPreviews = true
	}
}

//RequestPreserveResponseBody rewrite the body back to the http response for later inspection
func RequestPreserveResponseBody() RequestOption {
	return func(opts *requestOpts) {
		opts.preserveResponseBody = true
	}
}

type requestOpts struct {
	baseURL              url.URL
	userAgent            string
	requiredPreviews     bool
	allPreviews          bool
	preserveResponseBody bool
}

var defaultRequestOpts = requestOpts{
	baseURL: url.URL{
		Host:   "api.github.com",
		Scheme: "https",
	},
	userAgent: "octo-go",
}

func buildRequestOptions(opts []RequestOption) requestOpts {
	result := defaultRequestOpts
	for _, opt := range opts {
		opt(&result)
	}
	return result
}
