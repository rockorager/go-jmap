package identity

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail/emailsubmission"
)

// Get details identity details
// https://www.rfc-editor.org/rfc/rfc8621.html#section-6.1
type Get struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IDs []jmap.ID `json:"ids,omitempty"`

	Properties []string `json:"properties,omitempty"`

	ReferenceIDs *jmap.ResultReference `json:"#ids,omitempty"`

	ReferenceProperties *jmap.ResultReference `json:"#properties,omitempty"`
}

func (m *Get) Name() string { return "Identity/get" }

func (m *Get) Requires() []jmap.URI { return []jmap.URI{emailsubmission.URI} }

type GetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	State string `json:"state,omitempty"`

	List []*Identity `json:"list,omitempty"`

	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
