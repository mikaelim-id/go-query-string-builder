package query

import (
	"fmt"
	"strings"
)

type SelectQuery struct {
	SelectStatement []string
	FromCommand     string
	GroupByClause   []string
	OrderByClause   []OrderBy
	Limit           int
	Offset          *int
	WhereClause
}

type OrderBy struct {
	Field string
	Asc   bool
}

func (q *SelectQuery) Build() string {
	query := fmt.Sprintf(`select %s from %s`,
		strings.Join(q.SelectStatement[:], ","),
		q.FromCommand,
	)

	if len(q.WhereClause) > 0 {
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
