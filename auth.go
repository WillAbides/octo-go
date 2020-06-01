package octo

import (
	"context"
	"crypto/rsa"
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthProvider sets the Authorization header authenticate you with the API
type AuthProvider interface {
	AuthorizationHeader(ctx context.Context) (string, error)
}

type patAuthProvider struct {
	Token string
}

// AuthorizationHeader implements AuthProvider
func (p *patAuthProvider) AuthorizationHeader(_ context.Context) (string, error) {
	return "token " + p.Token, nil
}

// appAuthProvider provides authentication for a GitHub App. See also appInstallationAuthProvider
type appAuthProvider struct {
	AppID      int64
	PrivateKey *rsa.PrivateKey
}

// AuthorizationHeader implements AuthProvider
func (a *appAuthProvider) AuthorizationHeader(_ context.Context) (string, error) {
	now := time.Now()
	claims := &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute).Unix(),
		Issuer:    fmt.Sprintf("%d", a.AppID),
	}
	signed, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(a.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("can't sign claims: %v", err)
	}
	return "bearer " + signed, nil
}

// appInstallationAuthProvider provides authentication for a GitHub App Installation
type appInstallationAuthProvider struct {
	AppID          int64
	InstallationID int64
	PrivateKey     *rsa.PrivateKey

	// request options for requests made to the create-installation-token endpoint
	RequestOptions []RequestOption

	tkn         string
	tknExpiry   time.Time
	tknMux      sync.Mutex
	tokenClient *Client
}

func (a *appInstallationAuthProvider) getTokenClient() (*Client, error) {
	if a.tokenClient != nil {
		return a.tokenClient, nil
	}
	opts := append(a.RequestOptions, RequestAuthProvider(&appAuthProvider{
		AppID:      a.AppID,
		PrivateKey: a.PrivateKey,
	}))
	var err error
	a.tokenClient, err = NewClient(opts...)
	if err != nil {
		return nil, err
	}
	return a.tokenClient, nil
}

// AuthorizationHeader implements AuthProvider
func (a *appInstallationAuthProvider) AuthorizationHeader(ctx context.Context) (string, error) {
	a.tknMux.Lock()
	defer a.tknMux.Unlock()
	if time.Now().Before(a.tknExpiry.Add(-1 * time.Minute)) {
		return a.tkn, nil
	}
	tokenClient, err := a.getTokenClient()
	if err != nil {
		return "", err
	}
	resp, err := tokenClient.AppsCreateInstallationToken(ctx, &AppsCreateInstallationTokenReq{
		InstallationId:    a.InstallationID,
		MachineManPreview: true,
	})
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
