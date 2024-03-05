package main

import (
	"fmt"
	"log"
	"net/http"
	"orders/db"
	"orders/handler"
)

func main() {

	// Connect to the database and migrate models
	db.InitDb()

	// Register the handler for the API endpoint
	http.HandleFunc("/order", handler.ItemHandler)

	// Start the server
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
