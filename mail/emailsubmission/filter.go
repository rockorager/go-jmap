package emailsubmission

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

type Filter interface {
	implementsFilter()
}

// Determines the set of EmailSubmissions returned in the results. If null, all
// objects in the account of this type are included in the results.
type FilterOperator struct {
	// This MUST be one of the following strings: “AND” / “OR” / “NOT”
	Operator jmap.Operator `json:"operator,omitempty"`

	// The conditions to evaluate against each record.
	Conditions []Filter `json:"conditions,omitempty"`
}

func (fo *FilterOperator) implementsFilter() {}

// FilterCondition is an interface that represents FilterCondition
// objects. A filter condition object can be either a named struct, ie
// MailboxFilterConditionName, or a MailboxFilter itself. MailboxFilters can
// be used to create complex filtering ie return mailboxes which are subscribed
// and NOT named Inbox
type FilterCondition struct {
	// identityIds field must be in this list to match
	IdentityIDs []jmap.ID `json:"identityIds,omitempty"`

	// emailId field must be in this list to match
	EmailIDs []jmap.ID `json:"emailIds,omitempty"`

	// threadId field must be in this list to match
	ThreadIDs []jmap.ID `json:"threadIds,omitempty"`

	// The undoStatus property must exactly match this to match
	UndoStatus string `json:"undoStatus,omitempty"`

	// UTC. The sendAt property must be before this time to match
	Before *time.Time `json:"before,omitempty"`

	// UTC. The sendAt property must be after this time to match
	After *time.Time `json:"after,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}
