package main

import (
  "crypto/sha256"
  "fmt"
)

func main() {
  h := sha256.New()
  h.Write([]byte("Cracóvia."))

  //Another sintax could be:
  // sum := sha256.Sum256([]byte("this is a password."))
  // fmt.Printf("%x", sum)

  //Calculate and print the hash
  fmt.Printf("This the first example: %x", h.Sum(nil))
  fmt.Println("\n")

  //This here is the example of how two diferent strings by only one character can change the whole hash.

  sum := sha256.Sum256([]byte("this is password."))
  sumCap := sha256.Sum256([]byte("This is password."))
  fmt.Printf("Lowercase hash: %x", sum)
  fmt.Printf("\n\n")
  fmt.Printf("Capital hash: %x", sumCap)
}
