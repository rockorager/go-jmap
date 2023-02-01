package mailbox

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// This is a standard “/changes” method as described in [@!RFC8620], Section
// 5.2.
//
// When the state of the set of Foo records in an account changes on the server
// (whether due to creation, updates, or deletion), the state property of the
// Foo/get response will change. The Foo/changes method allows a client to
// efficiently update the state of its Foo cache to match the new state on the
// server.
type Changes struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId,omitempty"`

	// The current state of the client. This is the string that was
	// returned as the state argument in the Foo/get response. The server
	// will return the changes that have occurred since this state.
	SinceState string `json:"sinceState,omitempty"`

	// The maximum number of ids to return in the response. The server MAY
	// choose to return fewer than this value but MUST NOT return more. If
	// not given by the client, the server may choose how many to return.
	// If supplied by the client, the value MUST be a positive integer
	// greater than 0. If a value outside of this range is given, the
	// server MUST reject the call with an invalidArguments error.
	MaxChanges uint64 `json:"maxChanges,omitempty"`
}

func (m *Changes) Name() string { return "Mailbox/changes" }

func (m *Changes) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

// This is a standard “/changes” method as described in [@!RFC8620], Section
// 5.2 but with one extra argument to the response: updatedProperties
type ChangesResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// This is the sinceState argument echoed back; it’s the state from
	// which the server is returning changes.
	OldState string `json:"oldState,omitempty"`

	// This is the state the client will be in after applying the set of
	// changes to the old state.
	NewState string `json:"newState,omitempty"`

	// If true, the client may call Foo/changes again with the newState
	// returned to get further updates. If false, newState is the current
	// server state.
	HasMoreChanges bool `json:"hasMoreChanges,omitempty"`

	// An array of ids for records that have been created since the old
	// state.
	Created []jmap.ID `json:"created,omitempty"`

	// An array of ids for records that have been updated since the old
	// state.
	Updated []jmap.ID `json:"updated,omitempty"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []jmap.ID `json:"destroyed,omitempty"`

	// If only the “totalEmails”, “unreadEmails”, “totalThreads”, and/or
	// “unreadThreads” Mailbox properties have changed since the old state,
	// this will be the list of properties that may have changed. If the
	// server is unable to tell if only counts have changed, it MUST just
	// be null.
	UpdatedProperties []string `json:"updatedProperties,omitempty"`
}

func newChangesResponse() jmap.MethodResponse { return &ChangesResponse{} }
