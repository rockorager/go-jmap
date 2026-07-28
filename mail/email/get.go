package email

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Get email details
// https://www.rfc-editor.org/rfc/rfc8621.html#section-4.2
type Get struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IDs []jmap.ID `json:"ids,omitempty"`

	Properties []string `json:"properties,omitempty"`

	BodyProperties []string `json:"bodyProperties,omitempty"`

	FetchTextBodyValues bool `json:"fetchTextBodyValues,omitempty"`

	FetchHTMLBodyValues bool `json:"fetchHTMLBodyValues,omitempty"`

	FetchAllBodyValues bool `json:"fetchAllBodyValues,omitempty"`

	MaxBodyValueBytes uint64 `json:"maxBodyValueBytes,omitempty"`

	ReferenceIDs *jmap.ResultReference `json:"#ids,omitempty"`

	ReferenceProperties *jmap.ResultReference `json:"#properties,omitempty"`
}

func (m *Get) Name() string { return "Email/get" }

func (m *Get) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type GetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	State string `json:"state,omitempty"`

	List []*Email `json:"list,omitempty"`

	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
