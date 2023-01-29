package core

import "git.sr.ht/~rockorager/go-jmap"

const URI string = "urn:ietf:params:jmap:core"

func init() {
	jmap.RegisterCapability(&Core{})
	jmap.RegisterMethod("Core/echo", newEcho)
}

type Core struct {
	// The maximum file size, in octets, that the server will accept for a
	// single file upload (for any purpose).
	MaxSizeUpload uint64 `json:"maxSizeUpload"`

	// The maximum number of concurrent requests the server will accept to the
	// upload endpoint.
	MaxConcurrentUpload uint64 `json:"maxConcurrentUpload"`

	// The maximum size, in octets, that the server will accept for a single
	// request to the API endpoint.
	MaxSizeRequest uint64 `json:"maxSizeRequest"`

	// The maximum number of concurrent requests the server will accept to the
	// API endpoint.
	MaxConcurrentRequests uint64 `json:"maxConcurrentRequests"`

	// The maximum number of method calls the server will accept in a single
	// request to the API endpoint.
	MaxCallsInRequest uint64 `json:"maxCallsInRequest"`

	// The maximum number of objects that the client may request in a single
	// /get type method call.
	MaxObjectsInGet uint64 `json:"maxObjectsInGet"`

	// The maximum number of objects the client may send to create, update or
	// destroy in a single /set type method call. This is the combined total, e.g.
	// if the maximum is 10 you could not create 7 objects and destroy 6, as this
	// would be 13 actions, which exceeds the limit.
	MaxObjectsInSet uint64 `json:"maxObjectsInSet"`

	// A list of identifiers for algorithms registered in the collation
	// registry defined in RFC 4790 that the server supports for sorting
	// when querying records.
	CollationAlgorithms []jmap.CollationAlgo `json:"collationAlgorithms"`
}

func (c *Core) URI() string { return URI }

func (c *Core) New() jmap.Capability { return &Core{} }
