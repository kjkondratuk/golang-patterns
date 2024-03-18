package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	VALUE_1 = "my string"
	VALUE_2 = "my other string"
	VALUE_3 = "yet another string"

	MY_VALUE_1 = MyEnumVal(VALUE_1)
	MY_VALUE_2 = MyEnumVal(VALUE_2)
	MY_VALUE_3 = MyEnumVal(VALUE_3)

	JSON_STRING = `
{
	"stringVal": "my other string"
}
`

	JSON_STRING_1 = `
{
	"stringVal": "something else"
}
`
)

type MyEnumVal string

func (e *MyEnumVal) UnmarshalJSON(b []byte) error {
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

type payload struct {
	Value string `json:"stringVal"`
}

type enum_payload struct {
	Value MyEnumVal `json:"stringVal"`
}

func main() {
	// Can't do this because of constant keyword
	//VALUE_1 = "some other string"

	// The below isn't great because we can assign invalid values to our enum field during marshalling
	val := &payload{}
	if err := json.Unmarshal([]byte(JSON_STRING), val); err != nil {
		fmt.Println(fmt.Sprintf("val: WHOOPS! %s", err))
	}

	fmt.Printf("val: %+v\n", val)

	val1 := &payload{}
	if err := json.Unmarshal([]byte(JSON_STRING_1), val1); err != nil {
		fmt.Println(fmt.Sprintf("val1: WHOOPS! %s", err))
	}

	fmt.Printf("val1: %+v\n", val1)

	// The below is better because the caller can't modify valid enum values and invalid enum values are rejected when unmarshalling
	val2 := &enum_payload{}
	if err := json.Unmarshal([]byte(JSON_STRING), val2); err != nil {
		fmt.Println(fmt.Sprintf("val2: WHOOPS! %s", err))
	}

	fmt.Printf("val2: %+v\n", val2)

	val3 := &enum_payload{}
	if err := json.Unmarshal([]byte(JSON_STRING_1), val3); err != nil {
		fmt.Println(fmt.Sprintf("val3: WHOOPS! %s", err))
	}

	fmt.Printf("val3: %+v\n", val3)

	// The problem with this one is that you can do this though (create new values of the enum):
	myNewVal := MyEnumVal("some other value")
	ep := enum_payload{myNewVal}
	output, err := json.Marshal(ep)
	if err != nil {
		return
	}

	fmt.Printf("%s\n", output)
	val4 := &enum_payload{}
	err = json.Unmarshal(output, val4)
	if err != nil {
		fmt.Println(fmt.Sprintf("val4: WHOOPS! %s", err))
		return
	}
	fmt.Printf("%+v\n", val4)

	// the sealed (package implementation makes the interface a little cleaner

}
