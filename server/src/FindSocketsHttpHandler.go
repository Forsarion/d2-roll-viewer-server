package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	manifest "manifest"
	"net/http"
	"strconv"
)

// FindSockets Description
func FindSockets(writer http.ResponseWriter, request *http.Request) {
	var file, error = ioutil.ReadFile("../misc/forsarion-inventory.json")
	if error != nil {
		panic(error)
	}
	var characterInventory CharacterInventory
	error = json.Unmarshal(file, &characterInventory)
	if error != nil {
		panic(error)
	}

	// fmt.Printf("%+v\n", characterInventory.Response.ItemComponents.Sockets.Data["6917529046296926329"])

	file, error = ioutil.ReadFile("manifest.json")
	if error != nil {
		panic(error)
	}

	var manifestDatabase manifest.Database
	error = json.Unmarshal(file, &manifestDatabase)
	if error != nil {
		panic(error)
	}

	for _, sockets := range characterInventory.Response.ItemComponents.Sockets.Data {
		for _, socket := range sockets.Sockets {
			for _, hash := range socket.ReusablePlugHashes {
				id := strconv.FormatInt(hash, 10)
				fmt.Printf("%+v\n", hash)
				fmt.Printf("%+v\n", manifestDatabase.Data[id])
			}
		}
	}
}
