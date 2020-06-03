package query

import (
	"fmt"
	"strings"
)

type UpdateQuery struct {
	UpdateStatement string
	SetCommand      map[string]string
	SharedQuery
}

func (q *UpdateQuery) AppendSet(column string, value interface{}) {
	if q.SetCommand == nil {
		q.SetCommand = make(map[string]string)
	}

	if isNotNil(value) {
		stringValue, valid := value.(string)
		if valid && strings.Contains(stringValue, "jsonb_set") {
			q.SetCommand[column] = value.(string)
			return
		}

		stringPtrValue, valid := value.(*string)
		if valid && strings.Contains(*stringPtrValue, "jsonb_set") {
			q.SetCommand[column] = *stringPtrValue
			return
		}

		q.SetCommand[column] = fmt.Sprintf(":%s", column)
	}
}

func (q *UpdateQuery) Build() string {
	query := fmt.Sprintf(`update %s`,
		q.UpdateStatement,
	)

	setFlag := true
	for key, value := range q.SetCommand {
		if isNotNil(value) {
			if setFlag {
				query += " set "
				setFlag = false
			} else {
				query += ","
			}

			query += fmt.Sprintf(`%s=%s`, key, value)
		}
	}

	if isNotNil(q.WhereClause) {
		query += fmt.Sprintf(" where %s", q.WhereClause)
	}

	return query
}
