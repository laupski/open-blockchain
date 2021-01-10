package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

// Block to be added to an existing BlockChain.
type Block struct {
	Id           uuid.UUID
	timestamp    time.Time
	Transactions TransactionList
	Hash         []byte
	PreviousHash []byte
	nonce        int
}

var initNonce = 0

// NewBlock instantiates a new Block which can then be added to an existing BlockChain.
func NewBlock(t TransactionList) Block {
	id := uuid.New()
	timestamp := time.Now()

	return Block{
		Id:           id,
		timestamp:    timestamp,
		Transactions: t,
		Hash:         CalculateHash(t.String() + timestamp.String() + string(rune(initNonce))),
		nonce:        initNonce,
	}
}

// CalculateHash calculates the Hash of a generic string Value
func CalculateHash(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

// MineBlock recalculates the Hash to establish proof of work.
func (b *Block) MineBlock(d int32) {
	for fmt.Sprintf("%x", b.Hash)[:d] != strings.Repeat("0", int(d)) {
		b.nonce++
		b.Hash = CalculateHash(b.Transactions.String() + b.timestamp.String() + string(rune(b.nonce)))
	}
}

// ReCalculateHash is used for under the Verify function. Used for checking stored vs actual hashes.
func (b Block) ReCalculateHash() []byte {
	return CalculateHash(b.Transactions.String() + b.timestamp.String() + string(rune(b.nonce)))
}

// String outputs a Block in JSON string format.
func (b Block) String() string {
	s, _ := json.MarshalIndent(b, "", "    ")
	return string(s)
}

// CheckTransactions in the Block. Calls VerifyTransactions underneath.
func (b Block) CheckTransactions() bool {
	for _, t := range b.Transactions {
		if !t.VerifyTransaction() {
			return false
		}
	}
	return true
}
