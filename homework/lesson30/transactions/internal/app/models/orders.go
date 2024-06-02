package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    int
	User      User `gorm:"foreignKey:UserId"`
	ProductId int
	Product   Product `gorm:"foreignKey:ProductId"`
}

func (o Order) TableName() string {
	return "orders"
}
