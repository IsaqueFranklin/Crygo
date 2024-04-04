package main

import (
  "crypto/sha256"
  "fmt"
  "io"
  "os"
)

func main() {
  file, err := os.Open("test") //Open the file for reading
  if err != nil {
    panic(err)
  }
  defer file.Close() //Be sure to close your file.

  hash := sha256.New() //Use the Hash in crypto/sha256.

  if _, err := io.Copy(hash, file); err != nil {
    panic(err)
  }

  sum := fmt.Sprintf("%x", hash.Sum(nil)) //Get encoded hash sum
  fmt.Println(sum)
}
