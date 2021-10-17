package utils

import (
	"fmt"
	"math/rand"
	"os"
)

func FatalError(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func RandomId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
