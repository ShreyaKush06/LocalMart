package main

import (
	"backend/blockchain"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var bc *blockchain.Blockchain
var db *sql.DB
var jwtKey = []byte("your_secret_key") // Replace with a secure key

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Shop      string `json:"shop"`
	OnBlinkit bool   `json:"onBlinkit"`
	Location  string `json:"location,omitempty"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var products = []Product{
	{ID: 1, Name: "Campus T-Shirt", Price: "₹0", Shop: "Campus Store", OnBlinkit: false},
	{ID: 2, Name: "Special Chai Mix", Price: "₹150", Shop: "Canteen", OnBlinkit: false},
}

var requestedItems = []Product{}

func initializeDatabase() error {
	// Connect to MySQL server (without database)
	rootDB, err := sql.Open("mysql", "root:Ggoyat@15@tcp(127.0.0.1:3306)/")
	if err != nil {
		return err
	}
	defer rootDB.Close()

	// Create database if not exists
	_, err = rootDB.Exec("CREATE DATABASE IF NOT EXISTS `LocalMart`")
	if err != nil {
		return err
	}

	// Create tables
	queries := []string{
		`CREATE TABLE IF NOT EXISTS roles (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(50) NOT NULL UNIQUE
        )`,
		`CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            username VARCHAR(50) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL,
            role_id INT NOT NULL,
            FOREIGN KEY (role_id) REFERENCES roles(id)
        )`,
		`INSERT IGNORE INTO roles (id, name) VALUES (1, 'admin'), (2, 'user')`,
		`INSERT IGNORE INTO users (username, password, role_id) 
         VALUES ('admin1', 'admin_password', 1), ('user1', 'user_password', 2)
         ON DUPLICATE KEY UPDATE username=username`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Get DB credentials from env or use defaults
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASS", "Ggoyat@15")
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "LocalMart")

	// Connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// Connect
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize database
	if err := initializeDatabase(); err != nil {
		log.Fatal("Database initialization failed: ", err)
	}

	bc = blockchain.NewBlockchain()

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/signup", signup).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(authMiddleware)
	api.HandleFunc("/requests", getRequestedItems).Methods("GET")
	api.HandleFunc("/list-item", listItem).Methods("POST")

	// Product endpoints
	api.HandleFunc("/products", getProducts).Methods("GET")
	api.HandleFunc("/products", addProduct).Methods("POST")

	// Blockchain endpoints
	api.HandleFunc("/blockchain", getBlockchain).Methods("GET")

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Helper function to get env vars with defaults
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
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
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var hashedPassword, role string
	err := db.QueryRow("SELECT password, name FROM users INNER JOIN roles ON users.role_id = roles.id WHERE username = ?", creds.Username).Scan(&hashedPassword, &role)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare passwords (use bcrypt in production)
	if creds.Password != hashedPassword {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func signup(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if username already exists
	var exists bool
	err := db.QueryRow("SELECT 1 FROM users WHERE username = ?", creds.Username).Scan(&exists)
	if err == nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// In production, hash the password with bcrypt here

	// Insert new user (regular user role = 2)
	_, err = db.Exec("INSERT INTO users (username, password, role_id) VALUES (?, ?, 2)",
		creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add role-based access control
		if strings.HasPrefix(r.URL.Path, "/api/requests") && claims.Role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
