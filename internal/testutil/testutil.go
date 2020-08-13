package testutil

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/requests"
)

const (
	// test app id
	AppID int64 = 67080
	// test installation id
	AppInstallationID int64 = 9437013
)

func init() {
	root, err := projectRoot()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		fmt.Println("no .env file. tests can still run")
	}
}

func appPrivateKey(t *testing.T) []byte {
	t.Helper()
	if os.Getenv("APP_PRIVATE_KEY") == "" {
		return nil
	}
	got, err := base64.StdEncoding.DecodeString(os.Getenv("APP_PRIVATE_KEY"))
	require.NoError(t, err)
	return got
}

func projectRoot() (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	if filepath.Base(root) != "octo-go" {
		return "", fmt.Errorf("%s isn't octo-go", root)
	}
	return filepath.Abs(root)
}

// ProjectRoot returns the absolute path to the repo root for octo-go
func ProjectRoot(t *testing.T) string {
	t.Helper()
	got, err := projectRoot()
	assert.NoError(t, err)
	return got
}

// PATAuth returns auth
func PATAuth() requests.Option {
	return octo.WithPATAuth(os.Getenv("GITHUB_TOKEN"))
}

// AppAuth returns auth
func AppAuth(t *testing.T) requests.Option {
	key := appPrivateKey(t)
	if key == nil {
		return nil
	}
	return octo.WithAppAuth(AppID, key)
}

// AppInstallationAuth returns auth
func AppInstallationAuth(authClient []requests.Option) requests.Option {
	return octo.WithAppInstallationAuth(AppInstallationID, authClient, nil)
}

// VCRClient returns a vcr client
func VCRClient(t *testing.T, cas string, opts ...requests.Option) octo.Client {
	t.Helper()
	cas = strings.ReplaceAll(cas, "/", "_")
	cas = filepath.Join(filepath.FromSlash("testdata/vcr/"), cas)
	r, err := recorder.NewAsMode(cas, recorder.ModeReplaying, http.DefaultTransport)
	require.NoError(t, err)
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})
	r.SetMatcher(func(req *http.Request, cr cassette.Request) bool {
		if cr.Body != "" {
			assert.NotNil(t, req.Body)
			reqBody, err := ioutil.ReadAll(req.Body)
			assert.NoError(t, err)
			assert.Equal(t, cr.Body, string(reqBody))
		}

		reqHeader := req.Header.Clone()
		reqHeader.Del("Authorization")
		assert.Equal(t, cr.Headers, reqHeader)

		return cassette.DefaultMatcher(req, cr)
	})
	t.Cleanup(func() {
		require.NoError(t, r.Stop())
	})

	return octo.NewClient(append(opts, octo.WithHTTPClient(&http.Client{
		Transport: r,
	}))...)
}
