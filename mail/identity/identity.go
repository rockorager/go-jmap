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

// Information about an email address or domain the user may send from
// https://www.rfc-editor.org/rfc/rfc8621.html#section-6
type Identity struct {
	ID jmap.ID `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	ReplyTo []*mail.Address `json:"replyTo,omitempty"`

	Bcc []*mail.Address `json:"bcc,omitempty"`

	TextSignature string `json:"textSignature,omitempty"`

	HTMLSignature string `json:"htmlSignature,omitempty"`

	MayDelete bool `json:"mayDelete,omitempty"`
}
