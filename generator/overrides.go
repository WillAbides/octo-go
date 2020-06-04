package main

import (
	"fmt"
	"strings"

	"github.com/willabides/octo-go/generator/internal/model"
)

// the attributes here will be added for endpoints with matching IDs
var overrideAddAttrs = map[string][]endpointAttribute{
	"teams/check-manages-repo-in-org": {attrBoolean},
	"repos/upload-release-asset":      {attrExplicitURL},
}

func endpointWithOverrides(endpoint model.Endpoint) model.Endpoint {
	ptr := &endpoint
	for _, override := range endpointOverrides {
		override(ptr)
	}
	return *ptr
}

var endpointOverrides = []func(endpoint *model.Endpoint){

	// markdown/render-raw doesn't take a json body
	func(endpoint *model.Endpoint) {
		if endpoint.ID != "markdown/render-raw" {
			return
		}
		endpoint.Requests = []model.Request{
			{
				MimeType: "*/*",
				Schema: &model.ParamSchema{
					Type: model.ParamTypeString,
				},
			},
		}
	},

	// list-languages returns a map
	func(endpoint *model.Endpoint) {
		if endpoint.ID != "repos/list-languages" {
			return
		}
		respSchema := endpoint.Responses[200].Body
		respSchema.Ref = ""
		respSchema.ObjectParams = nil
		respSchema.ItemSchema = &model.ParamSchema{
			Type: model.ParamTypeInt,
		}
	},

	// attrExplicitURL has no path variables
	func(endpoint *model.Endpoint) {
		if !endpointHasAttribute(*endpoint, attrExplicitURL) {
			return
		}
		endpoint.PathParams = model.Params{}
	},
}

func schemaPathString(schemaPath []string) string {
	return strings.Join(schemaPath, "/")
}

var schemaOverrides = []func(schemaPath []string, schema *model.ParamSchema){

	// apps/create-installation-token/reqBody/permissions is map[string]string
	func(schemaPath []string, schema *model.ParamSchema) {
		if schemaPathString(schemaPath) != "apps/create-installation-token/reqBody/permissions" {
			return
		}
		schema.ItemSchema = &model.ParamSchema{
			Type: model.ParamTypeString,
		}
		fmt.Println(schemaPathString(schemaPath))
	},

	// permissions are maps
	func(schemaPath []string, schema *model.ParamSchema) {
		if !strings.HasSuffix(schemaPathString(schemaPath), "/permissions") || schema.Type != model.ParamTypeObject {
			return
		}
		if len(schema.ObjectParams) == 0 {
			return
		}
		for i := 1; i < len(schema.ObjectParams); i++ {
			if schema.ObjectParams[i].Schema.Type != schema.ObjectParams[i-1].Schema.Type {
				panic(fmt.Sprintf("%s violates the assumption that permissions maps will only have one type", schemaPathString(schemaPath)))
			}
		}
		schema.ItemSchema = &model.ParamSchema{
			Type: schema.ObjectParams[0].Schema.Type,
		}
		schema.ObjectParams = nil
	},

	// reactions are maps
	func(schemaPath []string, schema *model.ParamSchema) {
		if !strings.HasSuffix(schemaPathString(schemaPath), "/reactions") || schema.Type != model.ParamTypeObject {
			return
		}
		var found bool
		for _, objectParam := range schema.ObjectParams {
			if objectParam.Name == "+1" {
				found = true
				break
			}
		}
		if !found {
			return
		}
		schema.ObjectParams = nil
		schema.ItemSchema = &model.ParamSchema{
			Type: model.ParamTypeInt,
		}
	},
	// a lot of numbers should be integers
	func(schemaPath []string, schema *model.ParamSchema) {
		if schema.Type != model.ParamTypeNumber {
			return
		}
		suffixes := []string{
			"_count",
			"/count",
			"/id",
			"_id",
			"price_in_cents",
			"/comments",
			"/commits",
			"/total_private_repos",
			"/total_commits",
			"/total_ms",
			"/total",
			"/totalResults",
			"/additions",
			"/deletions",
			"_issues",
			"/line",
			"_line",
			"/number",
			"_number",
			"_repos",
			"/startIndex",
			"_position",
			"/position",
			"/jobs",
			"/ahead_by",
			"/behind_by",
			"/changed_files",
			"/changes",
			"/collaborators",
			"/contributions",
			"/duration",
			"/uniques",
			"/week",
			"/limit",
			"/itemsPerPage",
			"_gists",
			"/followers",
			"/remaining",
			"/following",
			"_column",
			"/reset",
			"/size_in_bytes",
			"_comments",
			"/run_duration_ms",
			"weeks/ITEM_SCHEMA/a",
			"weeks/ITEM_SCHEMA/c",
			"weeks/ITEM_SCHEMA/d",
			"repos/get-punch-card-stats/responseBody/200/ITEM_SCHEMA/ITEM_SCHEMA",
			"repos/get-code-frequency-stats/responseBody/200/ITEM_SCHEMA/ITEM_SCHEMA",
			"repos/get-participation-stats/responseBody/200/all/ITEM_SCHEMA",
			"repos/get-participation-stats/responseBody/200/owner/ITEM_SCHEMA",
			"repos/get-commit-activity-stats/responseBody/200/ITEM_SCHEMA/days/ITEM_SCHEMA",
		}
		sPath := schemaPathString(schemaPath)
		for _, suffix := range suffixes {
			if strings.HasSuffix(sPath, suffix) {
				schema.Type = model.ParamTypeInt
				return
			}
		}
	},
}

func overrideParamSchema(schemaPath []string, schema *model.ParamSchema) {
	if schema == nil {
		return
	}
	for _, override := range schemaOverrides {
		override(schemaPath, schema)
	}
}

func fixPreviewNote(note string) string {
	note = strings.TrimSpace(note)
	note = strings.Split(note, "```")[0]
	note = strings.TrimSpace(note)
	setThisFlagPhrases := []string{
		"provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header",
		"provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header",
		"provide the following custom [media type](https://developer.github.com/v3/media) in the `Accept` header",
	}
	for _, phrase := range setThisFlagPhrases {
		note = strings.ReplaceAll(note, phrase, "set this to true")
	}
	note = strings.TrimSpace(note)
	note = strings.TrimSuffix(note, ":")
	note = strings.TrimSpace(note)
	note = strings.TrimSuffix(note, ".") + "."
	return note
}
