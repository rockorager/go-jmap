package jmap

import (
	"encoding/json"
	"fmt"
)

// An Invocation represents method calls and responses
type Invocation struct {
	// The name of the method call or response
	Name string
	// Object containing the named arguments for the method or response
	Args interface{}
	// Arbitrary string set by client, echoed back with responses
	CallID string
}

func (i *Invocation) MarshalJSON() ([]byte, error) {
	j := []interface{}{
		i.Name,
		i.Args,
		i.CallID,
	}
	return json.Marshal(j)
}

func (i *Invocation) UnmarshalJSON(data []byte) error {
	raw := []json.RawMessage{}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	if len(raw) != 3 {
		return fmt.Errorf("Not enough values in invocation")
	}
	if err := json.Unmarshal(raw[0], &i.Name); err != nil {
		return err
	}
	newFn, ok := methods[i.Name]
	if !ok {
		return fmt.Errorf("method '%s' not registered", i.Name)
	}
	i.Args = newFn()
	if err := json.Unmarshal(raw[1], i.Args); err != nil {
		return err
	}
	if err := json.Unmarshal(raw[2], &i.CallID); err != nil {
		return err
	}
	return nil
}
