package mailbox

import "git.sr.ht/~rockorager/go-jmap"

type SortComparator struct {
	// The name of the property on the Mailbox objects to compare.
	Property string `json:"property,omitempty"`

	// If true, sort in ascending order.
	IsAscending bool `json:"isAscending"`

	// The identifier, as registered in the collation registry defined in
	// RFC4790, for the algorithm to use when comparing the order of
	// strings.
	Collation jmap.CollationAlgo `json:"collation,omitempty"`
}
