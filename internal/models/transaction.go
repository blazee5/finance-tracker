package models

type Transaction struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	UserID      string `json:"user_id" bson:"user_id"`
	Description string `json:"description" bson:"description"`
	Amount      int    `json:"amount" bson:"amount"`
	Type        string `json:"type" bson:"type"`
}
