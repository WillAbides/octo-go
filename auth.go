package octo

import (
	"context"
	"crypto/rsa"
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/willabides/octo-go/requests/apps"
)

type patAuthProvider struct {
	token string
}

// AuthorizationHeader implements AuthProvider
func (p *patAuthProvider) AuthorizationHeader(_ context.Context) (string, error) {
	return "token " + p.token, nil
}

type appAuthProvider struct {
	mux        sync.Mutex
	appID      int64
	privateKey []byte
	pk         *rsa.PrivateKey
}

// AuthorizationHeader implements AuthProvider
func (p *appAuthProvider) AuthorizationHeader(_ context.Context) (string, error) {
	p.mux.Lock()
	defer p.mux.Unlock()
	if p.pk == nil {
		pk, err := jwt.ParseRSAPrivateKeyFromPEM(p.privateKey)
		if err != nil {
			return "", fmt.Errorf("couldn't parse private key")
		}
		p.pk = pk
	}
	now := time.Now()
	claims := &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute).Unix(),
		Issuer:    fmt.Sprintf("%d", p.appID),
	}
	signed, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(p.pk)
	if err != nil {
		return "", fmt.Errorf("can't sign claims: %v", err)
	}
	return "bearer " + signed, nil
}

type appInstallationAuthProvider struct {
	mux            sync.Mutex
	installationID int64
	requestBody    *apps.CreateInstallationAccessTokenReqBody
	client         Client
	tkn            string
	expiry         time.Time
}

// AuthorizationHeader implements AuthProvider
func (p *appInstallationAuthProvider) AuthorizationHeader(ctx context.Context) (string, error) {
	p.mux.Lock()
	defer p.mux.Unlock()
	if time.Now().Before(p.expiry.Add(-1 * time.Minute)) {
		return p.tkn, nil
	}
	req := &apps.CreateInstallationAccessTokenReq{
		InstallationId: p.installationID,
	}
	if p.requestBody != nil {
		req.RequestBody = *p.requestBody
	}
	resp, err := p.client.Apps().CreateInstallationAccessToken(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error creating installation token: %v", err)
	}
	expiry, err := time.Parse(time.RFC3339, resp.Data.ExpiresAt)
	if err != nil {
		return "", fmt.Errorf("error parsing token expiry: %v", err)
	}
	p.expiry = expiry
	p.tkn = resp.Data.Token
	return "bearer " + p.tkn, nil
}
