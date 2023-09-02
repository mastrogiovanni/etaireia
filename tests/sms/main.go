package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	baseURL := "https://platform.clickatell.com"

	resource := "/messages/http/send"

	params := url.Values{}
	params.Add("apiKey", "P6jRLPAZSxSLf0MpvG9onA==")
	params.Add("to", "+393493147800")
	params.Add("content", "Ciao Michele come stai?")

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status code: %d", resp.StatusCode)
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
