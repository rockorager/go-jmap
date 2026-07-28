package thread

import "git.sr.ht/~rockorager/go-jmap"

func init() {
	jmap.RegisterMethod("Thread/get", newGetResponse)
	jmap.RegisterMethod("Thread/changes", newChangesResponse)
}

// See RFC8621, Section 3.
type Thread struct {
	ID jmap.ID `json:"id,omitempty"`

	EmailIDs []jmap.ID `json:"emailIds,omitempty"`
}
