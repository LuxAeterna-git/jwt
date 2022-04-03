package main

import (
	"context"
	"github.com/LuxAeterna-git/jwt"
	"github.com/LuxAeterna-git/jwt/pkg/handler"
	"github.com/LuxAeterna-git/jwt/pkg/repository"
	"github.com/LuxAeterna-git/jwt/pkg/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {

	//connect with mongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tmp, err := mongo.Connect(ctx1, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//check if db avaible

	err = tmp.Ping(ctx1, nil)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(tmp)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(jwt.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while runningserver: %s", err.Error())
	}
}
