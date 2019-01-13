package main

import (
	"encoding/json"
	"io/ioutil"
	manifest "manifest"
	"net/http"
	"strconv"
)

// FindHelmets Description
func FindHelmets(writer http.ResponseWriter, request *http.Request) {
	var file, error = ioutil.ReadFile("../misc/forsarion-inventory.json")
	if error != nil {
		panic(error)
	}
	var characterInventory CharacterInventory
	error = json.Unmarshal(file, &characterInventory)
	if error != nil {
		panic(error)
	}

	file, error = ioutil.ReadFile("manifest.json")
	if error != nil {
		panic(error)
	}

	var manifestDatabase manifest.Database
	error = json.Unmarshal(file, &manifestDatabase)
	if error != nil {
		panic(error)
	}

	var helmets []Helmet
	for _, item := range characterInventory.Response.ProfileInventory.Data.Items {
		id := strconv.FormatInt(item.Hash, 10)
		inventoryItem := manifestDatabase.Data[id]
		if inventoryItem.Type != manifest.Armor {
			continue
		}
		if inventoryItem.SubType != manifest.HelmetArmor {
			continue
		}

		var plugs []Plug
		for _, socket := range inventoryItem.Sockets.SocketEntries {
			for _, reusablePlugItem := range socket.ReusablePlugItems {
				id := strconv.FormatUint(reusablePlugItem.PlugItemHash, 10)
				inventoryItem := manifestDatabase.Data[id]
				plugs = append(plugs, makePlug(id, inventoryItem.DisplayProperties.Name, inventoryItem.DisplayProperties.Description))
			}
		}

		helmets = append(helmets, makeHelmet(
			id,
			inventoryItem.DisplayProperties.Name,
			inventoryItem.DisplayProperties.Description,
			plugs))
	}

	json, error := json.Marshal(HelmetResponse{helmets})
	if error != nil {
		panic(error)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}
