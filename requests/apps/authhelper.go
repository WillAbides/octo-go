package apps

import (
	"context"
	"fmt"
	"time"
)

// InstallationAuthHelper is a helper for octo.WithAppInstallationAuth
// use its GetInstallationToken function as octo.WithAppInstallationAuth's tokenGetter
type InstallationAuthHelper struct {
	installationID int64
	appAuthClient  Client
	requestBody    *CreateInstallationAccessTokenReqBody
}

// NewInstallationAuthHelper returns an InstallationAuthHelper
//  appAuthClient is a client that has octo.WithAppAuth set for your app
//  tokenReqBody is the body of the request to get the installation token. It can be nil if you don't want to restrict its scope.
func NewInstallationAuthHelper(installationID int64, appAuthClient Client, tokenReqBody *CreateInstallationAccessTokenReqBody) *InstallationAuthHelper {
	return &InstallationAuthHelper{
		requestBody:    tokenReqBody,
		installationID: installationID,
		appAuthClient:  appAuthClient,
	}
}

// GetInstallationToken implements octo.InstallationTokenGetter
func (a *InstallationAuthHelper) GetInstallationToken(ctx context.Context) (token string, expiry time.Time, err error) {
	req := &CreateInstallationAccessTokenReq{
		InstallationId: a.installationID,
	}
	if a.requestBody != nil {
		req.RequestBody = *a.requestBody
	}
	resp, err := a.appAuthClient.CreateInstallationAccessToken(ctx, req)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error getting installation token: %v", err)
	}
	expiry, err = time.Parse(time.RFC3339, resp.Data.ExpiresAt)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error parsing token expiry: %v", err)
	}
	return resp.Data.Token, expiry, nil
}
