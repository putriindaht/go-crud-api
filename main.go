package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// define structure
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// fake db
var items []Item
var idCounter = 1

// create an item
func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)

	newItem.ID = idCounter
	idCounter++
	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)

}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items", getItems).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}
