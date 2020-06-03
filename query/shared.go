package query

import (
	"fmt"
	"reflect"
)

type SharedQuery struct {
	WhereClause     string
}

func (q *SharedQuery) AppendAndEqualCondition(column string, value interface{}) {
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

func (q *SharedQuery) AppendAndCondition(query string) {
	if q.WhereClause == "" {
		q.WhereClause = query
	} else {
		q.WhereClause += fmt.Sprintf(` and %s`, query)
	}
}


func (q *SharedQuery) AppendOrEqualCondition(column string, value interface{}) {
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

func (q *SharedQuery) AppendOrCondition(query string) {
	if q.WhereClause == "" {
		q.WhereClause = query
	} else {
		q.WhereClause += fmt.Sprintf(` or %s`, query)
	}
}