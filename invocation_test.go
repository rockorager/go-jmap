package jmap

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvocationMarshal(t *testing.T) {
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
	data, err := json.Marshal(inv)
	assert.NoError(err)
	if err != nil {
		t.Logf("expected: no error")
		t.Logf("actual: %v", err)
		t.FailNow()
	}
	expected := `["Core/echo",{"Hello":"world"},"0"]`
	assert.Equal(expected, string(data))
}

func TestInvocationUnmarshal(t *testing.T) {
	assert := assert.New(t)
	raw := []byte(`["Core/echo",{"Hello":"world"},"0"]`)
	inv := &Invocation{}
	err := json.Unmarshal(raw, inv)
	assert.NoError(err)
	assert.Equal("Core/echo", inv.Name)
	assert.Equal("0", inv.CallID)

	echo, ok := inv.Args.(*Echo)
	assert.Truef(ok, "invocation arguments are not type *Echo")
	assert.Equal("world", echo.Hello)
}
