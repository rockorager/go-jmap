package mail

import "git.sr.ht/~rockorager/go-jmap"

// Replies are grouped together with the original message to form a Thread. In
// JMAP, a Thread is simply a flat list of Emails, ordered by date. Every Email
// MUST belong to a Thread, even if it is the only Email in the Thread.
type Thread struct {
	// The ID of the thread
	ID jmap.ID `json:"id,omitempty"`

	// The ids of the Emails in the Thread, sorted by the receivedAt date
	// of the Email, oldest first. If two Emails have an identical date,
	// the sort is server dependent but MUST be stable (sorting by id is
	// recommended).
	EmailIDs []jmap.ID `json:"emailIds,omitempty"`
}

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
type ThreadGetRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The ids of the Foo objects to return. If null, then all records of
	// the data type are returned, if this is supported for that data type
	// and the number of records does not exceed the maxObjectsInGet limit.
	IDs []jmap.ID `json:"ids,omitempty"`

	// If supplied, only the properties listed in the array are returned
	// for each Foo object. If null, all properties of the object are
	// returned. The id property of the object is always returned, even if
	// not explicitly requested. If an invalid property is requested, the
	// call MUST be rejected with an invalidArguments error.
	Properties []string `json:"properties,omitempty"`
}

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
type ThreadGetResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// A (preferably short) string representing the state on the server for
	// all the data of this type in the account (not just the objects
	// returned in this call). If the data changes, this string MUST
	// change. If the Foo data is unchanged, servers SHOULD return the same
	// state string on subsequent requests for this data type.
	//
	// When a client receives a response with a different state string to a
	// previous call, it MUST either throw away all currently cached
	// objects for the type or call Foo/changes to get the exact changes.
	State string `json:"state,omitempty"`

	// An array of the Foo objects requested. This is the empty array
	// if no objects were found or if the ids argument passed in was also
	// an empty array. The results MAY be in a different order to the ids
	// in the request arguments. If an identical id is included more than
	// once in the request, the server MUST only include it once in either
	// the list or the notFound argument of the response.
	//
	// Each specification must define it's own List property
	List []*Thread `json:"list,omitempty"`

	// This array contains the ids passed to the method for records that do
	// not exist. The array is empty if all requested ids were found or if
	// the ids argument passed in was either null or an empty array.
	NotFound []jmap.ID `json:"notFound,omitempty"`
}

// This is a standard “/changes” method as described in [@!RFC8620], Section 5.2.
type ThreadChangesRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The current state of the client. This is the string that was
	// returned as the state argument in the Foo/get response. The server
	// will return the changes that have occurred since this state.
	SinceState string `json:"sinceState,omitempty"`

	// The maximum number of ids to return in the response. The server MAY
	// choose to return fewer than this value but MUST NOT return more. If
	// not given by the client, the server may choose how many to return.
	// If supplied by the client, the value MUST be a positive integer
	// greater than 0. If a value outside of this range is given, the
	// server MUST reject the call with an invalidArguments error.
	MaxChanges jmap.UnsignedInt `json:"maxChanges,omitempty"`
}

// This is a standard “/changes” method as described in [@!RFC8620], Section 5.2.
type ThreadChangesResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// This is the sinceState argument echoed back; it’s the state from
	// which the server is returning changes.
	OldState string `json:"oldState,omitempty"`

	// This is the state the client will be in after applying the set of
	// changes to the old state.
	NewState string `json:"newState,omitempty"`

	// If true, the client may call Foo/changes again with the newState
	// returned to get further updates. If false, newState is the current
	// server state.
	HasMoreChanges bool `json:"hasMoreChanges,omitempty"`

	// An array of ids for records that have been created since the old
	// state.
	Created []jmap.ID `json:"created,omitempty"`

	// An array of ids for records that have been updated since the old
	// state.
	Updated []jmap.ID `json:"updated,omitempty"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []jmap.ID `json:"destroyed,omitempty"`
}
