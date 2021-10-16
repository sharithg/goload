package docker

import (
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func RandomId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
