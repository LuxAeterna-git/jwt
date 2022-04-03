package repository

import (
	"context"
	"github.com/LuxAeterna-git/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type mongoData struct {
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

func (m *AuthMongo) CreateUser(user jwt.User) (string, error) {
	data := &mongoData{user.Name, user.Username, user.Password}
	usersCollection := m.db.Database("db").Collection("users")
	_, err := usersCollection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println(err)
	}

	usersCollection.FindOne(context.TODO(), data).Decode(&user)
	return user.ID.String(), nil
}

func (m *AuthMongo) GetUser(username, password string) (jwt.User, error) {
	var user jwt.User
	data := bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	}
	usersCollection := m.db.Database("db").Collection("users")
	err := usersCollection.FindOne(context.TODO(), data).Decode(&user)
	if err != nil {
		return jwt.User{}, err
	}
	return jwt.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}, nil
}
