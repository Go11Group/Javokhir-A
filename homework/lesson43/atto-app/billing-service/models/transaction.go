package models

type Transaction struct {
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	CardID          string  `json:"card_id"`
	UserID          string  `json:"user_id"`
	TransactionDate string  `josn:"transactionDate"`
}
