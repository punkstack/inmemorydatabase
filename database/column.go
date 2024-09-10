package database

type ColumnType string

const (
	StringType ColumnType = "string"
	IntType    ColumnType = "int"
)

type Column struct {
	Name      string
	Type      ColumnType
	Validator Validator
}
