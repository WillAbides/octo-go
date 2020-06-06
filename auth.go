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
	token string
}

// AuthorizationHeader implements AuthProvider
func (p *patAuthProvider) AuthorizationHeader(_ context.Context) (string, error) {
	return "token " + p.token, nil
}

// appAuthProvider provides authentication for a GitHub App. See also appInstallationAuthProvider
type appAuthProvider struct {
	appID      int64
	privateKey *rsa.PrivateKey
}

// AuthorizationHeader implements AuthProvider
func (a *appAuthProvider) AuthorizationHeader(_ context.Context) (string, error) {
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

// appInstallationAuthProvider provides authentication for a GitHub App Installation
type appInstallationAuthProvider struct {
	appID          int64
	installationID int64
	privateKey     *rsa.PrivateKey
	requestOptions []RequestOption
	requestBody    *AppsCreateInstallationTokenReqBody
	tkn            string
	tknExpiry      time.Time
	tknMux         sync.Mutex
	tokenClient    Client
}

func (a *appInstallationAuthProvider) getTokenClient() Client {
	if a.tokenClient != nil {
		return a.tokenClient
	}
	opts := append(a.requestOptions, RequestAuthProvider(&appAuthProvider{
		appID:      a.appID,
		privateKey: a.privateKey,
	}))
	a.tokenClient = NewClient(append(opts, RequestEnableRequirePreviews())...)
	return a.tokenClient
}

// AuthorizationHeader implements AuthProvider
func (a *appInstallationAuthProvider) AuthorizationHeader(ctx context.Context) (string, error) {
	a.tknMux.Lock()
	defer a.tknMux.Unlock()
	if time.Now().Before(a.tknExpiry.Add(-1 * time.Minute)) {
		return a.tkn, nil
	}
	tokenClient := a.getTokenClient()
	req := &AppsCreateInstallationTokenReq{
		InstallationId: a.installationID,
	}
	if a.requestBody != nil {
		req.RequestBody = *a.requestBody
	}
	resp, err := tokenClient.AppsCreateInstallationToken(ctx, req)
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
