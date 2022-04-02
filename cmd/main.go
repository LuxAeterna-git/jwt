package main

import (
	"github.com/LuxAeterna-git/jwt"
	"github.com/LuxAeterna-git/jwt/pkg/handler"
	"github.com/LuxAeterna-git/jwt/pkg/repository"
	"github.com/LuxAeterna-git/jwt/pkg/service"
	"log"
)

func main() {
	repo := repository.NewSRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(jwt.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while runningserver: %s", err.Error())
	}
}
