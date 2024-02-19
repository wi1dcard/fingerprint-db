package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	url := "https://localhost:8443/json"
	method := "GET"

	tr := &http2.Transport{TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	}}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("fingerprint.json", body, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
