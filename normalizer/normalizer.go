package normalizer

type Normalizer interface {
	Transform(string) string
}
