package pseudonymizer

import (
	"github.com/stee040/pseudonymization_go/encrypted_salt"
	"github.com/stee040/pseudonymization_go/hasher"
	"github.com/stee040/pseudonymization_go/normalizer"
)

type pseudonymizer struct {
	storage_key []byte
	salt        []byte
	hasher      hasher.Hasher
	normalizer  normalizer.Normalizer
}

func DefaultPseudonymizer(storage_key []byte, serialized_encrypted_salt string) pseudonymizer {
	p := NewPseudonymizer(storage_key, serialized_encrypted_salt, nil, nil)
	return p
}

func NewPseudonymizer(storage_key []byte, serialized_encrypted_salt string, h hasher.Hasher, n normalizer.Normalizer) pseudonymizer {
	if h == nil {
		h = hasher.Scrypt{}
	}
	if n == nil {
		n = normalizer.CprNormalizer{}
	}
	salt := encrypted_salt.New().Load(serialized_encrypted_salt).Decrypt(storage_key)

	p := pseudonymizer{
		storage_key: storage_key,
		salt:        salt,
		hasher:      h,
		normalizer:  n,
	}
	return p
}

func (p pseudonymizer) Pseudonymize(plaintext string) []byte {
	pseudonym := p.hasher.Hash([]byte(p.normalizer.Transform(plaintext)), p.salt)
	return pseudonym
}
