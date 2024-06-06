package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string `gorm:"not nul" json:"user_name"`
	FirstName string `gorm:"not nul" json:"first_name"`
	LastName  string `gorm:"not nul" json:"last_name"`
	Email     string `gorm:"not nul; unique" json:"gmail"`
}
