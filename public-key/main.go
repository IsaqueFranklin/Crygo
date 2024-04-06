 //A diferença entre public-key para symmetric-key é que na public-key são usadas duas chaves diferentes, uma para encriptar e outra para desencriptar.

//Existem a chave pública, usada para encriptar, e a chave privada, usada para desencriptar.

//RSA é um exemplo de ecosistema de public-key e pode ser implementada usando o subpacote rsa.

 package main

 import (
  "fmt"
  "log"
  "crypto/sha256"
  "crypto/rsa"
  "encoding/hex"
  "crypto/rand"
  "crypto"
 )

 func main() {
   //Create an RSA key pair of size 2048 bits
   priv, err := rsa.GenerateKey(rand.Reader, 2048)
   if err != nil {
     log.Fatalln(err)
   }

   pub := priv.Public()

   options := rsa.OAEPOptions{
     crypto.SHA256,
     []byte("Label"),
   }

   message := "Secret message!"

   rsact, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub.(*rsa.PublicKey), []byte(message), options.Label)
 
 if err != nil {
   log.Fatalln(err)
 }

  fmt.Println("RSA ciphertext", hex.EncodeToString(rsact))
  

  rsapt, err := priv.Decrypt(rand.Reader, rsact, &options)
  if err != nil {
    log.Fatalln(err)
  }

  fmt.Println("RSA plaintext", string(rsapt))
}
