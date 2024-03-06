package db

import (
	"fmt"

	"orders/models"
)

// insert order (POST)
func InsertOrder(data *models.Order) error {
	db := GetDb()

	if err := db.Create(data).Error; err != nil {
		fmt.Println("Error inserting order:", err)
		return err
	}

	return nil
}

// update order status (PATCH)
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

func GetOrders(filters map[string]interface{}, page int, pageSize int, sortBy string, sortOrder string) ([]models.Order, error) {
	db := GetDb()

	query := db.Model(&models.Order{}).Preload("Items")

	// Apply filters
	for key, value := range filters {
		switch key {
		case "status":
			query = query.Where("status = ?", value)
		case "currencyUnit":
			query = query.Where("currency_unit = ?", value)
		default:
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	// Apply sorting
	if sortBy != "" {
		if sortOrder == "DESC" {
			query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
		} else {
			query = query.Order(fmt.Sprintf("%s ASC", sortBy))
		}
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// Execute query
	var orders []models.Order
	if err := query.Find(&orders).Error; err != nil {
		fmt.Println("Error retrieving orders:", err)
		return nil, err
	}

	return orders, nil
}

// get order by id (GET)
func GetOrderByID(id string) (*models.Order, error) {
	db := GetDb()

	var order models.Order
	if err := db.Preload("Items").First(&order, id).Error; err != nil {
		fmt.Println("Error retrieving order:", err)
		return nil, err
	}

	return &order, nil
}
