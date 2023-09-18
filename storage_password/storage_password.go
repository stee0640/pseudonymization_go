package storage_password

import (
	"fmt"
	"os"

	"github.com/stee0640/pseudonymization_go/hasher"

	"golang.org/x/term"
)

type storagePassword struct {
	storage_key_salt []byte
	hasher           hasher.Hasher
}

func New(storage_key_salt []byte, h hasher.Hasher) storagePassword {
	if h == nil {
		h = hasher.Scrypt{}
	}
	s := storagePassword{
		storage_key_salt: storage_key_salt,
		hasher:           h,
	}
	return s
}

func (s storagePassword) DeriveKey(password []byte) []byte {
	return s.hasher.Hash([]byte(password), s.storage_key_salt)
}

func (s storagePassword) GetpassDeriveKey() []byte {
	fmt.Print("Enter encryption password: ")
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))
	return s.DeriveKey(password)
}
