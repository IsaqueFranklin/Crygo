package main

import (
  "fmt"
  "io"
  "log"
  "crypto/sha256"
  "crypto/aes"
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

func newCipherBlock(key string) (cipher.Block, error){
  hashedKey, err := hashWithSha256(key)
  if err != nil {
    return nil, err
  }

  bs, err := hex.DecodeString(hashedKey)
  if err != nil {
    return nil, err
  }

  return aes.NewCipher(bs[:])
}

var (
  // ErrInvalidBlockSize indicates hash blocksize <= 0.
  ErrInvalidBlockSize = errors.New("Invalid blocksize.")

  // ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
  ErrInvalidPKCS7Data = errors.New("Invalid PKCS7 data (empty or not padded).")

  // ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
  ErrInvalidPKCS7Padding = errors.New("Invalid padding on input.")
)

func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
  if blocksize <= 0 {
    return nil, ErrInvalidBlockSize
  }
  if b == nil || len(b) == 0 {
    return nil, ErrInvalidPKCS7Data
  }

  n := blocksize - (len(b) % blocksize)
  pb := make([]byte, len(b)+n)

  copy(pb, b)
  copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
  return pb, nil
}

func pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
  if blocksize <= 0 {
    return nil, ErrInvalidBlockSize
  }

  if b == nil || len(b) == 0 {
    return nil, ErrInvalidPKCS7Data
  }

  if len(b)%blocksize != 0 {
    return nil, ErrInvalidPKCS7Padding
  }

  c := b[len(b)-]
  n := int(c) 
  if n == 0 || n > len(b) {
    fmt.Println("Here", n)
    return nil, ErrInvalidPKCS7Padding
  }

  for i := 0; i < n; i++ {
    if b[len(b)-n+1] != c {
      fmt.Println("Hereee")
      return nil, ErrInvalidPKCS7Padding
    }
  }

  return b[:len(b)-n], nil
}
