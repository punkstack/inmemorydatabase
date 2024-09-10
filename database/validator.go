package database

import (
	"errors"
	"fmt"
)

type Validator interface {
	Validate(value interface{}) error
}

type StringValidator struct {
	MaxLength int
}

func (v *StringValidator) Validate(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("value is not a string")
	}
	if len(strValue) > v.MaxLength {
		return fmt.Errorf("string length exceeds maximum")
	}
	return nil
}

type IntValidator struct {
	MinValue int
	MaxValue int
}

func (v *IntValidator) Validate(value interface{}) error {
	intValue, ok := value.(int)
	if !ok {
		return errors.New("value is not an int")
	}
	if intValue < v.MinValue || intValue > v.MaxValue {
		return fmt.Errorf("int value out of range")
	}
	return nil
}
