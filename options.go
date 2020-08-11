package octo

import (
	"net/http"
	"net/url"

	"github.com/willabides/octo-go/requests"
	"github.com/willabides/octo-go/requests/apps"
)

// WithBaseURL set the baseURL to use. Default is https://api.github.com
func WithBaseURL(baseURL url.URL) requests.Option {
	return func(opts *requests.Options) error {
		opts.SetBaseURL(baseURL)
		return nil
	}
}

// WithRequiredPreviews enables any previews that are required for your request
func WithRequiredPreviews() requests.Option {
	return func(opts *requests.Options) error {
		opts.SetRequiredPreviews(true)
		return nil
	}
}

// WithAllPreviews enables all previews that are available for your request
func WithAllPreviews() requests.Option {
	return func(opts *requests.Options) error {
		opts.SetAllPreviews(true)
		return nil
	}
}

// PreserveResponseBody rewrite the body back to the http common for later inspection
func PreserveResponseBody() requests.Option {
	return func(opts *requests.Options) error {
		opts.SetPreserveResponseBody(true)
		return nil
	}
}

// WithHTTPClient sets an http client to use for requests. If unset, http.DefaultClient is used
func WithHTTPClient(client *http.Client) requests.Option {
	return func(opts *requests.Options) error {
		opts.SetHttpClient(client)
		return nil
	}
}

// WithUserAgent sets the User-Agent header in requests
func WithUserAgent(userAgent string) requests.Option {
	return func(opts *requests.Options) error {
		opts.SetUserAgent(userAgent)
		return nil
	}
}

// WithPATAuth authenticates requests with a Personal Access Token
func WithPATAuth(token string) requests.Option {
	return WithAuthProvider(&patAuthProvider{
		token: token,
	})
}

// WithAppAuth provides authentication for a GitHub App. See also WithAppInstallationAuth
//
// appID is the GitHub App's id
// privateKey is the app's private key. It should be the content of a PEM file
func WithAppAuth(appID int64, privateKey []byte) requests.Option {
	return WithAuthProvider(&appAuthProvider{
		appID:      appID,
		privateKey: privateKey,
	})
}

// WithAppInstallationAuth provides authentication for a GitHub App installation
//  client is the client to use when fetching the installation token. It should use WithAppAuth.
//  requestBody is used to restrict access to the installation token. Leave it nil if you don't want to restrict access.
func WithAppInstallationAuth(installationID int64, client Client, requestBody *apps.CreateInstallationAccessTokenReqBody) requests.Option {
	return WithAuthProvider(&appInstallationAuthProvider{
		installationID: installationID,
		requestBody:    requestBody,
		opts:           client,
	})
}

// WithAuthProvider sets a provider to use in setting the Authentication header
//
// This is for custom providers. You will typically want to use WithPATAuth, WithAppAuth or WithAppInstallationAuth
// instead.
func WithAuthProvider(authProvider requests.AuthProvider) requests.Option {
	return func(opts *requests.Options) error {
		opts.SetAuthProvider(authProvider)
		return nil
	}
}
