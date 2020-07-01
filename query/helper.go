package query

import (
	"fmt"
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
	} else if reflect.ValueOf(v).Kind() == reflect.Slice {
		return reflect.ValueOf(v).Len() > 0
	}

	return v != nil
}

func BuildGroupedAndCondition(conditionList ...string) string {
	query := "("
	for _, cond := range conditionList {
		if query == "(" {
			query += cond
		} else {
			query += fmt.Sprintf(` and %s`, cond)
		}
	}

	query += ")"

	return query
}

func BuildGroupedOrCondition(conditionList ...string) string {
	query := "("
	for _, cond := range conditionList {
		if query == "(" {
			query += cond
		} else {
			query += fmt.Sprintf(` or %s`, cond)
		}
	}

	query += ")"

	return query
}
