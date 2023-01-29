// Package jmap implements JMAP Core protocol as defined in
// draft-ietf-jmap-core-17 (published March 2019).
//
// Documentation strings for most of the protocol objects are taken from (or
// based on) contents of draft-ietf-jmap-core-17 and is subject to the IETF
// Trust Provisions.
// See https://trustee.ietf.org/trust-legal-provisions.html for details.
// See included draft-ietf-jmap-core-17.txt for related copyright notices.
package jmap

import (
	"encoding/json"
	"time"
)

func init() {
	RegisterMethods(
		&BlobCopy{},
		&Echo{},
		&MethodError{},
	)
	RegisterMethods()
}

// Patch represents a patch which can be used in a set.Update call.
// All paths MUST also conform to the following restrictions; if there is any
// violation, the update MUST be rejected with an invalidPatch error:
//
//	The pointer MUST NOT reference inside an array (i.e., you MUST NOT
//	insert/delete from an array; the array MUST be replaced in its entirety
//	instead). All parts prior to the last (i.e., the value after the final
//	slash) MUST already exist on the object being patched. There MUST NOT be
//	two patches in the Patch where the pointer of one is the prefix of
//	the pointer of the other, e.g., “alerts/1/offset” and “alerts”.
type Patch struct {
	Path  string
	Value interface{}
}

func (p *Patch) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m[p.Path] = p.Value
	return json.Marshal(m)
}

func (p *Patch) UnmarshalJSON(data []byte) error {
	var patch map[string]interface{}
	err := json.Unmarshal(data, &patch)
	if err != nil {
		return err
	}
	for k, v := range patch {
		p.Path = k
		p.Value = v
	}
	return nil
}

type FilterOperator string

const (
	// All of the conditions must match for the filter to match.
	FilterOperatorAND FilterOperator = "AND"

	// At least one of the conditions must match for the filter to match.
	FilterOperatorOR FilterOperator = "OR"

	// None of the conditions must match for the filter to match.
	FilterOperatorNOT FilterOperator = "NOT"
)

// The id and index in the query results (in the new state) for every Foo that
// has been added to the results since the old state AND every Foo in the
// current results that was included in the removed array (due to a filter or
// sort based upon a mutable property).
//
// If the sort and filter are both only on immutable properties and an upToId
// is supplied and exists in the results, any ids that were added but have a
// higher index than upToId SHOULD be omitted.
//
// The array MUST be sorted in order of index, with the lowest index first.
type AddedItem struct {
	ID    string `json:"id"`
	Index uint64 `json:"index"`
}

// To allow clients to make more efficient use of the network and avoid round
// trips, an argument to one method can be taken from the result of a previous
// method call in the same request.
//
// To do this, the client prefixes the argument name with # (an octothorpe). The
// value is a ResultReference object as described below. When processing a
// method call, the server MUST first check the arguments object for any names
// beginning with #. If found, the result reference should be resolved and the
// value used as the “real” argument. The method is then processed as normal. If
// any result reference fails to resolve, the whole method MUST be rejected with
// an invalidResultReference error. If an arguments object contains the same
// argument name in normal and referenced form (e.g., foo and #foo), the method
// MUST return an invalidArguments error.
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

// Date is a time.Time that is serialized to JSON in RFC 3339 format (without
// the fractional part).
type Date time.Time

func (d Date) MarshalText() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339))
	b = time.Time(d).AppendFormat(b, time.RFC3339)
	return b, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*(*time.Time)(d), err = time.Parse(time.RFC3339, s)
	if err != nil {
		*(*time.Time)(d) = time.Time{}
	}
	return nil
}

// Date is a time.Time that is serialized to JSON in RFC 3339 format (without
// the fractional part) in UTC timezone.
//
// If UTCDate value is not in UTC, it will be converted to UTC during
// serialization.
type UTCDate time.Time

func (d UTCDate) MarshalText() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = time.Time(d).UTC().AppendFormat(b, time.RFC3339)
	return b, nil
}

func (d *UTCDate) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*(*time.Time)(d), err = time.ParseInLocation(time.RFC3339, s, time.UTC)
	return err
}
