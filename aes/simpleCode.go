//AES is the advanced encryption standard
//AES is a symmetric key encryption algorithm
//Symmetric key encyption allow two parties to encrypt and decrypt information using a shared secret.

package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
)

func main() {
  fmt.Println("Encryption Programa v0.01")

  text := []byte("My Super Secret Code Stuff")
  key := []byte("passprahsewithlajnlsknlkc!")

  //Generate a new aes cipher using ou 32 byte long key
  c, err := aes.NewCipher(key)

  //if there are any errors, handle them
  if err != nil {
    fmt.Println(err)
  }
  
  // gcm or Galois/Counter Mode, is a mode of operation
    // for symmetric key cryptographic block ciphers
    // - https://en.wikipedia.org/wiki/Galois/Counter_Mode

  gcm, err := cipher NewGCM(c)

  if err != nil {
    fmt.Println(err)
  }

  //creates a new byte array the size of the nonce
  //which must be passed to seal

  nonce := make([]byte, gcm.NonceSize())
  //Populates our nonce with a cryptographcally secure random sequence
  if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
    fmt.Println(err)
  }

  // here we encrypt our text using the Seal function
    // Seal encrypts and authenticates plaintext, authenticates the
    // additional data and appends the result to dst, returning the updated
    // slice. The nonce must be NonceSize() bytes long and unique for all
    // time, for a given key.
    fmt.Println(gcm.Seal(nonce, nonce, text, nil))
}
