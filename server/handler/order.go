package handler

// Importing necessary packages
import (
	"encoding/json"
	"fmt"
	"net/http"

	"orders/db"
	"orders/models"
	"orders/validations"
	"strconv"
)

// Handle requests based on HTTP methods
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		getOrder(w, r)
	case "POST":
		createOrder(w, r)
	case "PUT":
		updateOrderStatus(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Create order handler (POST
func createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order

	// Decode the request body into the newOrder struct
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the newOrder
	if err := validations.ValidateOrder(newOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert into database
	if err := db.InsertOrder(&newOrder); err != nil {
		http.Error(w, "Failed to insert order into database", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New order created successfully")
}

// update order handler (PUT)
func updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Parse order id and status from url parameter
	fmt.Println(r.URL.Query())
	id := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	// validation
	if id == "" {
		http.Error(w, "order id is required", http.StatusBadRequest)
		return
	}
	if status == "" {
		http.Error(w, "order status is required", http.StatusBadRequest)
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

// Get single order handler (GET)
func getOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Validation
	if id == "" {
		http.Error(w, "Order id is required", http.StatusBadRequest)
		return
	}

	// Convert order ID to uint
	_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid order id", http.StatusBadRequest)
		return
	}

	// Retrieve the order from the database
	order, err := db.GetOrderByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve order", http.StatusInternalServerError)
		return
	}

	// Convert order to JSON
	jsonOrder, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to convert order to JSON", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonOrder)
}

// Get orders handler (GET)
func GetOrders(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	sortBy := r.URL.Query().Get("sortBy")
	sortOrder := r.URL.Query().Get("sortOrder")
	status := r.URL.Query().Get("status")
	currencyUnit := r.URL.Query().Get("currencyUnit")

	// Create filters map
	filters := make(map[string]interface{})
	if status != "" {
		filters["status"] = status
	}
	if currencyUnit != "" {
		filters["currencyUnit"] = currencyUnit
	}

	// Retrieve orders
	orders, totalPages, err := db.GetOrders(filters, page, pageSize, sortBy, sortOrder)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	// Convert orders and total pages to JSON
	response := struct {
		Orders     []models.Order `json:"orders"`
		TotalPages int            `json:"totalPages"`
	}{
		Orders:     orders,
		TotalPages: totalPages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to convert response to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
