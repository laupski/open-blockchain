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
	data         string
	Hash         string
	PreviousHash string
	nonce        int
}

var initNonce = 0

// NewBlock instantiates a new Block which can then be added to an existing BlockChain.
func NewBlock(d string) Block {
	return Block{
		Id:        uuid.New(),
		timestamp: time.Now(),
		data:      d,
		Hash:      CalculateHash(d, time.Now(), initNonce),
		nonce:     initNonce,
	}
}

// CalculateHash calculates Hash string value of data, a timestamp and a nonce value.
func CalculateHash(d string, t time.Time, nonce int) string {
	h := sha256.New()
	h.Write([]byte(d + t.String() + string(rune(nonce))))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MineBlock recalculates the Hash to establish proof of work.
func (b *Block) MineBlock(d int) {
	for b.Hash[:d] != strings.Repeat("0", d) {
		b.nonce++
		b.Hash = CalculateHash(b.data, b.timestamp, b.nonce)
	}
}

func (b Block) ReCalculateHash() string {
	return CalculateHash(b.data, b.timestamp, b.nonce)
}

// String outputs a Block in JSON string format.
func (b Block) String() string {
	s, _ := json.Marshal(b)
	return string(s)
}
