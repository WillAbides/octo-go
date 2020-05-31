package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/mitchellh/go-wordwrap"
	"github.com/willabides/octo-go/generator/internal/model"
)

func reqBodyStructName(endpointID string) string {
	endpointID = strings.ReplaceAll(endpointID, "/", "-")
	return toExportedName(fmt.Sprintf("%s-req-body", endpointID))
}

func reqStructName(endpoint model.Endpoint) string {
	return toExportedName(fmt.Sprintf("%s-%s-req", endpoint.Concern, endpoint.Name))
}

func addRequestStruct(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	file.Commentf("%s is request data for Client.%s\n\n%s",
		structName,
		toExportedName(endpoint.ID),
		endpoint.DocsURL,
	)
	file.Type().Id(structName).StructFunc(func(group *jen.Group) {
		group.Id("pgURL").String()
		for _, param := range endpoint.PathParams {
			if param.HelpText != "" {
				group.Line().Comment(wordwrap.WrapString(param.HelpText, 80))
			}
			group.Id(toExportedName(param.Name)).Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "PATH_PARAMS"}, nil))
		}
		for _, param := range endpoint.QueryParams {
			if param.HelpText != "" {
				group.Line().Comment(wordwrap.WrapString(param.HelpText, 80))
			}
			group.Id(toExportedName(param.Name)).Op("*").Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "QUERY_PARAMS"}, &paramSchemaFieldTypeOptions{
				usePointers: true,
			}))
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
			group.Id(toExportedName(param.Name + "-header")).Op("*").Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "QUERY_PARAMS"}, &paramSchemaFieldTypeOptions{
				usePointers: true,
			}))
		}
		for _, preview := range endpoint.Previews {
			if preview.Note != "" {
				group.Line().Comment(wordwrap.WrapString(fixPreviewNote(preview.Note), 80))
			}
			group.Id(toExportedName(preview.Name + "-preview")).Bool()
		}
	})

	for _, fn := range []func(file *jen.File, endpoint model.Endpoint){
		func(fl *jen.File, endpoint model.Endpoint) {
			fl.Func().Params(jen.Id("r").Id("*" + structName)).Id("pagingURL() string").Block(
				jen.Id("return r.pgURL"),
			)
		},
		reqURLPathFunc,
		func(fl *jen.File, endpoint model.Endpoint) {
			fl.Func().Params(jen.Id("r").Id("*" + structName)).Id("method").Params().String().Block(
				jen.Return(jen.Lit(endpoint.Method)),
			)
		},
		reqURLQueryFunc,
		reqHeaderFunc,
		reqBodyFunc,
		reqDataStatusesFunc,
		reqValidStatusesFunc,
		reqEndpointTypeFunc,
		reqHTTPRequestFunc,
		reqRelReqFunc,
	} {
		fn(file, endpoint)
		file.Line()
	}

}

func reqRelReqFunc(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	comment := `Rel updates this request to point to a relative link from resp. Returns false if the link does not exist. Handy for paging.`
	file.Comment(wordwrap.WrapString(comment, 80))
	file.Func().Params(jen.Id("r").Id("*"+structName)).Id("Rel").Params(
		jen.Id("link RelName"),
		jen.Id("resp").Op("*").Id(respStructName(endpoint)),
	).Params(jen.Bool()).Block(
		jen.Id("u := resp.RelLink(link)"),
		jen.If(jen.Id("u").Op("==").Lit("")).Block(jen.Return(jen.False())),
		jen.Id("r.pgURL = u"),
		jen.Return(jen.True()),
	)
}

func reqDataStatusesFunc(file *jen.File, endpoint model.Endpoint) {
	structName := reqStructName(endpoint)
	file.Func().Params(jen.Id("r").Id("*" + structName)).Id("dataStatuses").Params().Params(
		jen.Op("[]").Int(),
	).Block(
		jen.Return().Op("[]").Int().ValuesFunc(func(group *jen.Group) {
			for _, code := range responseCodesWithBodies(endpoint) {
				group.Lit(code)
			}
		}),
	)
}

func reqEndpointTypeFunc(file *jen.File, endpoint model.Endpoint) {
	tp := getEndpointType(endpoint).String()
	structName := reqStructName(endpoint)
	file.Func().Params(
		jen.Id("r").Id("*" + structName),
	).Id("endpointType()").Params(
		jen.Id("endpointType"),
	).Block(
		jen.Return(jen.Id(tp)),
	)
}

func reqValidStatusesFunc(file *jen.File, endpoint model.Endpoint) {
	codes := make([]int, 0, len(endpoint.Responses))
	for code := range endpoint.Responses {
		if code < 400 {
			codes = append(codes, code)
		}
	}
	sort.Ints(codes)
	if getEndpointType(endpoint) == endpointTypeRedirect {
		codes = []int{-1}
	}
	structName := reqStructName(endpoint)
	file.Func().Params(jen.Id("r").Id("*" + structName)).Id("validStatuses").Params().Params(
		jen.Op("[]").Int(),
	).Block(
		jen.Return().Op("[]").Int().ValuesFunc(func(group *jen.Group) {
			for _, code := range codes {
				group.Lit(code)
			}
		}),
	)
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
	file.Comment("httpRequest creates an http request")
	file.Func().Params(jen.Id("r").Id("*"+structName)).Id("httpRequest").Params(
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
					dict[jen.Lit(header.Name)] = jen.Id("r").Dot(toExportedName(header.Name + "-header"))
				}
			}),
		)
		fnBlock.Id("previewVals").Op(":=").Map(jen.String()).Bool().Values(
			jen.DictFunc(func(dict jen.Dict) {
				for _, preview := range endpoint.Previews {
					dict[jen.Lit(preview.Name)] = jen.Id("r").Dot(toExportedName(preview.Name + "-preview"))
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
					group.Id("r").Dot(toExportedName(param.Name))
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
			paramArg := jen.Id("r").Dot(toExportedName(param.Name))
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