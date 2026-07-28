package mdn

import (
	"git.sr.ht/~rockorager/go-jmap"
)

// Parse blobs as messages in the style of RFC5322 to get MDN objects
// https://www.rfc-editor.org/rfc/rfc9007.html#section-2.2
type Parse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	BlobIDs []jmap.ID `json:"blobIds,omitempty"`
}

func (m *Parse) Name() string { return "MDN/parse" }

func (m *Parse) Requires() []jmap.URI { return []jmap.URI{URI} }

type ParseResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	Parsed map[jmap.ID]*MDN `json:"parsed,omitempty"`

	NotParsable []jmap.ID `json:"notParsable,omitempty"`

	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newParseResponse() jmap.MethodResponse { return &ParseResponse{} }
