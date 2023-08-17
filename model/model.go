package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Fname      string
	Username   string
	Email      string
	Address    string
	Province   string
	PostalCode string
	Country    string
	Phone      string
}
