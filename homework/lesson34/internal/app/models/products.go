package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(50);not null" json:"product_name"`
	Description string  `gorm:"text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
}

func (p Product) TableName() string {
	return "products"
}
