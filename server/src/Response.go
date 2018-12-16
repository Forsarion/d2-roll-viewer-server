package main

// CharacterInventory Root Level Character Inventory definition.
type CharacterInventory struct {
	Response struct {
		// ProfileInventory Profile Inventory definition.
		ProfileInventory struct {
			privacy int
			Data    struct {
				Items []Item `json:"items"`
			} `json:"data"`
		} `json:"profileInventory"`
	} `json:"Response"`
	ErrorCode int
}

// Item Item defintion.
type Item struct {
	Hash       int64  `json:"itemHash"`
	InstanceID string `json:"itemInstanceID"`
}
