package models

import "time"

type Card struct {
	CardID     string    `json:"card_id"`
	CardNumber string    `json:"card_number"`
	UserId     string    `json:"user_id"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}
