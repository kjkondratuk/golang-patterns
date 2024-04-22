package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("starting...")

	wg := sync.WaitGroup{}
	wg.Add(2)

	s := NewServer(&wg)
	//defer s.ShutdownServer()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				wg.Done()
				fmt.Println("recovering from: ", r)
				switch r.(type) {
				case mySpecialError:
					fmt.Println(fmt.Sprintf("the error type was mySpecialError: code - %d: error - %s", r.(mySpecialError).Code, r.(mySpecialError).Message))
				default:
					fmt.Println("unhandled exception")
				}
			}
		}()
		theThing(&wg, "something bad happened")
		//theThing(&wg, "baderr")
		// ...
		// ...
	}()

	go s.StartServer()

	wg.Wait()
	s.ShutdownServer()

	fmt.Println("finished.")
}

func theThing(wg *sync.WaitGroup, s string) {
	fmt.Println("executing.")
	switch s {
	case "ok":
		wg.Done()
		return
	case "baderr":
		panic("something really bad happened")
	default:
		panic(mySpecialError{
			Code:    500,
			Message: "something bad happened",
		})
	}
}

type mySpecialError struct {
	Code    int
	Message string
}

type server struct {
	wg *sync.WaitGroup
}

func NewServer(wg *sync.WaitGroup) *server {
	return &server{wg: wg}
}

func (s *server) StartServer() {
	fmt.Println("my server is running")
	time.Sleep(5 * time.Second)
	s.wg.Done()
}

func (s *server) ShutdownServer() {
	fmt.Println("server cleanup occurred")
}
