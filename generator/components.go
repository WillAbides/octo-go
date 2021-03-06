package main

import (
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model"
)

func compSchemaRefStmt(schema *model.ParamSchema, pq pkgQual) *jen.Statement {
	if !strings.HasPrefix(schema.Ref, "#/components/schemas/") {
		return nil
	}
	nm := strings.TrimPrefix(schema.Ref, "#/components/schemas/")
	return jen.Qual(pq.pkgPath("components"), toExportedName(nm))
}

func addComponentSchemas(file *jen.File, schemas map[string]*model.ParamSchema, pq pkgQual) {
	names := make([]string, 0, len(schemas))
	for name := range schemas {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		schema := schemas[name]
		structName := toExportedName(name)
		tp := paramSchemaFieldType(schema, []string{"components", "schemas", name}, pq, &paramSchemaFieldTypeOptions{
			noHelperRecursive: true,
		})
		file.Type().Id(structName).Add(tp)
		file.Line()
		if schema.Type == model.ParamTypeOneOf {
			file.Line()
			file.Add(oneOfValueFunc(structName, pq, schema))
			file.Line()
			file.Add(oneOfSetValueFunc(structName, pq, schema))
			file.Line()
			file.Add(oneOfMarshalJSONFunc(structName, schema))
			file.Line()
			file.Add(oneOfUnmarshalJSONFunc(structName, schema))
			file.Line()
		}
	}
}

func oneOfMarshalJSONFunc(structName string, schema *model.ParamSchema) jen.Code {
	return jen.Func().Params(
		jen.Id("c").Op("*").Id(structName),
	).Id("MarshalJSON() ([]byte, error)").Block(
		jen.Switch(jen.Id("c.oneOfField")).BlockFunc(func(group *jen.Group) {
			for _, param := range schema.ObjectParams {
				paramName := toUnexportedName(param.Name)
				group.Case(jen.Lit(paramName))
				group.Return(jen.Qual("encoding/json", "Marshal").Call(jen.Id("&c").Dot(paramName)))
			}
		}),
		jen.Return(jen.Qual("encoding/json", "Marshal").Call(jen.Id("interface{}(nil)"))),
	)
}

func oneOfUnmarshalJSONFunc(structName string, schema *model.ParamSchema) jen.Code {
	return jen.Func().Params(jen.Id("c").Op("*").Id(structName)).Id("UnmarshalJSON(data []byte) error").BlockFunc(
		func(group *jen.Group) {
			group.Id("var err error")
			for _, param := range schema.ObjectParams {
				paramName := toUnexportedName(param.Name)
				group.Id("err").Op("=").Qual("encoding/json", "Unmarshal").Call(
					jen.Id("data"),
					jen.Id("&c").Dot(paramName),
				)
				group.If(jen.Id("err == nil")).Block(
					jen.Id("c.oneOfField =").Lit(paramName),
					jen.Return(jen.Nil()),
				)
			}
			group.Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("could not unmarshal json")))
		},
	)
}

func orList(vals []string) string {
	var result string
	for i := range vals {
		result += vals[i]
		switch {
		case i < len(vals)-2:
			result += ", "
		case i < len(vals)-1:
			result += " or "
		}
	}

	return result
}

func oneOfReturnTypeForComment(name string, pq pkgQual, schema *model.ParamSchema) string {
	return oneOfQualifiedType(name, pq, schema).commentString()
}

type qualifiedType struct {
	pkg   string
	name  string
	slice bool
}

func (q *qualifiedType) commentString() string {
	result := ""
	if q.slice {
		result += "[]"
	}
	if q.pkg != "" {
		parts := strings.Split(q.pkg, "/")
		result += parts[len(parts)-1] + "."
	}
	result += q.name
	return result
}

func (q *qualifiedType) jenType(pq pkgQual) *jen.Statement {
	nm := jen.Id(q.name)
	if q.pkg != "" {
		nm = jen.Qual(pq.pkgPath(q.pkg), q.name)
	}
	if !q.slice {
		return nm
	}
	return jen.Op("[]").Add(nm)
}

