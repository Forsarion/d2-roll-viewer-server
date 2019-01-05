package manifest

// Database Database of all items
type Database struct {
	// Data id to data dictionary
	Data map[string]InventoryItem `json:"DestinyInventoryItemDefinition"`
}

// InventoryItem Description
type InventoryItem struct {
	ItemID            string
	SubType           SubType `json:"itemSubType"`
	Type              Type    `json:"itemType"`
	DisplayProperties struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"displayProperties"`
	Sockets struct {
		Detail        string   `json:"detail"`
		SocketEntries []Socket `json:"socketEntries"`
	} `json:"sockets"`
}

// Socket Description
type Socket struct {
	SocketTypeHash        uint64             `json:"socketTypeHash"`
	SingleInitialItemHash uint64             `json:"singleInitialItemHash"`
	ReusablePlugItems     []ReusablePlugItem `json:"reusablePlugItems"`
	PlugSources           uint               `json:"plugSources"`
}

// ReusablePlugItem Desription
type ReusablePlugItem struct {
	PlugItemHash uint64 `json:"plugItemHash"`
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
