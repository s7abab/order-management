package main

import (
	"fmt"
	"log"
	"net/http"
	"orders/db"
	"orders/handler"

	"github.com/rs/cors"
)

func main() {
	// initialize database connection and migrate models
	db.InitDb()

	// handlers
	http.HandleFunc("/api/v1/order", handler.OrderHandler)
	http.HandleFunc("/api/v1/orders", handler.GetOrders)

	// cors handlers
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// apply cors to all routes
	handler := c.Handler(http.DefaultServeMux)

	// start the server
	fmt.Println("Server is listening on port 8080 🪄")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
