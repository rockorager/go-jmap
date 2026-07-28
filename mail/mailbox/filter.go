package mailbox

import "git.sr.ht/~rockorager/go-jmap"

type Filter interface {
	implementsFilter()
}

type FilterOperator struct {
	Operator jmap.Operator `json:"operator,omitempty"`

	Conditions []Filter `json:"conditions,omitempty"`
}

func (fo *FilterOperator) implementsFilter() {}

// Filter criteria for mailbox queries
// https://www.rfc-editor.org/rfc/rfc8621.html#section-2.3
type FilterCondition struct {
	ParentID jmap.ID `json:"parentId,omitempty"`

	Name string `json:"name,omitempty"`

	Role Role `json:"role,omitempty"`

	HasAnyRole bool `json:"hasAnyRole,omitempty"`

	IsSubscribed bool `json:"isSubscribed,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}
