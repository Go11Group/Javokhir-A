package repositories

import (
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/internal/app/models"
	"gorm.io/gorm"
)

type Fetcher struct {
	Db *gorm.DB
}

func NewFetcher(db *gorm.DB) *Fetcher {
	return &Fetcher{
		Db: db,
	}
}

func (f *Fetcher) FetchAll(result interface{}) error {
	tableName := ""

	switch result.(type) {
	case *[]models.User:
		tableName = (&models.User{}).TableName()
	case *[]models.Product:
		tableName = (&models.Product{}).TableName()
	case *[]models.Order:
		tableName = (&models.Order{}).TableName()
	}

	if tableName == "" {
		return fmt.Errorf("invalid model type")
	}

	if err := f.Db.Table(tableName).Find(result).Error; err != nil {
		return fmt.Errorf("failed to fetch all records: %v", err)
	}

	return nil
}
