package models

import (
	"gorm.io/gorm"
)

// Order struct
type Order struct {
	gorm.Model
	Status       string      `json:"status"`
	Total        float64     `json:"total"`
	CurrencyUnit string      `json:"currencyUnit"`
	Items        []OrderItem `json:"items"`
}

// Order items struct
type OrderItem struct {
	gorm.Model
	OrderID     uint   `json:"orderId" gorm:"index"` // Foreign key
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
}

// Initialize database connection
func InitDB(database *gorm.DB) {
	db = database
	// Auto migrate models
	db.AutoMigrate(&Order{}, &OrderItem{})
}

var db *gorm.DB
