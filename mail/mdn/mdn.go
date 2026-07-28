package mdn

import "git.sr.ht/~rockorager/go-jmap"

const URI jmap.URI = "urn:ietf:params:jmap:mdn"

func init() {
	jmap.RegisterCapability(&Capability{})
	jmap.RegisterMethod("MDN/send", newSendResponse)
	jmap.RegisterMethod("MDN/parse", newParseResponse)
}

// The MDN Capability
type Capability struct{}

func (m *Capability) URI() jmap.URI { return URI }

func (m *Capability) New() jmap.Capability { return &Capability{} }

// A Message Delivery Notification (MDN) object
// https://www.rfc-editor.org/rfc/rfc9007.html#section-2
type MDN struct {
	ForEmailID jmap.ID `json:"forEmailId,omitempty"`

	Subject string `json:"subject,omitempty"`

	TextBody string `json:"textBody,omitempty"`

	IncludeOriginalmessage bool `json:"includeOriginalMessage,omitempty"`

	ReportingUA string `json:"reportinUA,omitempty"`

	Disposition *Disposition `json:"disposition,omitempty"`

	MDNGateway string `json:"mdnGateway,omitempty"`

	OriginalRecipient string `json:"originalRecipient,omitempty"`

	FinalRecipient string `json:"finalRecipient,omitempty"`

	OriginalMessageID string `json:"originalMessageId,omitempty"`

	Error []string `json:"error,omitempty"`

	ExtensionFields map[string]string `json:"extensionFields,omitempty"`
}

type Disposition struct {
	// This MUST be one of the following strings:
	// - "manual-action"
	// - "automatic-action"
	ActionMode string `json:"actionMode,omitempty"`

	// This MUST be one of the following strings:
	// - "mdn-sent-manually"
	// - "mdn-sent-automatically"
	SendingMode string `json:"sendingMode,omitempty"`

	// This MUST be one of the following strings:
	// - "deleted"
	// - "dispatched"
	// - "displayed"
	// - "processed"
	Type string `json:"type,omitempty"`
}
