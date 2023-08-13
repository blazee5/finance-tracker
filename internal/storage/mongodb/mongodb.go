package mongodb

import (
	"context"
	"fmt"
	"github.com/blazee5/finance-tracker/internal/config"
	"github.com/blazee5/finance-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (db *UserDAO) Create(ctx context.Context, user models.User) (interface{}, error) {
	res, err := db.c.InsertOne(ctx, models.User{Name: user.Name, Email: user.Email, Password: user.Password})
	if err != nil {
		return 0, err
	}

	return res.InsertedID, nil
}

func (db *UserDAO) GetUser(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	err := db.c.FindOne(ctx, bson.D{{"email", email}, {"password", password}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
