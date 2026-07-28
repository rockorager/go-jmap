package jmap

type Response struct {
	// Responses are the responses to the request, in the same order that
	// the request was made
	Responses []*Invocation `json:"methodResponses"`

	// A map of client-specified ID to server-assigned ID
	CreatedIDs map[ID]ID `json:"createdIds,omitempty"`

	// SessionState is the current state of the Session
	SessionState string `json:"sessionState"`
}
