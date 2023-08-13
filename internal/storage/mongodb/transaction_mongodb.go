package mongodb

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *TransactionDAO) GetTransactions(ctx context.Context, id string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	cursor, err := db.c.Find(ctx, bson.D{{"user_id", id}})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (db *TransactionDAO) Create(ctx context.Context, transaction models.Transaction) (interface{}, error) {
	res, err := db.c.InsertOne(ctx, models.Transaction{UserID: transaction.UserID, Type: transaction.Type, Amount: transaction.Amount, Description: transaction.Description})
	if err != nil {
		return 0, err
	}

	return res.InsertedID, nil
}

func (db *TransactionDAO) GetTransaction(ctx context.Context, id string) (models.Transaction, error) {
	var transaction models.Transaction

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Transaction{}, err
	}

	err = db.c.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&transaction)

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

func (db *TransactionDAO) Update(ctx context.Context, transaction models.Transaction) (int64, error) {
	objectId, err := primitive.ObjectIDFromHex(transaction.ID)
	if err != nil {
		return 0, err
	}
	res, err := db.c.UpdateOne(ctx, bson.D{{"_id", objectId}}, bson.D{{"$set", bson.D{{"type", transaction.Type}, {"amount", transaction.Amount}, {"description", transaction.Description}}}})
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

func (db *TransactionDAO) Delete(ctx context.Context, id string) (int64, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	res, err := db.c.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
