package jmap

import (
	"fmt"
	"sync"
)

// A registry of a resource name to a constructor for a resource object. The
// resource can be a capability, a response, etc.
type registry struct {
	sync.Mutex
	// map of resource name to resource constructor
	m map[string]func() interface{}
}

// register a resource
func (r *registry) register(name string, constructor func() interface{}) {
	r.Lock()
	r.m[name] = constructor
	r.Unlock()
}

func (r *registry) newObject(name string) (interface{}, error) {
	r.Lock()
	defer r.Unlock()
	resp, ok := r.m[name]
	if !ok {
		return nil, fmt.Errorf("unknown resource: '%s'", name)
	}
	return resp(), nil
}
