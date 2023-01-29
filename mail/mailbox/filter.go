package mailbox

import "git.sr.ht/~rockorager/go-jmap"

type Filter interface {
	implementsFilter()
}

// Determines the set of Mailboxes returned in the results. If null, all
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
	// The Mailbox parentId property must match the given value exactly.
	ParentID string `json:"parentId,omitempty"`
	// The Mailbox name property contains the given string.
	Name string `json:"name,omitempty"`
	// The Mailbox role property must match the given value exactly.
	Role Role `json:"role,omitempty"`
	// If true, a Mailbox matches if it has any non-null value for its role
	// property.
	HasAnyRole bool `json:"hasAnyRole,omitempty"`
	// The isSubscribed property of the Mailbox must be identical to the
	// value given to match the condition.
	IsSubscribed bool `json:"isSubscribed,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}
