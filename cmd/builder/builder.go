package main

import (
	"fmt"
	"github.com/kjkondratuk/golang-patterns/cmd/builder/builder"
	"strings"
)

func main() {

	// Syntactically this can be nice because we can split up the creation into multiple steps
	b := builder.New().WithType("cool")

	b.WithHandler(func(m *builder.Message) *builder.Response {
		if strings.Contains(string(m.Body), "cool") {
			return &builder.Response{
				ResultCode: 200,
				ResultMessage: &builder.Message{
					Headers: nil,
					Body:    []byte("this message is cool"),
				},
			}
		}

		return &builder.Response{
			ResultCode: 200,
			ResultMessage: &builder.Message{
				Headers: nil,
				Body:    []byte("this message is NOT cool"),
			},
		}
	})

	service, err := b.Build()
	if err != nil {
		panic(err)
	}

	resp1 := service.Handle(&builder.Message{
		Headers: map[string]string{"type": "cool"},
		Body:    []byte("some body"),
	})

	fmt.Printf("%+v\n", string(resp1.ResultMessage.Body))

	resp2 := service.Handle(&builder.Message{
		Headers: map[string]string{"type": "cool"},
		Body:    []byte("some cool body"),
	})

	fmt.Printf("%+v\n", string(resp2.ResultMessage.Body))

	// You can also combine all builder steps into one
	_, err = builder.New().WithHandler(func(m *builder.Message) *builder.Response {
		return nil
	}).Build()

	if err != nil {
		fmt.Printf("There was an error: %s\n", err)
	}

	resp3 := service.Handle(&builder.Message{
		Headers: map[string]string{"type": "something else"},
		Body:    []byte("some cool body"),
	})

	fmt.Printf("%+v\n", resp3.ResultCode)

}
