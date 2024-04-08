package main

import (
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "encoding/pem"
  "fmt"
  "io/ioutil"
)

func main() {
  publicKeyPEM, err := ioutil.ReadFile("public.pem")
  if err != nil {
    panic(err)
  }

  publicKeyBlock, _ := pem.Decode(publicKeyPEM)
  publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)

  if err != nil {
    panic(err)
  }

  plaintext := []byte("hello, world!")
  ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.publicKey), plaintext)
  if err != nil {
    panic(err)
  }

  fmt.Printf("Encrypted: %x\n", ciphertext)
}
