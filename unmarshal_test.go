package octo_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

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

func responseExample(t *testing.T, endpointPath, httpMethod string, statusCode int) []byte {
	t.Helper()
	endpointPath = strings.ReplaceAll(endpointPath, ".", `\.`)
	path := fmt.Sprintf("paths.%s.%s.responses.%d.content.application/json.example",
		endpointPath,
		strings.ToLower(httpMethod),
		statusCode,
	)
	ex := schemaGJSON(t, path)
	require.True(t, ex.Exists(), "example doesn't exist in schema")
	return []byte(ex.String())
}

type unmarshalResponseBodyTest struct {
	name           string
	endpointPath   string
	httpMethod     string
	httpStatusCode int
	decode         func(decoder *json.Decoder) error
}

var unmarshalResponseBodyTests []unmarshalResponseBodyTest

func TestUnmarshalResponseBody(t *testing.T) {
	for _, test := range unmarshalResponseBodyTests {
		t.Run(fmt.Sprintf("%s_%d", test.name, test.httpStatusCode), func(t *testing.T) {
			ex := responseExample(t, test.endpointPath, test.httpMethod, test.httpStatusCode)
			decoder := json.NewDecoder(bytes.NewReader(ex))
			decoder.DisallowUnknownFields()
			require.NoError(t, test.decode(decoder))
		})
	}
}
