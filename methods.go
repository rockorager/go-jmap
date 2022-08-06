package jmap

// The types with suffix Args are intended to satisfy the Invocation.Args
// object These types are the standard types. Individual specifications may add
// to them.

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
//
// Objects of type Foo are fetched via a call to Foo/get The ids
// argument may be null to fetch all at once.
type GetRequestArgs struct {
	// The id of the account to use.
	AccountID ID `json:"accountId"`

	// The ids of the Foo objects to return. If null, then all records of
	// the data type are returned, if this is supported for that data type
	// and the number of records does not exceed the maxObjectsInGet limit.
	IDs []ID `json:"ids"`

	// If supplied, only the properties listed in the array are returned
	// for each Foo object. If null, all properties of the object are
	// returned. The id property of the object is always returned, even if
	// not explicitly requested. If an invalid property is requested, the
	// call MUST be rejected with an invalidArguments error.
	Properties []string `json:"properties"`
}

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
//
// Objects of type Foo are fetched via a call to Foo/get
type GetResponseArgs struct {
	// The id of the account used for the call.
	AccountID ID `json:"accountId"`

	// A (preferably short) string representing the state on the server for
	// all the data of this type in the account (not just the objects
	// returned in this call). If the data changes, this string MUST
	// change. If the Foo data is unchanged, servers SHOULD return the same
	// state string on subsequent requests for this data type.
	//
	// When a client receives a response with a different state string to a
	// previous call, it MUST either throw away all currently cached
	// objects for the type or call Foo/changes to get the exact changes.
	State string `json:"state"`

	// An array of the Foo objects requested. This is the empty array
	// if no objects were found or if the ids argument passed in was also
	// an empty array. The results MAY be in a different order to the ids
	// in the request arguments. If an identical id is included more than
	// once in the request, the server MUST only include it once in either
	// the list or the notFound argument of the response.
	//
	// Each specification must define it's own List property
	List []*interface{} `json:"list"`

	// This array contains the ids passed to the method for records that do
	// not exist. The array is empty if all requested ids were found or if
	// the ids argument passed in was either null or an empty array.
	NotFound []ID `json:"notFound"`
}

// This is a standard “/changes” method as described in [@!RFC8620], Section
// 5.2.
//
// When the state of the set of Foo records in an account changes on the server
// (whether due to creation, updates, or deletion), the state property of the
// Foo/get response will change. The Foo/changes method allows a client to
// efficiently update the state of its Foo cache to match the new state on the
// server.
type ChangesRequestArgs struct {
	// The id of the account to use.
	AccountID ID `json:"accountId"`

	// The current state of the client. This is the string that was
	// returned as the state argument in the Foo/get response. The server
	// will return the changes that have occurred since this state.
	SinceState string `json:"sinceState"`

	// The maximum number of ids to return in the response. The server MAY
	// choose to return fewer than this value but MUST NOT return more. If
	// not given by the client, the server may choose how many to return.
	// If supplied by the client, the value MUST be a positive integer
	// greater than 0. If a value outside of this range is given, the
	// server MUST reject the call with an invalidArguments error.
	MaxChanges UnsignedInt `json:"maxChanges"`
}

