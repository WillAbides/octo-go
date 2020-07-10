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
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	for _, test := range unmarshalResponseBodyTests {
		t.Run(fmt.Sprintf("%s_%d", test.name, test.httpStatusCode), func(t *testing.T) {
			examples := responseExamples(t, test.endpointPath, test.httpMethod, test.httpStatusCode)
			require.Greater(t, len(examples), 0, "no examples found")
			for name, ex := range examples {
				t.Run(name, func(t *testing.T) {
					decoder := json.NewDecoder(bytes.NewReader(ex))
					decoder.DisallowUnknownFields()
					require.NoError(t, test.decode(decoder))
				})
			}
		})
	}
}
