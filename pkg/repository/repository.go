package repository

import (
	"github.com/LuxAeterna-git/jwt"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user jwt.User) (int, error)
	GetUser(username, password string) (jwt.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
	}
}
