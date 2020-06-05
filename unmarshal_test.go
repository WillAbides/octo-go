package octo_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

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
			examples := responseExamples(t, test.endpointPath, test.httpMethod, test.httpStatusCode)
			require.Greater(t, len(examples), 0, "no examples found")
			for name, ex := range examples {
				// skip until ReposGetContents works for directories
				if name == "response-if-content-is-a-directory" {
					continue
				}
				t.Run(name, func(t *testing.T) {
					decoder := json.NewDecoder(bytes.NewReader(ex))
					// comment out decoder.DisallowUnknownFields until https://github.com/github/openapi/issues/177 is fixed
					//decoder.DisallowUnknownFields()
					require.NoError(t, test.decode(decoder))
				})
			}
		})
	}
}
