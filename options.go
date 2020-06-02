package octo

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
)

// RequestOption is an option for building an http request
type RequestOption func(opts *requestOpts) error

func resetOptions(newOpts requestOpts) RequestOption {
	return func(opts *requestOpts) error {
		*opts = newOpts
		return nil
	}
}

// RequestBaseURL set the baseURL to use. Default is https://api.github.com
func RequestBaseURL(baseURL url.URL) RequestOption {
	return func(opts *requestOpts) error {
		opts.baseURL = baseURL
		return nil
	}
}

// RequestOptions is a convenience function for when you want to send the same set of options to multiple requests
func RequestOptions(option ...RequestOption) RequestOption {
	return func(opts *requestOpts) error {
		for _, requestOption := range option {
			err := requestOption(opts)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

// RequestEnableRequirePreviews enables any previews that are required for your request
func RequestEnableRequirePreviews() RequestOption {
	return func(opts *requestOpts) error {
		opts.requiredPreviews = true
		return nil
	}
}

// RequestEnableAllPreviews enables all previews that are available for your request
func RequestEnableAllPreviews() RequestOption {
	return func(opts *requestOpts) error {
		opts.allPreviews = true
		return nil
	}
}

// RequestPreserveResponseBody rewrite the body back to the http response for later inspection
func RequestPreserveResponseBody() RequestOption {
	return func(opts *requestOpts) error {
		opts.preserveResponseBody = true
		return nil
	}
}

// RequestHTTPClient sets an http client to use for requests. If unset, http.DefaultClient is used
func RequestHTTPClient(client *http.Client) RequestOption {
	return func(opts *requestOpts) error {
		opts.httpClient = client
		return nil
	}
}

// RequestAuthProvider sets a provider to use in setting the Authentication header
//
// This is for custom providers. You will typically want to use RequestPATAuth, RequestAppAuth or RequestAppInstallationAuth
// instead.
func RequestAuthProvider(authProvider AuthProvider) RequestOption {
	return func(opts *requestOpts) error {
		opts.authProvider = authProvider
		return nil
	}
}

// RequestPATAuth authenticates requests with a Personal Access Token
func RequestPATAuth(token string) RequestOption {
	return func(opts *requestOpts) error {
		opts.authProvider = &patAuthProvider{
			token: token,
		}
		return nil
	}
}

// RequestAppAuth provides authentication for a GitHub App. See also RequestAppInstallationAuth
//
// appID is the GitHub App's id
// privateKey is the app's private key. It should be the content of a PEM file
func RequestAppAuth(appID int64, privateKey []byte) RequestOption {
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

// RequestAppInstallationAuth provides authentication for a GitHub App installation
//
// appID is the GitHub App's id
// privateKey is the app's private key. It should be the content of a PEM file
// requestBody is the body to be sent when creating an installation token. It can be nil, or you can set it to limit the
//  scope of the token's authorizations.
// requestOptions are options to be use when requesting a token. They do not affect options for the main request.
func RequestAppInstallationAuth(appID, installationID int64, privateKey []byte, requestBody *AppsCreateInstallationTokenReqBody, opt ...RequestOption) RequestOption {
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
