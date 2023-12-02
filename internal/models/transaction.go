package models

import "time"

type Transaction struct {
	ID          string    `json:"id" redis:"id" bson:"_id,omitempty"`
	User        ShortUser `json:"user" redis:"user" bson:"user"`
	Description string    `json:"description" redis:"description" bson:"description"`
	Amount      float64   `json:"amount" redis:"amount" bson:"amount"`
	Type        string    `json:"type" redis:"type" bson:"type"`
	Category    string    `json:"category" redis:"category" bson:"category"`
	CreatedAt   time.Time `json:"created_at" redis:"created_at" bson:"created_at"`
}

type Analyze struct {
	TotalIncome  float64 `json:"total_income" bson:"total_income"`
	TotalExpense float64 `json:"total_expense" bson:"total_expense"`
	Total        float64 `json:"total" bson:"total"`
}
