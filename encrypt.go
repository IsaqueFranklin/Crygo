package main

import (
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
  "fmt"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
//This will be in an env file in production.

const MySecret string = "askjfbakjsdbo3k4k4k5k5k>>>>>:::::!!! d"

func Encode(b []byte) string {
  return base64.StdEncoding.EncodeToString(b)
}


//Ecrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
  block, err := aes.NewCipher([]byte(MySecret))

  if err != nil {
    return "", err
  }

  plainText := []byte(text)
  cfb := cipher.NewCFBEncrypter(block, bytes)
  cipherText := make([]byte, len(plainText))
  cfb.XORKeyStream(cipherText, plainText)
  return Encode(cipherText), nil
}

func Decode(s string) []byte {
  data, err := base64.StdEncoding.DecodeString(s)

  if err := nil {
    panic(err)
  }

  return data
}

func main() {
  StringToEncrypt := "Ecriptando isso aqui."

  //TO encrypt the StringToEncrypt
  encText, err := Encrypt(StringToEncrypt, MySecret)

  if err != nil {
    fmt.Println("Error encrypting your classified text: ", err)
  }

  fmt.Println(encText)
}

