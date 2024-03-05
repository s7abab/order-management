package main

import (
	"fmt"
	"log"
	"net/http"
	"orders/db"
	"orders/handler"
)

func main() {

	// inii database connection and migrate models
	db.InitDb()

	// endpoints
	http.HandleFunc("/api/v1/order", handler.OrderHandler)
	http.HandleFunc("/api/v1/orders", handler.GetOrders)

	// start the server
	fmt.Println("Server is listening on port 8080 🪄")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
