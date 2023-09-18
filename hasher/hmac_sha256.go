package hasher

import (
	"crypto/hmac"
	"crypto/sha256"
)

type HmacSHA256 struct{}

func (HmacSHA256) Hash(plaintext []byte, salt []byte) []byte {
	mac := hmac.New(sha256.New, salt)
	mac.Write(plaintext)
	result := mac.Sum(nil)
	return result
}
