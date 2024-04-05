package singleton

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type myService struct {
	url url.URL
	env string
}

type MyService interface {
	Get() (string, error)
}

// This is not threadsafe because updates to this value are not atomic, so we may not know which thread "wins" in setting this value
var instance MyService

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
	if instance == nil {
		instance = &myService{
			url: resolveUrl(),
			env: resolveEnv(),
		}
	}
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
