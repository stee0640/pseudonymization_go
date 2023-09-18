package encrypted_salt

import (
	"encoding/hex"

	"github.com/stee0640/pseudonymization_go/encryption"
)

const NONCE_BYTES uint = 12
const SECRET_SALT_BYTES uint = 16

type encryptedSalt struct {
	nonce      []byte
	ciphertext []byte
}

func New() *encryptedSalt {
	e := encryptedSalt{
		nonce:      nil,
		ciphertext: nil,
	}
	return &e
}

func (e *encryptedSalt) Encrypt(encryption_key []byte, secret_salt []byte) {
	nonce := encryption.GetRandomBytes(NONCE_BYTES)
	e.ciphertext = encryption.New(encryption_key).Encrypt(secret_salt, nonce)
}

func (e *encryptedSalt) Decrypt(encryption_key []byte) (secret_salt []byte) {
	return encryption.New(encryption_key).Decrypt(e.ciphertext, e.nonce)
}

func Generate(encryption_key []byte) *encryptedSalt {
	e := New()
	secret_salt := encryption.GetRandomBytes(SECRET_SALT_BYTES)
	e.Encrypt(encryption_key, secret_salt)
	return e
}

func (e *encryptedSalt) Load(serialized_encrypted_salt string) *encryptedSalt {
	encrypted_salt, _ := hex.DecodeString(serialized_encrypted_salt)
	e.nonce = encrypted_salt[:NONCE_BYTES]
	e.ciphertext = encrypted_salt[NONCE_BYTES:]
	return e
}

func (e *encryptedSalt) Dump() string {
	return hex.EncodeToString(append(e.nonce, e.ciphertext...))
}
