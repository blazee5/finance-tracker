package mongodb

import (
	"context"
	"fmt"
	"github.com/blazee5/finance-tracker/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run(cfg *config.Config) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/",
		cfg.DBHost, cfg.DBPort))
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		return nil, err
	}

	return client, nil
}
