//Asymmetric cryptography, or public-key cryptography.
//Here the key used to encrypt data is different from key used to decrypt data.

package main

import (
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "encoding/pem"
  "os"
)

func main() {
  privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
    panic(err)
  }
}
