package octo_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

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
		"ActivityListStargazersForRepoResponseBody":               true,
		"ActivityListReposStarredByAuthenticatedUserResponseBody": true,
		"ActivityListReposStarredByUserResponseBody":              true,
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
	require.NoError(t, os.MkdirAll("tmp", 0750))
	err = ioutil.WriteFile("tmp/example_errs.yml", y, 0o640)
	require.NoError(t, err)
}
