package database

import "errors"

type Database struct {
	tables map[string]*Table
}

func NewDatabase() *Database {
	return &Database{
		tables: make(map[string]*Table),
	}
}

func (db *Database) CreateTable(name string, columns []Column) error {
	if _, exists := db.tables[name]; exists {
		return errors.New("table already exists")
	}
	db.tables[name] = NewTable(columns)
	return nil
}

func (db *Database) DeleteTable(name string) error {
	if _, exists := db.tables[name]; !exists {
		return errors.New("table does not exist")
	}
	delete(db.tables, name)
	return nil
}

func (db *Database) GetTable(name string) (*Table, error) {
	table, exists := db.tables[name]
	if !exists {
		return nil, errors.New("table does not exist")
	}
	return table, nil
}
