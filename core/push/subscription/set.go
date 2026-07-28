package subscription

import (
	"git.sr.ht/~rockorager/go-jmap"
)

// Modify push subscription details
// https://www.rfc-editor.org/rfc/rfc8620.html#section-7.2.2
type Set struct {
	Create  map[jmap.ID]*PushSubscription `json:"create,omitempty"`
	Update  map[jmap.ID]*jmap.Patch       `json:"update,omitempty"`
	Destroy []jmap.ID                     `json:"destroy,omitempty"`
}

func (m *Set) Name() string { return "PushSubscription/set" }

func (m *Set) Requires() []jmap.URI { return nil }

type SetResponse struct {
	Created      map[jmap.ID]*PushSubscription `json:"created,omitempty"`
	Updated      map[jmap.ID]*PushSubscription `json:"updated,omitempty"`
	Destroyed    []jmap.ID                     `json:"destroyed,omitempty"`
	NotCreated   map[jmap.ID]*jmap.SetError    `json:"notCreated,omitempty"`
	NotUpdated   map[jmap.ID]*jmap.SetError    `json:"notUpdated,omitempty"`
	NotDestroyed map[jmap.ID]*jmap.SetError    `json:"notDestroyed,omitempty"`
}

func newSetResponse() jmap.MethodResponse { return &SetResponse{} }
