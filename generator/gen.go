package main

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/mitchellh/go-wordwrap"
	"github.com/willabides/octo-go/generator/internal/model"
)

func addClientMethod(file *jen.File, endpoint model.Endpoint) {
	file.Commentf("%s performs requests for \"%s\"\n\n%s.\n\n  %s %s\n\n%s",
		toExportedName(endpoint.ID),
		endpoint.ID,
		endpoint.Summary,
		endpoint.Method,
		endpoint.Path,
		endpoint.DocsURL,
	)
	file.Func().Params(jen.Id("c").Op("*").Id("Client")).Id(toExportedName(endpoint.ID)).Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("req").Op("*").Id(reqStructName(endpoint)),
		jen.Id("opt ...RequestOption"),
	).Params(
		jen.Op("*").Id(respStructName(endpoint)),
		jen.Id("error"),
	).BlockFunc(func(group *jen.Group) {
		group.Id("r, err := c.doRequest(ctx, req, opt...)")
		group.If(jen.Id("err != nil")).Block(jen.Id("return nil, err"))
		group.Id("resp").Op(":=").Op("&").Id(respStructName(endpoint)).Values(jen.Dict{
			jen.Id("response"): jen.Id("*r"),
			jen.Id("request"):  jen.Id("req"),
		})
		switch {
		case endpointHasAttribute(endpoint, attrBoolean):
			group.Id("err").Op("=").Id("r.setBoolResult(&resp.Data)")
			group.If(jen.Id("err != nil")).Block(jen.Id("return nil, err"))
			group.Id("err").Op("=").Id("r.decodeBody(nil)")
		case len(responseCodesWithBodies(endpoint)) > 0:
			group.Id("resp").Dot("Data").Op("=").New(jen.Id(respBodyStructName(endpoint)))
			group.Id("err").Op("=").Id("r.decodeBody(resp.Data)")
		default:
			group.Id("err").Op("=").Id("r.decodeBody(nil)")
		}
		group.If(jen.Id("err != nil")).Block(jen.Id("return nil, err"))
		group.Return(jen.Id("resp, nil"))
	})
}

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

func endpointJSONRequestSchema(endpoint model.Endpoint) *model.ParamSchema {
	for _, request := range endpoint.Requests {
		if request.MimeType == "application/json" {
			return request.Schema
		}
	}
	return nil
}

type endpointAttribute int

const (
	// endpoint that redirects via 302
	attrRedirectOnly endpointAttribute = iota
	// gives boolean status through 204 vs 404 status codes
	attrBoolean
	// attrInvalid is last so we can get a list of all valid types with a for loop
	attrInvalid
)

var attrNames = map[endpointAttribute]string{
	attrRedirectOnly: "attrRedirectOnly",
	attrBoolean:      "attrBoolean",
}

func (e endpointAttribute) String() string {
	return attrNames[e]
}

func addEndpointAttributes(file *jen.File) {
	file.Type().Id("endpointAttribute").Int()
	file.Const().Parens(jen.Do(func(statement *jen.Statement) {
		for i := endpointAttribute(0); i < attrInvalid; i++ {
			statement.Id(attrNames[i])
			if i == 0 {
				statement.Id("endpointAttribute").Op("=").Iota()
			}
			statement.Line()
		}
	}))
}

func endpointHasAttribute(endpoint model.Endpoint, attribute endpointAttribute) bool {
	for _, attr := range getEndpointAttributes(endpoint) {
		if attribute == attr {
			return true
		}
	}
	return false
}

func getEndpointAttributes(endpoint model.Endpoint) []endpointAttribute {
	override, ok := overrideEndpointAttrs[endpoint.ID]
	if ok {
		return override
	}
	var result []endpointAttribute
	for _, check := range attrChecks {
		check(endpoint, &result)
	}
	return result
}

type attrCheck func(endpoint model.Endpoint, attrs *[]endpointAttribute)

var attrChecks = []attrCheck{
	// attrBoolean if the endpoint has exatcly two responses: 204 and 404
	func(endpoint model.Endpoint, attrs *[]endpointAttribute) {
		if len(endpoint.Responses) != 2 {
			return
		}
		if _, ok := endpoint.Responses[204]; !ok {
			return
		}
		if _, ok := endpoint.Responses[404]; !ok {
			return
		}
		*attrs = append(*attrs, attrBoolean)
	},

	// attrRedirectOnly if the endpoint has onlly one response: 302
	func(endpoint model.Endpoint, attrs *[]endpointAttribute) {
		if len(endpoint.Responses) != 1 {
			return
		}
		_, ok := endpoint.Responses[302]
		if !ok {
			return
		}
		*attrs = append(*attrs, attrRedirectOnly)
	},
}
