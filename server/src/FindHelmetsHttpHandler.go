package main

import (
	"encoding/json"
	"fmt"
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

		for socketIndex, socket := range inventoryItem.Sockets.SocketEntries {
			var armorPerks *manifest.SocketCategory
			for _, socketCategory := range inventoryItem.Sockets.SocketCategories {
				if socketCategory.Hash == manifest.ArmorPerks {
					armorPerks = &socketCategory
					break
				}
			}

			if armorPerks == nil {
				panic("Missing ArmorPerks category on Helmets!")
			}

			isArmorPerk := false
			for _, value := range armorPerks.Indexes {
				if uint(socketIndex) == value {
					isArmorPerk = true
					break
				}
			}

			if isArmorPerk == false {
				continue
			}

			id := strconv.FormatUint(socket.SingleInitialItemHash, 10)
			inventoryItem := manifestDatabase.Data[id]
			plugs = append(plugs, makePlug(id, inventoryItem.DisplayProperties.Name, inventoryItem.DisplayProperties.Description))
		}

		for _, socket := range characterInventory.Response.ItemComponents.Sockets.Data["1664085089"].Sockets {
			fmt.Printf("socket: %+v\n", socket)
			id := strconv.FormatUint(uint64(socket.PlugHash), 10)
			inventoryItem := manifestDatabase.Data[id]
			fmt.Printf("name: %+v\n", inventoryItem.DisplayProperties.Name)
			fmt.Printf("description: %+v\n", inventoryItem.DisplayProperties.Description)
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
