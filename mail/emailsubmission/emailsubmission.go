package emailsubmission

import (
	"encoding/json"
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

const URI jmap.URI = "urn:ietf:params:jmap:submission"

func init() {
	jmap.RegisterCapability(&Capability{})
	jmap.RegisterMethod("EmailSubmission/get", newGetResponse)
	jmap.RegisterMethod("EmailSubmission/changes", newChangesResponse)
	jmap.RegisterMethod("EmailSubmission/query", newQueryResponse)
	jmap.RegisterMethod("EmailSubmission/queryChanges", newQueryChangesResponse)
	jmap.RegisterMethod("EmailSubmission/set", newSetResponse)
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
	SubmissionExtensions json.RawMessage `json:"submissionExtensions,omitempty"`
}

func (m *Capability) URI() jmap.URI { return URI }

func (m *Capability) New() jmap.Capability { return &Capability{} }

// Submission of an Email for delivery to one or more recipients.
// https://www.rfc-editor.org/rfc/rfc8621.html#section-7
type EmailSubmission struct {
	ID jmap.ID `json:"id,omitempty"`

	IdentityID jmap.ID `json:"identityId,omitempty"`

	EmailID jmap.ID `json:"emailId,omitempty"`

	ThreadID jmap.ID `json:"threadId,omitempty"`

	Envelope *Envelope `json:"envelope,omitempty"`

	SendAt *time.Time `json:"sendAt,omitempty"`

	UndoStatus string `json:"undoStatus,omitempty"`

	DeliveryStatus map[string]*DeliveryStatus `json:"deliveryStatus,omitempty"`

	DSNBlobIDs []jmap.ID `json:"dsnBlobIds,omitempty"`

	MDNBlobIDs []jmap.ID `json:"mdnBlobIds,omitempty"`
}

func (s *EmailSubmission) MarshalJSON() ([]byte, error) {
	if s.SendAt != nil && s.SendAt.Location() != time.UTC {
		utc := s.SendAt.UTC()
		s.SendAt = &utc
	}
	// create a type alias to avoid infinite recursion
	type Alias EmailSubmission
	return json.Marshal((*Alias)(s))
}

type Envelope struct {
	// The email address to use as the return address in the SMTP submission
	MailFrom *Address `json:"mailFrom,omitempty"`

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
