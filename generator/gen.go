package main

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/mitchellh/go-wordwrap"
	"github.com/willabides/octo-go/generator/internal/model"
)

func addRequestFunc(file *jen.File, endpoint *model.Endpoint) {
	file.Commentf("%s performs requests for \"%s\"\n\n%s.\n\n  %s %s\n\n%s",
		toExportedName(endpoint.ID),
		endpoint.ID,
		endpoint.Summary,
		endpoint.Method,
		endpoint.Path,
		endpoint.DocsURL,
	)
	file.Func().Id(toExportedName(endpoint.ID)).Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("req").Op("*").Id(reqStructName(endpoint)),
		jen.Id("opt ...RequestOption"),
	).Params(
		jen.Op("*").Id(respStructName(endpoint)),
		jen.Id("error"),
	).BlockFunc(func(group *jen.Group) {
		group.If(jen.Id("req == nil")).Block(
			jen.Id("req").Op("=").New(jen.Id(reqStructName(endpoint))),
		)
		group.Id("resp").Op(":=").Op("&").Id(respStructName(endpoint)).Values(jen.Dict{
			jen.Id("request"): jen.Id("req"),
		})
		group.Id("r, err := doRequest(ctx, req, opt...)")
		group.If(jen.Id("r != nil")).Block(
			jen.Id("resp.response = *r"),
		)
		group.If(jen.Id("err != nil")).Block(jen.Id("return resp, err"))

		switch {
		case endpointHasAttribute(endpoint, attrNoResponseBody):
			group.Id("err").Op("=").Id("r.decodeBody(nil)")
		case endpointHasAttribute(endpoint, attrBoolean):
			group.Id("err").Op("=").Id("r.setBoolResult(&resp.Data)")
			group.If(jen.Id("err != nil")).Block(jen.Id("return nil, err"))
			group.Id("err").Op("=").Id("r.decodeBody(nil)")
		case len(responseCodesWithBodies(endpoint)) > 0:
			group.Id("resp").Dot("Data").Op("=").Id(respBodyStructName(endpoint)).Block()
			group.Id("err").Op("=").Id("r.decodeBody(&resp.Data)")
		default:
			group.Id("err").Op("=").Id("r.decodeBody(nil)")
		}
		group.If(jen.Id("err != nil")).Block(jen.Id("return nil, err"))
		group.Return(jen.Id("resp, nil"))
	})
}

func addClientMethod(file *jen.File, endpoint *model.Endpoint) {
	file.Commentf("%s performs requests for \"%s\"\n\n%s.\n\n  %s %s\n\n%s",
		toExportedName(endpoint.ID),
		endpoint.ID,
		endpoint.Summary,
		endpoint.Method,
		endpoint.Path,
		endpoint.DocsURL,
	)
	file.Func().Params(jen.Id("c").Id("Client")).Id(toExportedName(endpoint.ID)).Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("req").Op("*").Id(reqStructName(endpoint)),
		jen.Id("opt ...RequestOption"),
	).Params(
		jen.Op("*").Id(respStructName(endpoint)),
		jen.Id("error"),
	).Block(
		jen.Return(
			jen.Id(toExportedName(endpoint.ID)).Call(
				jen.Id("ctx"),
				jen.Id("req"),
				jen.Id("append(c, opt...)..."),
			),
		),
	)
}

func toUnexportedName(in string) string {
	if _, ok := nameOverrides[in]; ok {
		in = nameOverrides[in]
	}
	out := in
	for _, separator := range []string{"_", "-", ".", "/"} {
		words := strings.Split(out, separator)
		for i := range words {
			if i == 0 {
				continue
			}
			words[i] = strings.Title(words[i])
		}
		out = strings.Join(words, "")
	}
	return out
}

func toExportedName(in string) string {
	if _, ok := nameOverrides[in]; ok {
		in = nameOverrides[in]
	}
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

	compSchemaRef := compSchemaRefStmt(schema)
	if compSchemaRef != nil {
		return compSchemaRef
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
	case model.ParamTypeOneOf:
		return paramSchemaOneOfFieldType(schema, schemaPath, opts)
	}
	return nil
}

func paramSchemaOneOfFieldType(schema *model.ParamSchema, schemaPath []string, opts *paramSchemaFieldTypeOptions) *jen.Statement {
	if opts == nil {
		opts = new(paramSchemaFieldTypeOptions)
	}
	if !opts.noHelperRecursive {
		opts.noHelper = false
	}
	paramFields := []jen.Code{jen.Id("oneOfField").String()}
	for _, param := range schema.ObjectParams {
		paramFields = append(paramFields, oneOfParamStmt(param, schemaPath, opts))
	}
	return jen.Struct(paramFields...)
}

func paramSchemaObjectFieldType(schema *model.ParamSchema, schemaPath []string, opts *paramSchemaFieldTypeOptions) *jen.Statement {
	if opts == nil {
		opts = new(paramSchemaFieldTypeOptions)
	}
	if !opts.noHelperRecursive {
		opts.noHelper = false
	}
	var paramFields []jen.Code
	for _, param := range schema.ObjectParams {
		paramFields = append(paramFields, objectParamStmt(param, schemaPath, opts))
	}
	if len(paramFields) > 0 {
		return jen.Struct(paramFields...)
	}
	if schema.ItemSchema != nil {
		stmt := jen.Map(jen.String())
		return stmt.Add(paramSchemaFieldType(schema.ItemSchema, append(schemaPath, "ITEM_SCHEMA"), opts))
	}
	return jen.Interface()
}

func oneOfParamStmt(param *model.Param, schemaPath []string, opts *paramSchemaFieldTypeOptions) *jen.Statement {
	stmt := jen.Id(toUnexportedName(param.Name))
	pType := paramSchemaFieldType(param.Schema, append(schemaPath, param.Name), opts)
	if needsPointer(param.Schema, opts.usePointers) {
		stmt.Op("*")
	}
	stmt.Add(pType)
	return prependCodeWithComment(param.HelpText, stmt)
}

// prependCodeWithComment returns a statement that is code with a comment prepended, or just the code if the comment is ""
func prependCodeWithComment(comment string, code ...jen.Code) *jen.Statement {
	if comment == "" {
		return jen.Add(code...)
	}
	if len(comment) > 120 {
		comment = wordwrap.WrapString(comment, 80)
	}
	return jen.Line().Comment(comment).Line().Add(code...)
}

func objectParamStmt(param *model.Param, schemaPath []string, opts *paramSchemaFieldTypeOptions) *jen.Statement {
	stmt := jen.Id(toExportedName(param.Name))
	if needsPointer(param.Schema, opts.usePointers) {
		stmt.Op("*")
	}
	stmt.Add(paramSchemaFieldType(param.Schema, append(schemaPath, param.Name), opts))
	jsonTag := param.Name
	if !param.Required {
		jsonTag += ",omitempty"
	}
	stmt.Tag(map[string]string{
		"json": jsonTag,
	})
	return prependCodeWithComment(param.HelpText, stmt)
}

func needsPointer(schema *model.ParamSchema, usePointers bool) bool {
	if !usePointers && !schema.Nullable {
		return false
	}
	switch schema.Type {
	case model.ParamTypeInterface, model.ParamTypeArray:
		return false
	case model.ParamTypeObject, model.ParamTypeOneOf:
		if len(schema.ObjectParams) == 0 {
			return false
		}
		if schema.Nullable {
			return true
		}
	}
	return usePointers
}
