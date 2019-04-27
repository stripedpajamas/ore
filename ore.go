package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// ORE represents an implementation of the order preserving encryption protocol
type ORE struct {
	sk  []byte
	aes cipher.BlockMode
}

// NewORE returns a pointer to an initialized ORE
func NewORE(sk []byte) (*ORE, error) {
	block, err := aes.NewCipher(sk)
	if err != nil {
		return nil, err
	}
	ecb := newECBEncrypter(block)
	return &ORE{sk: sk, aes: ecb}, nil
}

// GenerateSecretKey generates a random 32 byte key
func GenerateSecretKey() ([]byte, error) {
	key := make([]byte, 32) // 32 byte key selects AES-256
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

// encrypt src using secret key and put result in dst
func encrypt(sk, dst, src []byte) {

}
