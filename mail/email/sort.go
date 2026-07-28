package email

import "git.sr.ht/~rockorager/go-jmap"

// Email sort criteria
// https://www.rfc-editor.org/rfc/rfc8621.html#section-4.4.2
type SortComparator struct {
	Property string `json:"property,omitempty"`

	Keyword string `json:"keyword,omitempty"`

	IsAscending bool `json:"isAscending"`

	Collation jmap.CollationAlgo `json:"collation,omitempty"`
}
