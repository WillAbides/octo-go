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
		case getEndpointType(endpoint) == endpointTypeBoolean:
			group.Id("err").Op("=").Id("r.setBoolResult(&resp.Data)")
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

type endpointType int

const (
	// just a regular endpoint
	endpointTypeRegular endpointType = iota
	// endpoint that redirects via 302
	endpointTypeRedirect
	// gives boolean status through 204 vs 404 status codes
	endpointTypeBoolean
	// endpointTypeInvalid is last so we can get a list of all valid types with a for loop
	endpointTypeInvalid
)

var endpointTypeNames = map[endpointType]string{
	endpointTypeRegular:  "endpointTypeRegular",
	endpointTypeRedirect: "endpointTypeRedirect",
	endpointTypeBoolean:  "endpointTypeBoolean",
}

func (e endpointType) String() string {
	return endpointTypeNames[e]
}

func addEndpointTypes(file *jen.File) {
	file.Type().Id("endpointType").Int()
	file.Const().Parens(jen.Do(func(statement *jen.Statement) {
		for i := endpointType(0); i < endpointTypeInvalid; i++ {
			statement.Id(endpointTypeNames[i])
			if i == 0 {
				statement.Id("endpointType").Op("=").Iota()
			}
			statement.Line()
		}
	}))
}

func getEndpointType(endpoint model.Endpoint) endpointType {
	if isRedirectOnlyEndpoint(endpoint) {
		return endpointTypeRedirect
	}
	switch {
	case isRedirectOnlyEndpoint(endpoint):
		return endpointTypeRedirect
	case isBooleanEndpoint(endpoint):
		return endpointTypeBoolean
	}
	return endpointTypeRegular
}

func isBooleanEndpoint(endpoint model.Endpoint) bool {
	if len(endpoint.Responses) != 2 {
		return false
	}
	if _, ok := endpoint.Responses[204]; !ok {
		return false
	}
	if _, ok := endpoint.Responses[404]; !ok {
		return false
	}
	return true
}

func isRedirectOnlyEndpoint(endpoint model.Endpoint) bool {
	if len(endpoint.Responses) != 1 {
		return false
	}
	_, ok := endpoint.Responses[302]
	return ok
}
