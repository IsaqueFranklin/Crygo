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

}
