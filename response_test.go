package jmap

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseUnmarshal(t *testing.T) {
	RegisterMethod("Test/method", newTest)
	assert := assert.New(t)
	data := []byte(`{"sessionState": "state","methodResponses":[["Test/method",{"Hello":"world"},"0"]]}`)

	resp := &Response{}
	err := json.Unmarshal(data, resp)
	assert.NoError(err)
	assert.Equal("state", resp.SessionState)
	assert.Equal(1, len(resp.Responses))

	inv := resp.Responses[0]
	assert.Equal("Test/method", inv.Name)
	assert.Equal("0", inv.CallID)

	echo, ok := inv.Args.(*test)
	assert.Truef(ok, "invocation arguments are not type *Echo")
	assert.Equal("world", echo.Hello)
}

func TestResponseMarshal(t *testing.T) {
	assert := assert.New(t)
	resp := &Response{
		SessionState: "state",
		Responses: []*Invocation{
			{
				Name: "Test/method",
				Args: &struct {
					Hello string
				}{
					Hello: "world",
				},
				CallID: "0",
			},
		},
	}
	data, err := json.Marshal(resp)
	assert.NoError(err)
	expected := `{"methodResponses":[["Test/method",{"Hello":"world"},"0"]],"sessionState":"state"}`
	assert.Equal(expected, string(data))
}
