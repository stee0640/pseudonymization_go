package hasher

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

type Scrypt struct{}

func (Scrypt) Hash(plaintext []byte, salt []byte) []byte {
	result, err := scrypt.Key(plaintext, salt, 16384, 8, 1, 32)
	if err != nil {
		panic(fmt.Sprintf("scrypt failed with %#v", err))
	}
	return result
}
