# go-query-string-builder
go-query-string-builder is Go library for building query string

### Usage
```go
import "github.com/mikaelim-id/go-query-string-builder/query"
```

### SelectQuery
#### Init
```go
selectQuery := SelectQuery{
		SelectStatement: []string{"*",},
		FromCommand:     "test_table",
    WhereClause:     "name='adama",
		GroupByClause:   []string{"name"},
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
		Offset: nil,
	}
```

#### AppendAndEqualCondition
Append **Field = Value** and condition in SelectQuery, will be ignored if value equal nil.
```go
selectQuery.AppendAndEqualCondition("name","name_value")
```

#### AppendAndCondition
Append and condition in SelectQuery.
```go
selectQuery.AppendAndCondition("name=name_value")
```

#### AppendOrEqualCondition
Append **Field = Value** or condition in SelectQuery, will be ignored if value equal nil.
```go
selectQuery.AppendOrEqualCondition("name","name_value")
```

#### AppendOrCondition
Append or condition in SelectQuery.
```go
selectQuery.AppendOrCondition("name=name_value")
```

#### Get Query String
```go
selectQuery.build()
```


#### Example
```go
pointer := 1
offset := 1
  
selectQuery := query.SelectQuery{
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
selectQuery.AppendAndEqualCondition("value", 1)
selectQuery.AppendAndEqualCondition("pointer", &pointer)
selectQuery.AppendAndEqualCondition("address", nil)
selectQuery.AppendAndCondition("phone is not null")
selectQuery.AppendOrCondition(query.BuildGroupedOrCondition("location_id=1", "location_id=3"))
  
fmt.Println(selectQuery.build())
// select name,address,phone,date_of_birth,location_id from test_table where
// name='test_name' and value=1 and pointer='1' and phone is not null or (location_id=1 or location_id=3)
// group by name,address order by name desc,address asc limit 5 offset 1
```


### UpdateQuery
#### Init
```go
updateQuery := UpdateQuery{
		UpdateStatement: "test_table",
    SetCommand: map[string]string{
      "name": "new name value",
    },
    WhereClause: "name='adam'",
	}
```

#### AppendSet
Append **Field = Value** set in UpdateQuery, will be ignored if value equal nil.
```go
updateQuery.AppendSet("address","new avenue")
```

#### AppendAndEqualCondition
Append **Field = Value** and condition in SelectQuery, will be ignored if value equal nil.
```go
selectQuery.AppendAndEqualCondition("name","name_value")
```

#### AppendAndCondition
Append and condition in SelectQuery.
```go
selectQuery.AppendAndCondition("name=name_value")
```

#### AppendOrEqualCondition
Append **Field = Value** or condition in SelectQuery, will be ignored if value equal nil.
```go
selectQuery.AppendOrEqualCondition("name","name_value")
```

#### AppendOrCondition
Append or condition in SelectQuery.
```go
selectQuery.AppendOrCondition("name=name_value")
```

#### Get Query String
```go
updateQuery.build()
```

#### Example
```go
jsonbSetPtr := `jsonb_set(additional_data,'{review_by}','"test"')`

updateQuery := query.UpdateQuery{
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
  
fmt.Println(updateQuery.build())
// update test_table set name=:name,value=:value,location_id=:location_id,
// form_data=jsonb_set(form_data,'{additional_data,review_by}','"test"'),
// client_data=jsonb_set(additional_data,'{review_by}','"test"')
// where name='test' and location_id='test' and value=1
```

### Helper

#### BuildGroupedAndCondition
Create grouped and condition.
```go
groupeAndCondition := selectQuery.BuildGroupedAndCondition("name='adam'", "name='eve'")
fmt.Println(groupedAndCondition)
// (name='adam' and name='eve')
```

#### BuildGroupedOrCondition
Create grouped or condition.
```go
groupedOrCondition := selectQuery.BuildGroupedOrCondition("name='adam'", "name='eve'")
fmt.Println(groupedOrCondition)
// (name='adam' or name='eve')
```
