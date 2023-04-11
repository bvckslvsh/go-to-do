package main

import (
	"log"

	gotodo "github.com/bvckslvsh/go-to-do"
	"github.com/bvckslvsh/go-to-do/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(gotodo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running : %s", err.Error())
	}
}
