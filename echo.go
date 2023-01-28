package jmap

// The Core/echo method
type Echo struct {
	Hello string
}

func (e *Echo) Name() string {
	return "Core/echo"
}

func (e *Echo) Uses() string {
	return CoreCapabilityName
}

// An echo response is a mirror of the request
func (e *Echo) NewResponse() interface{} {
	return &Echo{}
}
