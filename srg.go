package srg

import (
	"math/rand/v2"
	"strings"
)

const (
	AlphaNums = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateAlphaNumberic(length uint) string {
	var sb strings.Builder

	for i := 0; i < int(length); i++ {
		rand := rand.IntN(len(AlphaNums))
		b := AlphaNums[rand]
		sb.WriteByte(b)
	}

	return sb.String()
}
