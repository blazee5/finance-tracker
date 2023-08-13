package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	Db *mongo.Client
}
