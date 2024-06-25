package service

import (
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/billing-service/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/billing-service/repository"
)

type CardService struct {
	CardRepository *repository.CardRepository
}

func NewCardService(cr *repository.CardRepository) *CardService {
	return &CardService{CardRepository: cr}
}

func (cs *CardService) CreateCard(user *repository.CreateCard) (*string, error) {
	if len(user.CardType) == 0 {
		return nil, fmt.Errorf("card type is empty")
	}

	if len(user.CardNumber) == 0 {
		return nil, fmt.Errorf("card number is empty")
	}

	if len(user.UserID) == 0 {
		return nil, fmt.Errorf("user id is empty")
	}

	return cs.CardRepository.CreateCard(*user)
}

func (cs *CardService) GetCard(cardID *string) (*models.Card, error) {

	if len(*cardID) == 0 {
		return nil, fmt.Errorf("card id is empty")
	}

	return cs.CardRepository.GetCard(cardID)
}
