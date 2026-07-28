package email

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Copy messages from one account to another
// https://www.rfc-editor.org/rfc/rfc8621.html#section-4.7
type Copy struct {
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	IfFromInState string `json:"ifFromInState,omitempty"`

	Account jmap.ID `json:"accountId,omitempty"`

	IfInState string `json:"ifInState,omitempty"`

	Create map[jmap.ID]*Email `json:"create,omitempty"`

	OnSuccessDestroyOriginal bool `json:"onSuccessDestroyOriginal,omitempty"`

	DestroyFromIfInState string `json:"destroyFromIfInState,omitempty"`
}

func (m *Copy) Name() string { return "Email/copy" }

func (m *Copy) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type CopyResponse struct {
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	Account jmap.ID `json:"accountId,omitempty"`

	OldState string `json:"oldState,omitempty"`

	NewState string `json:"newState,omitempty"`

	Created map[jmap.ID]*Email `json:"created,omitempty"`

	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`
}

func newCopyResponse() jmap.MethodResponse { return &CopyResponse{} }
