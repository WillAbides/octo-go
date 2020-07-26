package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/willabides/octo-go/generator/internal/model"
)

// the attributes here will be added for endpoints with matching IDs
var overrideAddAttrs = map[string][]endpointAttribute{
	"repos/upload-release-asset": {attrExplicitURL},
	"projects/move-card":         {attrNoResponseBody},
	"projects/move-column":       {attrNoResponseBody},
	"markdown/render-raw":        {attrBodyUploader},
}

func endpointWithOverrides(endpoint model.Endpoint) (model.Endpoint, error) {
	ptr := &endpoint
	for _, override := range endpointOverrides {
		err := override.override(ptr)
		if err != nil {
			return model.Endpoint{}, err
		}
	}

	if ptr.RequestBody != nil {
		err := oneOfCheck(ptr.ID, ptr.RequestBody.Schema)
		if err != nil {
			return model.Endpoint{}, err
		}
	}

	return *ptr, nil
}

var oneOfOverrides = []*oneOfOverride{
	{
		endpointID: "gists/create",
		paramName:  "public",
		oneOfPicker: func(param *model.Param) bool {
			return param.Schema.Type == model.ParamTypeBool
		},
	},
	{
		endpointID: "projects/create-card",
		paramName:  "",
		oneOfPicker: func(param *model.Param) bool {
			return len(param.Schema.ObjectParams) == 2
		},
	},
	{
		endpointID: "users/add-email-for-authenticated",
		paramName:  "",
		oneOfPicker: func(param *model.Param) bool {
			return param.Schema.Type == model.ParamTypeObject
		},
	},
	{
		endpointID: "users/delete-email-for-authenticated",
		paramName:  "",
		oneOfPicker: func(param *model.Param) bool {
			return param.Schema.Type == model.ParamTypeObject
		},
	},
	{
		endpointID: "scim/update-attribute-for-user",
		paramName:  "Operations/ItemSchema/value",
		oneOfPicker: func(param *model.Param) bool {
			return param.Schema.Type == model.ParamTypeObject
		},
	},
}

type endpointOverride struct {
	endpointID        string
	endpointAttribute *endpointAttribute
	fn                func(endpoint *model.Endpoint) error
}

func (e *endpointOverride) override(endpoint *model.Endpoint) error {
	if e.endpointID != "" && endpoint.ID != e.endpointID {
		return nil
	}
	if e.endpointAttribute != nil && !endpointHasAttribute(*endpoint, *e.endpointAttribute) {
		return nil
	}
	return e.fn(endpoint)
}

var endpointOverrides = []*endpointOverride{
	// Exclude the "exclude" query param until I write support for array query params.
	// "exclude" isn't documented, so I assume it isn't crucial.
	{
		endpointID: "migrations/get-status-for-authenticated-user",
		fn: func(endpoint *model.Endpoint) error {
			params := endpoint.QueryParams
			var excludeIdx *int
			for i := range params {
				param := params[i]
				if param.Name == "exclude" {
					excludeIdx = &i
					break
				}
			}
			if excludeIdx == nil {
				return nil
			}
			params = append(params[:*excludeIdx], params[*excludeIdx+1:]...)
			endpoint.QueryParams = params
			return nil
		},
	},

	// force labels to be strings in issues/update
	{
		endpointID: "issues/update",
		fn: func(endpoint *model.Endpoint) error {
			params := endpoint.RequestBody.Schema.ObjectParams
			for i := range params {
				param := params[i]
				if param.Name != "labels" {
					continue
				}
				itemSchema := param.Schema.ItemSchema.Clone()
				itemSchema.Type = model.ParamTypeString
				itemSchema.ObjectParams = nil
				param.Schema.ItemSchema = itemSchema
				params[i] = param
				break
			}
			return nil
		},
	},

	// force labels to be strings in issues/create
	{
		endpointID: "issues/create",
		fn: func(endpoint *model.Endpoint) error {
			params := endpoint.RequestBody.Schema.ObjectParams
			for i := range params {
				param := params[i]
				if param.Name != "labels" {
					continue
				}
				itemSchema := param.Schema.ItemSchema.Clone()
				itemSchema.Type = model.ParamTypeString
				itemSchema.ObjectParams = nil
				param.Schema.ItemSchema = itemSchema
				params[i] = param
				break
			}
			return nil
		},
	},

	// run all oneOfOverrides
	{
		fn: func(endpoint *model.Endpoint) error {
			for _, override := range oneOfOverrides {

				err := override.override(endpoint)
				if err != nil {
					return err
				}
			}
			return nil
		},
	},
	// repos/upload-release-asset requires a Content-Type header
	{
		endpointID: "repos/upload-release-asset",
		fn: func(endpoint *model.Endpoint) error {
			endpoint.Headers = append(endpoint.Headers, &model.Param{
				Required: true,
				Name:     "Content-Type",
				HelpText: "Content-Type for the uploaded file",
				Schema: &model.ParamSchema{
					Type: model.ParamTypeString,
				},
			})
			return nil
		},
	},

	// endpoints with attrExplicitURL should have no PathParams
	{
		endpointAttribute: attrExplicitURL.pointer(),
		fn: func(endpoint *model.Endpoint) error {
			endpoint.PathParams = []*model.Param{}
			return nil
		},
	},
}

