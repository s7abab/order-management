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
	case "GET":
		GetOrdersHandler(w, r)
	case "POST":
		createOrder(w, r)
	case "PATCH":
		updateOrderStatus(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Create order (POST)
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

// update order (PATCH)
func updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Parse order id and status from url parameter
	id := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	// validation
	if id == "" {
		http.Error(w, "Order id is required", http.StatusBadRequest)
		return
	}
	if status == "" {
		http.Error(w, "Order status is required", http.StatusBadRequest)
		return
	}

	// Convert order ID to uint
	_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid order id", http.StatusBadRequest)
		return
	}

	db.UpdateOrderStatus(id, status)

	// Respond with a success message
	fmt.Fprintf(w, "Order status updated successfully")
}

// get orders (GET)
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// parse params
	filters := make(map[string]interface{})
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	sortBy := r.URL.Query().Get("sortBy")
	sortOrder := r.URL.Query().Get("sortOrder")

	// Retrieve orders
	orders, err := db.GetOrders(filters, page, pageSize, sortBy, sortOrder)
	if err != nil {

		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	// Convert orders to JSON
	jsonOrders, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Failed to convert orders to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonOrders)
}
