package query

import (
	"fmt"
	"reflect"
	"strings"
)

type SelectQuery struct {
	SelectStatement []string
	FromCommand     string
	WhereClause     string
	GroupByClause   []string
	OrderByClause   []OrderBy
	Limit           int
	Offset          *int
}

type OrderBy struct {
	Field string
	Asc   bool
}

func (q *SelectQuery) AppendAndEqualCondition(column string, value interface{}) {
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

func (q *SelectQuery) AppendAndCondition(query string) {
	if q.WhereClause == "" {
		q.WhereClause = query
	} else {
		q.WhereClause += fmt.Sprintf(` and %s`, query)
	}
}

func (q *SelectQuery) AppendOrCondition(query string) {
	if q.WhereClause == "" {
		q.WhereClause = query
	} else {
		q.WhereClause += fmt.Sprintf(` or %s`, query)
	}
}

func (q *SelectQuery) BuildGroupedOrCondition(conditionList ...string) string {
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

func (q *SelectQuery) Build() string {
	query := fmt.Sprintf(`select %s from %s`,
		strings.Join(q.SelectStatement[:], ","),
		q.FromCommand,
	)

	if isNotNil(q.WhereClause) {
		query += fmt.Sprintf(" where %s", q.WhereClause)
	}

	if len(q.GroupByClause) > 0 {
		query += fmt.Sprintf(" group by %s",
			strings.Join(q.GroupByClause[:], ","))
	}

	if isNotNil(q.OrderByClause) {
		orderByFlag := true

		for _, value := range q.OrderByClause {
			if orderByFlag {
				query += " order by "
				orderByFlag = false
			} else {
				query += ","
			}

			orderBy := "desc"
			if value.Asc {
				orderBy = "asc"
			}

			query += fmt.Sprintf("%s %s", value.Field, orderBy)
		}
	}

	if q.Limit > 0 {
		query += fmt.Sprintf(" limit %d", q.Limit)
	}

	if q.Offset != nil {
		query += fmt.Sprintf(" offset %d", *q.Offset)
	}

	return query
}
