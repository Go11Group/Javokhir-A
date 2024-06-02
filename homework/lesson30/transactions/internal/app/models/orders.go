package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    int
	User      User
	ProductId int
	Product   Product
}
