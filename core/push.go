package core

// A StateChange object is sent to the client via Push mechanisms
type StateChange struct {
	// This MUST be the string "StateChange"
	Type string `json:"@type"`

	// Map of AccountID to TypeState. Only changed values will be in the map
	Changed map[string]TypeState
}

// TypeState is a map of Foo object names ("Mailbox", "Email", etc) to state
// property which would be returned by a call to Foo/get
type TypeState map[string]string
