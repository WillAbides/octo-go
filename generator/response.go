package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model"
)

func respStructName(endpoint model.Endpoint, code int) string {
	return toArgName(fmt.Sprintf("%s-%s-response-body%d", endpoint.Concern, endpoint.Name, code))
}

func addResponseBodies(file *jen.File, endpoint model.Endpoint) {
	if len(endpoint.Responses) == 0 {
		return
	}
	sortedCodes := make([]int, 0, len(endpoint.Responses))
	for code := range endpoint.Responses {
		if code < 300 {
			sortedCodes = append(sortedCodes, code)
		}
	}
	sort.Ints(sortedCodes)
	for _, respCode := range sortedCodes {
		schema := endpoint.Responses[respCode]
		tp := paramSchemaFieldType(schema, []string{endpoint.ID, "responseBody", strconv.Itoa(respCode)}, false, false)
		if tp == nil {
			continue
		}
		structName := respStructName(endpoint, respCode)
		file.Commentf("%s is a response body for %s\n\nAPI documentation: %s",
			structName,
			endpoint.ID,
			endpoint.DocsURL,
		)
		file.Type().Id(respStructName(endpoint, respCode)).Add(tp)
	}
}