var nameOverrides = map[string]string{
	"+1": "plus-one",
	"-1": "minus-one",
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

type oneOfOverride struct {
	endpointID  string
	paramName   string
	oneOfPicker func(param *model.Param) bool
}

func oneOfCheck(name string, schema *model.ParamSchema) error {
	if schema == nil {
		return nil
	}
	if schema.Type == model.ParamTypeOneOf {
		return errors.Errorf("%s has a oneOf schema", name)
	}
	err := oneOfCheck(path.Join(name, "ItemSchema"), schema.ItemSchema)
	if err != nil {
		return err
	}
	for i := range schema.ObjectParams {
		param := schema.ObjectParams[i]
		err = oneOfCheck(path.Join(name, param.Name), param.Schema)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *oneOfOverride) overrideParam(namePrefix string, param *model.Param) (*model.Param, error) {
	param = param.Clone()
	paramName := param.Name
	if namePrefix != "" {
		paramName = path.Join(namePrefix, paramName)
	}
	if paramName == o.paramName {
		if param.Schema.Type != model.ParamTypeOneOf {
			return nil, errors.New("param matches but isn't a oneOf")
		}
		var matches []*model.Param
		for i := range param.Schema.ObjectParams {
			oneOfParam := param.Schema.ObjectParams[i]
			if o.oneOfPicker(oneOfParam) {
				matches = append(matches, oneOfParam)
			}
		}
		if len(matches) == 0 || len(matches) > 1 {
			return nil, errors.New("oneOf picker didn't match exactly one param")
		}
		match := matches[0]
		param.Schema = match.Schema.Clone()
		if match.HelpText != "" {
			param.HelpText = match.HelpText
		}
	}
	params := param.Schema.ObjectParams
	for i := range params {
		pp, err := o.overrideParam(paramName, params[i])
		if err != nil {
			return nil, errors.WithStack(err)
		}
		params[i] = pp
	}
	if param.Schema.ItemSchema != nil {
		for i := range param.Schema.ItemSchema.ObjectParams {
			pp, err := o.overrideParam(path.Join(paramName, "ItemSchema"), param.Schema.ItemSchema.ObjectParams[i])
			if err != nil {
				return nil, errors.WithStack(err)
			}
			param.Schema.ItemSchema.ObjectParams[i] = pp
		}
	}
	return param, nil
}

func (o *oneOfOverride) override(endpoint *model.Endpoint) error {
	if endpoint.ID != o.endpointID {
		return nil
	}
	if o.paramName == "" {
		objectParams := endpoint.RequestBody.Schema.ObjectParams
		for i := 0; i < len(objectParams); i++ {
			param := objectParams[i]
			if !o.oneOfPicker(param) {
				continue
			}
			endpoint.RequestBody.Schema = param.Schema.Clone()
			break
		}
	}
	params := endpoint.RequestBody.Schema.ObjectParams
	for i := range params {
		p, err := o.overrideParam("", params[i])
		if err != nil {
			return err
		}
		params[i] = p
	}
	return nil
}
