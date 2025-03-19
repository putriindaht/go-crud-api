package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/putriindaht/go-crud-api/db"
)

// define structure
type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// fake db
var items []Item

// create an item
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem Item

	json.NewDecoder(r.Body).Decode(&newItem)

	collection := db.GetCollection("items")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newItem)
	if err != nil {
		println(err)
		http.Error(w, "Failed to insert item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)

}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get database and collection
	db.ConnectDb()

	router := mux.NewRouter()
	port := os.Getenv("PORT")

	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items", getItems).Methods("GET")

	fmt.Println("Server is running on port", port)
	http.ListenAndServe(":8080", router)
}
