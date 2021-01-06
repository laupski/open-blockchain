package blockchain

import (
	"encoding/json"
)

// BlockChain object to be instantiated and to have Blocks appended to its Genesis Block
type BlockChain struct {
	Chain               []Block
	Difficulty          int
	PendingTransactions TransactionList
	MiningReward        float64
}

// NewBlockChain instantiates a new BlockChain to be used. Can append existing Blocks to it.
func NewBlockChain(d int, mr float64) BlockChain {
	bc := make([]Block, 0)
	txl := TransactionList{}
	genesis := NewBlock(txl)
	bc = append(bc, genesis)

	return BlockChain{
		Chain:               bc,
		Difficulty:          d,
		PendingTransactions: txl,
		MiningReward:        mr,
	}
}

// GetLatestBlock returns the final Block in the BlockChain
func (bc BlockChain) GetLatestBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

// MineTransactions mines current pending Transactions on the BlockChain
func (bc *BlockChain) MineTransactions(address string) {
	b := NewBlock(bc.PendingTransactions)
	b.PreviousHash = bc.GetLatestBlock().Hash
	b.MineBlock(bc.Difficulty)
	bc.Chain = append(bc.Chain, b)

	// Reward the address who mined the transaction
	bc.PendingTransactions = TransactionList{NewTransaction("", address, bc.MiningReward)}
}

// PushTransactions to the BlockChain, use MineTransactions to process them.
func (bc *BlockChain) PushTransactions(tx ...Transaction) {
	bc.PendingTransactions = append(bc.PendingTransactions, tx...)
}

// GetBalance retrieves the address' current balance from Genesis Block -> last Block
func (bc *BlockChain) GetBalance(address string) float64 {
	balance := 0.00

	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			if t.FromAddress == address && t.ToAddress != address {
				balance -= t.Amount
			} else if t.ToAddress == address && t.FromAddress != address {
				balance += t.Amount
			}
		}
	}

	return balance
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
	if len(bc.Chain) == 0 {
		return true
	}

	for i := len(bc.Chain) - 2; i >= 0; i-- {
		currentBlock := bc.Chain[i]
		lastBlock := bc.Chain[i+1]
		if currentBlock.Hash != currentBlock.ReCalculateHash() {
			return false
		}
		if currentBlock.Hash != lastBlock.PreviousHash {
			return false
		}
	}

	return true
}
