package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hash"
	"time"
)

type Block struct {
	timestamp time.Time
	data string
	Hash string
	PreviousHash string
}

func NewBlock(d string, ph string) Block {
	h := sha256.New()
	h.Write([]byte(d + time.Now().String()))
	return Block {
		timestamp: time.Now(),
		data: d,
		Hash: fmt.Sprintf("%x",h.Sum(nil)),
		PreviousHash: ph,
	}
}

func (b Block) calculateHash() hash.Hash {
	h := sha256.New()
	h.Write([]byte(b.data + b.timestamp.String()))
	return h
}

func (b Block) String() string {
	s, _ := json.Marshal(b)
	return string(s)
}

func createGenesisBlock() Block {
	h := sha256.New()
	h.Write([]byte("genesis" + time.Now().String()))
	return Block {
		timestamp: time.Now(),
		data: "genesis",
		Hash: fmt.Sprintf("%x",h.Sum(nil)),
		PreviousHash: "",
	}
}
