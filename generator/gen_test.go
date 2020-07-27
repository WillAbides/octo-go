package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go/generator/internal/model"
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

func Test_respBodyType(t *testing.T) {
	endpoint := &model.Endpoint{
		Name:    "blah-blah",
		Concern: "puppies",
		Responses: map[int]*model.Response{
			200: {
				Body: &model.ParamSchema{
					Ref:  "",
					Type: model.ParamTypeArray,
					ItemSchema: &model.ParamSchema{
						Ref: "#/components/schemas/foo-bar",
					},
				},
			},
		},
	}
	got := respBodyType(endpoint)
	require.Equal(t, &qualifiedType{
		pkg:   "github.com/willabides/octo-go/components",
		name:  "FooBar",
		slice: true,
	}, got)
}
