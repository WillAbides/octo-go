package openapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// This is a placeholder for ad-hoc debugging
func TestOpenapi2Model(t *testing.T) {
	swaggerFile, err := os.Open(filepath.FromSlash("../../../../api.github.com.json"))
	require.NoError(t, err)
	got, err := Openapi2Model(swaggerFile)
	require.NoError(t, err)
	for _, endpoint := range got.Endpoints {
		_ = endpoint
	}
}
