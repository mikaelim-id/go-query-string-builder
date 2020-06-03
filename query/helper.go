package query

import (
	"reflect"

	"github.com/google/uuid"
)

func isNotNil(v interface{}) bool {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		return !reflect.ValueOf(v).IsNil()
	} else if reflect.ValueOf(v).Kind() == reflect.String {
		return v != ""
	} else if reflect.ValueOf(v).Kind() == reflect.Array {
		return v != uuid.Nil
	}

	return v != nil
}
