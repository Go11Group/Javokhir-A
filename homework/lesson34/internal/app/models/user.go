package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"not null;size:50" json:"user_name"`
	Email    string `gorm:"not null;unique;size:50" json:"email"`
	Password string `gorm:"not null;size:50" json:"password"`
}

func (User) TableName() string {
	return "users"
}
