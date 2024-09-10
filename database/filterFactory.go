package database

import "fmt"

type FilterFactory struct{}

func (ff *FilterFactory) CreateFilter(filterType FilterType, value interface{}) (Filter, error) {
	switch filterType {
	case EqualFilterType:
		return &EqualFilter{Value: value}, nil
	case GreaterFilterType:
		return &GreaterFilter{Value: value.(int)}, nil
	case LessFilterType:
		return &LessFilter{Value: value.(int)}, nil
	default:
		return nil, fmt.Errorf("unknown filter type: %s", filterType)
	}
}
