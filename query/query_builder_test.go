package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectQuery(t *testing.T) {
	expectation := "select name,address,phone,date_of_birth,location_id from test_table where" +
		" name='test_name' and address='test_address' and value=1 and pointer='1' or (location_id=1 or location_id=3) and (location_id=1 and location_id=3)" +
		" group by name,address order by name desc,address asc limit 5 offset 1"

	pointer := 1
	offset := 1

	selectQuery := SelectQuery{
		SelectStatement: []string{"name", "address", "phone", "date_of_birth", "location_id"},
		FromCommand:     "test_table",
		GroupByClause:   []string{"name", "address"},
		OrderByClause: []OrderBy{
			{
				Field: "name",
				Asc:   false,
			},
			{
				Field: "address",
				Asc:   true,
			},
		},
		Limit:  5,
		Offset: &offset,
	}

	selectQuery.AppendAndEqualCondition("name", "test_name")
	selectQuery.AppendAndEqualCondition("address", "test_address")
	selectQuery.AppendAndEqualCondition("value", 1)
	selectQuery.AppendAndEqualCondition("pointer", &pointer)
	selectQuery.AppendOrCondition(BuildGroupedOrCondition("location_id=1", "location_id=3"))
	selectQuery.AppendAndCondition(BuildGroupedAndCondition("location_id=1", "location_id=3"))
	assert.Equal(t, len(expectation), len(selectQuery.Build()))

	expectation2 := "select name,address,phone,date_of_birth,location_id from test_table where name='test_name'"

	selectQuery2 := SelectQuery{
		SelectStatement: []string{"name", "address", "phone", "date_of_birth", "location_id"},
		FromCommand:     "test_table",
	}

	selectQuery2.AppendOrEqualCondition("name", "test_name")

	assert.Equal(t, len(expectation2), len(selectQuery2.Build()))
}

func TestUpdateQuery(t *testing.T) {
	expectedResult := "update test_table set name=:name,value=:value,location_id=:location_id," +
		`form_data=jsonb_set(form_data,'{additional_data,review_by}','"test"'),` +
		`client_data=jsonb_set(additional_data,'{review_by}','"test"')` +
		" where name='test' and location_id='test' and value=1"

	jsonbSetPtr := `jsonb_set(additional_data,'{review_by}','"test"')`

	updateQuery := UpdateQuery{
		UpdateStatement: "test_table",
	}
	updateQuery.AppendSet("name", "nama")
	updateQuery.AppendSet("value", "val")
	updateQuery.AppendSet("location_id", "loc")
	updateQuery.AppendSet("form_data", `jsonb_set(form_data,'{additional_data,review_by}','"test"')`)
	updateQuery.AppendSet("client_data", &jsonbSetPtr)

	updateQuery.AppendAndEqualCondition("name", "test")
	updateQuery.AppendAndEqualCondition("location_id", "test")
	updateQuery.AppendAndEqualCondition("value", 1)

	result := updateQuery.Build()
	assert.Equal(t, len(expectedResult), len(result))
}
