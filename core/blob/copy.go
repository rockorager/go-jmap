package blob

import (
	"git.sr.ht/~rockorager/go-jmap"
)

// Copy a binary blob from one account to another
// https://www.rfc-editor.org/rfc/rfc8620.html#section-6.3
type Copy struct {
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	Account jmap.ID `json:"accountId,omitempty"`

	IDs []jmap.ID `json:"blobIds,omitempty"`
}

func (m *Copy) Name() string { return "Blob/copy" }

func (m *Copy) Requires() []jmap.URI { return nil }

type CopyResponse struct {
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	Account jmap.ID `json:"accountId,omitempty"`

	Copied map[jmap.ID]jmap.ID `json:"blobIds,omitempty"`

	NotCopied map[jmap.ID]*jmap.SetError `json:"notCopied,omitempty"`
}

func newCopyResponse() jmap.MethodResponse { return &CopyResponse{} }
