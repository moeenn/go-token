package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func GenerateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func main() {
	lengthPtr := flag.Int("length", 64, "Length of the token string")
	flag.Parse()

	if *lengthPtr == 0 {
		fmt.Fprintf(os.Stderr, "Length of the token string cannot be zero")
		os.Exit(1)
	}

	token, err := GenerateSecureToken(*lengthPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s\n", token)
}
