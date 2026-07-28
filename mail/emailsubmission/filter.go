package emailsubmission

import (
	"encoding/json"
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

type Filter interface {
	implementsFilter()
}

type FilterOperator struct {
	Operator jmap.Operator `json:"operator,omitempty"`

	Conditions []Filter `json:"conditions,omitempty"`
}

func (fo *FilterOperator) implementsFilter() {}

// Email submission filter criteria
// https://www.rfc-editor.org/rfc/rfc8621.html#section-7.3
type FilterCondition struct {
	IdentityIDs []jmap.ID `json:"identityIds,omitempty"`

	EmailIDs []jmap.ID `json:"emailIds,omitempty"`

	ThreadIDs []jmap.ID `json:"threadIds,omitempty"`

	UndoStatus string `json:"undoStatus,omitempty"`

	Before *time.Time `json:"before,omitempty"`

	After *time.Time `json:"after,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}

func (fc *FilterCondition) MarshalJSON() ([]byte, error) {
	if fc.Before != nil && fc.Before.Location() != time.UTC {
		utc := fc.Before.UTC()
		fc.Before = &utc
	}
	if fc.After != nil && fc.After.Location() != time.UTC {
		utc := fc.After.UTC()
		fc.After = &utc
	}
	// create a type alias to avoid infinite recursion
	type Alias FilterCondition
	return json.Marshal((*Alias)(fc))
}
