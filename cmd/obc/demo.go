package obc

import (
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "a demo run of a blockchain",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Creating Blockchain")
		bc := blockchain.NewBlockChain(4, 100)

		key1, _ := blockchain.NewKey("test1.key")
		address1 := crypto.FromECDSAPub(key1.Key.Public().(*ecdsa.PublicKey))
		fmt.Println("Address 1:" + base64.StdEncoding.EncodeToString(address1))
		key2, _ := blockchain.NewKey("test2.key")
		address2 := crypto.FromECDSAPub(key2.Key.Public().(*ecdsa.PublicKey))
		fmt.Println("Address 2:" + base64.StdEncoding.EncodeToString(address2))
		key3, _ := blockchain.NewKey("test3.key")
		address3 := crypto.FromECDSAPub(key3.Key.Public().(*ecdsa.PublicKey))
		fmt.Println("Address 3:" + base64.StdEncoding.EncodeToString(address3))
		key4, _ := blockchain.NewKey("test4.key")
		address4 := crypto.FromECDSAPub(key4.Key.Public().(*ecdsa.PublicKey))
		fmt.Println("Address 4:" + base64.StdEncoding.EncodeToString(address4))
		key5, _ := blockchain.NewKey("test5.key")
		address5 := crypto.FromECDSAPub(key5.Key.Public().(*ecdsa.PublicKey))
		fmt.Println("Address 5:" + base64.StdEncoding.EncodeToString(address5))

		t1 := blockchain.NewTransaction(address1, address2, 100)
		_ = t1.SignTransaction(key1)
		t2 := blockchain.NewTransaction(address1, address3, 50)
		_ = t2.SignTransaction(key1)
		t3 := blockchain.NewTransaction(address1, address4, 30)
		_ = t3.SignTransaction(key1)
		t4 := blockchain.NewTransaction(address1, address5, 20)
		_ = t4.SignTransaction(key1)

		txl := make(blockchain.TransactionList, 0)
		txl = append(txl, t1)
		txl = append(txl, t2)
		txl = append(txl, t3)
		txl = append(txl, t4)

		_ = bc.PushTransactions(txl...)
		fmt.Println("Before mining: ")
		fmt.Println(bc)
		bc.MineTransactions(address1)
		fmt.Println()
		fmt.Println("After mining: ")
		fmt.Println(bc)
		fmt.Println(bc.GetAllBalances())

		fmt.Println("Verifying Blockchain")
		fmt.Println(bc.Verify())

		t5 := blockchain.NewTransaction(address1, address2, 700)
		_ = t5.SignTransaction(key1)
		txl = blockchain.TransactionList{t5}
		_ = bc.PushTransactions(txl...)
		bc.MineTransactions(address1)
		fmt.Println()
		fmt.Println("After mining again: ")
		fmt.Println(bc)
		fmt.Println("Verifying Blockchain")
		fmt.Println(bc.Verify())

		fmt.Println(bc.GetAllBalances())

		_ = key1.Delete()
		_ = key2.Delete()
		_ = key3.Delete()
		_ = key4.Delete()
		_ = key5.Delete()
	},
}
