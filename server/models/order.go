package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Status       string      `json:"status"`
	Total        float64     `json:"total"`
	CurrencyUnit string      `json:"currencyUnit"`
	Items        []OrderItem `json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID     uint    `gorm:"index"` // Foreign key
	Order       Order   `gorm:"foreignKey:OrderID"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

// Initialize database connection
func InitDB(database *gorm.DB) {
	db = database
	// Auto migrate models
	db.AutoMigrate(&Order{}, &OrderItem{})
}

var db *gorm.DB
