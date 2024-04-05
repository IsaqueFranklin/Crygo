package main

import (
  "fmt"
  "crypto/sha256"
  "io"
  "log"
  "encoding/hex"
)

func hashWithSha256(plaintext string) (string, error) {
  h := sha256.New()
  if _, err := io.WriteString(h, plaintext); err != nil {
    return "", err
  }

  r := h.Sum(nil)
  return hex.EncodeToString(r), nil
}

func main() {
  var userInput string
  fmt.Println("Qual o conteúdo do seu hash? \n")
  fmt.Scanln(&userInput)
  fmt.Println("\n")

  hash, err := hashWithSha256(userInput)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(hash)
}
