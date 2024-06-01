package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string  `gorm:"not null"`
	Description   string  `gorm:"text"`
	Price         float64 `gorm:"not null"`
	StockQuantity int     `gorm:"not null"`
}
