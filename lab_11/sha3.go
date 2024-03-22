package main

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {

	input := "sha3 hello world"

	output := sha3Hash(input)

	fmt.Println(input, "=", output)
}

func sha3Hash(input string) string {

	// Create a new hash & write input string
	hash := sha3.New256()
	hash.Write([]byte(input))

	// Get the resulting encoded byte slice
	sha3 := hash.Sum(nil)

	// Convert the encoded byte slice to a string
	return fmt.Sprintf("%x", sha3)
}
