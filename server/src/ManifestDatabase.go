package main

// ManifestDatabase Database of all items
type ManifestDatabase struct {
	// Data id to data dictionary
	Data map[string]InventoryItem `json:"DestinyInventoryItemDefinition"`
}

// InventoryItem Description
type InventoryItem struct {
	ItemID            string  `json:"itemID"`
	SubType           SubType `json:"itemSubType"`
	Type              Type    `json:"itemType"`
	DisplayProperties struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"displayProperties"`
}

// SubType Description
type SubType uint

const (
	// HelmetArmor https://bungie-net.github.io/#/components/schemas/Destiny.DestinyItemSubType
	HelmetArmor SubType = 26
)

// Type Description
type Type uint

const (
	// Armor https://bungie-net.github.io/#/components/schemas/Destiny.DestinyItemType
	Armor Type = 2
)
