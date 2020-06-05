package octo_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"github.com/willabides/octo-go"
)

const (
	appID             int64 = 67080
	appInstallationID int64 = 9437013
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
	return octo.RequestAppInstallationAuth(appID, appInstallationID, key, nil)
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

	return octo.NewClient(append(opts, octo.RequestHTTPClient(&http.Client{
		Transport: r,
	}))...)
}

var schemaBytes []byte
var schemaBytesOnce sync.Once

func schemaGJSON(t *testing.T, path string) gjson.Result {
	schemaBytesOnce.Do(func() {
		var err error
		schemaBytes, err = ioutil.ReadFile("api.github.com.json")
		require.NoError(t, err)
	})
	return gjson.GetBytes(schemaBytes, path)
}

func responseExamples(t *testing.T, endpointPath, httpMethod string, statusCode int) map[string][]byte {
	examples := mappedResponseExamples(t, endpointPath, httpMethod, statusCode)
	single := singleResponseExample(t, endpointPath, httpMethod, statusCode)
	if single != nil {
		examples["_singleton_example"] = single
	}
	return examples
}

func singleResponseExample(t *testing.T, endpointPath, httpMethod string, statusCode int) []byte {
	t.Helper()
	endpointPath = strings.ReplaceAll(endpointPath, ".", `\.`)
	path := fmt.Sprintf("paths.%s.%s.responses.%d.content.application/json.example",
		endpointPath,
		strings.ToLower(httpMethod),
		statusCode,
	)
	ex := schemaGJSON(t, path)
	if ex.Exists() {
		return []byte(ex.String())
	}
	return nil
}

func mappedResponseExamples(t *testing.T, endpointPath, httpMethod string, statusCode int) map[string][]byte {
	t.Helper()
	endpointPath = strings.ReplaceAll(endpointPath, ".", `\.`)
	path := fmt.Sprintf("paths.%s.%s.responses.%d.content.application/json.examples",
		endpointPath,
		strings.ToLower(httpMethod),
		statusCode,
	)
	result := map[string][]byte{}
	for name, mapItem := range schemaGJSON(t, path).Map() {
		ref := mapItem.Get("$ref")
		if !ref.Exists() {
			continue
		}
		if !strings.HasPrefix(ref.String(), "#/components/examples/") {
			continue
		}
		jsonFile := strings.TrimPrefix(ref.String(), "#/") + ".json"
		b, err := ioutil.ReadFile(jsonFile)
		require.NoError(t, err)
		result[name] = b
	}

	return result
}
