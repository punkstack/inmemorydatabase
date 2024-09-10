package main

import (
	"fmt"
	"sqliteinmemory/database"
)

func main() {
	db := database.NewDatabase()

	columns := []database.Column{
		{Name: "name", Type: database.StringType, Validator: &database.StringValidator{MaxLength: 20}},
		{Name: "age", Type: database.IntType, Validator: &database.IntValidator{MinValue: 18, MaxValue: 110}},
	}

	err := db.CreateTable("users", columns)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	table, err := db.GetTable("users")
	if err != nil {
		fmt.Println("Error getting table:", err)
		return
	}

	err = table.Insert(map[string]interface{}{"id": 1, "name": "Manoj", "age": 30})
	if err != nil {
		fmt.Println("Error inserting row:", err)
		return
	}

	err = table.Insert(map[string]interface{}{"id": 2, "name": "Suhail", "age": 25})
	if err != nil {
		fmt.Println("Error inserting row:", err)
		return
	}

	err = table.Insert(map[string]interface{}{"id": 3, "name": "Aryaman", "age": 25})
	if err != nil {
		fmt.Println("Error inserting row:", err)
		return
	}

	table.FetchAll()

	filterFactory := &database.FilterFactory{}
	filters := map[string]database.Filter{}

	equalFilter, err := filterFactory.CreateFilter(database.EqualFilterType, "Manoj Bojja")
	if err != nil {
		fmt.Println("Error creating filter:", err)
		return
	}
	filters["name"] = equalFilter

	greaterFilter, err := filterFactory.CreateFilter(database.GreaterFilterType, 20)
	if err != nil {
		fmt.Println("Error creating filter:", err)
		return
	}

	filters["age"] = greaterFilter

	table.Filter(filters)
}
