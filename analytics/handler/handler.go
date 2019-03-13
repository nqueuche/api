package handler

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
	web "github.com/micro/go-web"
)

type Settings struct {
	Name    string
	Address string
}

func NewHandler(settings Settings) (web.Service, error) {
	if settings.Name == "" {
		return nil, fmt.Errorf("Name not provided")
	}

	if settings.Address == "" {
		return nil, fmt.Errorf("Address not provided")
	}

	router := httprouter.New()

	router.POST("/event", postEvent)
	router.POST("/user", postUser)

	service := web.NewService(
		web.Name(settings.Name),
		web.Address(settings.Address),
		web.Handler(router),
	)

	return service, nil
}
