package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/options"
	"github.com/willabides/octo-go/requests/apps"
)

// NewPATProvider returns a PATProvider
func NewPATProvider(token string) *PATProvider {
	return &PATProvider{
		token: token,
	}
}

// WithPATAuth authenticates requests with a Personal Access Token
func WithPATAuth(token string) options.Option {
	return octo.WithAuthProvider(NewPATProvider(token))
}

// PATProvider is a Personal Access Token authorization provider
type PATProvider struct {
	token string
}

// AuthorizationHeader implements AuthProvider
func (p *PATProvider) AuthorizationHeader(_ context.Context) (string, error) {
	return "token " + p.token, nil
}

// NewAppProvider returns a new AppProvider
func NewAppProvider(appID int64, privateKey []byte) (*AppProvider, error) {
	pk, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key")
	}
	return &AppProvider{
		appID:      appID,
		privateKey: pk,
	}, nil
}

// WithAppAuth provides authentication for a GitHub App. See also WithAppInstallationAuth
//
// appID is the GitHub App's id
// privateKey is the app's private key. It should be the content of a PEM file
func WithAppAuth(appID int64, privateKey []byte) options.Option {
	provider, err := NewAppProvider(appID, privateKey)
	if err != nil {
		return &errOption{err: err}
	}
	return octo.WithAuthProvider(provider)
}

// AppProvider provides authentication for a GitHub App
type AppProvider struct {
	appID      int64
	privateKey *rsa.PrivateKey
}

// AuthorizationHeader implements AuthProvider
func (a *AppProvider) AuthorizationHeader(_ context.Context) (string, error) {
	now := time.Now()
	claims := &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute).Unix(),
		Issuer:    fmt.Sprintf("%d", a.appID),
	}
	signed, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("can't sign claims: %v", err)
	}
	return "bearer " + signed, nil
}

// NewAppInstallationProvider returns an AppInstallationProvider
func NewAppInstallationProvider(appID, installationID int64, privateKey []byte,
	requestBody *apps.CreateInstallationAccessTokenReqBody,
	opt ...options.Option) (*AppInstallationProvider, error) {
	pk, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}
	return &AppInstallationProvider{
		appID:          appID,
		installationID: installationID,
		privateKey:     pk,
		requestBody:    requestBody,
		requestOptions: opt,
	}, nil
}

// WithAppInstallationAuth provides authentication for a GitHub App installation
//
// appID is the GitHub App's id
// privateKey is the app's private key. It should be the content of a PEM file
// requestBody is the body to be sent when creating an installation token. It can be nil, or you can set it to limit the
//  scope of the token's authorizations.
// requestOptions are options to be use when requesting a token. They do not affect options for the main request.
func WithAppInstallationAuth(appID, installationID int64, privateKey []byte,
	requestBody *apps.CreateInstallationAccessTokenReqBody,
	opt ...options.Option) options.Option {
	provider, err := NewAppInstallationProvider(appID, installationID, privateKey, requestBody, opt...)
	if err != nil {
		return &errOption{err: err}
	}
	return octo.WithAuthProvider(provider)
}

// AppInstallationProvider provides authentication for a GitHub App installation
type AppInstallationProvider struct {
	appID          int64
	installationID int64
	privateKey     *rsa.PrivateKey
	requestOptions []options.Option
	requestBody    *apps.CreateInstallationAccessTokenReqBody
	tkn            string
	tknExpiry      time.Time
	tknMux         sync.Mutex
	tokenClient    []options.Option
}

func (a *AppInstallationProvider) getTokenClient() []options.Option {
	if a.tokenClient != nil {
		return a.tokenClient
	}
	opts := append(a.requestOptions, octo.WithAuthProvider(&AppProvider{
		appID:      a.appID,
		privateKey: a.privateKey,
	}))
	opts = append(opts, octo.WithRequiredPreviews())
	a.tokenClient = opts
	return a.tokenClient
}

// AuthorizationHeader implements AuthProvider
func (a *AppInstallationProvider) AuthorizationHeader(ctx context.Context) (string, error) {
	a.tknMux.Lock()
	defer a.tknMux.Unlock()
	if time.Now().Before(a.tknExpiry.Add(-1 * time.Minute)) {
		return a.tkn, nil
	}
	tokenClient := a.getTokenClient()
	req := &apps.CreateInstallationAccessTokenReq{
		InstallationId: a.installationID,
	}
	if a.requestBody != nil {
		req.RequestBody = *a.requestBody
	}
	resp, err := apps.CreateInstallationAccessToken(ctx, req, tokenClient...)
	if err != nil {
		return "", fmt.Errorf("error getting installation token: %v", err)
	}
	expiry, err := time.Parse(time.RFC3339, resp.Data.ExpiresAt)
	if err != nil {
		return "", fmt.Errorf("error parsing token expiry: %v", err)
	}
	a.tkn = resp.Data.Token
	a.tknExpiry = expiry
	return "bearer " + a.tkn, nil
}

type errOption struct {
	err error
}

func (e *errOption) Apply(_ *options.Options) error {
	return e.err
}
