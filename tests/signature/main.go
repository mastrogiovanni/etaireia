package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jws"
)

func main() {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Printf("failed to generate private key: %s", err)
		return
	}

	buf, err := jws.Sign([]byte("Lorem ipsum"), jws.WithKey(jwa.RS256, privkey))
	if err != nil {
		log.Printf("failed to created JWS message: %s", err)
		return
	}

	fmt.Println(string(buf))

	// When you receive a JWS message, you can verify the signature
	// and grab the payload sent in the message in one go:
	verified, err := jws.Verify(buf, jws.WithKey(jwa.RS256, &privkey.PublicKey))
	if err != nil {
		log.Printf("failed to verify message: %s", err)
		return
	}

	log.Printf("signed message verified! -> %s", verified)
}
