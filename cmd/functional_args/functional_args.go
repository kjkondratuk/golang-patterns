package main

import (
	"fmt"
	"os"
)

type myStruct struct {
	// required
	structValue string

	// optional array
	supplementalValues []string

	debug bool

	// optional complex value
	cval complexValue
}

type complexValue struct {
	value1 string
	value2 string
}

type mutableOption func(s *myStruct)
type complexValueOption func(s complexValue) complexValue

type unmutableOption func(s myStruct) myStruct

func WithSupplementalValues(vals ...string) unmutableOption {
	return func(s myStruct) myStruct {
		s.supplementalValues = vals
		return s
	}
}

func WithComplexValue(cv ...complexValueOption) unmutableOption {
	return func(s myStruct) myStruct {
		for _, o := range cv {
			s.cval = o(s.cval)
		}
		return s
	}
}

func WithDebug() unmutableOption {
	return func(s myStruct) myStruct {
		ns := s
		ns.debug = true
		return ns
	}
}

func BadNew(sval string, cval complexValue, supplementalValues []string) *myStruct {
	return &myStruct{
		structValue:        sval,
		supplementalValues: supplementalValues,
		cval:               cval,
	}
}

func GoodNew(sval string, opts ...unmutableOption) *myStruct {
	s := myStruct{
		structValue: sval,
		cval: complexValue{
			value2: "default value",
		},
	}
	for _, opt := range opts {
		s = opt(s)
	}
	return &s
}

func main() {
	v1 := BadNew("my string value", complexValue{}, nil)
	fmt.Printf("%+v\n", v1)

	v2 := GoodNew("my string value")
	fmt.Printf("%+v\n", v2)

	v3 := GoodNew("my other string value", WithComplexValue(func(s complexValue) complexValue {
		s.value1 = "some value"
		return s
	}))
	fmt.Printf("%+v\n", v3)

	v4 := GoodNew("my other string value", WithSupplementalValues("a supplemental value", "a supplemental value 2"))
	fmt.Printf("%+v\n", v4)

	opts := []unmutableOption{
		WithSupplementalValues("a supplemental value", "a supplemental value 2"),
	}

	if os.Getenv("USE_COMPLEX") != "" {
		opts = append(opts, WithComplexValue(func(s complexValue) complexValue {
			return s
		}))
	}

	v5 := GoodNew("my other string value", opts...)
	fmt.Printf("%+v\n", v5)

	v6 := GoodNew("with debug", WithDebug())
	fmt.Printf("%+v\n", v6)

	v7 := myStruct{}
	fmt.Printf("%s", v7.supplementalValues)

}
