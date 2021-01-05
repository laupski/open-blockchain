package blockchain

import (
	"encoding/json"
)

type BlockChain struct {
	Name string
	Chain []Block
	Length int
}

func NewBlockChain(n string) BlockChain {
	bc := make([]Block, 0)
	bc = append(bc, createGenesisBlock())

	return BlockChain {
		Name: n,
		Chain: bc,
		Length: 1,
	}
}

func (bc BlockChain) GetLatestBlock() Block {
	return bc.Chain[bc.Length - 1]
}

func (bc *BlockChain) AppendBlock(b Block) {
	bc.Chain = append(bc.Chain, b)
	bc.Length++
}

func (bc BlockChain) String() string {
	s, err := json.Marshal(bc)
	_ = err
	return string(s)
}

func (bc BlockChain) Verify() bool {
	if bc.Length == 0 {
		return true
	}

	for i := len(bc.Chain) - 2; i >= 0; i-- {
		currentBlock := bc.Chain[i]
		lastBlock := bc.Chain[i + 1]
		if currentBlock.Hash != lastBlock.PreviousHash {
			return false
		}
	}

	return true
}