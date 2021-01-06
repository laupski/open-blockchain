package blockchain

import (
	"encoding/json"
)

// BlockChain object to be instantiated and to have Blocks appended to its Genesis Block
type BlockChain struct {
	Name   string
	Chain  []Block
	Length int
}

// NewBlockChain instantiates a new BlockChain to be used. Can append existing Blocks to it.
func NewBlockChain(n string) BlockChain {
	bc := make([]Block, 0)
	bc = append(bc, createGenesisBlock())

	return BlockChain{
		Name:   n,
		Chain:  bc,
		Length: 1,
	}
}

// GetLatestBlock returns the final Block in the BlockChain
func (bc BlockChain) GetLatestBlock() Block {
	return bc.Chain[bc.Length-1]
}

// AppendBlock adds the Block at the end of the BlockChain.
func (bc *BlockChain) AppendBlock(b Block) {
	bc.Chain = append(bc.Chain, b)
	bc.Length++
}

// String outputs the BlockChain and its underlying Blocks in JSON string format.
func (bc BlockChain) String() string {
	s, err := json.Marshal(bc)
	_ = err
	return string(s)
}

// Verify is used to verify the integrity of the BlockChain. From the last Block to Genesis, the PreviousHash field
// should be linked to the proper prior Block's Hash field. Similar to a singly linked list.
func (bc BlockChain) Verify() bool {
	if bc.Length == 0 {
		return true
	}

	for i := len(bc.Chain) - 2; i >= 0; i-- {
		currentBlock := bc.Chain[i]
		lastBlock := bc.Chain[i+1]
		if currentBlock.Hash != lastBlock.PreviousHash {
			return false
		}
	}

	return true
}
