package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://localhost:8443/"
	method := "GET"

	tr := &http.Transport{TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	}}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
