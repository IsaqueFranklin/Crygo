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

//Here on those are function for padding and unpadding the plaintext.
//Padding is simply the act of incrasing te length of the plaintext so that it can be a multiple of a fixed site (usually a block size). This is done by adding characters to the plaintext.

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

//Writing the functions for encryption and the decryption.

func encrypt(key, plaintext string) (string, error) {
  block, err := newCipherBlock(key)
  if err != nil {
    return "", err
  }

  //pad plaintext
  ptbs, _ := pkcs7Pad([]byte(plaintext), block.BlockSize())

  if len(ptbs)%aes.BlockSize != 0 {
    return "", errors.New("Plaintext is not a multiple of the block size.")
  }

  ciphertext := make([]byte, len(ptbs))
  
  //Create an initialization vector which is the length of the block size for AES.
  var iv []byte = make([]byte, aes.BlockSize)
  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    return "", err
  }

  //Ecrypt plaintext
  mode.CryptBlocks(cipher, ptbs)

  //Concatenate initialization vector and ciphertext.
  return hex.EncodeToString(iv) + ":" + hex.EncodeToString](ciphertext), nil
}
