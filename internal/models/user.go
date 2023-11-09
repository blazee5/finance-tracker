package models

type User struct {
	ID       string  `json:"id" bson:"_id,omitempty"`
	Name     string  `json:"name" bson:"name"`
	Email    string  `json:"email" bson:"email"`
	Balance  float64 `json:"balance" bson:"balance"`
	Password string  `json:"password" bson:"password"`
}
