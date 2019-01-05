package main

import (
	"encoding/json"
	"io/ioutil"
	manifest "manifest"
	"net/http"
	"strconv"
)

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

	var helmets []manifest.InventoryItem
	for _, item := range characterInventory.Response.ProfileInventory.Data.Items {
		id := strconv.FormatInt(item.Hash, 10)
		itemDefinition := manifestDatabase.Data[id]
		itemDefinition.ItemID = id
		if itemDefinition.Type != manifest.Armor {
			continue
		}
		if itemDefinition.SubType != manifest.HelmetArmor {
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
