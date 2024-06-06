package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Age       int
	Filed     string
	Geneder   string
	IsEmploee bool
}
