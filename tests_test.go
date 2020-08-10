package octo_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

var (
	schemaBytes     []byte
	schemaBytesOnce sync.Once
)

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
