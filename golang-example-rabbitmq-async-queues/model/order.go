package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerName string
	ProductName  string
	Quantity     int
	Status       string
}
