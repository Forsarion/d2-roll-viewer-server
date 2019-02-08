package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	manifest "github.com/Berk0ld/d2-roll-viewer-server/server/manifest"
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

		for socketIndex := 0; socketIndex < len(inventoryItem.Sockets.SocketEntries); socketIndex++ {
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

			inventoryData := characterInventory.Response.ItemComponents.Sockets.Data[item.InstanceID]
			randomPlug := inventoryData.Sockets[socketIndex]

			for _, plug := range randomPlug.ReusablePlugs {
				id := strconv.FormatInt(plug.Hash, 10)
				inventoryItem := manifestDatabase.Data[id]
				plugs = append(plugs, makePlug(
					id,
					inventoryItem.DisplayProperties.Name,
					inventoryItem.DisplayProperties.Description,
					plug.IsEnabled))
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
