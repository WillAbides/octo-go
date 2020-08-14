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
	endpointID = trimToFirstSlash(endpointID)
	endpointID = strings.ReplaceAll(endpointID, "/", "-")
	return toExportedName(fmt.Sprintf("%s-req-body", endpointID))
}

func reqStructName(endpoint *model.Endpoint) string {
	return toExportedName(fmt.Sprintf("%s-req", endpoint.Name))
}

func addRequestStruct(file *jen.File, pq pkgQual, endpoint *model.Endpoint) {
	structName := reqStructName(endpoint)
	file.Commentf(`%s is request data for Client.%s

%s

Non-nil errors will have the type *requests.RequestError, octo.ResponseError or url.Error.`,
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
		jen.Id("u := ").Id("getRelLink").Call(jen.Id("resp.HTTPResponse(), link")),
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
		return nil
	}
}

func addReqHTTPRequestFunc(file *jen.File, endpoint *model.Endpoint, pq pkgQual) {
	file.Add(reqHTTPRequestFunc(endpoint, pq))
}

func reqHTTPRequestFunc(endpoint *model.Endpoint, pq pkgQual) jen.Code {
	structName := reqStructName(endpoint)
	stmt := jen.Comment("HTTPRequest builds an *http.Request. Non-nil errors will have the type *requests.RequestError.")
	stmt.Line()
	stmt.Func().Params(jen.Id("r").Id("*"+structName)).Id("HTTPRequest").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("opt ...").Qual(pq.pkgPath("requests"), "Option"),
	).Params(
		jen.Op("*").Qual("net/http", "Request"), jen.Error(),
	).Block(
		reqURLQueryVal(endpoint),
		jen.Return(jen.Id("buildHTTPRequest").Call(
			jen.Id("ctx"),
			jen.Id("buildHTTPRequestOptions").Values(
				jen.DictFunc(func(dict jen.Dict) {
					dict[jen.Id("ExplicitURL")] = reqExplicitURLVal(endpoint)
					dict[jen.Id("Method")] = jen.Lit(endpoint.Method)
					if val := requiredPreviews(endpoint); len(val) > 0 {
						dict[jen.Id("RequiredPreviews")] = jen.Id("[]string").ValuesFunc(func(group *jen.Group) {
							for _, preview := range val {
								group.Lit(preview)
							}
						})
					}
					if len(endpoint.Previews) > 0 {
						dict[jen.Id("AllPreviews")] = jen.Op("[]").String().ValuesFunc(func(group *jen.Group) {
							for _, preview := range endpoint.Previews {
								group.Add(jen.Lit(preview.Name))
							}
						})
						dict[jen.Id("Previews")] = jen.Map(jen.String()).Bool().Values(jen.DictFunc(func(dict jen.Dict) {
							for _, preview := range endpoint.Previews {
								dict[jen.Lit(preview.Name)] = jen.Id("r").Dot(toExportedName(preview.Name + "-preview"))
							}
						}))
					}
					if val := reqHeaderMap(endpoint, pq); val != nil {
						dict[jen.Id("HeaderVals")] = val
					}
					if val := reqBodyValue(endpoint); val != nil {
						dict[jen.Id("Body")] = val
					}
					if len(endpoint.QueryParams) > 0 {
						dict[jen.Id("URLQuery")] = jen.Id("query")
					}
					dict[jen.Id("Options")] = jen.Id("opt")
					if val := reqURLPathVal(endpoint); val != nil {
						dict[jen.Id("URLPath")] = val
					}
					if endpointHasAttribute(endpoint, attrExplicitURL) {
						dict[jen.Id("RequireExplicitURL")] = jen.Lit(true)
					}
				}),
			),
		)),
	)
	return stmt
}

func requiredPreviews(endpoint *model.Endpoint) []string {
	result := make([]string, 0, len(endpoint.Previews))
	for _, preview := range endpoint.Previews {
		if preview.Required {
			result = append(result, preview.Name)
		}
	}
	return result
}

func validCodes(endpoint *model.Endpoint) []int {
	if endpointHasAttribute(endpoint, attrBoolean) {
		return []int{204, 404}
	}
	codes := make([]int, 0, len(endpoint.Responses))
	for code := range endpoint.Responses {
		if code < 400 {
			codes = append(codes, code)
		}
	}
	sort.Ints(codes)
	return codes
}

func reqHeaderMap(endpoint *model.Endpoint, pq pkgQual) *jen.Statement {
	var size int
	stmt := jen.Map(jen.String()).Op("*").String().Values(
		jen.DictFunc(func(dict jen.Dict) {
			headers := map[string]*jen.Statement{}
			if endpoint.SuccessMediaType != "" {
				size++
				headers["accept"] = jen.Id("strPtr").Call(jen.Lit(endpoint.SuccessMediaType))
			}

			for _, header := range endpoint.Headers {
				if header.Name == "accept" {
					continue
				}
				size++
				headers[strings.ToLower(header.Name)] = jen.Id("r").Dot(toExportedName(header.Name + "-header"))
			}

			if headers["content-type"] == nil && endpoint.RequestBody != nil && endpoint.RequestBody.MediaType != "" {
				size++
				headers["content-type"] = jen.Id("strPtr").Call(jen.Lit(endpoint.RequestBody.MediaType))
			}

			for k, v := range headers {
				size++
				dict[jen.Lit(k)] = v
			}
		}),
	)
	if size == 0 {
		return nil
	}
	return stmt
}

var bracesExp = regexp.MustCompile(`{[^}]+}`)

func reqURLPathVal(endpoint *model.Endpoint) jen.Code {
	if endpointHasAttribute(endpoint, attrExplicitURL) {
		return nil
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
	if len(endpoint.QueryParams) == 0 {
		return nil
	}
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
