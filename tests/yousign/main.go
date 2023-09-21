package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func initiateRequest() (string, error) {

	url := "https://api-sandbox.yousign.app/v3/signature_requests"

	payload := strings.NewReader(`
	{
		"delivery_mode": "email",
		"reminder_settings": {
		  "interval_in_days": 1,
		  "max_occurrences": 5
		},
		"signers_allowed_to_decline": false,
		"name": "test signature",
		"ordered_signers": true,
		"email_custom_note": "Please sign the email..."
	}
	`)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer ")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var content map[string]interface{}
	err = json.Unmarshal(body, &content)
	if err != nil {
		return "", err
	}

	id, ok := content["id"].(string)
	if !ok {
		return "", err
	}

	return id, nil

}

func uploadDocument(fileName string, idRequest string) (string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	r, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var fw io.Writer

	if fw, err = w.CreateFormFile("file", fileName); err != nil {
		return "", err
	}
	if _, err = io.Copy(fw, r); err != nil {
		return "", err
	}

	if fw, err = w.CreateFormField("nature"); err != nil {
		return "", err
	}
	if _, err = io.Copy(fw, strings.NewReader("signable_document")); err != nil {
		return "", err
	}

	if fw, err = w.CreateFormField("password"); err != nil {
		return "", err
	}
	if _, err = io.Copy(fw, strings.NewReader("pippo80")); err != nil {
		return "", err
	}

	if fw, err = w.CreateFormField("parse_anchors"); err != nil {
		return "", err
	}
	if _, err = io.Copy(fw, strings.NewReader("false")); err != nil {
		return "", err
	}

	// }
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	url := fmt.Sprintf("https://api-sandbox.yousign.app/v3/signature_requests/%s/documents", idRequest)

	req, _ := http.NewRequest("POST", url, &b)

	req.Header.Add("accept", "application/json")
	req.Header.Set("content-type", w.FormDataContentType())
	req.Header.Add("authorization", "Bearer ")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var content map[string]interface{}
	err = json.Unmarshal(body, &content)
	if err != nil {
		panic(err)
	}

	id, ok := content["id"].(string)
	if !ok {
		return "", errors.New("content id is not a string")
	}

	return id, nil
}

func addSigner(idRequest string, idDocument string) (string, error) {

	url := fmt.Sprintf("https://api-sandbox.yousign.app/v3/signature_requests/%s/signers", idRequest)

	payload := strings.NewReader(fmt.Sprintf(`
	{
		"info": {
		  "locale": "it",
		  "first_name": "Michele",
		  "last_name": "Mastrogiovanni",
		  "email": "michele.mastrogiovanni@gmail.com",
		  "phone_number": "+393493147800"
		},
		"signature_level": "advanced_electronic_signature",
		"fields": [
		  {
			"type": "signature",
			"document_id": "%s",
			"page": 1,
			"x": 0,
			"y": 0
		  }
		],
		"signature_authentication_mode": "otp_sms"
	  }
	`, idDocument))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer ")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// fmt.Printf("%+v", string(body))

	var content map[string]interface{}
	err = json.Unmarshal(body, &content)
	if err != nil {
		return "", err
	}

	id, ok := content["id"].(string)
	if !ok {
		return "", errors.New("content id is not a string")
	}

	return id, nil

}

func activate(idRequest string) {

	url := fmt.Sprintf("https://api-sandbox.yousign.app/v3/signature_requests/%s/activate", idRequest)

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer ")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}

func main() {
	idRequest, err := initiateRequest()
	if err != nil {
		panic(err)
	}
	fmt.Println("Request ID: " + idRequest)
	idDocument, err := uploadDocument("/home/michele/test.pdf", idRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Document ID: " + idDocument)

	// idRequest := "3cd06e45-ded8-47e5-8229-b99b7ccb1b22"
	// idDocument := "36f194cf-bd89-4ad9-9d72-db8ca5f36536"

	idSigner, err := addSigner(idRequest, idDocument)
	if err != nil {
		panic(err)
	}
	fmt.Println("Signer ID: " + idSigner)

	// idSigner := "21b1c688-326f-49b0-abc6-0f2e708f48a0"

	activate(idRequest)

}
