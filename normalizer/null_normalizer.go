package normalizer

type NullNormalizer struct{}

func (NullNormalizer) Transform(source string) string {
	return source
}
