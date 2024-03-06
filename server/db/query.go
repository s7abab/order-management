package db

import (
	"fmt"
	"math"
	"orders/models"
)

// Insert order in to db
func InsertOrder(data *models.Order) error {
	db := GetDb()

	if err := db.Create(data).Error; err != nil {
		fmt.Println("Error inserting order:", err)
		return err
	}

	return nil
}

// Update order status
func UpdateOrderStatus(id string, status string) error {
	db := GetDb()
	order := models.Order{}
	db.First(&order, "id = ?", id)
	fmt.Println(order)

	// Update
	order.Status = status
	db.Save(&order)
	return nil
}

// Fetch orders from db
func GetOrders(filters map[string]interface{}, page int, pageSize int, sortBy string, sortOrder string) ([]models.Order, int, error) {
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

	// Execute count query to get total number of orders
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		fmt.Println("Error counting orders:", err)
		return nil, 0, err
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
		return nil, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	return orders, totalPages, nil
}

// Get order by id
func GetOrderByID(id string) (*models.Order, error) {
	db := GetDb()

	var order models.Order
	if err := db.Preload("Items").First(&order, id).Error; err != nil {
		fmt.Println("Error retrieving order:", err)
		return nil, err
	}

	return &order, nil
}
