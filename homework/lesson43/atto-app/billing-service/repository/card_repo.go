package repository

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/billing-service/models"
	"github.com/google/uuid"
)

type CardRepository struct {
	Db *sql.DB
}

type CreateCard struct {
	CardType   string `json:"cardType"`
	CardNumber string `json:"card_number"`
	UserID     string `json:"user_id"`
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{Db: db}
}

func (cr *CardRepository) CreateCard(card CreateCard) (*string, error) {
	query := `
		INSERT INTO cards (card_id, user_id, card_number, card_type)
        VALUES ($1, $2, $3, $4)
    `
	smt, err := cr.Db.Prepare(query)

	if err != nil {
		return nil, fmt.Errorf("error preparing query: %v", err)
	}

	newCardID := uuid.New().String()

	if _, err := smt.Exec(newCardID, card.UserID, card.CardType, card.CardType); err != nil {
		return nil, fmt.Errorf("error creating card: %v", err)
	}

	return &newCardID, nil
}

func (cr *CardRepository) GetCard(cardID *string) (*models.Card, error) {
	var card models.Card
	card.CardID = *cardID

	query := `
        SELECT card_number, card_type, created_at, updated_at FROM cards WHERE card_id = $1 AND deleted_at IS NULL
    `
	smt, err := cr.Db.Prepare(query)

	if err != nil {
		return nil, fmt.Errorf("error preparing query: %v", err)
	}

	if err := smt.QueryRow(cardID).Scan(&card.CardNumber, &card.CardNumber, &card.CreatedAt, &card.UpdatedAt); err != nil {
		return nil, fmt.Errorf("error getting card: %v", err)
	}

	return &card, nil
}
