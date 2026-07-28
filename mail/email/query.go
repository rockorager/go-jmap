package email

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Get list of email IDs based on filter and sort criteria
// https://www.rfc-editor.org/rfc/rfc8621.html#section-4.4
type Query struct {
	Account jmap.ID `json:"accountId,omitempty"`

	Filter Filter `json:"filter,omitempty"`

	Sort []*SortComparator `json:"sort,omitempty"`

	Position int64 `json:"position,omitempty"`

	Anchor jmap.ID `json:"anchor,omitempty"`

	AnchorOffset int64 `json:"anchorOffset,omitempty"`

	Limit uint64 `json:"limit,omitempty"`

	CalculateTotal bool `json:"calculateTotal,omitempty"`

	CollapseThreads bool `json:"collapseThreads,omitempty"`
}

func (m *Query) Name() string { return "Email/query" }

func (m *Query) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type QueryResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	QueryState string `json:"queryState,omitempty"`

	CanCalculateChanges bool `json:"canCalculateChanges,omitempty"`

	Position uint64 `json:"position,omitempty"`

	IDs []jmap.ID `json:"ids,omitempty"`

	Total uint64 `json:"total,omitempty"`

	Limit uint64 `json:"limit,omitempty"`
}

func newQueryResponse() jmap.MethodResponse { return &QueryResponse{} }
