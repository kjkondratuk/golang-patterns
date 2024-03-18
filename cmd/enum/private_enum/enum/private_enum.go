package enum

import (
	"errors"
	"fmt"
	"strings"
)

const (
	VALUE_1 = "my string"
	VALUE_2 = "my other string"
	VALUE_3 = "yet another string"

	MY_VALUE_1 = myEnumVal(VALUE_1)
	MY_VALUE_2 = myEnumVal(VALUE_2)
	MY_VALUE_3 = myEnumVal(VALUE_3)
)

type myEnumVal string

type MyEnumVal interface {
	UnmarshalJSON(b []byte) error
	Value() string
}

func New(s string) (MyEnumVal, error) {
	switch s {
	case VALUE_1:
	case VALUE_2:
	case VALUE_3:
	default:
		return nil, fmt.Errorf("Invalid string specified for enum")
	}
	val := myEnumVal(s)
	return &val, nil
}

func (e *myEnumVal) Value() string {
	return string(*e)
}

func (e *myEnumVal) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")

	switch str {
	case VALUE_1:
		//v :=
		*e = MY_VALUE_1
	case VALUE_2:
		*e = MY_VALUE_2
	case VALUE_3:
		*e = MY_VALUE_3
	default:
		return errors.New(fmt.Sprintf("Invalid enum value: %s", str))
	}

	return nil
}
