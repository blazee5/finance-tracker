package mongodb

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *TransactionDAO) Create(ctx context.Context, user models.User, transaction domain.Transaction) (string, error) {
	res, err := db.c.InsertOne(ctx, models.Transaction{User: user, Type: transaction.Type, Amount: transaction.Amount, Description: transaction.Description})

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (db *TransactionDAO) GetTransactions(ctx context.Context, id string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	cursor, err := db.c.Find(ctx, bson.D{{"user._id", objectId}})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
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

func (db *TransactionDAO) Update(ctx context.Context, id string, transaction domain.Transaction) (int, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, err
	}

	res, err := db.c.UpdateOne(ctx, bson.D{
		{"_id", objectId}},
		bson.D{{"$set", bson.D{
			{"type", transaction.Type},
			{"amount", transaction.Amount},
			{"description", transaction.Description},
		}}})

	if err != nil {
		return 0, err
	}

	return int(res.ModifiedCount), nil
}

func (db *TransactionDAO) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = db.c.DeleteOne(ctx, bson.D{{"_id", objectId}})

	if err != nil {
		return err
	}

	return nil
}

func (db *TransactionDAO) GetAnalyze(ctx context.Context, id string) (models.Analyze, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Analyze{}, err
	}

	pipeline := []bson.D{
		{
			{"$match", bson.M{"user._id": objectId}},
		},
		{
			{"$group", bson.D{
				{"_id", 0},
				{"total_income", bson.D{{"$sum", bson.D{{"$cond", bson.A{bson.D{{"$eq", bson.A{"$type", "income"}}}, "$amount", 0}}}}}},
				{"total_expense", bson.D{{"$sum", bson.D{{"$cond", bson.A{bson.D{{"$eq", bson.A{"$type", "expense"}}}, "$amount", 0}}}}}},
			}},
		},
		{
			{"$project", bson.D{
				{"_id", 0},
				{"total_income", 1},
				{"total_expense", 1},
				{"total", bson.D{{"$subtract", bson.A{"$total_income", "$total_expense"}}}},
			}},
		},
	}

	var analyze models.Analyze

	cursor, err := db.c.Aggregate(ctx, pipeline)
	if err != nil {
		return models.Analyze{}, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if err := cursor.Decode(&analyze); err != nil {
			return models.Analyze{}, err
		}
	}

	return analyze, nil
}
