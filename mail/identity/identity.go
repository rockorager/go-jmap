package identity

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

func init() {
	jmap.RegisterMethod("Identity/get", newGetResponse)
	jmap.RegisterMethod("Identity/changes", newChangesResponse)
	jmap.RegisterMethod("Identity/set", newSetResponse)
}

type Identity struct {
	// The ID of the Identity
	//
	// immutable;server-set
	ID jmap.ID `json:"id,omitempty"`

	// The "From" name the client SHOULD use when creating a new Email from
	// this identity
	Name string `json:"name,omitempty"`

	// The "From" email address the client MUST use when creating a new
	// Email from this Identity. If the mailbox part of the address (the
	// section before the "@") is the single character "*", then the client
	// may use any valid address ending in that domain
	//
	// immutable
	Email string `json:"email,omitempty"`

	// The Reply-To value the client SHOULD set when creating a new Email
	// from this identity
	ReplyTo []*mail.Address `json:"replyTo,omitempty"`

	// The Bcc value the client SHOULD set when creating a new Email from
	// this Identity
	Bcc []*mail.Address `json:"bcc,omitempty"`

	// A signature the client SHOULD insert into new plaintext messages that
	// will be sent from this identity
	TextSignature string `json:"textSignature,omitempty"`

	// A signature the client SHOULD insert into new html messages that
	// will be sent from this identity
	HTMLSignature string `json:"htmlSignature,omitempty"`

	// If the user is allowed to delete this identity
	//
	// server-set
	MayDelete bool `json:"mayDelete,omitempty"`
}
