package core

import "git.sr.ht/~rockorager/go-jmap"

func init() {
	jmap.RegisterCapability(&Core{})
	jmap.RegisterMethod("Core/echo", newEcho)
}

type Core struct {
	// The maximum file size, in bytes
	MaxSizeUpload       uint64 `json:"maxSizeUpload"`
	MaxConcurrentUpload uint64 `json:"maxConcurrentUpload"`

	// The maximum size, in bytes, that the server will accept for a request
	MaxSizeRequest        uint64 `json:"maxSizeRequest"`
	MaxConcurrentRequests uint64 `json:"maxConcurrentRequests"`
	MaxCallsInRequest     uint64 `json:"maxCallsInRequest"`

	// The maximum number of objects that the client may request in a single
	// /get type method call.
	MaxObjectsInGet     uint64               `json:"maxObjectsInGet"`
	MaxObjectsInSet     uint64               `json:"maxObjectsInSet"`
	CollationAlgorithms []jmap.CollationAlgo `json:"collationAlgorithms"`
}

func (c *Core) URI() jmap.URI { return jmap.CoreURI }

func (c *Core) New() jmap.Capability { return &Core{} }
