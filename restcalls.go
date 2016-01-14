package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsonPayload struct {
	Id    int
	Title string
}

var (
	client *http.Client
)

func main_rest() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{Transport: tr}

	stringCall()
	jsonCall()
	anonymousCall()
}

func stringCall() {

	resp, _ := client.Get("http://jsonplaceholder.typicode.com/posts/1")
	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(contents))
}

func jsonCall() {

	resp, _ := client.Get("http://jsonplaceholder.typicode.com/posts/1")
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var jsonData JsonPayload
	decoder.Decode(&jsonData)

	fmt.Printf("We got this: %#v\n", jsonData)
}

func anonymousCall() {

	resp, _ := client.Get("http://jsonplaceholder.typicode.com/posts/1")
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var jsonData map[string]interface{}
	decoder.Decode(&jsonData)

	fmt.Printf("We got this: %#v\n", jsonData["body"])
}
