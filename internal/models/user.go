package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	Avatar   string             `json:"avatar" bson:"avatar,omitempty"`
}

type ShortUser struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty" redis:"id"`
	Name   string             `json:"name" bson:"name" redis:"name"`
	Email  string             `json:"email" bson:"email" redis:"email"`
	Avatar string             `json:"avatar" bson:"avatar" redis:"avatar"`
}
