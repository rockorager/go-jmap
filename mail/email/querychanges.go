package email

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// This is a standard "/queryChanges" method as described in [RFC8620], Section
// 5.6 with the following additional request argument: collapseThreads
type QueryChanges struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId,omitempty"`

	// The filter argument that was used with Foo/query.
	//
	// Each implementation must supply it's own Filter property
	Filter Filter `json:"filter,omitempty"`

	// The sort argument that was used with Foo/query.
	//
	// Each implementation must supply it's own Sort property
	Sort []*SortComparator `json:"sort,omitempty"`

	// The current state of the query in the client. This is the string
	// that was returned as the queryState argument in the Foo/query
	// response with the same sort/filter. The server will return the
	// changes made to the query since this state.
	SinceQueryState string `json:"sinceQueryState,omitempty"`

	// The maximum number of changes to return in the response. See error
	// descriptions below for more details.
	MaxChanges uint64 `json:"maxChanges,omitempty"`

	// The last (highest-index) id the client currently has cached from the
	// query results. When there are a large number of results, in a common
	// case, the client may have only downloaded and cached a small subset
	// from the beginning of the results. If the sort and filter are both
	// only on immutable properties, this allows the server to omit changes
	// after this point in the results, which can significantly increase
	// efficiency. If they are not immutable, this argument is ignored.
	UpToID jmap.ID `json:"upToId,omitempty"`

	// Does the client wish to know the total number of results now in the
	// query? This may be slow and expensive for servers to calculate,
	// particularly with complex filters, so clients should take care to
	// only request the total when needed.
	CalculateTotal bool `json:"calculateTotal,omitempty"`

	// If true, Emails in the same Thread as a previous Email in the list
	// (given the filter and sort order) will be removed from the list.
	// This means only one Email at most will be included in the list for
	// any given Thread.
	CollapseThreads bool `json:"collapseThreads,omitempty"`
}

func (m *QueryChanges) Name() string { return "Email/queryChanges" }

func (m *QueryChanges) Requires() string { return mail.URI }

// This is a standard "/queryChanges" method as described in [RFC8620], Section
// 5.6
type QueryChangesResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// This is the sinceQueryState argument echoed back; that is, the state
	// from which the server is returning changes.
	OldQueryState string `json:"oldQueryState,omitempty"`

	// This is the state the query will be in after applying the set of
	// changes to the old state.
	NewQueryState string `json:"newQueryState,omitempty"`

	// The id for every Foo that was in the query results in the old state
	// and that is not in the results in the new state.
	//
	// If the server cannot calculate this exactly, the server MAY return
	// the ids of extra Foos in addition that may have been in the old
	// results but are not in the new results.
	//
	// If the sort and filter are both only on immutable properties and an
	// upToId is supplied and exists in the results, any ids that were
	// removed but have a higher index than upToId SHOULD be omitted.
	//
	// If the filter or sort includes a mutable property, the server MUST
	// include all Foos in the current results for which this property may
	// have changed. The position of these may have moved in the results,
	// so must be reinserted by the client to ensure its query cache is
	// correct.
	Removed []jmap.ID `json:"removed,omitempty"`

	// The id and index in the query results (in the new state) for every
	// Foo that has been added to the results since the old state AND every
	// Foo in the current results that was included in the removed array
	// (due to a filter or sort based upon a mutable property).
	//
	// If the sort and filter are both only on immutable properties and an
	// upToId is supplied and exists in the results, any ids that were
	// added but have a higher index than upToId SHOULD be omitted.
	//
	// The array MUST be sorted in order of index, with the lowest index
	// first.
	Added []jmap.AddedItem `json:"added,omitempty"`
}

func newQueryChangesResponse() interface{} { return &QueryChangesResponse{} }
