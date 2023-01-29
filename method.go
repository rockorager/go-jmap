package jmap

import (
	"fmt"
	"sync"
)

// A JMAP method. The method object will be marshaled as the arguments to an
// invocation.
type Method interface {
	// The name of the method, ie "Core/echo"
	Name() string

	// The JMAP capabilities the method uses, ie "urn:ietf:params:jmap:core"
	Uses() string

	// NewResponse returns a new response object for this method, suitable
	// for passing into a json.Unmarshal call
	NewResponse() interface{}
}

// Register a method for use with the library. The method must be registered in
// order for respoonses to be unmarshaled properly. During the unmarshaling of a
// response to a method, the NewResponse method of the Method will be called.
// This object will be passed to json.Unmarshal and ultimately returned as part
// of the response
func RegisterMethods(ms ...Method) {
	for _, m := range ms {
		methods.register(m.Name(), m.NewResponse)
	}
}

var methods = &registry{
	m: make(map[string]func() interface{}),
}

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
