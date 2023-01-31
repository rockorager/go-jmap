package jmap

type Response struct {
	// An array of responses, in the same format as the Calls on the
	// Request object. The output of the methods will be added to the
	// methodResponses array in the same order as the methods are processed.
	Responses []*Invocation `json:"methodResponses"`

	// A map of (client-specified) creation id to the id the server assigned
	// when a record was successfully created.
	CreatedIDs map[ID]ID `json:"createdIds,omitempty"`

	// The current value of the “state” string on the JMAP Session object, as
	// described in section 2. Clients may use this to detect if this object
	// has changed and needs to be refetched.
	SessionState string `json:"sessionState"`
}
