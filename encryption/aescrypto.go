package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type aesGcm struct {
	cipher cipher.AEAD
}

func New(key []byte) *aesGcm {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	return &aesGcm{cipher: aead}
}

func GetRandomBytes(n uint) []byte {
	buf := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
	return buf
}

func (crypto *aesGcm) Encrypt(plaintext []byte, iv []byte) (ciphertext []byte) {
	ciphertext = crypto.cipher.Seal(nil, iv, plaintext, nil)
	return ciphertext
}

func (crypto *aesGcm) Decrypt(ciphertext []byte, iv []byte) (plaintext []byte) {
	plaintext, err := crypto.cipher.Open(nil, iv, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
