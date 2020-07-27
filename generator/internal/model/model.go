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

type RequestBody struct {
	MediaType string
	Schema    *ParamSchema
}

type Response struct {
	MediaType  string
	Body       *ParamSchema
	HasExample bool
}

type Preview struct {
	Required bool
	Name     string
	Note     string
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
