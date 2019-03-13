package main

import (
	"api/analytics/handler"
	"fmt"
	"log"
)

func init() {
	_, err := handler.NewDatabase()
	if err != nil {
		fmt.Println("ERR:", err)
	}
}

func main() {
	service, err := handler.NewHandler(handler.Settings{
		Name:    "cours",
		Address: ":1337",
	})

	if err != nil {
		log.Fatal(err)
	}

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