// This is a standard “/changes” method as described in [@!RFC8620], Section
// 5.2
type ChangesResponseArgs struct {
	// The id of the account used for the call.
	AccountID ID `json:"accountId"`

	// This is the sinceState argument echoed back; it’s the state from
	// which the server is returning changes.
	OldState string `json:"oldState"`

	// This is the state the client will be in after applying the set of
	// changes to the old state.
	NewState string `json:"newState"`

	// If true, the client may call Foo/changes again with the newState
	// returned to get further updates. If false, newState is the current
	// server state.
	HasMoreChanges bool `json:"hasMoreChanges"`

	// An array of ids for records that have been created since the old
	// state.
	Created []ID `json:"created"`

	// An array of ids for records that have been updated since the old
	// state.
	Updated []ID `json:"updated"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []ID `json:"destroyed"`
}

// Modifying the state of Foo objects on the server is done via the Foo/set
// method. This encompasses creating, updating, and destroying Foo records.
// This allows the server to sort out ordering and dependencies that may exist
// if doing multiple operations at once (for example, to ensure there is always
// a minimum number of a certain record type).
type SetRequestArgs struct {
	// The id of the account to use.
	AccountID ID `json:"accountId"`

	// This is a state string as returned by the Foo/get method
	// (representing the state of all objects of this type in the account).
	// If supplied, the string must match the current state; otherwise, the
	// method will be aborted and a stateMismatch error returned. If null,
	// any changes will be applied to the current state.
	IfInState string `json:"ifInState"`

	// A map of a creation id (a temporary id set by the client) to Foo
	// objects, or null if no objects are to be created.
	//
	// The Foo object type definition may define default values for
	// properties. Any such property may be omitted by the client.
	//
	// The client MUST omit any properties that may only be set by the
	// server (for example, the id property on most object types).
	Create map[ID]interface{} `json:"create"`

	// A map of an id to a Patch object to apply to the current Foo object
	// with that id, or null if no objects are to be updated.
	//
	// A PatchObject is of type String[*] and represents an unordered set
	// of patches. The keys are a path in JSON Pointer Format [@!RFC6901],
	// with an implicit leading “/” (i.e., prefix each key with “/” before
	// applying the JSON Pointer evaluation algorithm).
	//
	// All paths MUST also conform to the following restrictions; if there
	// is any violation, the update MUST be rejected with an invalidPatch
	// error:
	//
	//     The pointer MUST NOT reference inside an array (i.e., you MUST
	//     NOT insert/delete from an array; the array MUST be replaced in
	//     its entirety instead). All parts prior to the last (i.e., the
	//     value after the final slash) MUST already exist on the object
	//     being patched. There MUST NOT be two patches in the PatchObject
	//     where the pointer of one is the prefix of the pointer of the
	//     other, e.g., “alerts/1/offset” and “alerts”.
	//
	// The value associated with each pointer determines how to apply that
	// patch:
	//
	//     If null, set to the default value if specified for this
	//     property; otherwise, remove the property from the patched
	//     object. If the key is not present in the parent, this a no-op.
	//     Anything else: The value to set for this property (this may be a
	//     replacement or addition to the object being patched).
	//
	// Any server-set properties MAY be included in the patch if their
	// value is identical to the current server value (before applying the
	// patches to the object). Otherwise, the update MUST be rejected with
	// an invalidProperties SetError.
	//
	// This patch definition is designed such that an entire Foo object is
	// also a valid PatchObject. The client may choose to optimise network
	// usage by just sending the diff or may send the whole object; the
	// server processes it the same either way.
	Update map[ID]interface{} `json:"update"`

	// A list of ids for Foo objects to permanently delete, or null if no
	// objects are to be destroyed.
	Destroy []ID `json:"destroy"`
}

// Modifying the state of Foo objects on the server is done via the Foo/set
// method. This encompasses creating, updating, and destroying Foo records.
// This allows the server to sort out ordering and dependencies that may exist
// if doing multiple operations at once (for example, to ensure there is always
// a minimum number of a certain record type).
type SetResponseArgs struct {
	// The id of the account used for the call.
	AccountID ID `json:"accountId"`

	// The state string that would have been returned by Foo/get before
	// making the requested changes, or null if the server doesn’t know
	// what the previous state string was.
	OldState string `json:"oldState"`

	// The state string that will now be returned by Foo/get.
	NewState string `json:"newState"`

	// A map of the creation id to an object containing any properties of
	// the created Foo object that were not sent by the client. This
	// includes all server-set properties (such as the id in most object
	// types) and any properties that were omitted by the client and thus
	// set to a default by the server.
	//
	// This argument is null if no Foo objects were successfully created.
	Created map[ID]interface{} `json:"created"`

	// The keys in this map are the ids of all Foos that were successfully
	// updated.
	//
	// The value for each id is a Foo object containing any property that
	// changed in a way not explicitly requested by the PatchObject sent to
	// the server, or null if none. This lets the client know of any
	// changes to server-set or computed properties.
	//
	// This argument is null if no Foo objects were successfully updated.
	Updated map[ID]interface{} `json:"updated"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []ID `json:"destroyed"`
}

