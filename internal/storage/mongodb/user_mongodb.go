package mongodb

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

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
