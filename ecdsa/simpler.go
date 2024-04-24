package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "fmt"
    "math/big"
)

func main() {
    // Geração da chave privada
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        fmt.Println("Erro ao gerar a chave privada:", err)
        return
    }

    // Geração da chave pública
    publicKey := &privateKey.PublicKey

    // Mensagem a ser assinada
    message := []byte("Exemplo de mensagem a ser assinada")

    // Assinatura da mensagem
    r, s, err := ecdsa.Sign(rand.Reader, privateKey, message)
    if err != nil {
        fmt.Println("Erro ao assinar a mensagem:", err)
        return
    }

    // Verificação da assinatura
    valid := ecdsa.Verify(publicKey, message, r, s)
    fmt.Println("Assinatura válida?", valid)
}
