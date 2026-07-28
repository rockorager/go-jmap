package mdn

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Sends an RFC5322 message from an MDN object
// https://www.rfc-editor.org/rfc/rfc9007.html#section-2.1
type Send struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IdentityID jmap.ID `json:"identityId,omitempty"`

	Send map[jmap.ID]*MDN `json:"send,omitempty"`

	OnSuccessUpdateEmail map[jmap.ID]*jmap.Patch `json:"onSuccessUpdateEmail,omitempty"`
}

func (m *Send) Name() string { return "MDN/send" }

func (m *Send) Requires() []jmap.URI { return []jmap.URI{mail.URI, URI} }

type SendResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	Sent map[jmap.ID]*MDN `json:"sent,omitempty"`

	NotSent map[jmap.ID]*jmap.SetError `json:"notSent,omitempty"`
}

func newSendResponse() jmap.MethodResponse { return &SendResponse{} }
