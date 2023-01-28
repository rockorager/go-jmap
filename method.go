package jmap

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
