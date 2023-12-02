package mongodb

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/config"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TransactionRepository struct {
	db *mongo.Collection
}

func NewTransactionRepository(cfg *config.Config, client *mongo.Client) *TransactionRepository {
	return &TransactionRepository{
		db: client.Database(cfg.DBName).Collection("transactions"),
	}
}

func (repo *TransactionRepository) Create(ctx context.Context, user models.ShortUser, transaction domain.Transaction) (string, error) {
	if transaction.Date.IsZero() {
		transaction.Date = time.Now()
	}

	res, err := repo.db.InsertOne(ctx, models.Transaction{User: user, Type: transaction.Type, Category: transaction.Category, Amount: transaction.Amount, Description: transaction.Description, CreatedAt: transaction.Date})

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *TransactionRepository) GetTransactions(ctx context.Context, id, category string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"user._id", objectId},
	}

	if category != "" {
		filter = append(filter, bson.E{"category", category})
	}

	cursor, err := repo.db.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repo *TransactionRepository) GetTransaction(ctx context.Context, id string) (models.Transaction, error) {
	var transaction models.Transaction

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Transaction{}, err
	}

	err = repo.db.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&transaction)

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

func (repo *TransactionRepository) Update(ctx context.Context, id string, transaction domain.Transaction) (int, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, err
	}

	res, err := repo.db.UpdateOne(ctx, bson.D{
		{"_id", objectId}},
		bson.D{{"$set", bson.D{
			{"type", transaction.Type},
			{"amount", transaction.Amount},
			{"category", transaction.Category},
			{"description", transaction.Description},
			{"created_at", transaction.Date},
		}}})

	if err != nil {
		return 0, err
	}

	return int(res.ModifiedCount), nil
}

func (repo *TransactionRepository) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = repo.db.DeleteOne(ctx, bson.D{{"_id", objectId}})

	if err != nil {
		return err
	}

	return nil
}

func (repo *TransactionRepository) GetAnalyze(ctx context.Context, id string) (models.Analyze, error) {
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

	cursor, err := repo.db.Aggregate(ctx, pipeline)
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
