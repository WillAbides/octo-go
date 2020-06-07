package octo

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
)

// RequestOption is an option for building an http request
type RequestOption func(opts *requestOpts) error

// WithBaseURL set the baseURL to use. Default is https://api.github.com
func WithBaseURL(baseURL url.URL) RequestOption {
	return func(opts *requestOpts) error {
		opts.baseURL = baseURL
		return nil
	}
}

// WithRequiredPreviews enables any previews that are required for your request
func WithRequiredPreviews() RequestOption {
	return func(opts *requestOpts) error {
		opts.requiredPreviews = true
		return nil
	}
}

// WithAllPreviews enables all previews that are available for your request
func WithAllPreviews() RequestOption {
	return func(opts *requestOpts) error {
		opts.allPreviews = true
		return nil
	}
}

// PreserveResponseBody rewrite the body back to the http response for later inspection
func PreserveResponseBody() RequestOption {
	return func(opts *requestOpts) error {
		opts.preserveResponseBody = true
		return nil
	}
}

// WithHTTPClient sets an http client to use for requests. If unset, http.DefaultClient is used
func WithHTTPClient(client *http.Client) RequestOption {
	return func(opts *requestOpts) error {
		opts.httpClient = client
		return nil
	}
}

// WithUserAgent sets the User-Agent header in requests
func WithUserAgent(userAgent string) RequestOption {
	return func(opts *requestOpts) error {
		opts.userAgent = userAgent
		return nil
	}
}

// WithAuthProvider sets a provider to use in setting the Authentication header
//
// This is for custom providers. You will typically want to use WithPATAuth, WithAppAuth or WithAppInstallationAuth
// instead.
func WithAuthProvider(authProvider AuthProvider) RequestOption {
	return func(opts *requestOpts) error {
		opts.authProvider = authProvider
		return nil
	}
}

// WithPATAuth authenticates requests with a Personal Access Token
func WithPATAuth(token string) RequestOption {
	return func(opts *requestOpts) error {
		opts.authProvider = &patAuthProvider{
			token: token,
		}
		return nil
	}
}

// WithAppAuth provides authentication for a GitHub App. See also WithAppInstallationAuth
//
// appID is the GitHub App's id
// privateKey is the app's private key. It should be the content of a PEM file
func WithAppAuth(appID int64, privateKey []byte) RequestOption {
	return func(opts *requestOpts) error {
		pk, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
		if err != nil {
			return fmt.Errorf("error parsing private key: %v", pk)
		}
		opts.authProvider = &appAuthProvider{
			appID:      appID,
			privateKey: pk,
		}
		return nil
	}
}

// WithAppInstallationAuth provides authentication for a GitHub App installation
//
// appID is the GitHub App's id
// privateKey is the app's private key. It should be the content of a PEM file
// requestBody is the body to be sent when creating an installation token. It can be nil, or you can set it to limit the
//  scope of the token's authorizations.
// requestOptions are options to be use when requesting a token. They do not affect options for the main request.
func WithAppInstallationAuth(appID, installationID int64, privateKey []byte, requestBody *AppsCreateInstallationTokenReqBody, opt ...RequestOption) RequestOption {
	return func(opts *requestOpts) error {
		pk, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
		if err != nil {
			return fmt.Errorf("error parsing private key: %v", pk)
		}
		opts.authProvider = &appInstallationAuthProvider{
			appID:          appID,
			installationID: installationID,
			privateKey:     pk,
			requestBody:    requestBody,
			requestOptions: opt,
		}
		return nil
	}
}

type requestOpts struct {
	baseURL              url.URL
	userAgent            string
	requiredPreviews     bool
	allPreviews          bool
	preserveResponseBody bool
	authProvider         AuthProvider
	httpClient           *http.Client
}

var defaultRequestOpts = requestOpts{
	baseURL: url.URL{
		Host:   "api.github.com",
		Scheme: "https",
	},
	userAgent:  "octo-go",
	httpClient: http.DefaultClient,
}

func buildRequestOptions(opts []RequestOption) (requestOpts, error) {
	result := defaultRequestOpts
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		err := opt(&result)
		if err != nil {
			return requestOpts{}, err
		}
	}
	return result, nil
}
