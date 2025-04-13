# Blinket Gap Filler 🛍️

This is a hackathon project that supports local campus shopkeepers by showcasing items **not available on Blinkit**, but available inside our campus.

## 💡 Problem

Due to Blinkit delivery being widely available, students are ignoring local shops inside the campus. This is reducing their sales even though they offer unique and essential products.

## ✅ Solution

Our website helps:
- Highlight products that **are not on Blinkit**
- Promote **in-campus shops**
- Allow students to discover what’s available nearby

## 📁 Features

- 🛒 Student Login Page
- 🏠 Homepage to browse local shop products
- 🔍 Filter products not available on Blinkit
- 📍 Shop location (optional)
- 🧑‍💻 Admin/Shopkeeper panel (future enhancement)

## 🛠️ Tech Stack

- **Frontend:** HTML, CSS, JavaScript
- **Backend:** Golang
- **Database:** SQLite or PostgreSQL (TBD)

## 📂 Project Structure

blinket-gap-filler/
├── backend/
│   ├── main.go
│   ├── routes/
│   ├── models/
│   └── db/
├── frontend/
│   ├── index.html
│   ├── home.html
│   └── js/
├── .gitignore
├── LICENSE
└── README.md

## 🚀 Getting Started

To get started with the project locally, follow these steps:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/<your-username>/blinket-gap-filler.git

2. **Install dependencies**:

   #### For the Backend (Go):
   ```bash
   cd backend
   go mod tidy

3. **Initialize Go module**:
    ```bash
   go mod init github.com/<your-username>/blinket-gap-filler

4. **Run the Backend Server**:
   ```bash
   go run main.go

## 🌟 Features

- User login and homepage
- Listings of campus-only items
- Basic inventory and comparison with Blinkit

## 🚧 Future Scope

- Admin portal for shopkeepers
- Real-time inventory sync
- Mobile-friendly UI and delivery tracking

## 📂 New Project Structure

/blinket-gap-filler
├── backend/
│   ├── go.mod
│   ├── main.go
│   ├── blockchain/
│   │   └── blockchain.go
│   └── products/
│       └── products.go
└── frontend/
    ├── public/
    ├── src/
    │   ├── components/
    │   │   ├── Login.js
    │   │   ├── ProductList.js
    │   │   └── Navbar.js
    │   ├── App.js
    │   └── index.js
    ├── package.json
    └── .env

## 🔗 Blockchain Code (`blockchain.go`)
```go
 package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"sync"
)

type Block struct {
	Index        int
	Timestamp    string
	ProductID    int
	PrevHash     string
	Hash         string
}

type Blockchain struct {
	blocks []*Block
	mutex  sync.Mutex
}

func NewBlockchain() *Blockchain {
	genesisBlock := &Block{
		Index:     0,
		Timestamp: time.Now().String(),
		ProductID: 0,
		PrevHash:  "",
		Hash:      "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return &Blockchain{
		blocks: []*Block{genesisBlock},
	}
}

func (bc *Blockchain) AddBlock(productID int) *Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().Format(time.RFC3339),
		ProductID: productID,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	bc.blocks = append(bc.blocks, newBlock)
	return newBlock
}

func calculateHash(block *Block) string {
	record := string(block.Index) + block.Timestamp + string(block.ProductID) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (bc *Blockchain) Validate() bool {
	for i := 1; i < len(bc.blocks); i++ {
		prevBlock := bc.blocks[i-1]
		currentBlock := bc.blocks[i]
		
		if currentBlock.Hash != calculateHash(currentBlock) {
			return false
		}
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}

func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}
```


## 🔗 Backend/main.go code
```go
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
```


