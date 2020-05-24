package main

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/mitchellh/go-wordwrap"
	"github.com/willabides/octo-go/generator/internal/model"
)

func toExportedName(in string) string {
	out := in
	for _, separator := range []string{"_", "-", ".", "/"} {
		words := strings.Split(out, separator)
		for i, word := range words {
			words[i] = strings.Title(word)
		}
		out = strings.Join(words, "")
	}
	return out
}

func removeValFromStringSlice(sl []string, val string) []string {
	result := make([]string, 0, len(sl))
	for _, s := range sl {
		if s != val {
			result = append(result, s)
		}
	}
	return result
}

type paramSchemaFieldTypeOptions struct {
	usePointers, noHelper, noHelperRecursive bool
}

func paramSchemaFieldType(schema *model.ParamSchema, schemaPath []string, opts *paramSchemaFieldTypeOptions) *jen.Statement {
	if opts == nil {
		opts = new(paramSchemaFieldTypeOptions)
	}
	overrideParamSchema(schemaPath, schema)

	componentRef := componentRefStmt(schema)
	if componentRef != nil {
		return componentRef
	}

	helperStruct := reqBodyNestedStructName(schemaPath, schema)
	if opts.noHelperRecursive {
		opts.noHelper = true
	}
	if !opts.noHelper && helperStruct != "" {
		return jen.Id(helperStruct)
	}

	if schema == nil {
		return nil
	}
	switch schema.Type {
	case model.ParamTypeString:
		return jen.String()
	case model.ParamTypeInt:
		return jen.Int64()
	case model.ParamTypeBool:
		return jen.Bool()
	case model.ParamTypeNumber:
		return jen.Qual("encoding/json", "Number")
	case model.ParamTypeInterface:
		return jen.Interface()
	case model.ParamTypeArray:
		return jen.Id("[]").Add(paramSchemaFieldType(schema.ItemSchema, append(schemaPath, "ITEM_SCHEMA"), opts))
	case model.ParamTypeObject:
		return paramSchemaObjectFieldType(schema, schemaPath, opts)
	}
	return nil
}

func paramSchemaObjectFieldType(schema *model.ParamSchema, schemaPath []string, opts *paramSchemaFieldTypeOptions) *jen.Statement {
	if opts == nil {
		opts = new(paramSchemaFieldTypeOptions)
	}
	if !opts.noHelperRecursive {
		opts.noHelper = false
	}
	if len(schema.ObjectParams) > 0 {
		return jen.StructFunc(func(group *jen.Group) {
			for _, param := range schema.ObjectParams {
				if param.HelpText != "" {
					group.Line()
				}
				gStmt := jen.Id(toExportedName(param.Name))
				pType := paramSchemaFieldType(param.Schema, append(schemaPath, param.Name), opts)
				if opts.usePointers && needsPointer(param.Schema) {
					gStmt.Op("*")
				}
				jsonTag := param.Name
				if !param.Required {
					jsonTag += ",omitempty"
				}
				if param.HelpText != "" {
					group.Comment(wordwrap.WrapString(param.HelpText, 80))
				}
				group.Add(gStmt.Add(pType).Tag(map[string]string{
					"json": jsonTag,
				}))
			}
		})
	}
	if schema.ItemSchema != nil {
		stmt := jen.Map(jen.String())
		if opts.usePointers && needsPointer(schema.ItemSchema) {
			stmt.Op("*")
		}
		return stmt.Add(paramSchemaFieldType(schema.ItemSchema, append(schemaPath, "ITEM_SCHEMA"), opts))
	}
	return jen.Interface()
}

func needsPointer(schema *model.ParamSchema) bool {
	if schema.Type == model.ParamTypeInterface {
		return false
	}
	if schema.Type == model.ParamTypeArray {
		return false
	}
	if schema.Type == model.ParamTypeObject && len(schema.ObjectParams) == 0 {
		return false
	}
	return true
}
