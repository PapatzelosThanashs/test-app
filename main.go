package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Item represents a simple data structure with an ID and Name
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func main() {
	var err error

	// Open a database connection with proper error handling
	db, err = sql.Open("mysql", "root:example@tcp(service-db.jenkins.svc.cluster.local:3306)/test_db")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Wait for the database to be ready with retry logic
	waitForDB()

	// Close database when main function exits
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	// Start HTTP server
	http.HandleFunc("/items", itemsHandler)  // Handle requests to /items
	http.HandleFunc("/items/", itemHandler) // Handle requests to /items/{id}
	fmt.Println("Starting server on :8082...")
	log.Fatal(http.ListenAndServe(":8082", nil)) // Start server on port 8082
}

// waitForDB waits for the database to be available, retries on failure
func waitForDB() {
	for {
		// Ping the database to check for connectivity
		err := db.Ping()
		if err == nil {
			fmt.Println
