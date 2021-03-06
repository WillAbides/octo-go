package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model"
)

func reqBodyNestedStructName(schemaPath []string, schema *model.ParamSchema) string {
	// We don't want ITEM_SCHEMA in the name, and removing it doesn't cause duplicate struct names
	sp := removeValFromStringSlice(schemaPath, "ITEM_SCHEMA")

	if len(sp) < 3 {
		return ""
	}
	if schemaPath[1] != "reqBody" {
		return ""
	}
	if len(schema.ObjectParams) == 0 {
		return ""
	}
	suffix := toExportedName(strings.Join(sp[2:], "-"))
	return reqBodyStructName(sp[0]) + suffix
}

func addRequestBody(file *jen.File, pq pkgQual, endpoint *model.Endpoint) {
	if endpointHasAttribute(endpoint, attrBodyUploader) {
		return
	}
	addReqBodyNestedStructs(file, pq, endpoint)
	if endpoint.RequestBody == nil {
		return
	}
	schema := endpoint.RequestBody.Schema
	tp := paramSchemaFieldType(schema, []string{endpoint.ID, "reqBody"}, pq, &paramSchemaFieldTypeOptions{
		usePointers: true,
	})
	if tp == nil {
		return
	}

	structName := reqBodyStructName(endpoint.ID)
	file.Commentf("%s is a request body for %s\n\n%s",
		structName,
		endpoint.ID,
		endpoint.DocsURL,
	)
	file.Type().Id(structName).Add(tp)
}

func reqBodyNestedStructs(schemaPath []string, pq pkgQual, schema *model.ParamSchema) []*jen.Statement {
	var result []*jen.Statement
	helperName := reqBodyNestedStructName(schemaPath, schema)
	if helperName != "" {
		tp := paramSchemaFieldType(schema, schemaPath, pq, &paramSchemaFieldTypeOptions{
			usePointers: true,
			noHelper:    true,
		})
		sp := removeValFromStringSlice(schemaPath, "ITEM_SCHEMA")
		parentStructName := toExportedName(strings.Join(sp[:len(sp)-1], "-"))
		parentValueName := toExportedName(sp[len(sp)-1])
		comment := fmt.Sprintf("%s is a value for %s's %s field", helperName, parentStructName, parentValueName)
		result = append(result, jen.Comment(comment))
		result = append(result, jen.Type().Id(helperName).Add(tp))
	}
	if schema.ItemSchema != nil {
		result = append(result, reqBodyNestedStructs(append(schemaPath, "ITEM_SCHEMA"), pq, schema.ItemSchema)...)
	}
	for _, param := range schema.ObjectParams {
		nr := reqBodyNestedStructs(append(schemaPath, param.Name), pq, param.Schema)
		result = append(result, nr...)
	}
	return result
}

func addReqBodyNestedStructs(file *jen.File, pq pkgQual, endpoint *model.Endpoint) {
	if endpoint.RequestBody == nil {
		return
	}
	stmts := reqBodyNestedStructs([]string{endpoint.ID, "reqBody"}, pq, endpoint.RequestBody.Schema)
	for _, stmt := range stmts {
		file.Add(stmt)
	}
}
