package wauth

// Whitelist is the list of allowed origins
type Whitelist struct {
	Systems []System `json:"systems"`
}

// System is the object to be used for the system
type System struct {
	ID          string `json:"ID"`
	Description string `json:"Description"`
}
