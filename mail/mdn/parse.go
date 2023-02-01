package mdn

import (
	"git.sr.ht/~rockorager/go-jmap"
)

// Sends an RFC5322 message from an MDN object
type Parse struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId,omitempty"`

	// The IDs of blobs to parse as MDNs
	BlobIDs []jmap.ID `json:"blobIds,omitempty"`
}

func (m *Parse) Name() string { return "MDN/parse" }

func (m *Parse) Requires() []jmap.URI { return []jmap.URI{URI} }

type ParseResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// A map of the blob ID to the MDN resulting from the parse
	Parsed map[jmap.ID]*MDN `json:"parsed,omitempty"`

	// A list blob IDs that could not be parsed as MDNs
	NotParsable []jmap.ID `json:"notParsable,omitempty"`

	// A list of blob IDs that couldn't be found
	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newParseResponse() jmap.MethodResponse { return &ParseResponse{} }
