package subscription

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/core"
)

// Get the active Push Subscriptions that were created with the same
// authentication credentials used to make the call
type Get struct {
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

func (m *Get) Name() string { return "PushSubscription/get" }

func (m *Get) Requires() string { return core.URI }

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
type GetResponse struct {
	// An array of the Foo objects requested. This is the empty array
	// if no objects were found or if the ids argument passed in was also
	// an empty array. The results MAY be in a different order to the ids
	// in the request arguments. If an identical id is included more than
	// once in the request, the server MUST only include it once in either
	// the list or the notFound argument of the response.
	//
	// Each specification must define it's own List property
	List []*PushSubscription `json:"list,omitempty"`

	// This array contains the ids passed to the method for records that do
	// not exist. The array is empty if all requested ids were found or if
	// the ids argument passed in was either null or an empty array.
	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newGetResponse() interface{} { return &GetResponse{} }
