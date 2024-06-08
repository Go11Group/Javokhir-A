package handlers

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson34/repositories"
	"gorm.io/gorm"
)

type Handlers struct {
	repos *repositories.UniRepo
}

func NewHandlers(db *gorm.DB) *Handlers {
	Repos := repositories.NewUniRepo(db)
	return &Handlers{repos: Repos}
}
