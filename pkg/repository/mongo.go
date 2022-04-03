package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	*mongo.Client
	collection mongo.Collection
}
