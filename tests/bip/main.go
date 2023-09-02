package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/islishude/bip39"
)

func main() {

	// s, err := bip39.NewMnemonic(12, bip39.Italian)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(s)

	s := "piattino fucsia ninfa incrocio alimento scenico nebulosa sentire toscano asta finire ausilio"

	mnemonic := bip39.MnemonicToSeed(s, "pluto50")

	privateKey := ed25519.NewKeyFromSeed(mnemonic[:32])

	publicKey := privateKey[32:]

	fmt.Printf("Public Key: %s\n", hex.EncodeToString(publicKey))

	fmt.Printf("Private Key: %s\n", hex.EncodeToString(privateKey))

	document := []byte("Hello World!")

	signature := ed25519.Sign(privateKey, document)

	fmt.Printf("Signature: %s\n", hex.EncodeToString(signature))

	verified := ed25519.Verify([]byte(publicKey), document, signature)

	fmt.Printf("Verified: %+v\n", verified)

}
