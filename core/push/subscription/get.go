package subscription

import (
	"git.sr.ht/~rockorager/go-jmap"
)

// Get push subscription details
// https://www.rfc-editor.org/rfc/rfc8620.html#section-7.2.1
type Get struct {
	IDs        []jmap.ID `json:"ids,omitempty"`
	Properties []string  `json:"properties,omitempty"`
}

func (m *Get) Name() string { return "PushSubscription/get" }

func (m *Get) Requires() []jmap.URI { return nil }

type GetResponse struct {
	List     []*PushSubscription `json:"list,omitempty"`
	NotFound []jmap.ID           `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
