package validations

import (
	"errors"
	"orders/models"
	"strconv"
)

// Validate order
func ValidateOrder(order models.Order) error {
	// Check if required fields are empty
	if len(order.Items) == 0 {
		return errors.New("order must have at least one item")
	}

	// Validate each item in the order
	for _, item := range order.Items {
		if item.Name == "" {
			return errors.New("item name is required")
		}

		// Convert Price and Quantity to number types
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return errors.New("invalid item price")
		}
		quantity, err := strconv.Atoi(item.Quantity)
		if err != nil {
			return errors.New("invalid item quantity")
		}

		if price <= 0 {
			return errors.New("item price must be greater than zero")
		}
		if quantity <= 0 {
			return errors.New("item quantity must be greater than zero")
		}
	}
	return nil
}
