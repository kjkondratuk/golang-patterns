package factory

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
