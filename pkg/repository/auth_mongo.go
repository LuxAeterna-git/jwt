package repository

import (
	"context"
	"fmt"
	"github.com/LuxAeterna-git/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type mongoData struct {
	id       int
	Name     string
	Username string
	Password string
}

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (m *AuthMongo) CreateUser(user jwt.User) (int, error) {
	data := &mongoData{user.Id, user.Name, user.Username, user.Password}
	usersCollection := m.db.Database("db").Collection("users")
	_, err := usersCollection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println(err)
	}
	res := usersCollection.FindOne(context.TODO(), data)
	fmt.Println(res.DecodeBytes())
	return user.Id, nil
}

func (m *AuthMongo) GetUser(username, password string) (jwt.User, error) {
	var md mongoData
	data := bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	}
	usersCollection := m.db.Database("db").Collection("users")
	err := usersCollection.FindOne(context.TODO(), data).Decode(&md)
	if err != nil {
		return jwt.User{}, err
	}
	return jwt.User{
		Id:       md.id,
		Name:     md.Name,
		Username: md.Username,
		Password: md.Password,
	}, nil
}
