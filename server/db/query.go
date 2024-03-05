package db

import (
	"fmt"

	"orders/models"
)

// insert order
func InsertOrder(data *models.Order) error {
	db := GetDb()

	if err := db.Create(data).Error; err != nil {
		fmt.Println("Error inserting order:", err)
		return err
	}

	return nil
}

// update order status
func UpdateOrderStatus(id string, status string) error {
	db := GetDb()
	order := models.Order{}
	db.First(&order, "id = ?", id)
	fmt.Println(order)

	// update
	order.Status = status
	db.Save(&order)
	return nil
}
