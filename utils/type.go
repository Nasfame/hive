package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func getTypeOf(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

// GetTypeString returns the type name of a value.
// It returns without '.' notations
func GetTypeString[T any](val T) string {
	typeString := fmt.Sprint("%T", val)

	typeParts := strings.Split(typeString, ".")
	typeName := typeParts[len(typeParts)-1]

	return typeName
}
