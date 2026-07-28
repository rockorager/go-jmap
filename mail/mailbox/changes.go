package mailbox

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Get mailbox changes for the whole account
// https://www.rfc-editor.org/rfc/rfc8621.html#section-2.2
type Changes struct {
	Account jmap.ID `json:"accountId,omitempty"`

	SinceState string `json:"sinceState,omitempty"`

	MaxChanges uint64 `json:"maxChanges,omitempty"`
}

func (m *Changes) Name() string { return "Mailbox/changes" }

func (m *Changes) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type ChangesResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	OldState string `json:"oldState,omitempty"`

	NewState string `json:"newState,omitempty"`

	HasMoreChanges bool `json:"hasMoreChanges,omitempty"`

	Created []jmap.ID `json:"created,omitempty"`

	Updated []jmap.ID `json:"updated,omitempty"`

	Destroyed []jmap.ID `json:"destroyed,omitempty"`

	UpdatedProperties []string `json:"updatedProperties,omitempty"`
}

func newChangesResponse() jmap.MethodResponse { return &ChangesResponse{} }
