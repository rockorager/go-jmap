// Package mdn is an implementation of RFC 9007: Handling Message Disposition
// Notification with the JSON Meta Application Protocol (JMAP). In plain terms,
// it handles read receipts of emails.
//
// Documentation strings for most of the protocol objects are taken from (or
// based on) contents of RFC 9007 and is subject to the IETF Trust Provisions.
// See https://trustee.ietf.org/license-info for details.
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
type MDN struct {
	// The Email ID of the received message to which this MDN is related
	ForEmailID jmap.ID `json:"forEmailId,omitempty"`

	// The Subject of the MDN
	Subject string `json:"subject,omitempty"`

	// The human-readable part of the MDN, as plain text
	TextBody string `json:"textBody,omitempty"`

	// If true, the content of the original message will appear in the third
	// component of the multipart/report generated for the MDN
	IncludeOriginalmessage bool `json:"includeOriginalMessage,omitempty"`

	// The name of the Mail User Agent (MUA) creating this MDN
	ReportingUA string `json:"reportinUA,omitempty"`

	// The object containing the diverse MDN disposition options
	Disposition *Disposition `json:"disposition,omitempty"`

	// The name of the gateway or MTA that translated a foreign
	// (non-internet) MDN into this MDN
	//
	// server-set
	MDNGateway string `json:"mdnGateway,omitempty"`

	// The original recipient address specified by the sender of the message
	// which the MDN is for
	//
	// server-set
	OriginalRecipient string `json:"originalRecipient,omitempty"`

	// The recipient for which the MDN is issued
	//
	// server-set
	FinalRecipient string `json:"finalRecipient,omitempty"`

	// The "Message-ID" header field of the message this MDN is for
	//
	// server-set
	OriginalMessageID string `json:"originalMessageId,omitempty"`

	// Additional information in the form of text messages when the "error"
	// disposition modifier appears
	//
	// server-set
	Error []string `json:"error,omitempty"`

	// The object where keys are extension-field names and values are
	// extension-field values
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
