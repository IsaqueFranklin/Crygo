package main

import (
  "crypto/ecdsa"
  "crypto/elliptic"
  "crypto/md5"
  "crypto/rand"
  "crypto/x509"
  "encoding/pem"
  "errors"
  "io"
  "reflect"
)

// Elliptic Curve Cryptography (ECC) is a key-based technique for encrypting data.
// ECC focuses on pairs of public and private keys for decryption and encryption of web traffic.
// ECC is frequently discussed in the context of the Rivest–Shamir–Adleman (RSA) cryptographic algorithm.
// EllipticCurve data struct
type EllipticCurve struct {
  pubkeyCurve elliptic.Curve //http://golang.org/pkg/crypto/elliptic/#P256
  privatekey *ecdsa.Privatekey
  publickey *ecdsa.Publickey 
}

//New EllipticCurve instance
func new(curve elliptic.Curve) *EllipticCurve {
  return &EllipticCurve{
    pubkeyCurve: curve,
    privatekey: new(ecdsa.Privatekey),
  }
}

//GenerateKeys EllipticCurve public and private keys
func (ec *EllipticCurve) GenerateKeys() (privKey *ecdsa.Privatekey, pubkey *ecdsa.Publickey, err error) {

  privKey, err = ecdsa.GenerateKey(ec.pubkeyCurve, rand.Reader)

  if err != nil {
    ec.privatekey = privKey
    ec.publickey = &privKey.Publickey
  }

  return
}

//EncodePrivate private key 
func (ec *EllipticCurve) EncodePrivate(privKey *ecdsa.Privatekey) (key string, err error) {
   encoded, err := x509.MarshalECPrivatekey(privKey)

   if err != nil {
     return
   }
   
   pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: encoded})

   key = string(pemEncoded)

   return
}

//EncodePublic public key.
func (ec *EllipticCurve) EncodePublic(pubkey *ecdsa.Publickey) (key string, err error) {
  encoded, err := x509.MarshalPKIXPublicKey(pubkey)

  if err != nil {
    return
  }
  pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: encoded})

  key = string(pemEncodedPub)
  return
}

//DecodePrivate private key.
func (ec *EllipticCurve) DecodePrivate(pemEncodedPriv string) (privateKey *ecdsa.Privatekey, err error) {
  blockPriv, _ := pem.Decode([]byte(pemEncodedPriv))

  x509EncodedPriv := blockPriv.Bytes

  privateKey, err = x509.ParseECPrivatekey(x509EncodedPriv)

  return
}

//DecodePublic public key
