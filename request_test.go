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
		Using: []URI{"urn:ietf:params:jmap:core"},
		Calls: []*Invocation{inv},
	}
	data, err := json.Marshal(req)
	assert.NoError(err)
	expected := `{"using":["urn:ietf:params:jmap:core"],"methodCalls":[["Core/echo",{"Hello":"world"},"0"]]}`
	assert.Equal(expected, string(data))
}

func TestMergeURIs(t *testing.T) {
	assert := assert.New(t)
	target := []URI{"one", "two", "three"}
	opts := []URI{"one", "four"}

	res := mergeURIs(target, opts)
	assert.Equal(4, len(res))
	assert.Equal([]URI{"one", "two", "three", "four"}, res)
}
