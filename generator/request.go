package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/mitchellh/go-wordwrap"
	"github.com/willabides/octo-go/generator/internal/model"
)

func reqBodyStructName(endpointID string) string {
	endpointID = strings.ReplaceAll(endpointID, "/", "-")
	return toArgName(fmt.Sprintf("%s-req-body", endpointID))
}

func reqStructName(endpoint model.Endpoint) string {
	return toArgName(fmt.Sprintf("%s-%s-req", endpoint.Concern, endpoint.Name))
}

func addRequestStruct(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	file.Commentf("%s builds requests for \"%s\"\n\n%s.\n\n  %s %s\n\n%s",
		structName,
		endpoint.ID,
		endpoint.Summary,
		endpoint.Method,
		endpoint.Path,
		endpoint.DocsURL,
	)
	file.Type().Id(structName).StructFunc(func(group *jen.Group) {
		for _, param := range endpoint.PathParams {
			if param.HelpText != "" {
				group.Line().Comment(wordwrap.WrapString(param.HelpText, 80))
			}
			group.Id(toArgName(param.Name)).Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "PATH_PARAMS"}, false, false))
		}
		for _, param := range endpoint.QueryParams {
			if param.HelpText != "" {
				group.Line().Comment(wordwrap.WrapString(param.HelpText, 80))
			}
			group.Id(toArgName(param.Name)).Op("*").Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "QUERY_PARAMS"}, true, false))
		}
		if endpoint.JSONBodySchema != nil {
			group.Id("RequestBody").Id(reqBodyStructName(endpoint.ID))
		}
		for _, param := range endpoint.Headers {
			if param.Name == "accept" {
				continue
			}
			if param.HelpText != "" {
				group.Line().Comment(wordwrap.WrapString(param.HelpText, 80))
			}
			group.Id(toArgName(param.Name + "-header")).Op("*").Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "QUERY_PARAMS"}, true, false))
		}
		for _, preview := range endpoint.Previews {
			if preview.Note != "" {
				group.Line().Comment(wordwrap.WrapString(fixPreviewNote(preview.Note), 80))
			}
			group.Id(toArgName(preview.Name + "-preview")).Bool()
		}
	})

	reqURLPathFunc(file, endpoint)
	file.Line()

	file.Func().Params(jen.Id("r").Id("*" + structName)).Id("method").Params().String().Block(
		jen.Return(jen.Lit(endpoint.Method)),
	)
	file.Line()

	reqURLQueryFunc(file, endpoint)
	file.Line()

	reqHeaderFunc(file, endpoint)
	file.Line()

	reqBodyFunc(file, endpoint)
	file.Line()

	reqHTTPRequestFunc(file, endpoint)
}

func reqBodyFunc(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	file.Func().Params(jen.Id("r").Id("*" + structName)).Id("body").
		Params().
		Interface().
		Block(jen.Do(func(stmt *jen.Statement) {
			if endpoint.JSONBodySchema == nil {
				stmt.Return(jen.Nil())
				return
			}
			stmt.Return(jen.Id("r.RequestBody"))
		}))
}

func reqHTTPRequestFunc(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	file.Comment("HTTPRequest creates an http request")
	file.Func().Params(jen.Id("r").Id("*"+structName)).Id("HTTPRequest").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("opt ...RequestOption"),
	).Params(
		jen.Op("*").Qual("net/http", "Request"), jen.Error(),
	).Block(
		jen.Id("return buildHTTPRequest(ctx, r, opt)"),
	)
}

func reqHeaderFunc(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	stmt := file.Func().Params(jen.Id("r").Id("*"+structName)).Id("header").Params(jen.Id("requiredPreviews, allPreviews bool")).Qual("net/http", "Header")
	hasRequiredPreviews := false
	for _, preview := range endpoint.Previews {
		if preview.Required {
			hasRequiredPreviews = true
			break
		}
	}
	stmt.BlockFunc(func(fnBlock *jen.Group) {
		fnBlock.Id("headerVals").Op(":=").Map(jen.String()).Op("*").String().Values(
			jen.DictFunc(func(dict jen.Dict) {
				for _, header := range endpoint.Headers {
					if header.Name == "accept" {
						continue
					}
					dict[jen.Lit(header.Name)] = jen.Id("r").Dot(toArgName(header.Name + "-header"))
				}
			}),
		)
		fnBlock.Id("previewVals").Op(":=").Map(jen.String()).Bool().Values(
			jen.DictFunc(func(dict jen.Dict) {
				for _, preview := range endpoint.Previews {
					dict[jen.Lit(preview.Name)] = jen.Id("r").Dot(toArgName(preview.Name + "-preview"))
				}
			}),
		)

		if hasRequiredPreviews {
			fnBlock.If(jen.Id("requiredPreviews")).BlockFunc(func(ifGroup *jen.Group) {
				for _, preview := range endpoint.Previews {
					if !preview.Required {
						continue
					}
					ifGroup.Id("previewVals").Index(jen.Lit(preview.Name)).Op("=").True()
				}
			})
		}

		if len(endpoint.Previews) > 0 {
			fnBlock.If(jen.Id("allPreviews")).BlockFunc(func(ifGroup *jen.Group) {
				for _, preview := range endpoint.Previews {
					ifGroup.Id("previewVals").Index(jen.Lit(preview.Name)).Op("=").True()
				}
			})
		}

		fnBlock.Return(jen.Id("requestHeaders")).Params(
			jen.Id("headerVals"),
			jen.Id("previewVals"),
		)
	})
}

var bracesExp = regexp.MustCompile(`{[^}]+}`)

func reqURLPathFunc(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	file.Func().Params(jen.Id("r").Id("*" + structName)).Id("urlPath").Params().String().
		BlockFunc(func(group *jen.Group) {
			pth := bracesExp.ReplaceAllString(endpoint.Path, "%v")
			group.Return(jen.Qual("fmt", "Sprintf").ParamsFunc(func(group *jen.Group) {
				group.Lit(pth)
				for _, param := range endpoint.PathParams {
					group.Id("r").Dot(toArgName(param.Name))
				}
			}))
		})
}

func reqURLQueryFunc(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	stmt := file.Func().Params(jen.Id("r").Id("*" + structName)).Id("urlQuery").Params()
	stmt.Qual("net/url", "Values")
	stmt.BlockFunc(func(group *jen.Group) {
		group.Id("query").Op(":=").Qual("net/url", "Values").Block()

		for _, param := range endpoint.QueryParams {
			paramArg := jen.Id("r").Dot(toArgName(param.Name))
			group.If(paramArg.Clone().Op("!=").Nil()).BlockFunc(func(ifGroup *jen.Group) {
				var valStmt *jen.Statement
				switch param.Schema.Type {
				case model.ParamTypeString:
					valStmt = jen.Op("*").Add(paramArg)
				case model.ParamTypeInt:
					valStmt = jen.Qual("strconv", "FormatInt").Params(jen.Op("*").Add(paramArg), jen.Lit(10))
				case model.ParamTypeBool:
					valStmt = jen.Qual("strconv", "FormatBool").Params(jen.Op("*").Add(paramArg))
				default:
					fmt.Printf("UNEXPECTED %v\n", param)
				}
				ifGroup.Id("query").Dot("Set").Params(jen.Lit(param.Name), valStmt)
			})
		}
		group.Return(jen.Id("query"))
	})
}
