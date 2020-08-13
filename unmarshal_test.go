package octo_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v2"
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

type unmarshalResponseBodyTest struct {
	name           string
	endpointPath   string
	httpMethod     string
	httpStatusCode int
	operationID    string
	decode         func(decoder *json.Decoder) error
}

var unmarshalResponseBodyTests []unmarshalResponseBodyTest

func TestUnmarshalResponseBody(t *testing.T) {
	skippers := map[string]bool{
		"[]components.StarredRepository": true,
		"[]components.Stargazer":         true,
	}
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	type OutputTp struct {
		OperationID string
		Error       string
	}
	var output []*OutputTp
	for _, test := range unmarshalResponseBodyTests {
		t.Run(fmt.Sprintf("%s_%d", test.name, test.httpStatusCode), func(t *testing.T) {
			if skippers[test.name] {
				t.Skip("known bad example")
			}
			examples := responseExamples(t, test.endpointPath, test.httpMethod, test.httpStatusCode)
			if len(examples) == 0 {
				t.Skip("no examples found")
			}
			require.Greater(t, len(examples), 0, "no examples found")
			for name, ex := range examples {
				t.Run(name, func(t *testing.T) {
					decoder := json.NewDecoder(bytes.NewReader(ex))
					decoder.DisallowUnknownFields()
					fmt.Println(string(ex))
					err := test.decode(decoder)
					if !assert.NoError(t, err) {
						decoder2 := json.NewDecoder(bytes.NewReader(ex))
						err2 := test.decode(decoder2)
						if err2 != nil {
							return
						}
						output = append(output, &OutputTp{
							OperationID: test.operationID,
							Error:       err.Error(),
						})
					}
				})
			}
		})
	}
	y, err := yaml.Marshal(output)
	require.NoError(t, err)
	require.NoError(t, os.MkdirAll("tmp", 0o750))
	err = ioutil.WriteFile("tmp/example_errs.yml", y, 0o640)
	require.NoError(t, err)
}
