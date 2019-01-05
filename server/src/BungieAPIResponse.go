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
		ItemComponents struct {
			Sockets struct {
				Data map[string]Sockets `json:"data"`
			} `json:"sockets"`
		} `json:"itemComponents"`
	} `json:"Response"`
	ErrorCode int
}

// Item Item defintion.
type Item struct {
	Hash       int64  `json:"itemHash"`
	InstanceID string `json:"itemInstanceID"`
}

// Sockets Sockets definition.
type Sockets struct {
	Sockets []Socket `json:"sockets"`
}

// Socket Socket definition.
type Socket struct {
	PlugHash           int64          `json:"plugHash"`
	IsEnabled          bool           `json:"isEnabled"`
	IsVisible          bool           `json:"isVisible"`
	ReusablePlugHashes []int          `json:"reusablePlugHashes"`
	ReusablePlugs      []ReusablePlug `json:"reusablePlugs"`
}

// ReusablePlug Reusable Plug definition.
type ReusablePlug struct {
	PlugItemHash int  `json:"plugItemHash"`
	CanInsert    bool `json:"canInsert"`
	IsEnabled    bool `json:"enabled"`
}
