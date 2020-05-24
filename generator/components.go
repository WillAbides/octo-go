package main

import (
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model"
)

func componentRefStmt(schema *model.ParamSchema) *jen.Statement {
	if !strings.HasPrefix(schema.Ref, "#/components/schemas/") {
		return nil
	}
	nm := strings.TrimPrefix(schema.Ref, "#/components/schemas/")
	return jen.Struct(jen.Qual("github.com/willabides/octo-go/components", toExportedName(nm)))
}

func addComponentSchemas(file *jen.File, schemas map[string]*model.ParamSchema) {

	names := make([]string, 0, len(schemas))
	for name := range schemas {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		schema := schemas[name]
		structName := toExportedName(name)
		tp := paramSchemaFieldType(schema, []string{"components", "schemas", "name"}, &paramSchemaFieldTypeOptions{
			noHelperRecursive: true,
		})
		file.Type().Id(structName).Add(tp)
		file.Line()
	}
}
