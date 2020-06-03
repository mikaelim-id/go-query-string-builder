package query

import (
	"fmt"
	"reflect"
)

type WhereClause string

func (q *WhereClause) AppendAndEqualCondition(column string, value interface{}) {
	if isNotNil(value) {
		if reflect.ValueOf(value).Kind() == reflect.String ||
			reflect.ValueOf(value).Kind() == reflect.Array {
			q.AppendAndCondition(fmt.Sprintf(`%s='%s'`, column, value))
		} else if reflect.ValueOf(value).Kind() == reflect.Int {
			q.AppendAndCondition(fmt.Sprintf(`%s=%d`, column, value))
		} else if reflect.ValueOf(value).Kind() == reflect.Ptr {
			q.AppendAndCondition(fmt.Sprintf(`%s='%v'`, column, reflect.ValueOf(value).Elem()))
		}
	}
}

func (q *WhereClause) AppendAndCondition(query string) {
	if q == nil {
		return
	}

	currentValue := *q

	if currentValue == "" {
		currentValue = WhereClause(query)
	} else {
		currentValue += WhereClause(fmt.Sprintf(` and %s`, query))
	}

	*q = currentValue
}

func (q *WhereClause) AppendOrEqualCondition(column string, value interface{}) {
	if isNotNil(value) {
		if reflect.ValueOf(value).Kind() == reflect.String ||
			reflect.ValueOf(value).Kind() == reflect.Array {
			q.AppendOrCondition(fmt.Sprintf(`%s='%s'`, column, value))
		} else if reflect.ValueOf(value).Kind() == reflect.Int {
			q.AppendOrCondition(fmt.Sprintf(`%s=%d`, column, value))
		} else if reflect.ValueOf(value).Kind() == reflect.Ptr {
			q.AppendOrCondition(fmt.Sprintf(`%s='%v'`, column, reflect.ValueOf(value).Elem()))
		}
	}
}

func (q *WhereClause) AppendOrCondition(query string) {
	if q == nil {
		return
	}

	currentValue := *q

	if currentValue == "" {
		currentValue = WhereClause(query)
	} else {
		currentValue += WhereClause(fmt.Sprintf(` or %s`, query))
	}

	*q = currentValue
}
