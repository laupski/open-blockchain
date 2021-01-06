package blockchain

import "encoding/json"

// Transaction type to be used within a Block inside the BlockChain
type Transaction struct {
	FromAddress string
	ToAddress   string
	Amount      float64
}

// TransactionList used to process Transactions in a BlockChain
type TransactionList []Transaction

// NewTransaction instantiates a new transaction to be mined in the BlockChain
func NewTransaction(f string, t string, a float64) Transaction {
	return Transaction{
		FromAddress: f,
		ToAddress:   t,
		Amount:      a,
	}
}

// String outputs a Transaction in JSON string format.
func (t Transaction) String() string {
	s, _ := json.Marshal(t)
	return string(s)
}

func (tx TransactionList) String() string {
	s, _ := json.Marshal(tx)
	return string(s)
}
