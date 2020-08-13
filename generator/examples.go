package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"
)

func componentExampleFiles(schemaPath, outputPath string) error {
	schema, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(schemaPath)
	if err != nil {
		return err
	}
	for exName, exRef := range schema.Components.Examples {
		val := exRef.Value.Value
		fileName := filepath.Join(outputPath, fmt.Sprintf("%s.json", exName))
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(&val)
		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
