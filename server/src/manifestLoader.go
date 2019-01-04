package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadManifest() []byte {
	url := "https://bungie.net/common/destiny2_content/json/ru/aggregate-4ba29e22-81f5-435e-b09c-2e0ea0f3dd41.json"

	request, error := http.NewRequest("GET", url, nil)
	if error != nil {
		panic(error)
	}

	request.Header.Add("X-API-KEY", "5f9e82c6d37b46a29e370268fffd93f1")
	request.Header.Add("cache-control", "no-cache")
	request.Header.Add("Postman-Token", "de809f00-07a5-444a-852c-465f993b18ff")

	result, error := http.DefaultClient.Do(request)
	if error != nil {
		panic(error)
	}
	defer result.Body.Close()

	fmt.Println("Response status code:", result.StatusCode)

	body, error := ioutil.ReadAll(result.Body)
	if error != nil {
		panic(error)
	}

	return body
}
