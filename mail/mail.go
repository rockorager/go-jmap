package mail

import (
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
	// The maximum number of Mailboxes (see Section 2) that can be can
	// assigned to a single Email object (see Section 4). This MUST be an
	// integer >= 1, or null for no limit (or rather, the limit is always
	// the number of Mailboxes in the account).
	MaxMailboxesPerEmail uint64 `json:"maxMailboxesPerEmail"`

	// The maximum depth of the Mailbox hierarchy (i.e., one more than the
	// maximum number of ancestors a Mailbox may have), or null for no
	// limit.
	MaxMailboxDepth uint64 `json:"maxMailboxDepth"`

	// The maximum length, in (UTF-8) octets, allowed for the name of a
	// Mailbox. This MUST be at least 100, although it is recommended
	// servers allow more.
	MaxSizeMailboxName uint64 `json:"maxSizeMailboxName"`

	// The maximum total size of attachments, in octets, allowed for a
	// single Email object. A server MAY still reject the import or
	// creation of an Email with a lower attachment size total (for
	// example, if the body includes several megabytes of text, causing the
	// size of the encoded MIME structure to be over some server-defined
	// limit).
	//
	// Note that this limit is for the sum of unencoded attachment sizes.
	// Users are generally not knowledgeable about encoding overhead, etc.,
	// nor should they need to be, so marketing and help materials normally
	// tell them the “max size attachments”. This is the unencoded size
	// they see on their hard drive, so this capability matches that and
	// allows the client to consistently enforce what the user understands
	// as the limit.
	//
	// The server may separately have a limit for the total size of the
	// message [@!RFC5322], created by combining the attachments (often
	// base64 encoded) with the message headers and bodies. For example,
	// suppose the server advertises maxSizeAttachmentsPerEmail: 50000000
	// (50 MB). The enforced server limit may be for a message size of
	// 70000000 octets. Even with base64 encoding and a 2 MB HTML body, 50
	// MB attachments would fit under this limit.
	MaxSizeAttachmentsPerEmail uint64 `json:"maxSizeAttachmentsPerEmail"`

	// A list of all the values the server supports for the “property”
	// field of the Comparator object in an Email/query sort (see Section
	// 4.4.2). This MAY include properties the client does not recognise
	// (for example, custom properties specified in a vendor extension).
	// Clients MUST ignore any unknown properties in the list.
	EmailQuerySortOptions []string `json:"emailQuerySortOptions"`

	// If true, the user may create a Mailbox (see Section 2) in this
	// account with a null parentId. (Permission for creating a child of an
	// existing Mailbox is given by the myRights property on that Mailbox.)
	MayCreateTopLevelMailbox bool `json:"mayCreateTopLevelMailbox"`
}

func (m *Mail) URI() jmap.URI { return URI }

func (m *Mail) New() jmap.Capability { return &Mail{} }

// An Email address
type Address struct {
	// The display-name of the mailbox [@!RFC5322]. If this is a
	// quoted-string:
	//
	//     The surrounding DQUOTE characters are removed. Any quoted-pair
	//     is decoded. White space is unfolded, and then any leading and
	//     trailing white space is removed.
	//
	// If there is no display-name but there is a comment immediately
	// following the addr-spec, the value of this SHOULD be used instead.
	// Otherwise, this property is null.
	Name string `json:"name,omitempty"`

	// The addr-spec of the mailbox [@!RFC5322].
	Email string `json:"email,omitempty"`
}
