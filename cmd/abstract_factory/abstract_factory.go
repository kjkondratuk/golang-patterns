package main

import (
	"fmt"
	"github.com/kjkondratuk/golang-patterns/cmd/abstract_factory/factory"
	"os"
)

func main() {
	//services := []GenericService{
	//	&genericService{"default implementation"},
	//	&localService{genericService{"local implementation"}},
	//	&deployedService{genericService{"deployed implementation"}},
	//}

	//services := []GenericService{
	//	New("default"),
	//	New("local"),
	//	New("deployed"),
	//}
	//
	//for _, s := range services {
	//	fmt.Printf("%s\n", s.GetValue())
	//}

	v := os.Getenv("ENV")

	s := factory.New(v)
	fmt.Printf("%s\n", s.GetValue())
}
