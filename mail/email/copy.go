package email

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// This is a standard "/copy" method as described in [RFC8620], Section 5.4,
// except only the "mailboxIds", "keywords", and "receivedAt" properties may be
// set during the copy.  This method cannot modify the message represented by
// the Email.
type Copy struct {
	// The id of the account to copy records from.
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	// This is a state string as returned by the Foo/get method. If
	// supplied, the string must match the current state of the account
	// referenced by the fromAccountId when reading the data to be copied;
	// otherwise, the method will be aborted and a stateMismatch error
	// returned. If null, the data will be read from the current state.
	IfFromInState string `json:"ifFromInState,omitempty"`

	// The id of the account to copy records to. This MUST be different to
	// the fromAccountId.
	Account jmap.ID `json:"accountId,omitempty"`

	// This is a state string as returned by the Foo/get method. If
	// supplied, the string must match the current state of the account
	// referenced by the accountId; otherwise, the method will be aborted
	// and a stateMismatch error returned. If null, any changes will be
	// applied to the current state.
	IfInState string `json:"ifInState,omitempty"`

	// A map of the creation id to a Foo object. The Foo object MUST
	// contain an id property, which is the id (in the fromAccount) of the
	// record to be copied. When creating the copy, any other properties
	// included are used instead of the current value for that property on
	// the original.
	Create map[jmap.ID]*Email `json:"create,omitempty"`

	// If true, an attempt will be made to destroy the original records
	// that were successfully copied: after emitting the Foo/copy response,
	// but before processing the next method, the server MUST make a single
	// call to Foo/set to destroy the original of each successfully copied
	// record; the output of this is added to the responses as normal, to
	// be returned to the client.
	OnSuccessDestroyOriginal bool `json:"onSuccessDestroyOriginal,omitempty"`

	// This argument is passed on as the ifInState argument to the implicit
	// Foo/set call, if made at the end of this request to destroy the
	// originals that were successfully copied.
	DestroyFromIfInState string `json:"destroyFromIfInState,omitempty"`
}

func (m *Copy) Name() string { return "Email/copy" }

func (m *Copy) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type CopyResponse struct {
	// The id of the account records were copied from.
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	// The id of the account records were copied to.
	Account jmap.ID `json:"accountId,omitempty"`

	// The state string that would have been returned by Foo/get on the
	// account records that were copied to before making the requested
	// changes, or null if the server doesn’t know what the previous state
	// string was.
	OldState string `json:"oldState,omitempty"`

	// The state string that will now be returned by Foo/get on the account
	// records were copied to.
	NewState string `json:"newState,omitempty"`

	// A map of the creation id to an object containing any properties of
	// the copied Foo object that are set by the server (such as the id in
	// most object types; note, the id is likely to be different to the id
	// of the object in the account it was copied from).
	//
	// This argument is null if no Foo objects were successfully copied.
	Created map[jmap.ID]*Email `json:"created,omitempty"`

	// A map of the creation id to a SetError object for each record that
	// failed to be copied, or null if none.
	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`
}

func newCopyResponse() jmap.MethodResponse { return &CopyResponse{} }
