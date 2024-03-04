package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kjkondratuk/golang-patterns/cmd/abstract_factory/factory"
)

type customService struct {
	msg string
}

func (cs *customService) GetValue() string {
	return cs.msg
}

func main() {

	err := factory.Register("custom", func() factory.GenericService {
		return &customService{msg: "custom implementation"}
	})
	if err != nil {
		log.Printf("error registering custom service: %v", err)
		return
	}

	err = factory.Register("stage", func() factory.GenericService {
		return &customService{msg: "stage implementation"}
	})
	if err != nil {
		log.Printf("error registering custom service: %v", err)
		return
	}

	envServices := []factory.GenericService{
		factory.New("default"),
		factory.New("local"),
		factory.New("deployed"),
		factory.New("custom"),
	}

	for _, s := range envServices {
		fmt.Printf("%s\n", s.GetValue())
	}

	v := os.Getenv("ENV")

	s := factory.New(v)
	fmt.Printf("ENV: %s\n", v)
	fmt.Printf("%s\n", s.GetValue())
}
