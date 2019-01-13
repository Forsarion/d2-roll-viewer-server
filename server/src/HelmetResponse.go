package main

// HelmetResponse Description
type HelmetResponse struct {
	Helmets []Helmet `json:"helmets"`
}

// Helmet Description
type Helmet struct {
	ItemID      string `json:"itemID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Plugs       []Plug `json:"plugs"`
}

// makeHelmet Constructor
func makeHelmet(itemID string, name string, description string, plugs []Plug) Helmet {
	return Helmet{itemID, name, description, plugs}
}

// Plug Constructor
type Plug struct {
	ItemID      string `json:"itemID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// makePlug Constructor
func makePlug(itemID string, name string, description string) Plug {
	return Plug{itemID, name, description}
}
