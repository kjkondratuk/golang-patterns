package main

import (
	"fmt"
)

const (
	VALUE_1 = iota
	VALUE_2
	VALUE_3
	VALUE_4
)

func main() {
	value := 1
	switch value {
	case VALUE_2:
		fmt.Println("The value was value 2!")
	default:
		fmt.Println("The value was some other value.")
	}
}