// A query on the set of Foos in an account is made by calling Foo/query. This
// takes a number of arguments to determine which records to include, how they
// should be sorted, and which part of the result should be returned (the full
// list may be very long). The result is returned as a list of Foo ids.
type QueryRequestArgs struct {
	// The id of the account to use.
	AccountID ID `json:"accountId"`

	// Determines the set of Foos returned in the results. If null, all
	// objects in the account of this type are included in the results.
	//
	// Each implementation must implement it's own Filter
	// Filter interface{} `json:"filter"`

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
	Position Int `json:"position"`

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
	Anchor ID `json:"anchor"`

	// The index of the first result to return relative to the index of the
	// anchor, if an anchor is given. This MAY be negative. For example, -1
	// means the Foo immediately preceding the anchor is the first result
	// in the list returned (see below for more details).
	AnchorOffset Int `json:"anchorOffset"`

	// The maximum number of results to return. If null, no limit presumed.
	// The server MAY choose to enforce a maximum limit argument. In this
	// case, if a greater value is given (or if it is null), the limit is
	// clamped to the maximum; the new limit is returned with the response
	// so the client is aware. If a negative value is given, the call MUST
	// be rejected with an invalidArguments error.
	Limit UnsignedInt `json:"limit"`

	// Does the client wish to know the total number of results in the
	// query? This may be slow and expensive for servers to calculate,
	// particularly with complex filters, so clients should take care to
	// only request the total when needed.
	CalculateTotal bool `json:"calculateTotal"`
}

type FilterOperator string

const (
	// All of the conditions must match for the filter to match.
	FilterOperatorAND FilterOperator = "AND"

	// At least one of the conditions must match for the filter to match.
	FilterOperatorOR FilterOperator = "OR"

	// None of the conditions must match for the filter to match.
	FilterOperatorNOT FilterOperator = "NOT"
)

// Lists the names of properties to compare between two Foo records, and how to
// compare them, to determine which comes first in the sort. If two Foo records
// have an identical value for the first comparator, the next comparator will
// be considered, and so on. If all comparators are the same (this includes the
// case where an empty array or null is given as the sort argument), the sort
// order is server dependent, but it MUST be stable between calls to Foo/query.
type SortComparator struct {
	// The name of the property on the Foo objects to compare.
	Property string `json:"property"`

	// If true, sort in ascending order. If false, reverse the comparator’s
	// results to sort in descending order.
	IsAscending bool `json:"isAscending"`

	// The identifier, as registered in the collation registry defined in
	// [@!RFC4790], for the algorithm to use when comparing the order of
	// strings. The algorithms the server supports are advertised in the
	// capabilities object returned with the Session object (see Section
	// 2).
	//
	// If omitted, the default algorithm is server-dependent, but:
	//
	//     It MUST be unicode-aware. It MAY be selected based on an
	//     Accept-Language header in the request (as defined in
	//     [@!RFC7231], Section 5.3.5), or out-of-band information about
	//     the user’s language/locale. It SHOULD be case insensitive where
	//     such a concept makes sense for a language/locale. Where the
	//     user’s language is unknown, it is RECOMMENDED to follow the
	//     advice in Section 5.2.3 of [@!RFC8264].
	//
	// The “i;unicode-casemap” collation [@!RFC5051] and the Unicode
	// Collation Algorithm (http://www.unicode.org/reports/tr10/) are two
	// examples that fulfil these criterion and provide reasonable
	// behaviour for a large number of languages.
	//
	// When the property being compared is not a string, the collation
	// property is ignored, and the following comparison rules apply based
	// on the type. In ascending order:
	//
	//     Boolean: false comes before true. Number: A lower number comes
	//     before a higher number. Date/UTCDate: The earlier date comes
	//     first.
	Collation string `json:"collation"`
}

