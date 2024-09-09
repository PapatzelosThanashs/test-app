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
			fmt.Println("Database is ready")
			return
		}
		fmt.Printf("Waiting for database to be ready... Error: %v\n", err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}
}

// itemsHandler handles requests for the items collection
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		createItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// itemHandler handles requests for individual items
func itemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):] // Get the ID from the URL path
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getItem(w, r, id)
	case http.MethodPut:
		updateItem(w, r, id)
	case http.MethodDelete:
		deleteItem(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getItems retrieves all items from the database
func getItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var itemsList []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		itemsList = append(itemsList, item)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemsList)
}

// createItem adds a new item to the database
func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO items (name) VALUES (?)", newItem.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newItem.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// getItem retrieves a single item by ID
func getItem(w http.ResponseWriter, r *http.Request, id int) {
	var item Item
	err := db.QueryRow("SELECT id, name FROM items WHERE id = ?", id).Scan(&item.ID, &item.Name)
	if err == sql.ErrNoRows {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// updateItem modifies an existing item in the database
func updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil || updatedItem.ID != id {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE items SET name = ? WHERE id = ?", updatedItem.Name, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedItem)
}

// deleteItem removes an item by ID from the database
func deleteItem(w http.ResponseWriter, r *http.Request, id int) {
	_, err := db.Exec("DELETE FROM items WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
