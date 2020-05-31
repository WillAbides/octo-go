package main

import (
	"strings"

	"github.com/willabides/octo-go/generator/internal/model"
)

var overrideEndpointTypes = map[string]endpointAttribute{
	"teams/check-manages-repo-in-org": attrBoolean,
}

func schemaPathString(schemaPath []string) string {
	return strings.Join(schemaPath, "/")
}

var schemaOverrides = []func(schemaPath []string, schema *model.ParamSchema){
	// list-languages returns a map
	func(schemaPath []string, schema *model.ParamSchema) {
		if schemaPathString(schemaPath) != "repos/list-languages/responseBody/200" {
			return
		}
		schema.Type = model.ParamTypeObject
		schema.ObjectParams = nil
		schema.ItemSchema = &model.ParamSchema{
			Type: model.ParamTypeInt,
		}
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
