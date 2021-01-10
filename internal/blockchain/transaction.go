package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"os"
)

// Transaction type to be used within a Block inside the BlockChain
type Transaction struct {
	FromAddress []byte
	ToAddress   []byte
	Amount      float32
	Signature   []byte
}

// TransactionList used to process Transactions in a BlockChain
type TransactionList []Transaction

// NewTransaction instantiates a new transaction to be mined in the BlockChain
func NewTransaction(f, t []byte, a float32) Transaction {
	return Transaction{
		FromAddress: f,
		ToAddress:   t,
		Amount:      a,
	}
}

// SignTransaction signs a Transaction with a Key
func (t *Transaction) SignTransaction(k *Key) error {
	if !bytes.Equal(t.FromAddress, crypto.FromECDSAPub(k.Key.Public().(*ecdsa.PublicKey))) {
		return errors.New("can't sign for a transaction that's doesn't match your key")
	}

	h := sha256.New()
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, t.Amount)
	binaryAmount := b.Bytes()
	slice := append(t.FromAddress, t.ToAddress...)
	slice = append(slice, binaryAmount...)
	h.Write(slice)
	hash := h.Sum(nil)
	signature, err := crypto.Sign(hash, k.Key)
	if err != nil {
		return err
	}

	t.Signature = signature
	return nil
}

// VerifyTransaction checks the signature, calculated hash and public key to see if they match.
func (t *Transaction) VerifyTransaction() bool {
	if t.FromAddress == nil {
		return true
	}
	if t.Signature == nil {
		return false
	}

	h := sha256.New()
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, t.Amount)
	binaryAmount := b.Bytes()
	slice := append(t.FromAddress, t.ToAddress...)
	slice = append(slice, binaryAmount...)
	h.Write(slice)
	hash := h.Sum(nil)
	return crypto.VerifySignature(t.FromAddress, hash, t.Signature[:len(t.Signature)-1])
}

// String outputs a Transaction in JSON string format.
func (t Transaction) String() string {
	s, _ := json.MarshalIndent(t, "", "    ")
	return string(s)
}

// String outputs a TransactionList in JSON string format.
func (tx TransactionList) String() string {
	s, _ := json.MarshalIndent(tx, "", "    ")
	return string(s)
}

// SaveTransactionToJSON saves a Transaction to the working directory named transaction.json. Only 1 Transaction can be saved.
func (t Transaction) SaveTransactionToJSON() error {
	j, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("transaction.json", j, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadTransactionFromJSON reads a Transaction from the working directory in transaciotn.json and loads it into memory.
func ReadTransactionFromJSON() (*Transaction, error) {
	jsonFile, err := os.Open("transaction.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	j, err := ioutil.ReadAll(jsonFile)
	var t Transaction

	err = json.Unmarshal(j, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
