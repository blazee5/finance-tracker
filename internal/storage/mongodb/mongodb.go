package mongodb

import (
	"context"
	"fmt"
	"github.com/blazee5/finance-tracker/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Db             *mongo.Client
	UserDAO        *UserDAO
	TransactionDAO *TransactionDAO
}

type UserDAO struct {
	c *mongo.Collection
}

type TransactionDAO struct {
	c *mongo.Collection
}

func NewUserDAO(client *mongo.Client, cfg *config.Config) (*UserDAO, error) {
	return &UserDAO{
		c: client.Database(cfg.DBName).Collection("users"),
	}, nil
}

func NewTransactionDAO(client *mongo.Client, cfg *config.Config) (*TransactionDAO, error) {
	return &TransactionDAO{
		c: client.Database(cfg.DBName).Collection("transactions"),
	}, nil
}

func Run(cfg *config.Config) (*Storage, error) {
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/",
		cfg.DBHost, cfg.DBPort))
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		return nil, err
	}

	return &Storage{Db: client}, nil
}
