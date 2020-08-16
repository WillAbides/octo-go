package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/inflection"
	"github.com/pkg/errors"
	"github.com/willabides/octo-go/generator/internal/model"
)

type Model struct {
	Endpoints        []*model.Endpoint
	ComponentSchemas map[string]*model.ParamSchema
	ComponentHeaders map[string]*model.ParamSchema
}

func Openapi2Model(schemaSrc io.Reader) (*Model, error) {
	data, err := ioutil.ReadAll(schemaSrc)
	if err != nil {
		return nil, errors.Errorf("could not read from schemaSrc")
	}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(data)
	if err != nil {
		return nil, errors.Errorf("could not load openapiDef")
	}
	mdl := new(Model)
	mdl.Endpoints, err = buildEndpoints(swagger)
	if err != nil {
		return nil, err
	}
	mdl.ComponentSchemas, err = buildComponentSchemas(swagger)
	if err != nil {
		return nil, err
	}
	mdl.ComponentHeaders, err = buildComponentHeaders(swagger)
	if err != nil {
		return nil, err
	}
	return mdl, nil
}

func buildComponentHeaders(swagger *openapi3.Swagger) (map[string]*model.ParamSchema, error) {
	result := make(map[string]*model.ParamSchema, len(swagger.Components.Headers))
	var err error
	for name, ref := range swagger.Components.Headers {
		result[name], err = opParamSchema(ref.Value.Schema)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func prepareComponentSchemas(swagger *openapi3.Swagger) {
	topNames := make([]string, 0, len(swagger.Components.Schemas))
	for nm := range swagger.Components.Schemas {
		topNames = append(topNames, nm)
	}
	sort.Strings(topNames)
	for _, nm := range topNames {
		prepareComponentSchemaObj(swagger, nm, swagger.Components.Schemas[nm])
	}
}

func prepareComponentSchemaObj(swagger *openapi3.Swagger, parentName string, schemaRef *openapi3.SchemaRef) {
	schema := schemaRef.Value
	oneOfNames := map[string]bool{}
	prepareOneOf(swagger, parentName, schema, oneOfNames)
	if schema.Items != nil && schemaRef.Ref == "" {
		prepareComponentSchemaObj(swagger, parentName, schema.Items)
		return
	}
	for propName, ref := range schema.Properties {
		propName = overrideComponentSchemaName(propName)
		if ref.Ref != "" {
			continue
		}
		val := ref.Value
		switch opSchemaType(val) {
		case model.ParamTypeObject, model.ParamTypeOneOf:
			if val.AdditionalProperties != nil {
				break
			}
			if len(val.AllOf) == 1 {
				break
			}

			fullName := fmt.Sprintf("%s-%s", parentName, propName)
			swagger.Components.Schemas[fullName] = openapi3.NewSchemaRef("", val)
			ref.Ref = "#/components/schemas/" + fullName
			prepareComponentSchemaObj(swagger, fullName, ref)
		case model.ParamTypeArray:
			itemsRef := val.Items
			itemsVal := itemsRef.Value
			fullName := arrayName(parentName, propName)
			if strings.HasPrefix(itemsRef.Ref, "#/components/schemas/") {
				fullName = strings.TrimPrefix(itemsRef.Ref, "#/components/schemas/")
			}
			wrongType := false
			switch opSchemaType(itemsVal) {
			case model.ParamTypeObject, model.ParamTypeOneOf:
			default:
				wrongType = true
			}
			if wrongType {
				break
			}
			if itemsVal.AdditionalProperties != nil {
				break
			}
			swagger.Components.Schemas[fullName] = openapi3.NewSchemaRef("", itemsVal)
			itemsRef.Ref = "#/components/schemas/" + fullName
			prepareComponentSchemaObj(swagger, fullName, itemsRef)
		}
	}
}

func arrayName(parentName, propName string) string {
	nameParts := strings.Split(propName, "-")
	lastIdx := len(nameParts) - 1
	nameParts[lastIdx] = inflection.Singular(nameParts[lastIdx])
	nm := strings.Join(nameParts, "-")
	return fmt.Sprintf("%s-%s", parentName, nm)
}

func prepareOneOf(swagger *openapi3.Swagger, parentName string, schema *openapi3.Schema, oneOfNames map[string]bool) {
	for _, ref := range schema.OneOf {
		if ref.Ref != "" {
			continue
		}
		val := ref.Value
		switch opSchemaType(val) {
		case model.ParamTypeObject, model.ParamTypeOneOf:
			if val.AdditionalProperties != nil {
				break
			}
			propName := oneOfPropName(ref, oneOfNames)
			fullName := fmt.Sprintf("%s-%s", parentName, propName)
			swagger.Components.Schemas[fullName] = openapi3.NewSchemaRef("", val)
			ref.Ref = "#/components/schemas/" + fullName
			prepareComponentSchemaObj(swagger, fullName, ref)
		}
	}
}

func overrideComponentSchemaName(name string) string {
	switch name {
	case "+1":
		return "plus-one"
	case "-1":
		return "minus-one"
	}
	return name
}

func buildComponentSchemas(swagger *openapi3.Swagger) (map[string]*model.ParamSchema, error) {
	prepareComponentSchemas(swagger)
	result := make(map[string]*model.ParamSchema, len(swagger.Components.Schemas))
	var err error
	for name, ref := range swagger.Components.Schemas {
		result[name], err = opParamSchema(ref)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func buildEndpoints(swagger *openapi3.Swagger) ([]*model.Endpoint, error) {
	var endpoints []*model.Endpoint
	for opPath, pathItem := range swagger.Paths {
		for method, op := range pathItem.Operations() {
			endpoint, err := buildEndpoint(opPath, method, op)
			if err != nil {
				return nil, err
			}
			endpoints = append(endpoints, endpoint)
		}
	}
	return endpoints, nil
}

func buildEndpoint(opPath, httpMethod string, op *openapi3.Operation) (*model.Endpoint, error) {
	endpoint := model.Endpoint{
		Path:       opPath,
		Method:     httpMethod,
		Name:       path.Base(op.OperationID),
		Concern:    path.Dir(op.OperationID),
		Summary:    op.Summary,
		HelpText:   op.Description,
		Deprecated: op.Deprecated,
		ID:         op.OperationID,
	}
	ext, err := opExtGithub(op)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if ext != nil {
		endpoint.EnabledForApps = ext.EnabledForApps
		endpoint.GithubCloudOnly = ext.GithubCloudOnly
		endpoint.Legacy = ext.Legacy
		for _, preview := range ext.Previews {
			endpoint.Previews = append(endpoint.Previews, &model.Preview{
				Required: preview.Required,
				Name:     preview.Name,
				Note:     preview.Note,
			})
		}
	}
	if op.ExternalDocs != nil {
		endpoint.DocsURL = op.ExternalDocs.URL
	}
	for _, pRef := range op.Parameters {
		var param *model.Param
		param, err = buildParam(pRef.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		switch pRef.Value.In {
		case openapi3.ParameterInQuery:
			endpoint.QueryParams = append(endpoint.QueryParams, param)
		case openapi3.ParameterInHeader:
			endpoint.Headers = append(endpoint.Headers, param)
		case openapi3.ParameterInPath:
			endpoint.PathParams = append(endpoint.PathParams, param)
		}
	}
	endpoint.Responses, err = responses(op)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error processing responses from operation %q", op.OperationID))
	}
	endpoint.SuccessMediaType, err = successMediaType(endpoint.Responses)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	reqBody, err := requestsBody(op)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	endpoint.RequestBody = reqBody
	return &endpoint, nil
}

func successMediaType(responses map[int]*model.Response) (string, error) {
	var result string
	for code, response := range responses {
		if code > 299 {
			continue
		}
		mediaType := response.MediaType
		if mediaType == "" {
			continue
		}
		if result != "" && result != mediaType {
			return "", errors.New("endpoint has multiple success mediaTypes")
		}
		result = mediaType
	}
	return result, nil
}

var op2modelTypes = map[string]model.ParamType{
	"string":  model.ParamTypeString,
	"integer": model.ParamTypeInt,
	"boolean": model.ParamTypeBool,
	"object":  model.ParamTypeObject,
	"":        model.ParamTypeObject,
	"array":   model.ParamTypeArray,
	"number":  model.ParamTypeNumber,
}

func opSchemaType(opSchema *openapi3.Schema) model.ParamType {
	if strings.HasPrefix(opSchema.Type, "[]") {
		return model.ParamTypeArray
	}
	switch {
	case strings.HasPrefix(opSchema.Type, "[]"):
		return model.ParamTypeArray
	case len(opSchema.OneOf) > 0:
		return model.ParamTypeOneOf
	}
	return op2modelTypes[opSchema.Type]
}

func opParamSchema(schemaRef *openapi3.SchemaRef) (*model.ParamSchema, error) {
	opSchema := schemaRef.Value

	if len(opSchema.AllOf) == 1 {
		sr := &openapi3.SchemaRef{
			Ref:   opSchema.AllOf[0].Ref,
			Value: opSchema.AllOf[0].Value,
		}
		sr.Value.Nullable = opSchema.Nullable
		sr.Value.Description = opSchema.Description
		return opParamSchema(sr)
	}

	schema := model.ParamSchema{
		Nullable: opSchema.Nullable,
		Ref:      schemaRef.Ref,
		Type:     opSchemaType(opSchema),
	}
	var err error
	switch schema.Type {
	case model.ParamTypeInvalid:
		return nil, errors.Errorf("unknown schema type %s", opSchema.Type)
	case model.ParamTypeArray:
		if opSchema.Items == nil {
			return nil, errors.New("opSchema.Items is nil")
		}
		schema.ItemSchema, err = opParamSchema(opSchema.Items)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	case model.ParamTypeOneOf:
		schema.ObjectParams, err = oneOfObjectParams(schemaRef, opSchema.OneOf)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	case model.ParamTypeObject:
		if opSchema.AdditionalProperties != nil {
			schema.ItemSchema, err = opParamSchema(opSchema.AdditionalProperties)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}
		props := make(map[string]*openapi3.SchemaRef, len(opSchema.Properties))
		for k, v := range opSchema.Properties {
			props[k] = v
		}

		err = appendParams(schemaRef, opSchema, props, &schema)
		if err != nil {
			return nil, err
		}
	}
	return &schema, nil
}

func appendParams(schemaRef *openapi3.SchemaRef, opSchema *openapi3.Schema, props map[string]*openapi3.SchemaRef, schema *model.ParamSchema) error {
	err := addAllOfProps(opSchema, props)
	if err != nil {
		return errors.WithStack(err)
	}

	propNames := make([]string, 0, len(props))
	for name := range props {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)
	for _, name := range propNames {
		ref := props[name]
		var objParam *model.Param
		objParam, err = opObjectParam(schemaRef, ref, name)
		if err != nil {
			return errors.WithStack(err)
		}
		schema.ObjectParams = append(schema.ObjectParams, objParam)
	}
	return nil
}

func oneOfObjectParams(opSchemaRef *openapi3.SchemaRef, oneOf []*openapi3.SchemaRef) ([]*model.Param, error) {
	result := make([]*model.Param, 0, len(oneOf))
	names := make(map[string]bool, len(oneOf))
	for _, ref := range oneOf {
		name := oneOfPropName(ref, names)
		objParam, err := opObjectParam(opSchemaRef, ref, name)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		result = append(result, objParam)
	}
	return result, nil
}

func oneOfPropName(ref *openapi3.SchemaRef, names map[string]bool) string {
	name := strings.TrimPrefix(ref.Ref, "#")
	name = strings.TrimSpace(name)
	nameParts := strings.Split(name, "/")
	name = nameParts[len(nameParts)-1]
	if name != "" {
		names[name] = true
		return name
	}
	if name == "" {
		typeName := opSchemaType(ref.Value).String()
		name = fmt.Sprintf("as-%s", typeName)
	}
	if names[name] {
		for i := 2; ; i++ {
			tryName := fmt.Sprintf("%s-%d", name, i)
			if !names[tryName] {
				name = tryName
				break
			}
		}
	}
	names[name] = true
	return name
}

func addAllOfProps(opSchema *openapi3.Schema, props map[string]*openapi3.SchemaRef) error {
	allOfProps := map[string]*openapi3.SchemaRef{}
	for _, ref := range opSchema.AllOf {
		names := make([]string, 0, len(ref.Value.Properties))
		for name := range ref.Value.Properties {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			val := ref.Value.Properties[name]
			if allOfProps[name] != nil && !eqSchemaRef(val, allOfProps[name]) {
				return errors.Errorf("duplicating property name from allOf: %q", name)
			}
			allOfProps[name] = val
		}
	}
	for k, v := range allOfProps {
		props[k] = v
	}
	return nil
}

// strips examples then compares json
func eqSchemaRef(a, b *openapi3.SchemaRef) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	aa := new(openapi3.SchemaRef)
	aa.Value = new(openapi3.Schema)
	*aa.Value = *a.Value
	aa.Value.Example = nil
	bb := new(openapi3.SchemaRef)
	bb.Value = new(openapi3.Schema)
	*bb.Value = *b.Value
	bb.Value.Example = nil
	aBytes, err := json.Marshal(aa)
	if err != nil {
		panic(err)
	}
	bBytes, err := json.Marshal(bb)
	if err != nil {
		panic(err)
	}
	ok := bytes.Equal(aBytes, bBytes)
	return ok
}

func opObjectParam(opSchemaRef, propSchemaRef *openapi3.SchemaRef, name string) (*model.Param, error) {
	propSchema := propSchemaRef.Value
	opSchema := opSchemaRef.Value
	param := model.Param{
		Name:     name,
		HelpText: propSchema.Description,
	}
	for _, reqName := range opSchema.Required {
		if name == reqName {
			param.Required = true
			break
		}
	}
	var err error
	param.Schema, err = opParamSchema(propSchemaRef)
	if err != nil {
		return nil, errors.Wrapf(err, "error running opParamSchema on param named %s", name)
	}
	return &param, nil
}

func buildParam(opParam *openapi3.Parameter) (*model.Param, error) {
	schema, err := opParamSchema(opParam.Schema)
	if err != nil {
		return nil, err
	}
	param := model.Param{
		Required: opParam.Required,
		Name:     opParam.Name,
		HelpText: opParam.Description,
		Schema:   schema,
	}
	return &param, nil
}

func opExtGithub(op *openapi3.Operation) (*extGithub, error) {
	xMsg, ok := op.Extensions["x-github"].(json.RawMessage)
	if !ok {
		return nil, nil
	}
	var ext extGithub
	err := json.Unmarshal(xMsg, &ext)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ext, nil
}

type extGithub struct {
	Legacy          bool
	EnabledForApps  bool
	GithubCloudOnly bool
	Previews        []struct {
		Name     string
		Required bool
		Note     string
	}
}

func getJsonResponse(content openapi3.Content) (name string, mediaType *openapi3.MediaType, err error) {
	preferredTypes := []string{
		"application/scim+json",
		"application/vnd.github.v3.star+json",
	}
	var gotPreferredNames []string
	var gotPreferred []*openapi3.MediaType

	for _, preferredName := range preferredTypes {
		got := content.Get(preferredName)
		if got == nil {
			continue
		}
		gotPreferredNames = append(gotPreferredNames, preferredName)
		gotPreferred = append(gotPreferred, got)
	}

	switch len(gotPreferred) {
	case 0:
	case 1:
		if len(gotPreferredNames) == 0 {
			panic("oops")
		}
		return gotPreferredNames[0], gotPreferred[0], nil
	default:
		return "", nil, errors.Errorf("got multiple preffered json mediatypes: %v", gotPreferredNames)
	}

	var mts []*openapi3.MediaType
	var mtNames []string
	for nm, mediaType := range content {
		if strings.HasSuffix(nm, "json") {
			mts = append(mts, mediaType)
			mtNames = append(mtNames, nm)
		}
	}
	switch len(mts) {
	case 0:
		return "", nil, nil
	case 1:
		if len(mtNames) == 0 {
			panic("oops")
		}

		return mtNames[0], mts[0], nil
	default:
		return "", nil, errors.Errorf("got multiple json mediatypes: %v", mtNames)
	}
}

func responses(op *openapi3.Operation) (map[int]*model.Response, error) {
	result := make(map[int]*model.Response, len(op.Responses))
	for responseCode, responseRef := range op.Responses {
		code, err := strconv.Atoi(responseCode)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		var hasExample bool
		var schema *model.ParamSchema
		mediaType, jsonResponse, err := getJsonResponse(responseRef.Value.Content)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if jsonResponse != nil {
			schema, err = opParamSchema(jsonResponse.Schema)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("OperationID: %s, responseCode: %d", op.OperationID, code))
			}
			if jsonResponse.Example != nil || len(jsonResponse.Examples) > 0 {
				hasExample = true
			}
		}

		result[code] = &model.Response{
			MediaType:  mediaType,
			Body:       schema,
			HasExample: hasExample,
		}
	}
	return result, nil
}

func requestBodyContent(op *openapi3.Operation) (string, *openapi3.MediaType, error) {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return "", nil, nil
	}
	explicitMediaTypes := map[string]string{
		"markdown/render-raw": "text/x-markdown",
	}
	contents := op.RequestBody.Value.Content
	wantMediaType := explicitMediaTypes[op.OperationID]
	if wantMediaType == "" {
		switch len(contents) {
		case 0:
			return "", nil, nil
		case 1:
			for mediaType, content := range contents {
				return mediaType, content, nil
			}
		default:
			return "", nil, errors.Errorf("requestBody for %s has more than one content and no explicitly set media type", op.OperationID)
		}
	}
	errMsg := fmt.Sprintf("requestBody for %s is explicitly set to %s, but that media type is not present", op.OperationID, wantMediaType)
	if contents == nil {
		return "", nil, errors.New(errMsg)
	}
	content := contents[wantMediaType]
	if content == nil {
		return "", nil, errors.New(errMsg)
	}
	return wantMediaType, content, nil
}

func requestsBody(op *openapi3.Operation) (*model.RequestBody, error) {
	mediaType, content, err := requestBodyContent(op)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if content == nil {
		return nil, nil
	}
	schema, err := opParamSchema(content.Schema)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &model.RequestBody{
		MediaType: mediaType,
		Schema:    schema,
	}, nil
}
