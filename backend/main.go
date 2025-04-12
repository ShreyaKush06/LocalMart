// main.go
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var client *ethclient.Client
var contract *CampusShop

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Ethereum client
	client, err = ethclient.Dial(os.Getenv("ETHEREUM_NODE"))
	if err != nil {
		log.Fatal(err)
	}

	// Load contract
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	contract, err = NewCampusShop(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Set up Gin router
	r := gin.Default()

	// API routes
	r.GET("/products", getProducts)
	r.POST("/products", addProduct)
	r.POST("/orders", placeOrder)
	r.PUT("/orders/:id/fulfill", fulfillOrder)
	r.PUT("/orders/:id/pay", confirmPayment)

	// Start server
	r.Run(":8080")
}

func getProducts(c *gin.Context) {
	// Call smart contract to get products
	opts := &bind.CallOpts{Context: context.Background()}
	
	productCount, err := contract.ProductCount(opts)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var products []Product
	for i := 1; i <= int(productCount.Int64()); i++ {
		product, err := contract.Products(opts, big.NewInt(int64(i)))
		if err != nil {
			continue
		}
		products = append(products, Product{
			ID:          uint(i),
			Name:        product.Name,
			Price:       product.Price.Uint64(),
			Description: product.Description,
			IPFSHash:    product.IpfsHash,
			ShopOwner:   product.ShopOwner.Hex(),
			IsAvailable: product.IsAvailable,
		})
	}

	c.JSON(200, products)
}

func addProduct(c *gin.Context) {
	// Implementation to add product to blockchain
}

func placeOrder(c *gin.Context) {
	// Implementation to place order on blockchain
}

func fulfillOrder(c *gin.Context) {
	// Implementation to mark order as fulfilled
}

func confirmPayment(c *gin.Context) {
	// Implementation to confirm payment
}

// Helper function to get transaction auth
func getAuth() (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth, nil
}