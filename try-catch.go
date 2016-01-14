package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main_try() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	//callSomeJson()

	someVal, err := callSomeJson2()

	if err != nil {
		// do something in my app
	}

	fmt.Printf("Here's what we got: %s\n", someVal)
}

func callSomeJson() (json string) {

	panic("Oh No it didn't work")

	return "{never: 'gonna get it'}"
}

func callSomeJson2() (jsonStr string, err error) {

	jsonVal, err := json.Marshal("gopher")

	if err != nil {
		log.Fatal(err.Error())

		return "", err
	}

	return string(jsonVal), nil
}
