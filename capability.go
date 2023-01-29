package jmap

// A Capability broadcasts that the server supports underlying methods
type Capability interface {
	// The URI of the capability, eg "urn:ietf:params:jmap:core"
	URI() string

	// Generates a pointer to a new Capability object
	New() Capability
}

// Register a Capability
func RegisterCapability(c Capability) {
	capabilities[c.URI()] = c
}

var capabilities = make(map[string]Capability)
