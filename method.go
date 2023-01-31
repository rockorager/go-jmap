package jmap

// A JMAP method. The method object will be marshaled as the arguments to an
// invocation.
type Method interface {
	// The name of the method, ie "Core/echo"
	Name() string

	// The JMAP capabilities required for the method, ie "urn:ietf:params:jmap:core"
	Requires() string
}

// A response to a method call
type MethodResponse interface{}

// A Factory function which produces a new MethodResponse object
type MethodResponseFactory func() MethodResponse

// Registered method results
var methods = map[string]MethodResponseFactory{}

// Register a method. The Name parameter will be used when unmarshalling
// responses to call the responseConstructor, which should generate a pointer to
// an empty Response object of that method. This object will be returned in the
// result set (unless there is an error)
func RegisterMethod(name string, factory MethodResponseFactory) {
	methods[name] = factory
}
