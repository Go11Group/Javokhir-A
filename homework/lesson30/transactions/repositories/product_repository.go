package repositories

import "gorm.io/gorm"

type ProductRepository struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Db: db,
	}
}
