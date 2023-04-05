package email

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// The "Email/import" method adds messages [RFC5322] to the set of Emails in an
// account.  The server MUST support messages with Email Address
// Internationalization (EAI) headers [RFC6532].  The messages must first be
// uploaded as blobs using the standard upload mechanism.
type Import struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// This is a state string as returned by the Foo/get method
	// (representing the state of all objects of this type in the account).
	// If supplied, the string must match the current state; otherwise, the
	// method will be aborted and a stateMismatch error returned. If null,
	// any changes will be applied to the current state.
	IfInState string `json:"ifInState,omitempty"`

	// A map of creation id (client specified) to EmailImport objects.
	Emails map[string]*EmailImport `json:"emails,omitempty"`
}

func (m *Import) Name() string { return "Email/import" }

func (m *Import) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type EmailImport struct {
	// The id of the blob containing the raw message [RFC5322].
	BlobID jmap.ID `json:"blobId,omitempty"`

	// The ids of the Mailboxes to assign this Email to.  At least one
	// Mailbox MUST be given.
	MailboxIDs map[jmap.ID]bool `json:"mailboxIds,omitempty"`

	// The keywords to apply to the Email.
	Keywords map[string]bool `json:"keywords,omitempty"`

	// The "receivedAt" date to set on the Email. The value must be in UTC
	ReceivedAt *time.Time `json:"receivedAt,omitempty"`
}

type ImportResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// The state string that would have been returned by Foo/get before
	// making the requested changes, or null if the server doesn’t know
	// what the previous state string was.
	OldState string `json:"oldState,omitempty"`

	// The state string that will now be returned by Foo/get.
	NewState string `json:"newState,omitempty"`

	// A map of the creation id to an object containing any properties of
	// the created Foo object that were not sent by the client. This
	// includes all server-set properties (such as the id in most object
	// types) and any properties that were omitted by the client and thus
	// set to a default by the server.
	//
	// This argument is null if no Foo objects were successfully created.
	Created map[jmap.ID]*Email `json:"created,omitempty"`

	// A map of the creation id to a SetError object for each Email that
	// failed to be created, or null if all successful.  The possible
	// errors are defined above.
	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`
}

func newImportResponse() jmap.MethodResponse { return &ImportResponse{} }
