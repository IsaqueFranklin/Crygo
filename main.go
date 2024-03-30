 package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("this is a password"))
	// Calculate and print the hash
	fmt.Printf("%x", h.Sum(nil))
}
