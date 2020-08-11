package apps_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/internal/testutil"
	"github.com/willabides/octo-go/requests"
	"github.com/willabides/octo-go/requests/apps"
)

func vcrClient(t *testing.T, cas string, opts ...requests.Option) apps.Client {
	return apps.NewClient(testutil.VCRClient(t, cas, opts...))
}

func TestCreateInstallationAccessToken(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), testutil.AppAuth(t), octo.WithRequiredPreviews())
	token, err := client.CreateInstallationAccessToken(ctx, &apps.CreateInstallationAccessTokenReq{
		InstallationId: testutil.AppInstallationID,
	})
	require.NoError(t, err)
	_, err = client.RevokeInstallationAccessToken(ctx, nil, octo.WithPATAuth(token.Data.Token))
	require.NoError(t, err)
}

func TestGetRepoInstallation(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), testutil.AppAuth(t))
		installation, err := client.GetRepoInstallation(ctx, &apps.GetRepoInstallationReq{
			Owner:             "WillAbides",
			Repo:              "octo-go",
			MachineManPreview: true,
		})
		require.NoError(t, err)
		require.Equal(t, testutil.AppInstallationID, installation.Data.Id)
	})

	t.Run("no installation", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), testutil.AppAuth(t))
		installation, err := client.GetRepoInstallation(ctx, &apps.GetRepoInstallationReq{
			Owner:             "torvalds",
			Repo:              "linux",
			MachineManPreview: true,
		})
		require.Error(t, err)
		require.Empty(t, installation.Data)
	})
}

func TestGetUserInstallation(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), testutil.AppAuth(t))
		installation, err := client.GetUserInstallation(ctx, &apps.GetUserInstallationReq{
			Username:          "WillAbides",
			MachineManPreview: true,
		})
		require.NoError(t, err)
		require.Equal(t, testutil.AppInstallationID, installation.Data.Id)
	})
}
