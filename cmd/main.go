package main

import (
	"log"

	gotodo "github.com/bvckslvsh/go-to-do"
	"github.com/bvckslvsh/go-to-do/pkg/handler"
	"github.com/bvckslvsh/go-to-do/pkg/repository"
	"github.com/bvckslvsh/go-to-do/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gotodo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running : %s", err.Error())
	}
}
