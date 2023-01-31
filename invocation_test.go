package jmap

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	Hello string
}

func newTest() MethodResponse {
	return &test{}
}

func TestInvocationMarshal(t *testing.T) {
	RegisterMethod("Test/method", newTest)
	assert := assert.New(t)
	inv := &Invocation{
		Name: "Test/method",
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
	expected := `["Test/method",{"Hello":"world"},"0"]`
	assert.Equal(expected, string(data))
}

func TestInvocationUnmarshal(t *testing.T) {
	RegisterMethod("Test/method", newTest)
	assert := assert.New(t)
	raw := []byte(`["Test/method",{"Hello":"world"},"0"]`)
	inv := &Invocation{}
	err := json.Unmarshal(raw, inv)
	assert.NoError(err)
	assert.Equal("Test/method", inv.Name)
	assert.Equal("0", inv.CallID)

	test, ok := inv.Args.(*test)
	t.Logf("TIMBUG %T", inv.Args)
	assert.Truef(ok, "invocation arguments are not type test")
	assert.Equal("world", test.Hello)
}
