package mongodb

import (
	"context"
	"fmt"
	"github.com/blazee5/finance-tracker/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run(cfg *config.Config) *mongo.Client {
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/",
		cfg.DBHost, cfg.DBPort))
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	return client
}
