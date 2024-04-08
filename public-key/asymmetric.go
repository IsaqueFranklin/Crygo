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

  publicKey := &privateKey.PublicKey

  privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
  privateKeyPEM := pem.EncodeToMemory(&pem.Block{
    Type: "RSA PRIVATE KEY",
    Bytes: privateKeyBytes,
  })

  err = os.WriterFile("private.pem", privateKeyPEM, 0644)
  if err != nil {
    panic(err)
  }

  publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
  if err != nil {
    panic(err)
  }
  publicKeyPEM := pem.EncodeToMemory(&pem.Block{
    Type: "RSA PUBLIC KEY",
    Bytes: publicKeyBytes,
  })

  err = os.WriterFile("public.pem", publicKeyPEM, 0644)
  if err != nil {
    panic(err)
  }
}
