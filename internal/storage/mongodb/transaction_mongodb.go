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

func (db *TransactionDAO) GetAnalyze(ctx context.Context, id string) ([]models.Analyze, error) {
	var analyze []models.Analyze

	cursor, err := db.c.Aggregate(ctx, []bson.M{
		{"$match": bson.M{"user_id": id}},
		{"$group": bson.M{
			"_id":           nil,
			"total_income":  bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []interface{}{"$type", "income"}}, "$amount", 0}}},
			"total_expense": bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []interface{}{"$type", "expense"}}, "$amount", 0}}},
		}},
		{"$project": bson.M{
			"_id":           0,
			"total_income":  1,
			"total_expense": 1,
			"total":         bson.M{"$subtract": []interface{}{"$total_income", "$total_expense"}},
		}},
	})

	if err != nil {
		return []models.Analyze{}, err
	}

	if err = cursor.All(ctx, &analyze); err != nil {
		return []models.Analyze{}, err
	}

	return analyze, nil
}
