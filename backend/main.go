package main

import (
	"backend/blockchain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var bc *blockchain.Blockchain

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Shop      string `json:"shop"`
	OnBlinkit bool   `json:"onBlinkit"`
	Location  string `json:"location,omitempty"`
}

var products = []Product{
	{ID: 1, Name: "Campus T-Shirt", Price: "₹0", Shop: "Campus Store", OnBlinkit: false},
	{ID: 2, Name: "Special Chai Mix", Price: "₹150", Shop: "Canteen", OnBlinkit: false},
}

var requestedItems = []Product{}

func main() {
	bc = blockchain.NewBlockchain()

	r := mux.NewRouter()

	// Product endpoints
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products", addProduct).Methods("POST")

	// Requested items endpoints
	r.HandleFunc("/requests", requestItem).Methods("POST")
	r.HandleFunc("/requests", getRequestedItems).Methods("GET")
	r.HandleFunc("/list-item", listItem).Methods("POST")

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

	// Add complete product to blockchain
	bc.AddBlock(blockchain.Product{
		ID:        newProduct.ID,
		Name:      newProduct.Name,
		Price:     newProduct.Price,
		Shop:      newProduct.Shop,
		OnBlinkit: newProduct.OnBlinkit,
		Location:  newProduct.Location,
	})

	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
}

func getRequestedItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requestedItems)
}

func requestItem(w http.ResponseWriter, r *http.Request) {
	var requestedItem Product
	if err := json.NewDecoder(r.Body).Decode(&requestedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestedItems = append(requestedItems, requestedItem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item requested successfully"})
}

func listItem(w http.ResponseWriter, r *http.Request) {
	var itemToBeListed Product
	if err := json.NewDecoder(r.Body).Decode(&itemToBeListed); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add item to blockchain
	bc.AddBlock(blockchain.Product{
		ID:        itemToBeListed.ID,
		Name:      itemToBeListed.Name,
		Price:     itemToBeListed.Price,
		Shop:      itemToBeListed.Shop,
		OnBlinkit: itemToBeListed.OnBlinkit,
		Location:  itemToBeListed.Location,
	})

	// Add item to products list
	products = append(products, itemToBeListed)

	// Remove from requested items
	for i, item := range requestedItems {
		if item.ID == itemToBeListed.ID {
			requestedItems = append(requestedItems[:i], requestedItems[i+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item listed successfully"})
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
