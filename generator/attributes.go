package main

import (
	"sort"
	"strings"

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
	// endpoint that shouldn't have any response body
	attrNoResponseBody
	// attrInvalid is last so we can get a list of all valid types with a for loop
	attrInvalid
)

var attrNames = map[endpointAttribute]string{
	attrRedirectOnly:    "AttrRedirectOnly",
	attrBoolean:         "AttrBoolean",
	attrBodyUploader:    "AttrBodyUploader",
	attrJSONRequestBody: "AttrJSONRequestBody",
	attrExplicitURL:     "AttrExplicitURL",
	attrNoResponseBody:  "AttrNoResponseBody",
}

func (e endpointAttribute) pointer() *endpointAttribute {
	return &e
}

func (e endpointAttribute) String() string {
	return attrNames[e]
}

func endpointHasAttribute(endpoint *model.Endpoint, attribute endpointAttribute) bool {
	for _, attr := range getEndpointAttributes(endpoint) {
		if attribute == attr {
			return true
		}
	}
	return false
}

func getEndpointAttributes(endpoint *model.Endpoint) []endpointAttribute {
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

type attrCheck func(endpoint *model.Endpoint, attrs *[]endpointAttribute)

var attrChecks = []attrCheck{
	// attrJSONRequestBody if the endpoint has an application/json request
	func(endpoint *model.Endpoint, attrs *[]endpointAttribute) {
		if endpoint.RequestBody == nil {
			return
		}
		if strings.HasSuffix(endpoint.RequestBody.MediaType, "json") {
			*attrs = append(*attrs, attrJSONRequestBody)
		}
	},

	// attrBodyUploader if the endpoint has a */* request
	func(endpoint *model.Endpoint, attrs *[]endpointAttribute) {
		if endpoint.RequestBody == nil {
			return
		}
		if endpoint.RequestBody.MediaType == "*/*" {
			*attrs = append(*attrs, attrBodyUploader)
		}
	},

	// attrBoolean if the endpoint has exatcly two responses: 204 and 404
	func(endpoint *model.Endpoint, attrs *[]endpointAttribute) {
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

	// attrRedirectOnly if the endpoint has only one response: 302
	func(endpoint *model.Endpoint, attrs *[]endpointAttribute) {
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
