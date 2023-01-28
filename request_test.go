package jmap

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestMarshal(t *testing.T) {
	assert := assert.New(t)
	inv := &Invocation{
		Name: "Core/echo",
		Args: &struct {
			Hello string
		}{
			Hello: "world",
		},
		CallID: "0",
	}
	req := &Request{
		Using:       []string{"urn:ietf:params:jmap:core"},
		Calls: []*Invocation{inv},
	}
	data, err := json.Marshal(req)
	assert.NoError(err)
	expected := `{"using":["urn:ietf:params:jmap:core"],"methodCalls":[["Core/echo",{"Hello":"world"},"0"]]}`
	assert.Equal(expected, string(data))
}
