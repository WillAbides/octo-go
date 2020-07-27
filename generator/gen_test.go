package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllEndpointAttributesHaveName(t *testing.T) {
	for i := endpointAttribute(0); i < attrInvalid; i++ {
		require.NotEmpty(t, i.String())
	}
}

// Generates to a temp directory.  This is primarily to accommodate easier step-through debugging.
func Test_run(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, os.RemoveAll(tmpDir))
	})
	outputPath := tmpDir
	fmt.Println(outputPath)
	schemaPath := filepath.FromSlash("../api.github.com.json")
	pkgPath := "github.com/willabides/octo-go"
	pkgName := "octo"
	err = run(schemaPath, outputPath, pkgPath, pkgName)
	require.NoError(t, err)
}
