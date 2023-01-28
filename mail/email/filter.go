package email

import "git.sr.ht/~rockorager/go-jmap"

type Filter interface {
	implementsFilter()
}

// Determines the set of Emails returned in the results. If null, all objects
// in the account of this type are included in the results.
type FilterOperator struct {
	// This MUST be one of the following strings: “AND” / “OR” / “NOT”
	Operator jmap.FilterOperator `json:"operator,omitempty"`

	// The conditions to evaluate against each record.
	Conditions []Filter `json:"conditions,omitempty"`
}

func (fo *FilterOperator) implementsFilter() {}

// EmailFilterCondition is an interface that represents FilterCondition
// objects. A filter condition object can be either a named struct, ie
// EmailFilterConditionName, or an EmailFilter itself. EmailFilters can
// be used to create complex filtering
type FilterCondition struct {
	// A Mailbox id.  An Email must be in this Mailbox to match the condition.
	InMailbox string `json:"inMailbox,omitempty"`

	// A list of Mailbox ids.  An Email must be in at least one Mailbox not in this
	// list to match the condition.  This is to allow messages solely in trash/spam
	// to be easily excluded from a search.
	InMailboxOtherThan []string `json:"inMailboxOtherThan,omitempty"`

	// The "receivedAt" date-time of the Email must be before this date- time to
	// match the condition.
	Before jmap.Date `json:"before,omitempty"`

	// The "receivedAt" date-time of the Email must be the same or after this
	// date-time to match the condition.
	After jmap.Date `json:"after,omitempty"`

	// The "size" property of the Email must be equal to or greater than this
	// number to match the condition.
	MinSize uint64 `json:"minSize,omitempty"`

	// The "size" property of the Email must be less than this number to match the
	// condition.
	MaxSize uint64 `json:"maxSize,omitempty"`

	// All Emails (including this one) in the same Thread as this Email must have
	// the given keyword to match the condition.
	AllInThreadHaveKeyword string `json:"allInThreadHaveKeyword,omitempty"`

	// At least one Email (possibly this one) in the same Thread as this Email must
	// have the given keyword to match the condition.
	SomeInThreadHaveKeyword string `json:"someInThreadHaveKeyword,omitempty"`

	// All Emails (including this one) in the same Thread as this Email must *not*
	// have the given keyword to match the condition.
	NoneInThreadHaveKeyword string `json:"noneInThreadHaveKeyword,omitempty"`

	// This Email must have the given keyword to match the condition.
	HasKeyword string `json:"hasKeyword,omitempty"`

	// This Email must not have the given keyword to match the condition.
	NotKeyword string `json:"notKeyword,omitempty"`

	// The "hasAttachment" property of the Email must be identical to the value
	// given to match the condition.
	HasAttachment bool `json:"hasAttachment,omitempty"`

	// Looks for the text in Emails.  The server MUST look up text in the From, To,
	// Cc, Bcc, and Subject header fields of the message and SHOULD look inside any
	// "text/*" or other body parts that may be converted to text by the server.
	// The server MAY extend the search to any additional textual property.
	Text string `json:"text,omitempty"`

	// Looks for the text in the From header field of the message.
	From string `json:"from,omitempty"`

	// Looks for the text in the To header field of the message.
	To string `json:"to,omitempty"`

	// Looks for the text in the Cc header field of the message.
	Cc string `json:"cc,omitempty"`

	// Looks for the text in the Bcc header field of the message.
	Bcc string `json:"bcc,omitempty"`

	// Looks for the text in the Subject header field of the message.
	Subject string `json:"subject,omitempty"`

	// Looks for the text in one of the body parts of the message.  The server MAY
	// exclude MIME body parts with content media types other than "text/*" and
	// "message/*" from consideration in search matching.  Care should be taken to
	// match based on the text content actually presented to an end user by viewers
	// for that media type or otherwise identified as appropriate for search
	// indexing. Matching document metadata uninteresting to an end user (e.g.,
	// markup tag and attribute names) is undesirable.
	Body string `json:"body,omitempty"`

	// The array MUST contain either one or two elements.  The first element is the
	// name of the header field to match against.  The second (optional) element is
	// the text to look for in the header field value.  If not supplied, the
	// message matches simply if it has a header field of the given name.
	Header []string `json:"header,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}
