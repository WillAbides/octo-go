package octo

import (
	"net/http"
	"net/url"

	"github.com/willabides/octo-go/options"
)

// WithBaseURL set the baseURL to use. Default is https://api.github.com
func WithBaseURL(baseURL url.URL) options.Option {
	return optionFunc(func(opts *options.Options) error {
		opts.SetBaseURL(baseURL)
		return nil
	})
}

// WithRequiredPreviews enables any previews that are required for your request
func WithRequiredPreviews() options.Option {
	return optionFunc(func(opts *options.Options) error {
		opts.SetRequiredPreviews(true)
		return nil
	})
}

// WithAllPreviews enables all previews that are available for your request
func WithAllPreviews() options.Option {
	return optionFunc(func(opts *options.Options) error {
		opts.SetAllPreviews(true)
		return nil
	})
}

// PreserveResponseBody rewrite the body back to the http common for later inspection
func PreserveResponseBody() options.Option {
	return optionFunc(func(opts *options.Options) error {
		opts.SetPreserveResponseBody(true)
		return nil
	})
}

// WithHTTPClient sets an http client to use for requests. If unset, http.DefaultClient is used
func WithHTTPClient(client *http.Client) options.Option {
	return optionFunc(func(opts *options.Options) error {
		opts.SetHttpClient(client)
		return nil
	})
}

// WithUserAgent sets the User-Agent header in requests
func WithUserAgent(userAgent string) options.Option {
	return optionFunc(func(opts *options.Options) error {
		opts.SetUserAgent(userAgent)
		return nil
	})
}

// WithAuthProvider sets a provider to use in setting the Authentication header
//
// This is for custom providers. You will typically want to use auth.WithPATAuth, auth.WithAppAuth or auth.WithAppInstallationAuth
// instead.
func WithAuthProvider(authProvider options.AuthProvider) options.Option {
	return optionFunc(func(opts *options.Options) error {
		opts.SetAuthProvider(authProvider)
		return nil
	})
}

type optionFunc func(opts *options.Options) error

func (fn optionFunc) Apply(opts *options.Options) error {
	return fn(opts)
}
