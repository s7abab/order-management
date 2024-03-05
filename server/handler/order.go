package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"orders/db"
	"orders/models"
	"strconv"
)

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	// Set response content type
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		createOrder(w, r)
	case "PATCH":
		updateOrderStatus(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Create order
func createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order

	// Decode the request body into the newOrder struct
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert into database
	if err := db.InsertOrder(&newOrder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New order created successfully")
}

// update order
func updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Parse order id and status from url parameter
	id := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	// validation
	if id == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}
	if status == "" {
		http.Error(w, "Order Status is required", http.StatusBadRequest)
		return
	}

	// Convert order ID to uint
	_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	db.UpdateOrderStatus(id, status)

	// Respond with a success message
	fmt.Fprintf(w, "Order status updated successfully")
}
