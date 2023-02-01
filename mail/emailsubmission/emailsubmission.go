package emailsubmission

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

const URI jmap.URI = "urn:ietf:params:jmap:submission"

func init() {
	jmap.RegisterCapability(&Capability{})
}

// The EmailSubmission Capability
type Capability struct {
	// The maximum number of seconds the server supports for delayed
	// sending. A value of 0 indicates delayed sending is not supported
	MaxDelayedSend uint64 `json:"maxDelayedSend,omitempty"`

	// The set of SMTP submission extensions supported by the server, which
	// the client may use when creating an EmailSubmission object (see
	// Section 7). Each key in the object is the ehlo-name, and the value is
	// a list of ehlo-args.
	SubmissionExtensions map[string]string `json:"submissionExtensions,omitempty"`
}

func (m *Capability) URI() jmap.URI { return URI }

func (m *Capability) New() jmap.Capability { return &Capability{} }

type EmailSubmission struct {
	// The ID of the [EmailSubmission]
	//
	// immutable;server-set
	ID jmap.ID `json:"id,omitempty"`

	// The ID of the Identity to associate with this submission
	//
	// immutable
	IdentityID jmap.ID `json:"identityId,omitempty"`

	// The ID of the Email to send
	//
	// immutable
	EmailID jmap.ID `json:"emailId,omitempty"`

	// The Thread ID of the Email to send
	//
	// immutable;server-set
	ThreadID jmap.ID `json:"threadId,omitempty"`

	// The Envelope used for SMTP
	//
	// immutable
	Envelope *Envelope `json:"envelope,omitempty"`

	// The date the submission was/will be released for delivery
	//
	// immutable;server-set
	SendAt time.Time `json:"sendAt,omitempty"`

	// A status indicating if the send can be undone. One of:
	// - "pending": it may be possible to cancel
	// - "final": the message has been sent
	// - "canceled": the submission was canceled
	//
	// If this is "pending", a client can attempt to cancel by issuing a set
	// method with this set to canceled
	UndoStatus string `json:"undoStatus,omitempty"`

	// The delivery status for each recipient
	DeliveryStatus map[string]*DeliveryStatus `json:"deliveryStatus,omitempty"`

	// A list of blob IDs for DSNs received for this submission
	//
	// server-set
	DSNBlobIDs []jmap.ID `json:"dsnBlobIds,omitempty"`

	// A list of blob IDs for MDNs received for this submission
	//
	// server-set
	MDNBlobIDs []jmap.ID `json:"mdnBlobIds,omitempty"`
}

type Envelope struct {
	// The email address to use as the return address in the SMTP submission
	MailFrom *Address `json:"mailfrom,omitempty"`

	// The email address to send the message to
	RcptTo []*Address `json:"rcptTo,omitempty"`
}

type Address struct {
	// The email address
	Email string `json:"email,omitempty"`

	// Parameters to send with the email submission, if any SMTP extensions
	// are used
	Parameters interface{} `json:"parameters,omitempty"`
}

type DeliveryStatus struct {
	// The SMTP reply returned for the recipient
	SMTPReply string `json:"smtpReply,omitempty"`

	// Represents whether the message has been successfully delivered to the
	// recipient. Will be one of:
	// - "queued": In a local mail queue
	// - "yes": Delivered
	// - "no": Delivery failed
	// - "unknown": Final delivery status is unknown
	Delivered string `json:"delivered,omitempty"`

	// Whether the message has been displayed by the recipient. One of:
	// - "unknown"
	// - "yes"
	Displayed string `json:"displayed,omitempty"`
}
