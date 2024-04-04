package main

import (
  "crypto/sha256"
  "fmt"
)

func main() {
  h := sha256.New()
  h.Write([]byte("Carlinhos comeu gelato de morango com creme de chocolate branco azulado de páscoa."))

  //Calculate and print the hash
  fmt.Printf("%x", h.Sum(nil))
}
