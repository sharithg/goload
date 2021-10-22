package utils

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/fatih/color"
)

func FatalError(message string) {
	red := color.New(color.FgRed).SprintFunc()

	errorMessage := fmt.Sprintf("%s: %s", red("Error"), message)

	fmt.Fprintln(os.Stderr, errorMessage)
	os.Exit(1)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

// utility function to create a random id given a length
func RandomId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
