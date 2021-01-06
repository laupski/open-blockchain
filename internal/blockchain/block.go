package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hash"
	"time"
)

// Block to be added to an existing BlockChain.
type Block struct {
	timestamp    time.Time
	data         string
	Hash         string
	PreviousHash string
}

// NewBlock instantiates a new Block which can then be added to an existing BlockChain.
func NewBlock(d string, ph string) Block {
	h := sha256.New()
	h.Write([]byte(d + time.Now().String()))
	return Block{
		timestamp:    time.Now(),
		data:         d,
		Hash:         fmt.Sprintf("%x", h.Sum(nil)),
		PreviousHash: ph,
	}
}

// calculateHash re-calculates the Hash object used in a Block
func (b Block) calculateHash() hash.Hash {
	h := sha256.New()
	h.Write([]byte(b.data + b.timestamp.String()))
	return h
}

// String outputs a Block in JSON string format.
func (b Block) String() string {
	s, _ := json.Marshal(b)
	return string(s)
}

// createGenesisBlock creates the genesis block for a BlockChain
func createGenesisBlock() Block {
	h := sha256.New()
	h.Write([]byte("genesis" + time.Now().String()))
	return Block{
		timestamp:    time.Now(),
		data:         "genesis",
		Hash:         fmt.Sprintf("%x", h.Sum(nil)),
		PreviousHash: "",
	}
}