func oneOfQualifiedType(name string, pq pkgQual, schema *model.ParamSchema) *qualifiedType {
	switch schema.Type {
	case model.ParamTypeInt:
		return &qualifiedType{
			name: "int64",
		}
	case model.ParamTypeNumber:
		return &qualifiedType{
			name: "string",
		}
	case model.ParamTypeInterface:
		return &qualifiedType{
			name: "interface{}",
		}
	case model.ParamTypeObject, model.ParamTypeOneOf:
		pkg := "octo"
		if strings.HasPrefix(schema.Ref, "#/components") {
			pkg = "components"
		}
		return &qualifiedType{
			pkg:  pkg,
			name: toExportedName(name),
		}
	case model.ParamTypeArray:
		if schema.ItemSchema.Type == model.ParamTypeObject &&
			len(schema.ItemSchema.ObjectParams) == 0 {
			return &qualifiedType{
				name:  "interface{}",
				slice: true,
			}
		}
		rv := oneOfQualifiedType(name, pq, schema.ItemSchema)
		if schema.ItemSchema.Ref == "" {
			if strings.HasPrefix(schema.Ref, "#/components") {
				rv.pkg = "components"
			}
		}
		if schema.Ref == "" {
			rv.slice = true
		}
		return rv
	default:
		return &qualifiedType{
			name: schema.Type.String(),
		}
	}
}

func oneOfSetValueFunc(structName string, pq pkgQual, schema *model.ParamSchema) jen.Code {
	paramNames := make([]string, 0, len(schema.ObjectParams))
	paramCommentTypes := make([]string, 0, len(schema.ObjectParams))
	paramTypes := make([]*qualifiedType, 0, len(schema.ObjectParams))
	for _, param := range schema.ObjectParams {
		paramNames = append(paramNames, toUnexportedName(param.Name))
		paramCommentTypes = append(paramCommentTypes, oneOfReturnTypeForComment(param.Name, pq, param.Schema))
		paramTypes = append(paramTypes, oneOfQualifiedType(param.Name, pq, param.Schema))
	}
	stmt := jen.Commentf("SetValue sets %s's value. The type must be one of %s.", structName, orList(paramCommentTypes)).Line()
	stmt.Add(
		jen.Func().Params(jen.Id("c").Op("*").Id(structName)).Id("SetValue(value interface{})").Block(
			jen.Switch(jen.Id("v := value").Dot("(type)")).BlockFunc(func(group *jen.Group) {
				for i := 0; i < len(paramNames); i++ {
					if len(paramTypes) <= i {
						panic("something has gone horribly wrong")
					}
					paramType := paramTypes[i]
					tp := paramType.jenType(pq)
					paramName := paramNames[i]

					group.Case(tp)
					group.Id("c").Dot(paramName).Id("=v")
				}
				group.Default()
				group.Panic(jen.Lit("type not acceptable"))
			}),
		),
	)
	return stmt
}

func oneOfValueFunc(structName string, pq pkgQual, schema *model.ParamSchema) jen.Code {
	paramNames := make([]string, 0, len(schema.ObjectParams))
	paramTypes := make([]string, 0, len(schema.ObjectParams))
	for _, param := range schema.ObjectParams {
		paramNames = append(paramNames, toUnexportedName(param.Name))
		paramTypes = append(paramTypes, oneOfReturnTypeForComment(param.Name, pq, param.Schema))
	}
	stmt := jen.Commentf("Value returns %s's value. The type will be one of %s.", structName, orList(paramTypes)).Line()
	stmt.Add(

		jen.Func().Params(jen.Id("c").Op("*").Id(structName)).Id("Value").Params().Interface().Block(
			jen.Switch(jen.Id("c.oneOfField")).BlockFunc(func(group *jen.Group) {
				for _, paramName := range paramNames {
					group.Case(jen.Lit(paramName))
					group.Return(jen.Id("c").Dot(paramName))
				}
			}),
			jen.Return(jen.Nil()),
		),
	)
	return stmt
}
