package main

import (
	"github.com/LuxAeterna-git/jwt"
	"github.com/LuxAeterna-git/jwt/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(jwt.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while runningserver: %s", err.Error())
	}
}
