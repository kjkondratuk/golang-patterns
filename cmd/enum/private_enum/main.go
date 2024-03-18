package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kjkondratuk/golang-patterns/cmd/enum/private_enum/enum"
)

type enum_payload struct {
	Value enum.MyEnumVal `json:"stringVal"`
}

// The downside of this pattern is that, since we're depending on the enum.New function being called, anyone using the
// value in a payload will have to manually unmarshal it.
func (p *enum_payload) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	strVal, ok := m["stringVal"].(string)
	if !ok {
		return errors.New("invalid MyEnumVal type")
	}
	val, err := enum.New(strVal)
	if err != nil {
		return err
	}
	p.Value = val
	return nil
}

func main() {
	// Can't do this because the struct is private
	//myNewVal := myEnumVal("some other value")

	myNewVal, err := enum.New("some other value")
	if err != nil {
		fmt.Println("Expected error: ", err)
	}

	v := enum.MY_VALUE_3
	myNewVal = &v

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
	fmt.Printf("%s\n", val4.Value.Value())
}
