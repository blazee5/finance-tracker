package mongodb

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/config"
	"github.com/blazee5/finance-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(cfg *config.Config, client *mongo.Client) *UserRepository {
	return &UserRepository{
		db: client.Database(cfg.DBName).Collection("users"),
	}
}

func (repo *UserRepository) Create(ctx context.Context, user models.User) (string, error) {
	res, err := repo.db.InsertOne(ctx, models.User{Name: user.Name, Email: user.Email, Password: user.Password})

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *UserRepository) GetUser(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	err := repo.db.FindOne(ctx, bson.D{
		{"email", email},
		{"password", password},
	}).Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, id string) (models.ShortUser, error) {
	var user models.ShortUser

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.ShortUser{}, err
	}

	err = repo.db.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&user)

	if err != nil {
		return models.ShortUser{}, err
	}

	return user, nil
}
