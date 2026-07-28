// Package mail is an implementation of JSON Metal Application Protocol (JMAP)
// for MAIL (RFC 8621)
package mail

import (
	"fmt"

	"git.sr.ht/~rockorager/go-jmap"
)

// urn:ietf:params:jmap:mail represents support for the Mailbox, Thread, Email,
// and SearchSnippet data types and associated API methods
const URI jmap.URI = "urn:ietf:params:jmap:mail"

const (
	// The Email event type
	EmailEvent jmap.EventType = "Email"

	// The EmailDelivery event type. This is a subset of an EmailEvent
	// subscription and only sends notifications when a new email has been
	// delivered, as opposed to any change for objects of type Email
	EmailDeliveryEvent jmap.EventType = "EmailDelivery"

	// The EmailSubmission event type
	EmailSubmissionEvent jmap.EventType = "EmailSubmission"

	// The Identity event type
	IdentityEvent jmap.EventType = "Identity"

	// The Mailbox event type
	MailboxEvent jmap.EventType = "Mailbox"

	// The Thread event type
	ThreadEvent jmap.EventType = "Thread"

	// The VacationResponse event type
	VacationResponseEvent jmap.EventType = "VacationResponse"
)

func init() {
	jmap.RegisterCapability(&Mail{})
}

type Mail struct {
	MaxMailboxesPerEmail       uint64 `json:"maxMailboxesPerEmail"`
	MaxMailboxDepth            uint64 `json:"maxMailboxDepth"`
	MaxSizeMailboxName         uint64 `json:"maxSizeMailboxName"`
	MaxSizeAttachmentsPerEmail uint64 `json:"maxSizeAttachmentsPerEmail"`

	// A list of all values the server supports for sorting
	EmailQuerySortOptions []string `json:"emailQuerySortOptions"`

	MayCreateTopLevelMailbox bool `json:"mayCreateTopLevelMailbox"`
}

func (m *Mail) URI() jmap.URI { return URI }

func (m *Mail) New() jmap.Capability { return &Mail{} }

// An Email address
type Address struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (a *Address) String() string {
	if a.Name == "" {
		return a.Email
	}
	return fmt.Sprintf("%s <%s>", a.Name, a.Email)
}
