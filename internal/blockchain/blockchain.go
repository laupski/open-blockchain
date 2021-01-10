package blockchain

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
)

// BlockChain object to be instantiated and to have Blocks appended to its Genesis Block
type BlockChain struct {
	Chain               []Block
	Difficulty          int32
	PendingTransactions TransactionList
	MiningReward        float32
}

// NewBlockChain instantiates a new BlockChain to be used. Can append existing Blocks to it.
func NewBlockChain(d int32, mr float32) *BlockChain {
	bc := make([]Block, 0)
	txl := TransactionList{}
	genesis := NewBlock(txl)
	bc = append(bc, genesis)

	return &BlockChain{
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
func (bc *BlockChain) MineTransactions(address []byte) {
	b := NewBlock(bc.PendingTransactions)
	b.PreviousHash = bc.GetLatestBlock().Hash
	b.MineBlock(bc.Difficulty)
	bc.Chain = append(bc.Chain, b)

	// Reward the address who mined the transaction
	bc.PendingTransactions = TransactionList{NewTransaction(nil, address, bc.MiningReward)}
}

// PushTransactions to the BlockChain, use MineTransactions to process them.
func (bc *BlockChain) PushTransactions(tx ...Transaction) error {
	for _, t := range tx {
		if t.FromAddress == nil || t.ToAddress == nil {
			return errors.New("must include to and from address")
		}
		if !t.VerifyTransaction() {
			return errors.New("unverified transaction")
		}
	}
	bc.PendingTransactions = append(bc.PendingTransactions, tx...)
	return nil
}

// GetBalance retrieves the address' current balance from Genesis Block -> last Block
func (bc *BlockChain) GetBalance(address []byte) float32 {
	balance := float32(0.00)

	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			if bytes.Equal(t.FromAddress, address) && !bytes.Equal(t.ToAddress, address) {
				balance -= t.Amount
			} else if bytes.Equal(t.ToAddress, address) && !bytes.Equal(t.FromAddress, address) {
				balance += t.Amount
			}
		}
	}

	return balance
}

// GetAllBalances in the BlockChain from Genesis.
func (bc *BlockChain) GetAllBalances() string {
	balances := make(map[string]float32)

	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			if t.FromAddress == nil {
				balances["Mining rewards awarded"] -= t.Amount
			} else {
				balances[base64.StdEncoding.EncodeToString(t.FromAddress)] -= t.Amount
			}
			balances[base64.StdEncoding.EncodeToString(t.ToAddress)] += t.Amount
		}
	}

	j, _ := json.MarshalIndent(balances, "", "    ")
	return string(j)
}

// String outputs the BlockChain and its underlying Blocks in JSON string format.
func (bc BlockChain) String() string {
	s, err := json.MarshalIndent(bc, "", "    ")
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
		if !bytes.Equal(currentBlock.Hash, currentBlock.ReCalculateHash()) {
			return false
		}
		if !bytes.Equal(currentBlock.Hash, lastBlock.PreviousHash) {
			return false
		}
		if !currentBlock.CheckTransactions() {
			return false
		}
	}

	return true
}
