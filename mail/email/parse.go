package email

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Parse binary blobs as RFC5322 messages
// https://www.rfc-editor.org/rfc/rfc8621.html#section-4.9
type Parse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	BlobIDs []jmap.ID `json:"blobIds,omitempty"`

	Properties []string `json:"properties,omitempty"`

	BodyProperties []string `json:"bodyProperties,omitempty"`

	FetchTextBodyValues bool `json:"fetchTextBodyValues,omitempty"`

	FetchHTMLBodyValues bool `json:"fetchHTMLBodyValues,omitempty"`

	FetchAllBodyValues bool `json:"fetchAllBodyValues,omitempty"`

	MaxBodyValueBytes uint64 `json:"maxBodyValueBytes,omitempty"`
}

func (m *Parse) Name() string { return "Email/parse" }

func (m *Parse) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type ParseResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	Parsed map[jmap.ID]*Email `json:"parsed,omitempty"`

	NotParsable []jmap.ID `json:"notParsable,omitempty"`

	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newParseResponse() jmap.MethodResponse { return &ParseResponse{} }
