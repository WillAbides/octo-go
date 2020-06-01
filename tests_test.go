package octo_test

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
)

const (
	appID             int64 = 67080
	appInstallationID int64 = 9385805
)

func appPrivateKey(t *testing.T) []byte {
	t.Helper()
	if os.Getenv("APP_PRIVATE_KEY") == "" {
		return nil
	}
	got, err := base64.StdEncoding.DecodeString(os.Getenv("APP_PRIVATE_KEY"))
	require.NoError(t, err)
	return got
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("no .env file. tests can still run")
	}
}

func patAuth() octo.RequestOption {
	return octo.RequestPATAuth(os.Getenv("GITHUB_TOKEN"))
}

func appAuth(t *testing.T) octo.RequestOption {
	key := appPrivateKey(t)
	if key == nil {
		return nil
	}
	return octo.RequestAppAuth(appID, key)
}

func appInstallationAuth(t *testing.T) octo.RequestOption {
	key := appPrivateKey(t)
	if key == nil {
		return nil
	}
	return octo.RequestAppInstallationAuth(appID, appInstallationID, key)
}

func vcrClient(t *testing.T, cas string, opts ...octo.RequestOption) *octo.Client {
	t.Helper()
	cas = strings.ReplaceAll(cas, "/", "_")
	cas = filepath.Join(filepath.FromSlash("testdata/vcr/"), cas)
	r, err := recorder.NewAsMode(cas, recorder.ModeReplaying, http.DefaultTransport)
	require.NoError(t, err)
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})
	t.Cleanup(func() {
		require.NoError(t, r.Stop())
	})

	cl, err := octo.NewClient(append(opts, octo.RequestHTTPClient(&http.Client{
		Transport: r,
	}))...)
	require.NoError(t, err)
	return cl
}
