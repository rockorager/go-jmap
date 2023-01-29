package jmap

import "fmt"

// A RequestError occurs when there is an error with the HTTP request
type RequestError struct {
	// The type of request error, eg "urn:ietf:params:jmap:error:limit"
	Type string `json:"type"`

	// The HTTP status code of the response
	Status int `json:"status"`

	// The description of the error
	Detail string `json:"detail"`

	// If the error is of type ErrLimit, Limit will contain the name of the
	// limit the request would have exceeded
	Limit *string `json:"limit,omitempty"`
}

func (e *RequestError) Error() string {
	if e.Limit != nil {
		return fmt.Sprintf("%s: %s", e.Detail, *e.Limit)
	}
	return fmt.Sprintf(e.Detail)
}

// A MethodError is returned when an error occurred while the server was
// processing a method. Instead of the Response of that method, a MethodError
// invocation will be in it's place
type MethodError struct {
	// The type of error that occurred. Always present
	Type string `json:"type"`
}

func (m *MethodError) Name() string { return "error" }

func (m *MethodError) Uses() string { return "" }

func (m *MethodError) NewResponse() interface{} { return &MethodError{} }