type QueryResponseArgs struct {
	// The id of the account used for the call.
	AccountID ID `json:"accountId"`

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
	QueryState string `json:"queryState"`

	// This is true if the server supports calling Foo/queryChanges with
	// these filter/sort parameters. Note, this does not guarantee that the
	// Foo/queryChanges call will succeed, as it may only be possible for a
	// limited time afterwards due to server internal implementation
	// details.
	CanCalculateChanges bool `json:"canCalculateChanges"`

	// The zero-based index of the first result in the ids array within the
	// complete list of query results.
	Position UnsignedInt `json:"position"`

	// The list of ids for each Foo in the query results, starting at the
	// index given by the position argument of this response and continuing
	// until it hits the end of the results or reaches the limit number of
	// ids. If position is >= total, this MUST be the empty list.
	IDs []ID `json:"ids"`

	// The total number of Foos in the results (given the filter). This
	// argument MUST be omitted if the calculateTotal request argument is
	// not true.
	Total UnsignedInt `json:"total"`

	// The limit enforced by the server on the maximum number of results to
	// return. This is only returned if the server set a limit or used a
	// different limit than that given in the request.
	Limit UnsignedInt `json:"limit"`
}

// The Foo/queryChanges method allows a client to efficiently update the state
// of a cached query to match the new state on the server. It takes the
// following arguments:
type QueryChangesRequestArgs struct {
	// The id of the account to use.
	AccountID ID `json:"accountId"`

	// The filter argument that was used with Foo/query.
	//
	// Each implementation must supply it's own Filter property
	// Filter interface{} `json:"filter"`

	// The sort argument that was used with Foo/query.
	//
	// Each implementation must supply it's own Sort property
	// Sort interface{} `json:"sort"`

	// The current state of the query in the client. This is the string
	// that was returned as the queryState argument in the Foo/query
	// response with the same sort/filter. The server will return the
	// changes made to the query since this state.
	SinceQueryState string `json:"sinceQueryState"`

	// The maximum number of changes to return in the response. See error
	// descriptions below for more details.
	MaxChanges UnsignedInt `json:"maxChanges"`

	// The last (highest-index) id the client currently has cached from the
	// query results. When there are a large number of results, in a common
	// case, the client may have only downloaded and cached a small subset
	// from the beginning of the results. If the sort and filter are both
	// only on immutable properties, this allows the server to omit changes
	// after this point in the results, which can significantly increase
	// efficiency. If they are not immutable, this argument is ignored.
	UpToID ID `json:"upToId"`

	// Does the client wish to know the total number of results now in the
	// query? This may be slow and expensive for servers to calculate,
	// particularly with complex filters, so clients should take care to
	// only request the total when needed.
	CalculateTotal bool `json:"calculateTotal"`
}

// The Foo/queryChanges method allows a client to efficiently update the state
// of a cached query to match the new state on the server. It takes the
// following arguments:
type QueryChangesResponseArgs struct {
	// The id of the account used for the call.
	AccountID ID `json:"accountId"`

	// This is the sinceQueryState argument echoed back; that is, the state
	// from which the server is returning changes.
	OldQueryState string `json:"oldQueryState"`

	// This is the state the query will be in after applying the set of
	// changes to the old state.
	NewQueryState string `json:"newQueryState"`

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
	Removed []ID `json:"removed"`

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
	Added []AddedItem `json:"added"`
}

type AddedItem struct {
	ID    ID          `json:"id"`
	Index UnsignedInt `json:"index"`
}
