package hasher

type Hasher interface {
	Hash(plaintext []byte, salt []byte) []byte
}
