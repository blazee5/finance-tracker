package domain

type Transaction struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
}
