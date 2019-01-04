package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func findHelmets(writer http.ResponseWriter, request *http.Request) {
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

	var manifestDatabase ManifestDatabase
	error = json.Unmarshal(file, &manifestDatabase)
	if error != nil {
		panic(error)
	}

	var helmets []InventoryItem
	for _, item := range characterInventory.Response.ProfileInventory.Data.Items {
		id := strconv.FormatInt(item.Hash, 10)
		itemDefinition := manifestDatabase.Data[id]
		itemDefinition.ItemID = id
		if itemDefinition.Type != Armor {
			continue
		}
		if itemDefinition.SubType != HelmetArmor {
			continue
		}
		helmets = append(helmets, itemDefinition)
	}

	json, error := json.Marshal(helmets)
	if error != nil {
		panic(error)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}
