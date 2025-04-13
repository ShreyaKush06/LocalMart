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