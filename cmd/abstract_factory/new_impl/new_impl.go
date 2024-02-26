package new_impl

import "github.com/kjkondratuk/golang-patterns/cmd/abstract_factory/factory"

type myService struct {
}

var _ factory.GenericService = &myService{}
