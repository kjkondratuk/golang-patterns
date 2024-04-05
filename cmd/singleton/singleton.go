package main

import (
	"fmt"
	"github.com/kjkondratuk/golang-patterns/cmd/singleton/almost_singleton"
	"github.com/kjkondratuk/golang-patterns/cmd/singleton/singleton"
	"github.com/kjkondratuk/golang-patterns/cmd/singleton/sync_singleton"
	"github.com/kjkondratuk/golang-patterns/cmd/singleton/sync_singleton_err"
	"log"
)

func main() {
	// so this works
	resp, err := almost_singleton.MyServ.Get()
	if err != nil {
		log.Fatalf("There was an error: %s", err)
	}
	fmt.Println(resp)

	// but...
	almost_singleton.MyServ = &NotMyService{}

	// This is not thread-safe because we're using a non-atomic value and no mutex to lock updates to the instance
	var single singleton.MyService
	single = singleton.New()
	fmt.Printf("single: %s\n", &single)

	// this is the same memory address, because this constructor call doesn't actually create a new instance, just returns the existing one
	single = singleton.New()
	fmt.Printf("single: %s\n", &single)

	// using the sync package, we can ensure that calls to the constructor are thread-safe
	var syncSingle sync_singleton.MyService
	syncSingle = sync_singleton.New()
	fmt.Printf("single: %s\n", &syncSingle)

	single = sync_singleton.New()
	fmt.Printf("single: %s\n", &syncSingle)

	// using the sync package, but returning a sentinel error when the value has already been initialized
	var syncSingleErr sync_singleton_err.MyService
	syncSingleErr, err = sync_singleton_err.New()
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("single: %s\n", &syncSingleErr)

	syncSingleErr, err = sync_singleton_err.New()
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("single: %s\n", &syncSingleErr)

}

type NotMyService struct {
}

func (s *NotMyService) Get() (string, error) {
	return "", nil
}
