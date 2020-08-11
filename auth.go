package octo

import (
	"context"
	"crypto/rsa"
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type patProvider struct {
	token string
}

// AuthorizationHeader implements AuthProvider
func (p *patProvider) AuthorizationHeader(_ context.Context) (string, error) {
	return "token " + p.token, nil
}

type appProvider struct {
	appID      int64
	privateKey *rsa.PrivateKey
}

// AuthorizationHeader implements AuthProvider
func (a *appProvider) AuthorizationHeader(_ context.Context) (string, error) {
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

type appInstallationProvider struct {
	tokenGetter GetInstallationToken
	tkn         string
	expiry      time.Time
	mux         sync.Mutex
}

// AuthorizationHeader implements AuthProvider
func (a *appInstallationProvider) AuthorizationHeader(ctx context.Context) (string, error) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if time.Now().Before(a.expiry.Add(-1 * time.Minute)) {
		return a.tkn, nil
	}
	tkn, expiry, err := a.tokenGetter(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting installation token: %v", err)
	}
	a.tkn = tkn
	a.expiry = expiry
	return "bearer " + a.tkn, nil
}
