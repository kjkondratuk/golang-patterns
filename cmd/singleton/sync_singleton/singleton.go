package sync_singleton

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type myService struct {
	url url.URL
	env string
}

type MyService interface {
	Get() (string, error)
}

var (
	instance MyService
	oncer    *sync.Once = &sync.Once{}
)

func resolveUrl() url.URL {
	u, err := url.Parse("https://google.com")
	if err != nil {
		log.Fatal("could not wire up service URL")
	}
	return *u
}

func resolveEnv() string {
	return "prod"
}

func New() MyService {
	oncer.Do(func() {
		instance = &myService{
			url: resolveUrl(),
			env: resolveEnv(),
		}
	})
	return instance
}

func (s *myService) Get() (string, error) {
	response, err := http.DefaultClient.Get(s.url.String())
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	if response.Body != nil {
		all, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(all), nil
	}

	return "", nil
}
