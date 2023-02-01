package mdn

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Sends an RFC5322 message from an MDN object
type Send struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId,omitempty"`

	// The ID of the Identity to associate with these MDNs
	IdentityID jmap.ID `json:"identityId,omitempty"`

	// A map of client-specified creation ID to MDN object
	Send map[jmap.ID]*MDN `json:"send,omitempty"`

	// A map of the ID to a patch of update the Email object referenced by
	// MDN/send, if the sending succeeds. The ID will always be a backward
	// reference to the creation ids
	OnSuccessUpdateEmail map[jmap.ID]*jmap.Patch `json:"onSuccessUpdateEmail,omitempty"`
}

func (m *Send) Name() string { return "MDN/send" }

func (m *Send) Requires() []jmap.URI { return []jmap.URI{mail.URI, URI} }

type SendResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// A map of the creation ID to an MDN containing any properties that
	// were not set by the client
	Sent map[jmap.ID]*MDN `json:"sent,omitempty"`

	// A map of creation ID to a SetError for each MDN not sent
	NotSent map[jmap.ID]*jmap.SetError `json:"notSent,omitempty"`
}

func newSendResponse() jmap.MethodResponse { return &SendResponse{} }
