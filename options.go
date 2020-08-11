package octo

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/willabides/octo-go/requests"
	"github.com/willabides/octo-go/requests/apps"
)

// WithBaseURL set the baseURL to use. Default is https://api.github.com
func WithBaseURL(baseURL url.URL) requests.Option {
	return optionFunc(func(opts *requests.Options) error {
		opts.SetBaseURL(baseURL)
		return nil
	})
}

// WithRequiredPreviews enables any previews that are required for your request
func WithRequiredPreviews() requests.Option {
	return optionFunc(func(opts *requests.Options) error {
		opts.SetRequiredPreviews(true)
		return nil
	})
}

// WithAllPreviews enables all previews that are available for your request
func WithAllPreviews() requests.Option {
	return optionFunc(func(opts *requests.Options) error {
		opts.SetAllPreviews(true)
		return nil
	})
}

// PreserveResponseBody rewrite the body back to the http common for later inspection
func PreserveResponseBody() requests.Option {
	return optionFunc(func(opts *requests.Options) error {
		opts.SetPreserveResponseBody(true)
		return nil
	})
}

// WithHTTPClient sets an http client to use for requests. If unset, http.DefaultClient is used
func WithHTTPClient(client *http.Client) requests.Option {
	return optionFunc(func(opts *requests.Options) error {
		opts.SetHttpClient(client)
		return nil
	})
}

// WithUserAgent sets the User-Agent header in requests
func WithUserAgent(userAgent string) requests.Option {
	return optionFunc(func(opts *requests.Options) error {
		opts.SetUserAgent(userAgent)
		return nil
	})
}

// WithPATAuth authenticates requests with a Personal Access Token
func WithPATAuth(token string) requests.Option {
	return WithAuthProvider(&PATAuthProvider{
		token: token,
	})
}

// WithAppAuth provides authentication for a GitHub App. See also WithAppInstallationAuth
//
// appID is the GitHub App's id
// privateKey is the app's private key. It should be the content of a PEM file
func WithAppAuth(appID int64, privateKey []byte) *AppAuthProvider {
	return &AppAuthProvider{
		appID:      appID,
		privateKey: privateKey,
	}
}

// GetInstallationToken is a function that provides an app installation token.
//  See apps.InstallationAuthHelper for an implementation.
type GetInstallationToken func(ctx context.Context) (token string, expiry time.Time, err error)

// WithAppInstallationAuth provides authentication for a GitHub App installation
//  appAuthProvider is the auth provider used to create the installation token.
//  requestBody is used to restrict access to the installation token. Leave it nil if you don't want to restrict access.
//  opt is additional request options for the installation token request.
func WithAppInstallationAuth(installationID int64, appAuthProvider *AppAuthProvider,
	requestBody *apps.CreateInstallationAccessTokenReqBody, opt ...requests.Option) *AppInstallationAuthProvider {
	return &AppInstallationAuthProvider{
		appAuthProvider: appAuthProvider,
		installationID:  installationID,
		requestBody:     requestBody,
		opts:            opt,
	}
}

// WithAuthProvider sets a provider to use in setting the Authentication header
//
// This is for custom providers. You will typically want to use WithPATAuth, WithAppAuth or WithAppInstallationAuth
// instead.
func WithAuthProvider(authProvider requests.AuthProvider) requests.Option {
	return optionFunc(func(opts *requests.Options) error {
		opts.SetAuthProvider(authProvider)
		return nil
	})
}

type optionFunc func(opts *requests.Options) error

func (fn optionFunc) Apply(opts *requests.Options) error {
	return fn(opts)
}
