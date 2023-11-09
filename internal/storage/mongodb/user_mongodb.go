package mongodb

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *UserDAO) Create(ctx context.Context, user models.User) (string, error) {
	res, err := db.c.InsertOne(ctx, models.User{Name: user.Name, Email: user.Email, Password: user.Password})

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (db *UserDAO) AddBalance(ctx context.Context, userId string, amount float64) error {
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	_, err = db.c.UpdateByID(ctx, objectId, bson.D{
		{"balance", bson.D{
			{
				"$sum", amount,
			}},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (db *UserDAO) GetUser(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	err := db.c.FindOne(ctx, bson.D{
		{"email", email},
		{"password", password},
	}).Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
