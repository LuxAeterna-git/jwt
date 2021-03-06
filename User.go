package jwt

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}
