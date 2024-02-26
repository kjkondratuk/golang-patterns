package main

import (
	"fmt"
	"os"
)

type genericService struct {
	msg string
}

type GenericService interface {
	GetValue() string
}

func (s *genericService) GetValue() string {
	return s.msg
}

// var _ GenericService = localService{}
type localService struct {
	genericService
}

func (s *localService) GetValue() string {
	return s.msg
}

// var _ GenericService = deployedService{}
type deployedService struct {
	genericService
}

func (s *deployedService) GetValue() string {
	return s.msg
}

// TODO: implement an extensible New function that allows the caller to define a new GenericService implementation
func New(env string) GenericService {
	if env == "local" {
		return &localService{genericService{"local implementation"}}
	} else if env == "deployed" {
		return &deployedService{genericService{"deployed implementation"}}
	} else {
		return &genericService{"default implementation"}
	}
}

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

	s := New(v)
	fmt.Printf("%s\n", s.GetValue())
}
