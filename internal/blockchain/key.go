package blockchain

import (
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

type Key struct {
	Name string
	Key  *ecdsa.PrivateKey
}

// NewKey will create a new ECDSA key pair. If one already exists, it will return the current pair.
// Working directory is current directory.
func NewKey(n string) (*Key, error) {

	// Only accept keys with the .key extension
	if n[len(n)-4:] != ".key" {
		return nil, errors.New("only accepting .key extension")
	}

	// Tru loading the Key first
	key, err := loadKey(n)
	if err == nil {
		return &Key{
			Name: n,
			Key:  key,
		}, nil
	}

	// If not found, generate a new Key
	key, err = crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	// Save the Key to local directory
	err = saveKey(n, key)
	if err != nil {
		return nil, err
	}

	return &Key{
		Name: n,
		Key:  key,
	}, nil
}

// Delete the Key in the working directory.
func (k *Key) Delete() error {
	return os.Remove(k.Name)
}

func LoadKey(n string) (*Key, error) {
	// Only accept keys with the .key extension
	if n[len(n)-4:] != ".key" {
		return nil, errors.New("only accepting .key extension")
	}

	key, err := loadKey(n)
	if err == nil {
		return &Key{
			Name: n,
			Key:  key,
		}, nil
	} else {
		return nil, err
	}
}

func PrintPublicAddress(k string) (string, error) {
	key, err := LoadKey(k)
	if err != nil {
		return "", err
	}

	address := crypto.FromECDSAPub(key.Key.Public().(*ecdsa.PublicKey))
	return base64.StdEncoding.EncodeToString(address), nil
}

func loadKey(n string) (*ecdsa.PrivateKey, error) {
	return crypto.LoadECDSA(n)
}

func saveKey(n string, pk *ecdsa.PrivateKey) error {
	return crypto.SaveECDSA(n, pk)
}
