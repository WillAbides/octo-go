package internal

import (
	"reflect"
)

var (
	reqOperationIDs     map[string]string //nolint:unused // save for later
	operationAttributes map[string][]EndpointAttribute
)

// EndpointAttribute is an attribute for an endpoint
type EndpointAttribute int

// OperationAttributes returns the EndpointAttributes associated with an operation id
func OperationAttributes(id string) []EndpointAttribute {
	return operationAttributes[id]
}

// OperationHasAttribute returns true if the operation id the given attribute
func OperationHasAttribute(id string, attr EndpointAttribute) bool {
	attrs := operationAttributes[id]
	for _, a := range attrs {
		if attr == a {
			return true
		}
	}
	return false
}

// structName returns the name of a struct from its reflect type or a pointer
//nolint:unused // save for later
func structName(tp reflect.Type) string {
	if tp.Kind() == reflect.Ptr {
		return structName(tp.Elem())
	}
	return tp.Name()
}
