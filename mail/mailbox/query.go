package mailbox

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// This is a standard “/query” method as described in [@!RFC8620], Section 5.5,
// but with the following additional request argument: sortAsTree, filterAsTree
type Query struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId,omitempty"`

	// Determines the set of Foos returned in the results. If null, all
	// objects in the account of this type are included in the results.
	//
	// Each implementation must implement it's own Filter
	Filter Filter `json:"filter,omitempty"`

	// Lists the names of properties to compare between two Foo records,
	// and how to compare them, to determine which comes first in the sort.
	// If two Foo records have an identical value for the first comparator,
	// the next comparator will be considered, and so on. If all
	// comparators are the same (this includes the case where an empty
	// array or null is given as the sort argument), the sort order is
	// server dependent, but it MUST be stable between calls to Foo/query.
	//
	// Each implementation must define it's own Sort property. The
	// SortComparator object can be used as a basis
	Sort []*SortComparator `json:"sort,omitempty"`

	// The zero-based index of the first id in the full list of results to
	// return.
	//
	// If a negative value is given, it is an offset from the end of the
	// list. Specifically, the negative value MUST be added to the total
	// number of results given the filter, and if still negative, it’s
	// clamped to 0. This is now the zero-based index of the first id to
	// return.
	//
	// If the index is greater than or equal to the total number of objects
	// in the results list, then the ids array in the response will be
	// empty, but this is not an error.
	Position int64 `json:"position,omitempty"`

	// A Foo id. If supplied, the position argument is ignored. The index
	// of this id in the results will be used in combination with the
	// anchorOffset argument to determine the index of the first result to
	// return (see below for more details).
	//
	// If an anchor argument is given, the anchor is looked for in the
	// results after filtering and sorting. If found, the anchorOffset is
	// then added to its index. If the resulting index is now negative, it
	// is clamped to 0. This index is now used exactly as though it were
	// supplied as the position argument. If the anchor is not found, the
	// call is rejected with an anchorNotFound error.
	//
	// If an anchor is specified, any position argument supplied by the
	// client MUST be ignored. If no anchor is supplied, any anchorOffset
	// argument MUST be ignored.
	//
	// A client can use anchor instead of position to find the index of an
	// id within a large set of results.
	Anchor jmap.ID `json:"anchor,omitempty"`

	// The index of the first result to return relative to the index of the
	// anchor, if an anchor is given. This MAY be negative. For example, -1
	// means the Foo immediately preceding the anchor is the first result
	// in the list returned (see below for more details).
	AnchorOffset int64 `json:"anchorOffset,omitempty"`

	// The maximum number of results to return. If null, no limit presumed.
	// The server MAY choose to enforce a maximum limit argument. In this
	// case, if a greater value is given (or if it is null), the limit is
	// clamped to the maximum; the new limit is returned with the response
	// so the client is aware. If a negative value is given, the call MUST
	// be rejected with an invalidArguments error.
	Limit uint64 `json:"limit,omitempty"`

	// Does the client wish to know the total number of results in the
	// query? This may be slow and expensive for servers to calculate,
	// particularly with complex filters, so clients should take care to
	// only request the total when needed.
	CalculateTotal bool `json:"calculateTotal,omitempty"`

	// If true, when sorting the query results and comparing Mailboxes A
	// and B:
	//
	// - If A is an ancestor of B, it always comes first regardless of the
	// sort comparators. Similarly, if A is descendant of B, then B always
	// comes first.
	//
	// - Otherwise, if A and B do not share a parentId, find the nearest
	// ancestors of each that do have the same parentId and compare the
	// sort properties on those Mailboxes instead.
	//
	// The result of this is that the Mailboxes are sorted as a tree
	// according to the parentId properties, with each set of children with
	// a common parent sorted according to the standard sort comparators.
	SortAsTree bool `json:"sortAsTree,omitempty"`

	// If true, a Mailbox is only included in the query if all its
	// ancestors are also included in the query according to the filter.
	FilterAsTree bool `json:"filterAsTree,omitempty"`
}

func (m *Query) Name() string { return "Mailbox/query" }

func (m *Query) Requires() string { return mail.URI }

type QueryResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// A string encoding the current state of the query on the server. This
	// string MUST change if the results of the query (i.e., the matching
	// ids and their sort order) have changed. The queryState string MAY
	// change if something has changed on the server, which means the
	// results may have changed but the server doesn’t know for sure.
	//
	// The queryState string only represents the ordered list of ids that
	// match the particular query (including its sort/filter). There is no
	// requirement for it to change if a property on an object matching the
	// query changes but the query results are unaffected (indeed, it is
	// more efficient if the queryState string does not change in this
	// case). The queryState string only has meaning when compared to
	// future responses to a query with the same type/sort/filter or when
	// used with /queryChanges to fetch changes.
	//
	// Should a client receive back a response with a different queryState
	// string to a previous call, it MUST either throw away the currently
	// cached query and fetch it again (note, this does not require
	// fetching the records again, just the list of ids) or call
	// Foo/queryChanges to get the difference.
	QueryState string `json:"queryState,omitempty"`

	// This is true if the server supports calling Foo/queryChanges with
	// these filter/sort parameters. Note, this does not guarantee that the
	// Foo/queryChanges call will succeed, as it may only be possible for a
	// limited time afterwards due to server internal implementation
	// details.
	CanCalculateChanges bool `json:"canCalculateChanges,omitempty"`

	// The zero-based index of the first result in the ids array within the
	// complete list of query results.
	Position uint64 `json:"position,omitempty"`

	// The list of ids for each Foo in the query results, starting at the
	// index given by the position argument of this response and continuing
	// until it hits the end of the results or reaches the limit number of
	// ids. If position is >= total, this MUST be the empty list.
	IDs []jmap.ID `json:"ids,omitempty"`

	// The total number of Foos in the results (given the filter). This
	// argument MUST be omitted if the calculateTotal request argument is
	// not true.
	Total int64 `json:"total,omitempty"`

	// The limit enforced by the server on the maximum number of results to
	// return. This is only returned if the server set a limit or used a
	// different limit than that given in the request.
	Limit uint64 `json:"limit,omitempty"`
}

func newQueryResponse() jmap.MethodResponse { return &QueryResponse{} }
