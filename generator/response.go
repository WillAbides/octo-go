package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/google/go-cmp/cmp"
	"github.com/willabides/octo-go/generator/internal/model"
)

func respBodyStructName(endpoint *model.Endpoint) string {
	return toExportedName(fmt.Sprintf("%s-%s-response-body", endpoint.Concern, endpoint.Name))
}

func respStructName(endpoint *model.Endpoint) string {
	return toExportedName(fmt.Sprintf("%s-%s-response", endpoint.Concern, endpoint.Name))
}

func sortedResponseCodes(endpoint *model.Endpoint) []int {
	sortedCodes := make([]int, 0, len(endpoint.Responses))
	for code := range endpoint.Responses {
		if code < 300 {
			sortedCodes = append(sortedCodes, code)
		}
	}
	sort.Ints(sortedCodes)
	return sortedCodes
}

func addResponse(file *jen.File, pq pkgQual, endpoint *model.Endpoint) {
	structName := respStructName(endpoint)
	file.Commentf("%s is a response for %s\n\n%s",
		structName,
		toExportedName(endpoint.ID),
		endpoint.DocsURL,
	)
	file.Type().Id(structName).StructFunc(func(group *jen.Group) {
		group.Qual(pq.pkgPath("common"), "Response")
		group.Id("request").Op("*").Id(reqStructName(endpoint))
		if endpointHasAttribute(endpoint, attrNoResponseBody) {
			return
		}
		bodyType := respBodyType(endpoint, pq)
		if bodyType != nil {
			group.Id("Data").Do(func(stmt *jen.Statement) {
				if bodyType.slice {
					stmt.Op("[]")
				}
			}).Qual(bodyType.pkg, bodyType.name)
		}
		if endpointHasAttribute(endpoint, attrBoolean) {
			group.Id("Data").Bool()
		}
	})
}

func respBodyType(endpoint *model.Endpoint, pq pkgQual) *qualifiedType {
	codeBodies := responseCodesWithBodies(endpoint)
	if len(codeBodies) == 0 {
		return nil
	}
	body := endpoint.Responses[codeBodies[0]].Body
	if body.Type == model.ParamTypeArray {
		if strings.HasPrefix(body.ItemSchema.Ref, "#/components/schemas/") {
			nm := strings.TrimPrefix(body.ItemSchema.Ref, "#/components/schemas/")
			return &qualifiedType{
				pkg:   pq.pkgPath("components"),
				name:  toExportedName(nm),
				slice: true,
			}
		}
	}

	if strings.HasPrefix(body.Ref, "#/components/schemas/") {
		nm := strings.TrimPrefix(body.Ref, "#/components/schemas/")
		return &qualifiedType{
			pkg:  pq.pkgPath("components"),
			name: toExportedName(nm),
		}
	}
	return &qualifiedType{
		pkg:  pq.pkgPath("octo"),
		name: toExportedName(fmt.Sprintf("%s-%s-response-body", endpoint.Concern, endpoint.Name)),
	}
}

func responseCodesWithBodies(endpoint *model.Endpoint) []int {
	sortedCodes := sortedResponseCodes(endpoint)
	bodyCodes := make([]int, 0, len(sortedCodes))
	for _, respCode := range sortedCodes {
		resp := endpoint.Responses[respCode]
		if resp.Body == nil {
			continue
		}

		// skip 202 Accepted bodies where the body is just a message field
		if len(bodyCodes) > 0 && respCode == 202 {
			if len(resp.Body.ObjectParams) == 1 && resp.Body.ObjectParams[0].Name == "message" {
				continue
			}
		}
		bodyCodes = append(bodyCodes, respCode)
	}
	for i := 1; i < len(bodyCodes); i++ {
		if !cmp.Equal(endpoint.Responses[bodyCodes[0]].Body, endpoint.Responses[bodyCodes[i]].Body) {
			panic(fmt.Sprintf("%s has broken our assumption that success bodies will all be equal", endpoint.ID))
		}
	}
	return bodyCodes
}

func addResponseBody(file *jen.File, pq pkgQual, endpoint *model.Endpoint) {
	if endpointHasAttribute(endpoint, attrNoResponseBody) {
		return
	}
	bt := respBodyType(endpoint, pq)
	if bt != nil && bt.pkg == "github.com/willabides/octo-go/components" {
		return
	}
	bodyCodes := responseCodesWithBodies(endpoint)
	if len(bodyCodes) == 0 {
		return
	}
	respCode := bodyCodes[0]
	resp := endpoint.Responses[respCode]
	tp := paramSchemaFieldType(resp.Body, []string{endpoint.ID, "responseBody"}, pq, nil)
	if tp == nil {
		return
	}
	structName := respBodyStructName(endpoint)
	file.Commentf("%s is a response body for %s\n\n%s",
		structName,
		toExportedName(endpoint.ID),
		endpoint.DocsURL,
	)
	file.Type().Id(respBodyStructName(endpoint)).Add(tp)

	if resp.Body.Type == model.ParamTypeOneOf {
		file.Line()
		file.Add(oneOfValueFunc(structName, pq, resp.Body))
		file.Line()
		file.Add(oneOfSetValueFunc(structName, pq, resp.Body))
		file.Line()
		file.Add(oneOfMarshalJSONFunc(structName, resp.Body))
		file.Line()
		file.Add(oneOfUnmarshalJSONFunc(structName, resp.Body))
		file.Line()
	}
}
