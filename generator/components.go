package main

import (
	"sort"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model"
)

func addComponentSchemas(file *jen.File, schemas map[string]*model.ParamSchema) {

	names := make([]string, 0, len(schemas))
	for name := range schemas {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		schema := schemas[name]
		structName := toExportedName(name)
		tp := paramSchemaFieldType(schema, []string{"components", "schemas", "name"}, false, true, true)
		file.Type().Id(structName).Add(tp)
		file.Line()
	}
}
