package jmap

import "fmt"

type Request struct {
	// The JMAP capabilities the request should use
	Using []URI `json:"using"`

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

	r.Using = mergeURIs(r.Using, m.Requires())
	return i.CallID
}

func mergeURIs(target []URI, opts []URI) []URI {
	m := make(map[URI]bool)
	for _, k := range target {
		m[k] = true
	}
	for _, k := range opts {
		m[k] = true
	}

	uris := []URI{}
	for k := range m {
		uris = append(uris, k)
	}
	return uris
}
