package models

type Transaction struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	User        User    `json:"user" bson:"user"`
	Description string  `json:"description" bson:"description"`
	Amount      float64 `json:"amount" bson:"amount"`
	Type        string  `json:"type" bson:"type"`
}

type Analyze struct {
	TotalIncome  float64 `json:"total_income" bson:"total_income"`
	TotalExpense float64 `json:"total_expense" bson:"total_expense"`
	Total        float64 `json:"total" bson:"total"`
}
