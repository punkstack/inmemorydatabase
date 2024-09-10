package database

import (
	"errors"
	"fmt"
	"sync"
)

type Table struct {
	columns map[string]Column
	rows    map[int]*Row
	nextID  int
	mu      sync.Mutex
}

func NewTable(columns []Column) *Table {
	colMap := make(map[string]Column)
	for _, col := range columns {
		colMap[col.Name] = col
	}
	return &Table{
		columns: colMap,
		rows:    make(map[int]*Row),
		nextID:  0,
	}
}

func (t *Table) Insert(rowData map[string]interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	row := &Row{data: make(map[string]interface{})}
	for colName, col := range t.columns {
		value, exists := rowData[colName]
		if !exists {
			return errors.New("missing required column: " + colName)
		}
		if col.Validator != nil {
			if err := col.Validator.Validate(value); err != nil {
				return fmt.Errorf("invalid value for column %s: %v", colName, err)
			}
		}
		row.data[colName] = value
	}
	t.rows[t.nextID] = row
	t.nextID++
	return nil
}

func (t *Table) FetchAll() {
	t.mu.Lock()
	defer t.mu.Unlock()

	for id, row := range t.rows {
		fmt.Printf("Row %d: %v \n", id, row.data)
	}
}

func (t *Table) Update(rowID int, data map[string]interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	row, exists := t.rows[rowID]
	if !exists {
		return errors.New("row not found")
	}

	for colName, col := range t.columns {
		if value, exists := data[colName]; exists {
			if col.Validator != nil {
				if err := col.Validator.Validate(value); err != nil {
					return fmt.Errorf("invalid value for column %s: %v", colName, err)
				}
			}
			row.data[colName] = value
		}
	}

	return nil
}

func (t *Table) Filter(filters map[string]Filter) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for id, row := range t.rows {
		match := true
		for columnName, filter := range filters {
			if !filter.Apply(row.data[columnName]) {
				match = false
				break
			}
		}
		if match {
			fmt.Printf("Row %d: %v \n", id, row.data)
		}
	}
}
