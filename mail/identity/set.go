package identity

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail/emailsubmission"
)

// Modify identities
// https://www.rfc-editor.org/rfc/rfc8621.html#section-6.3
type Set struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IfInState string `json:"ifInState,omitempty"`

	Create map[jmap.ID]*Identity `json:"create,omitempty"`

	Update map[jmap.ID]jmap.Patch `json:"update,omitempty"`

	Destroy []jmap.ID `json:"destroy,omitempty"`
}

func (m *Set) Name() string { return "Identity/set" }

func (m *Set) Requires() []jmap.URI { return []jmap.URI{emailsubmission.URI} }

type SetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	OldState string `json:"oldState,omitempty"`

	NewState string `json:"newState,omitempty"`

	Created map[jmap.ID]*Identity `json:"created,omitempty"`

	Updated map[jmap.ID]*Identity `json:"updated,omitempty"`

	Destroyed []jmap.ID `json:"destroyed,omitempty"`

	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`

	NotUpdated map[jmap.ID]*jmap.SetError `json:"notUpdated,omitempty"`

	NotDestroyed map[jmap.ID]*jmap.SetError `json:"notDestroyed,omitempty"`
}

func newSetResponse() jmap.MethodResponse { return &SetResponse{} }
