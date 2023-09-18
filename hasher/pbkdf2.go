package hasher

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

type Pbkdf2 struct{}

func (Pbkdf2) Hash(plaintext []byte, salt []byte) []byte {
	result := pbkdf2.Key(plaintext, salt, 100000, 32, sha256.New)
	return result
}
