package maskedemail

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

type MaskedEmail struct {
	// The ID of the MaskedEmail
	//
	// immutable; server-set
	ID jmap.ID `json:"id,omitempty"`
	// The email address
	//
	// immutable; server-set
	Email string `json:"email,omitempty"`

	// One of the following:
	// 	pending
	//	enabled - the address is active
	// 	disabled - the address is active, but mail goes to trash
	//	deleted - inactive, mail is bounced
	State string `json:"state,omitempty"`

	// The protocol and domain of the site the user is using the masked
	// email for (eg https://www.example.com)
	ForDomain string `json:"forDomain,omitempty"`

	// A short user-supplied description of what this address is for. If the
	// user does not supply, leave as the empty string
	Description string `json:"description"`

	// The UTC time the most recent message was received
	//
	// server-set
	LastMessageAt *time.Time `json:"lastMessageAt,omitempty"`

	// The time the address was created
	//
	// immutable; server-set
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// A deep link to the credential or other record related to this address
	URL string `json:"url,omitempty"`

	// Create only
	// This is only used during a Set, otherwise ignored. It is not returned
	// in a Get call. If supplied, the address will start with the prefix.
	// It must be <= 64 chars and [a-z0-9_]
	EmailPrefix string `json:"emailPrefix,omitempty"`
}

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
//
// Objects of type MaskedEmail are fetched via a call to MaskedEmail/get The ids
// argument may be null to fetch all at once.
type Get struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId,omitempty"`

	// The ids of the Foo objects to return. If null, then all records of
	// the data type are returned, if this is supported for that data type
	// and the number of records does not exceed the maxObjectsInGet limit.
	IDs []jmap.ID `json:"ids,omitempty"`

	// If supplied, only the properties listed in the array are returned
	// for each Mailbox object. If null, all properties of the object are
	// returned. The id property of the object is always returned, even if
	// not explicitly requested. If an invalid property is requested, the
	// call MUST be rejected with an invalidArguments error.
	Properties []string `json:"properties,omitempty"`

	// Use IDs from a previous call
	ReferenceIDs *jmap.ResultReference `json:"#ids,omitempty"`

	// Use Properties from a previous call
	ReferenceProperties *jmap.ResultReference `json:"#properties,omitempty"`
}
