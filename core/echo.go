package core

// The Core/echo method
type Echo struct {
	Hello string
}

func (e Echo) Name() string {
	return "Core/echo"
}

func (e Echo) Requires() string {
	return URI
}

func newEcho() interface{} {
	return &Echo{}
}
