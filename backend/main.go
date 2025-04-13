package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"blinket/blockchain"
	"github.com/gorilla/mux"
)

var bc *blockchain.Blockchain

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Shop        string `json:"shop"`
	OnBlinkit   bool   `json:"onBlinkit"`
	Location    string `json:"location,omitempty"`
}

var products = []Product{
	{ID: 1, Name: "Campus T-Shirt", Price: "₹499", Shop: "Campus Store", OnBlinkit: false},
	{ID: 2, Name: "Special Chai Mix", Price: "₹150", Shop: "Canteen", OnBlinkit: false},
}

func main() {
	bc = blockchain.NewBlockchain()
	
	r := mux.NewRouter()
	
	// Product endpoints
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products", addProduct).Methods("POST")
	
	// Blockchain endpoints
	r.HandleFunc("/blockchain", getBlockchain).Methods("GET")
	
	// Login endpoint
	r.HandleFunc("/login", login).Methods("POST")
	
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Add to blockchain
	bc.AddBlock(newProduct.ID)
	
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bc.GetBlocks())
}

func login(w http.ResponseWriter, r *http.Request) {
	// Simplified login (add proper auth later)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}