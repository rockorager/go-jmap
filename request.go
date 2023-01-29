package jmap

import "fmt"

type Request struct {
	// The JMAP capabilities the request should use
	Using []string `json:"using"`

	// A slice of methods the server will process. These will be processed
	// sequentially
	Calls []*Invocation `json:"methodCalls"`

	// A map of (client-specified) creation ID to the ID the server assigned
	// when a record was successfully created.
	CreatedIDs map[ID]ID `json:"createdIds,omitempty"`
}

// Invoke a method. Each call to Invoke will add the passed Method to the
// Request. The Requires method will be called and added to the request. The 
// CallID of the Method is returned. CallIDs are assigned as the hex 
// representation of the index of the call, eg "0"
func (r *Request) Invoke(m Method) string {
	i := &Invocation{
		Name:   m.Name(),
		Args:   m,
		CallID: fmt.Sprintf("%x", len(r.Calls)),
	}
	r.Calls = append(r.Calls, i)
	using := false
	for _, uses := range r.Using {
		if uses == m.Requires() {
			using = true
		}
	}
	if !using {
		r.Using = append(r.Using, m.Requires())
	}
	return i.CallID
}
