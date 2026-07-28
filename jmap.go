// Package jmap implements JMAP Core protocol
// as defined in RFC 8620 published on July 2019.
package jmap

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func init() {
	RegisterMethod("error", newMethodError)
}

// URI is an identifier of a capability, eg "urn:ietf:params:jmap:core"
type URI string

// ID is a unique identifier assigned by the server
type ID string

var idRegexp = regexp.MustCompile(`^[A-Za-z0-9\-_]+$`)

// Valid checks to make sure that the given ID is valid according to the
// specification.
func (id ID) Valid() (bool, error) {
	if len(string(id)) < 1 {
		return false, fmt.Errorf("invalid ID: too short")
	}
	if len(string(id)) > 255 {
		return false, fmt.Errorf("invalid ID: too long")
	}

	return true, nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	if _, err := id.Valid(); err != nil {
		return nil, err
	}
	return json.Marshal(string(id))
}

// Patch is a JMAP patch object which can be used in set.Update calls. The keys
// are json pointer paths, and the value is the value to set the path to.
type Patch map[string]interface{}

// Operator is used when constructing FilterOperator. It MUST be "AND", "OR", or
// "NOT"
type Operator string

const (
	// All of the conditions must match for the filter to match.
	OperatorAND Operator = "AND"

	// At least one of the conditions must match for the filter to match.
	OperatorOR Operator = "OR"

	// None of the conditions must match for the filter to match.
	OperatorNOT Operator = "NOT"
)

// AddedItem is an item that has been added to the results of a Query
type AddedItem struct {
	ID    ID     `json:"id"`
	Index uint64 `json:"index"`
}

// ResultReference is a reference to a previous Invocations' result
type ResultReference struct {
	// The method call id (see Section 3.1.1) of a previous method call in
	// the current request.
	ResultOf string `json:"resultOf"`

	// The required name of a response to that method call.
	Name string `json:"name"`

	// A pointer into the arguments of the response selected via the name
	// and resultOf properties. This is a JSON Pointer [@!RFC6901], except
	// it also allows the use of * to map through an array (see the
	// description below).
	Path string `json:"path"`
}

type CollationAlgo string

const (
	// Defined in RFC 4790.
	ASCIINumeric CollationAlgo = "i;ascii-numeric"

	// Defined in RFC 4790.
	ASCIICasemap = "i;ascii-casemap"

	// Defined in RFC 5051.
	UnicodeCasemap = "i;unicode-casemap"
)
