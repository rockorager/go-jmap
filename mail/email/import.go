package email

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Import email from binary blobs
// https://www.rfc-editor.org/rfc/rfc8621.html#section-4.8
type Import struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IfInState string `json:"ifInState,omitempty"`

	Emails map[string]*EmailImport `json:"emails,omitempty"`
}

func (m *Import) Name() string { return "Email/import" }

func (m *Import) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type EmailImport struct {
	BlobID jmap.ID `json:"blobId,omitempty"`

	MailboxIDs map[jmap.ID]bool `json:"mailboxIds,omitempty"`

	Keywords map[string]bool `json:"keywords,omitempty"`

	ReceivedAt *time.Time `json:"receivedAt,omitempty"`
}

type ImportResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	OldState string `json:"oldState,omitempty"`

	NewState string `json:"newState,omitempty"`

	Created map[jmap.ID]*Email `json:"created,omitempty"`

	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`
}

func newImportResponse() jmap.MethodResponse { return &ImportResponse{} }
