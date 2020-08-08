package internal

import (
	"reflect"
)

// EndpointAttribute is an attribute for an endpoint
type EndpointAttribute int

var endpointAttributes map[string][]EndpointAttribute

// ReqAttributes returns all attributes endpoint attributes for an endpoint request struct
func ReqAttributes(req interface{}) []EndpointAttribute {
	return endpointAttributes[structName(reflect.TypeOf(req))]
}

// ReqHasEndpointAttribute determines whether a request's endspoint has a given attribute.
func ReqHasEndpointAttribute(req interface{}, attr EndpointAttribute) bool {
	attrs := ReqAttributes(req)
	for _, a := range attrs {
		if attr == a {
			return true
		}
	}
	return false
}

// structName returns the name of a struct from its reflect type or a pointer
func structName(tp reflect.Type) string {
	if tp.Kind() == reflect.Ptr {
		return structName(tp.Elem())
	}
	return tp.Name()
}
