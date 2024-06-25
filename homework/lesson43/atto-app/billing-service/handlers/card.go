package handlers

import (
	"database/sql"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/billing-service/repository"
	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/billing-service/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	CardService *service.CardService
}

func NewHandler(db *sql.DB) *Handler {

	cardSer := service.NewCardService(repository.NewCardRepository(db))

	return &Handler{CardService: cardSer}
}

func (h *Handler) CreateCard(ctx *gin.Context) {
	userID := ctx.Param("card_id")

	var card repository.CreateCard
	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cardID, err := h.CardService.CreateCard(&card)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Card created successfully with ID: " + *cardID})
}

func (h *Handler) GetCard(ctx *gin.Context) {
	cardID := ctx.Param("card_id")

	card, err := h.CardService.GetCard(&cardID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"card": card})
}
