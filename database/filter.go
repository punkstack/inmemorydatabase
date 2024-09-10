package database

type FilterType string

const (
	EqualFilterType   FilterType = "equal"
	GreaterFilterType FilterType = "greater"
	LessFilterType    FilterType = "less"
)

type Filter interface {
	Apply(value interface{}) bool
}

type EqualFilter struct {
	Value interface{}
}

func (f *EqualFilter) Apply(value interface{}) bool {
	return value == f.Value
}

type GreaterFilter struct {
	Value int
}

func (f *GreaterFilter) Apply(value interface{}) bool {
	return value.(int) > f.Value
}

type LessFilter struct {
	Value int
}

func (f *LessFilter) Apply(value interface{}) bool {
	return value.(int) < f.Value
}
