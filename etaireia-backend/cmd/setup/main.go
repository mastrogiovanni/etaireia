package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

func createSubscription(url string) (ed25519.PublicKey, ed25519.PrivateKey) {

	document := []byte("Hello World!")

	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	publicKeyString := hex.EncodeToString([]byte(publicKey))

	var b bytes.Buffer

	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("publicKey")
	if err != nil {
		panic(err)
	}
	io.Copy(fw, strings.NewReader(publicKeyString))

	fw, err = w.CreateFormField("name")
	if err != nil {
		panic(err)
	}
	io.Copy(fw, strings.NewReader("Michele"))

	fw, err = w.CreateFormField("surname")
	if err != nil {
		panic(err)
	}
	io.Copy(fw, strings.NewReader("Mastrogiovanni"))

	fw, err = w.CreateFormFile("document", "test.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(fw, strings.NewReader(string(document)))
	w.Close()

	// log.Printf("%s\n", string(b.String()))

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	log.Printf("Response: %s\n", bodyString)

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		panic(err)
	}

	return publicKey, privateKey

}

func approveSubscription(url string, publicKey ed25519.PublicKey) {

	body := map[string]string{
		"approverPublicKey":     hex.EncodeToString(publicKey),
		"subscriptionPublicKey": hex.EncodeToString(publicKey),
	}

	bodyString, _ := json.Marshal(body)

	bodyReader := bytes.NewReader(bodyString)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	log.Println(res)

}

func main() {

	publicKey, _ := createSubscription("https://digital-signature.mastrogiovanni.cloud/api/v1/subscription")

	approveSubscription("https://digital-signature.mastrogiovanni.cloud/api/v1/subscription/approve", publicKey)

	// signature := ed25519.Sign(privateKey, document)

	// signatureString := hex.EncodeToString([]byte(signature))

}
