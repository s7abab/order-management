package main

// Importing necessary packages
import (
	"fmt"
	"log"
	"net/http"
	"orders/db"
	"orders/handler"

	"github.com/rs/cors"
)

func main() {
	// Initialize database connection and migrate models
	db.InitDb()

	// API Request Handlers
	http.HandleFunc("/api/v1/order", handler.OrderHandler)
	http.HandleFunc("/api/v1/orders", handler.GetOrders)

	// Cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Apply cors to all routes
	handler := c.Handler(http.DefaultServeMux)

	// Start the server
	fmt.Println("Server is listening on port 8080 🪄")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
