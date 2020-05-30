package main

import (
	"fmt"
	"sort"

	"github.com/dave/jennifer/jen"
	"github.com/google/go-cmp/cmp"
	"github.com/willabides/octo-go/generator/internal/model"
)

func respBodyStructName(endpoint model.Endpoint) string {
	return toExportedName(fmt.Sprintf("%s-%s-response-body", endpoint.Concern, endpoint.Name))
}

func respStructName(endpoint model.Endpoint) string {
	return toExportedName(fmt.Sprintf("%s-%s-response", endpoint.Concern, endpoint.Name))
}

func sortedResponseCodes(endpoint model.Endpoint) []int {
	sortedCodes := make([]int, 0, len(endpoint.Responses))
	for code := range endpoint.Responses {
		if code < 300 {
			sortedCodes = append(sortedCodes, code)
		}
	}
	sort.Ints(sortedCodes)
	return sortedCodes
}

func addResponse(file *jen.File, endpoint model.Endpoint) {
	structName := respStructName(endpoint)
	file.Commentf("%s is a response for %s\n\n%s",
		structName,
		toExportedName(endpoint.ID),
		endpoint.DocsURL,
	)
	file.Type().Id(structName).StructFunc(func(group *jen.Group) {
		group.Id("response")
		group.Id("request").Op("*").Id(reqStructName(endpoint))
		switch {
		case len(responseCodesWithBodies(endpoint)) > 0:
			group.Id("Data").Id("*" + respBodyStructName(endpoint))
		case getEndpointType(endpoint) == endpointTypeBoolean:
			group.Id("Data").Bool()
		}
	})
}

func responseCodesWithBodies(endpoint model.Endpoint) []int {
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
		if !cmp.Equal(endpoint.Responses[bodyCodes[0]], endpoint.Responses[bodyCodes[i]]) {
			panic(fmt.Sprintf("%s has broken our assumption that success bodies will all be equal", endpoint.ID))
		}
	}
	return bodyCodes
}

func addResponseBody(file *jen.File, endpoint model.Endpoint) {
	bodyCodes := responseCodesWithBodies(endpoint)
	if len(bodyCodes) == 0 {
		return
	}
	respCode := bodyCodes[0]
	resp := endpoint.Responses[respCode]
	tp := paramSchemaFieldType(resp.Body, []string{endpoint.ID, "responseBody"}, nil)
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
}
