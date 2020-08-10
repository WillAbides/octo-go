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

func reqStructName(endpoint *model.Endpoint) string {
	return toExportedName(fmt.Sprintf("%s-%s-req", endpoint.Concern, endpoint.Name))
}

func addRequestStruct(file *jen.File, pq pkgQual, endpoint *model.Endpoint) {
	structName := reqStructName(endpoint)
	file.Commentf("%s is request data for Client.%s\n\n%s",
		structName,
		toExportedName(endpoint.ID),
		endpoint.DocsURL,
	)
	file.Type().Id(structName).StructFunc(func(group *jen.Group) {
		group.Id("_url").String()
		for _, param := range endpoint.PathParams {
			if param.HelpText != "" {
				group.Line().Comment(wordwrap.WrapString(param.HelpText, 80))
			}
			group.Id(toExportedName(param.Name)).Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "PATH_PARAMS"}, pq, nil))
		}
		if endpointHasAttribute(endpoint, attrExplicitURL) {
			group.Line().Comment("URL to query. This must be explicitly set for this endpoint and any base URL set in options will be ignored.")
			group.Id("URL").String()
		}
		for _, param := range endpoint.QueryParams {
			if param.HelpText != "" {
				group.Line().Comment(wordwrap.WrapString(param.HelpText, 80))
			}
			group.Id(toExportedName(param.Name)).Op("*").Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "QUERY_PARAMS"}, pq, &paramSchemaFieldTypeOptions{
				usePointers: true,
			}))
		}
		switch {
		case endpointHasAttribute(endpoint, attrJSONRequestBody):
			group.Id("RequestBody").Id(reqBodyStructName(endpoint.ID))
		case endpointHasAttribute(endpoint, attrBodyUploader):
			group.Line().Comment("http request's body")
			group.Id("RequestBody").Qual("io", "Reader")
		}
		for _, param := range endpoint.Headers {
			if param.Name == "accept" {
				continue
			}
			if param.HelpText != "" {
				group.Line().Comment(wordwrap.WrapString(param.HelpText, 80))
			}
			group.Id(toExportedName(param.Name + "-header")).Op("*").Add(paramSchemaFieldType(param.Schema, []string{endpoint.ID, "QUERY_PARAMS"}, pq, &paramSchemaFieldTypeOptions{
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

	for _, fn := range []func(file *jen.File, endpoint *model.Endpoint, pq pkgQual){
		addReqHTTPRequestFunc,
		addReqBuilderFunc,
		reqRelReqFunc,
	} {
		fn(file, endpoint, pq)
		file.Line()
	}
}

func reqRelReqFunc(file *jen.File, endpoint *model.Endpoint, pq pkgQual) {
	structName := reqStructName(endpoint)
	comment := `Rel updates this request to point to a relative link from resp. Returns false if the link does not exist. Handy for paging.`
	file.Comment(wordwrap.WrapString(comment, 80))
	file.Func().Params(jen.Id("r").Id("*"+structName)).Id("Rel").Params(
		jen.Id("link string"),
		jen.Id("resp").Op("*").Id(respStructName(endpoint)),
	).Params(jen.Bool()).Block(
		jen.Id("u := resp.RelLink(string(link))"),
		jen.If(jen.Id("u").Op("==").Lit("")).Block(jen.Return(jen.False())),
		jen.Id("r._url = u"),
		jen.Return(jen.True()),
	)
}

func reqBodyValue(endpoint *model.Endpoint) jen.Code {
	switch {
	case endpointHasAttribute(endpoint, attrJSONRequestBody):
		return jen.Id("r.RequestBody")
	case endpointHasAttribute(endpoint, attrBodyUploader):
		return jen.Id("r.RequestBody")
	default:
		return jen.Nil()
	}
}

func addReqHTTPRequestFunc(file *jen.File, endpoint *model.Endpoint, pq pkgQual) {
	file.Add(reqHTTPRequestFunc(endpoint, pq))
}

func reqHTTPRequestFunc(endpoint *model.Endpoint, pq pkgQual) jen.Code {
	structName := reqStructName(endpoint)
	stmt := jen.Comment("HTTPRequest builds an *http.Request")
	stmt.Line()
	stmt.Func().Params(jen.Id("r").Id("*"+structName)).Id("HTTPRequest").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("opt ...").Qual(pq.pkgPath("options"), "Option"),
	).Params(
		jen.Op("*").Qual("net/http", "Request"), jen.Error(),
	).Block(
		jen.Id("opts, err := ").Qual(pq.pkgPath("options"), "BuildOptions").Call(jen.Id("opt...")),
		jen.Id("if err != nil {return nil, err}"),
		jen.Id("return r.requestBuilder().HTTPRequest(ctx, opts)"),
	)
	return stmt
}

func addReqBuilderFunc(file *jen.File, endpoint *model.Endpoint, pq pkgQual) {
	file.Add(reqBuilderFunc(endpoint, pq))
}

func reqBuilderFunc(endpoint *model.Endpoint, pq pkgQual) jen.Code {
	structName := reqStructName(endpoint)
	return jen.Func().Params(jen.Id("r").Id("*"+structName)).Id("requestBuilder()").Params(
		jen.Op("*").Qual(pq.pkgPath("internal"), "RequestBuilder"),
	).Block(
		reqURLQueryVal(endpoint),
		jen.Id("builder := &").Qual(pq.pkgPath("internal"), "RequestBuilder").Values(
			jen.Dict{
				jen.Id("OperationID"): jen.Lit(endpoint.ID),
				jen.Id("ExplicitURL"): reqExplicitURLVal(endpoint),
				jen.Id("Method"):      jen.Lit(endpoint.Method),
				jen.Id("DataStatuses"): jen.Op("[]").Int().ValuesFunc(func(group *jen.Group) {
					for _, code := range responseCodesWithBodies(endpoint) {
						group.Lit(code)
					}
				}),
				jen.Id("ValidStatuses"): jen.Op("[]").Int().ValuesFunc(func(group *jen.Group) {
					for _, code := range validCodes(endpoint) {
						group.Lit(code)
					}
				}),
				jen.Id("RequiredPreviews"): jen.Op("[]").String().ValuesFunc(func(group *jen.Group) {
					for _, preview := range endpoint.Previews {
						if !preview.Required {
							continue
						}
						group.Add(jen.Lit(preview.Name))
					}
				}),
				jen.Id("AllPreviews"): jen.Op("[]").String().ValuesFunc(func(group *jen.Group) {
					for _, preview := range endpoint.Previews {
						group.Add(jen.Lit(preview.Name))
					}
				}),
				jen.Id("HeaderVals"): reqHeaderMap(endpoint),
				jen.Id("Previews"): jen.Map(jen.String()).Bool().Values(jen.DictFunc(func(dict jen.Dict) {
					for _, preview := range endpoint.Previews {
						dict[jen.Lit(preview.Name)] = jen.Id("r").Dot(toExportedName(preview.Name + "-preview"))
					}
				})),
				jen.Id("Body"):     reqBodyValue(endpoint),
				jen.Id("URLQuery"): jen.Id("query"),
				jen.Id("URLPath"):  reqURLPathVal(endpoint),
			},
		),
		jen.Return(jen.Id("builder")),
	)
}

func validCodes(endpoint *model.Endpoint) []int {
	codes := make([]int, 0, len(endpoint.Responses))
	for code := range endpoint.Responses {
		if code < 400 {
			codes = append(codes, code)
		}
	}
	sort.Ints(codes)
	return codes
}

func reqHeaderMap(endpoint *model.Endpoint) *jen.Statement {
	return jen.Map(jen.String()).Op("*").String().Values(
		jen.DictFunc(func(dict jen.Dict) {
			headers := map[string]*jen.Statement{}
			if endpoint.SuccessMediaType != "" {
				headers["accept"] = jen.Id("String").Call(jen.Lit(endpoint.SuccessMediaType))
			}

			for _, header := range endpoint.Headers {
				if header.Name == "accept" {
					continue
				}
				headers[strings.ToLower(header.Name)] = jen.Id("r").Dot(toExportedName(header.Name + "-header"))
			}

			if headers["content-type"] == nil && endpoint.RequestBody != nil && endpoint.RequestBody.MediaType != "" {
				headers["content-type"] = jen.Id("String").Call(jen.Lit(endpoint.RequestBody.MediaType))
			}

			for k, v := range headers {
				dict[jen.Lit(k)] = v
			}
		}),
	)
}

var bracesExp = regexp.MustCompile(`{[^}]+}`)

func reqURLPathVal(endpoint *model.Endpoint) jen.Code {
	if endpointHasAttribute(endpoint, attrExplicitURL) {
		return jen.Lit("")
	}
	pth := bracesExp.ReplaceAllString(endpoint.Path, "%v")
	return jen.Qual("fmt", "Sprintf").ParamsFunc(func(group *jen.Group) {
		group.Lit(pth)
		for _, param := range endpoint.PathParams {
			group.Id("r").Dot(toExportedName(param.Name))
		}
	})
}

func reqExplicitURLVal(endpoint *model.Endpoint) jen.Code {
	if endpointHasAttribute(endpoint, attrExplicitURL) {
		return jen.Id("r.URL")
	}
	return jen.Id("r._url")
}

func reqURLQueryVal(endpoint *model.Endpoint) jen.Code {
	stmt := jen.Id("query := ").Qual("net/url", "Values").Op("{}")
	stmt.Line()
	for _, param := range endpoint.QueryParams {
		paramArg := jen.Id("r").Dot(toExportedName(param.Name))
		stmt.If(paramArg.Clone().Op("!= nil")).BlockFunc(func(ifGroup *jen.Group) {
			var valStmt *jen.Statement
			switch param.Schema.Type {
			case model.ParamTypeString:
				valStmt = jen.Op("*").Add(paramArg)
			case model.ParamTypeInt:
				valStmt = jen.Qual("strconv", "FormatInt").Params(jen.Op("*").Add(paramArg), jen.Lit(10))
			case model.ParamTypeBool:
				valStmt = jen.Qual("strconv", "FormatBool").Params(jen.Op("*").Add(paramArg))
			default:
				fmt.Println(endpoint.ID)
				fmt.Printf("UNEXPECTED %v, %s\n", param, param.Schema.Type)
			}
			ifGroup.Id("query").Dot("Set").Params(jen.Lit(param.Name), valStmt)
		})
		stmt.Line()
	}
	return stmt
}
