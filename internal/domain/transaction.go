package domain

import "time"

type Transaction struct {
	Description string    `json:"description" validate:"required"`
	Amount      float64   `json:"amount" validate:"required,number"`
	Type        string    `json:"type" validate:"required"`
	Date        time.Time `json:"date"`
}
