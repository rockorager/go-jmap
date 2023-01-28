package thread

import "git.sr.ht/~rockorager/go-jmap"

const MailCapability = "urn:ietf:params:jmap:mail"

func init() {
	jmap.RegisterMethods(
		&Get{},
		&Changes{},
	)
}

// Replies are grouped together with the original message to form a Thread. In
// JMAP, a Thread is simply a flat list of Emails, ordered by date. Every Email
// MUST belong to a Thread, even if it is the only Email in the Thread.
type Thread struct {
	// The ID of the thread
	ID string `json:"id,omitempty"`

	// The ids of the Emails in the Thread, sorted by the receivedAt date
	// of the Email, oldest first. If two Emails have an identical date,
	// the sort is server dependent but MUST be stable (sorting by id is
	// recommended).
	EmailIDs []string `json:"emailIds,omitempty"`
}
