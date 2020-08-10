package model

type Endpoint struct {
	GithubCloudOnly  bool
	EnabledForApps   bool
	Legacy           bool
	Deprecated       bool
	ID               string
	Path             string
	Method           string
	Name             string
	Concern          string
	DocsURL          string
	Summary          string
	HelpText         string
	PathParams       []*Param
	QueryParams      []*Param
	Headers          []*Param
	Previews         []*Preview
	RequestBody      *RequestBody
	Responses        map[int]*Response
	SuccessMediaType string
}

func (e *Endpoint) Clone() *Endpoint {
	result := new(Endpoint)
	*result = *e
	result.PathParams = cloneParams(result.PathParams)
	result.QueryParams = cloneParams(result.QueryParams)
	result.Headers = cloneParams(result.Headers)
	result.Previews = clonePreviews(result.Previews)
	result.RequestBody = result.RequestBody.Clone()
	result.Responses = cloneResponseMap(result.Responses)
	return result
}

type RequestBody struct {
	MediaType string
	Schema    *ParamSchema
}

func (b *RequestBody) Clone() *RequestBody {
	if b == nil {
		return nil
	}
	result := new(RequestBody)
	*result = *b
	result.Schema = result.Schema.Clone()
	return result
}

type Response struct {
	MediaType  string
	Body       *ParamSchema
	HasExample bool
}

func (r *Response) Clone() *Response {
	result := new(Response)
	*result = *r
	result.Body = result.Body.Clone()
	return result
}

func cloneResponseMap(responses map[int]*Response) map[int]*Response {
	if responses == nil {
		return nil
	}
	result := make(map[int]*Response, len(responses))
	for k, v := range responses {
		result[k] = v.Clone()
	}
	return result
}

type Preview struct {
	Required bool
	Name     string
	Note     string
}

func clonePreviews(previews []*Preview) []*Preview {
	result := make([]*Preview, len(previews))
	for i := range previews {
		preview := *previews[i]
		result[i] = &preview
	}
	return result
}

type ParamType int

const (
	ParamTypeInvalid ParamType = iota
	ParamTypeString
	ParamTypeInt
	ParamTypeBool
	ParamTypeNumber
	ParamTypeInterface
	ParamTypeObject
	ParamTypeArray
	ParamTypeOneOf
)

func (pt ParamType) String() string {
	switch pt {
	case ParamTypeString:
		return "string"
	case ParamTypeInt:
		return "int"
	case ParamTypeBool:
		return "bool"
	case ParamTypeNumber:
		return "number"
	case ParamTypeInterface:
		return "interface"
	case ParamTypeObject:
		return "object"
	case ParamTypeArray:
		return "array"
	case ParamTypeOneOf:
		return "oneOf"
	}
	return "invalid"
}

type ParamSchema struct {
	Nullable     bool
	Ref          string
	Type         ParamType
	ItemSchema   *ParamSchema
	ObjectParams []*Param
}

func (p *ParamSchema) Clone() *ParamSchema {
	if p == nil {
		return nil
	}
	result := ParamSchema{
		Type: p.Type,
		Ref:  p.Ref,
	}
	if p.ObjectParams != nil {
		result.ObjectParams = make([]*Param, len(p.ObjectParams))
		for i, objectParam := range p.ObjectParams {
			result.ObjectParams[i] = objectParam.Clone()
		}
	}
	if p.ItemSchema != nil {
		result.ItemSchema = p.ItemSchema.Clone()
	}
	return &result
}

type Param struct {
	Required bool
	Name     string
	HelpText string
	Schema   *ParamSchema
}

func (p *Param) Clone() *Param {
	q := new(Param)
	*q = *p
	q.Schema = q.Schema.Clone()
	return q
}

func cloneParams(params []*Param) []*Param {
	if params == nil {
		return nil
	}
	result := make([]*Param, len(params))
	for i := range params {
		result[i] = params[i].Clone()
	}
	return result
}
