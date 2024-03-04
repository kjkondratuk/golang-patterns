package factory

import (
	"errors"
	"fmt"
	"sync"
)

var ConstructorAlreadyRegisteredError = errors.New("constructor already registered")

type genericService struct {
	msg string
}

type GenericService interface {
	GetValue() string
}

func (s *genericService) GetValue() string {
	return s.msg
}

var _ GenericService = &localService{}

type localService struct {
	genericService
}

var _ GenericService = &deployedService{}

type deployedService struct {
	genericService
}

var (
	constructors = make(map[string]func() GenericService)
	mutex        = sync.RWMutex{}
)

func New(env string) GenericService {
	mutex.RLock()
	defer mutex.RUnlock()

	if constructor, ok := constructors[env]; ok {
		return constructor()
	}

	switch env {
	case "local":
		return &localService{genericService{"local implementation"}}
	case "deployed":
		return &deployedService{genericService{"deployed implementation"}}
	default:
		return &genericService{"default implementation"}
	}
}

func Register(env string, constructor func() GenericService) error {

	if constructor == nil {
		panic("constructor is nil")
	}

	if _, exists := constructors[env]; exists {
		return ConstructorAlreadyRegisteredError
	}

	mutex.Lock()
	defer mutex.Unlock()

	constructors[env] = constructor

	fmt.Printf("Registered %s environment\n", env)
	return nil
}
