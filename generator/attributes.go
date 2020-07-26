package main

import (
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model"
)

type endpointAttribute int

const (
	// endpoint that redirects via 302
	attrRedirectOnly endpointAttribute = iota
	// gives boolean status through 204 vs 404 status codes
	attrBoolean
	// endpoint where the request body is a reader to be uploaded instead of something to be marshalled to JSON
	attrBodyUploader
	// endpoint with a json request body
	attrJSONRequestBody
	// requires a URL parameter to be set explicitly
	attrExplicitURL
	// endpoint that may need some coercing to return an array response
	attrForceArrayResponse
	// endpoint that shouldn't have any response body
	attrNoResponseBody
	// attrInvalid is last so we can get a list of all valid types with a for loop
	attrInvalid
)

var attrNames = map[endpointAttribute]string{
	attrRedirectOnly:       "attrRedirectOnly",
	attrBoolean:            "attrBoolean",
	attrBodyUploader:       "attrBodyUploader",
	attrJSONRequestBody:    "attrJSONRequestBody",
	attrExplicitURL:        "attrExplicitURL",
	attrForceArrayResponse: "attrForceArrayResponse",
	attrNoResponseBody:     "attrNoResponseBody",
}

func (e endpointAttribute) pointer() *endpointAttribute {
	return &e
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
	var result []endpointAttribute
	override, ok := overrideAddAttrs[endpoint.ID]
	if ok {
		result = append(result, override...)
	}
	for _, check := range attrChecks {
		check(endpoint, &result)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

type attrCheck func(endpoint model.Endpoint, attrs *[]endpointAttribute)

var attrChecks = []attrCheck{
	// attrJSONRequestBody if the endpoint has an application/json request
	func(endpoint model.Endpoint, attrs *[]endpointAttribute) {
		if endpoint.RequestBody == nil {
			return
		}
		if strings.HasSuffix(endpoint.RequestBody.MediaType, "json") {
			*attrs = append(*attrs, attrJSONRequestBody)
		}
	},

	// attrBodyUploader if the endpoint has a */* request
	func(endpoint model.Endpoint, attrs *[]endpointAttribute) {
		if endpoint.RequestBody == nil {
			return
		}
		if endpoint.RequestBody.MediaType == "*/*" {
			*attrs = append(*attrs, attrBodyUploader)
		}
	},

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
