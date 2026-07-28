package jmap

// An EventType is the name of a Type provided by a capability which may be
// subscribed to using a PushSubscription or an EventSource connection. Each
// specification may define their own types and events
//
// EventType is the type of object the Event is for ("Mailbox", "Email")
type EventType string

// Subscribe to all events
const AllEvents EventType = "*"

// A StateChange object is sent to the client via Push mechanisms. It
// communicates when a change has occurred
type StateChange struct {
	// This MUST be the string "StateChange"
	Type string `json:"@type"`

	// Map of AccountID to TypeState. Only changed values will be in the map
	Changed map[ID]TypeState `json:"changed"`
}

// TypeState is a map of Foo object names ("Mailbox", "Email", etc) to state
// property which would be returned by a call to Foo/get
type TypeState map[string]string
