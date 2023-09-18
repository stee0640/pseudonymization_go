package normalizer

import (
	"strings"
)

type CprNormalizer struct{}

func (CprNormalizer) Transform(cpr string) string {
	return strings.ReplaceAll(cpr, "-", "")
}
