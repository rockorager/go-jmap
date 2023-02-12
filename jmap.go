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
	"fmt"
	"regexp"
)

func init() {
	RegisterMethod("error", newMethodError)
}

// URI is an identifier of a capability, eg "urn:ietf:params:jmap:core"
type URI string

// ID is a unique identifier assigned by the server. The character set must
// contain only ASCII alphanumerics, hyphen, or underscore and the ID must be
// between 1 and 255 octets long.
type ID string

var idRegexp = regexp.MustCompile(`^[A-Za-z0-9\-_]+$`)

func (id ID) MarshalJSON() ([]byte, error) {
	if len(string(id)) < 1 {
		return nil, fmt.Errorf("invalid ID: too short")
	}
	if len(string(id)) > 255 {
		return nil, fmt.Errorf("invalid ID: too long")
	}
	if !idRegexp.MatchString(string(id)) {
		return nil, fmt.Errorf("invalid ID: invalid characters")
	}
	return json.Marshal(string(id))
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
//
// The keys are a JSON pointer path, and the value is the value to set the path
// to
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
	ID    ID     `json:"id"`
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

type CollationAlgo string

const (
	// The ASCIINumeric collation is a simple collation intended for use
	// with arbitrary sized unsigned decimal integer numbers stored as octet
	// strings. US-ASCII digits (0x30 to 0x39) represent digits of the numbers.
	// Before converting from string to integer, the input string is truncated
	// at the first non-digit character. All input is valid; strings which do
	// not start with a digit represent positive infinity.
	//
	// Defined in RFC 4790.
	ASCIINumeric CollationAlgo = "i;ascii-numeric"

	// The ASCIICasemap collation is a simple collation which operates on
	// octet strings and treats US-ASCII letters case-insensitively. It provides
	// equality, substring and ordering operations. All input is valid. Note that
	// letters outside ASCII are not treated case- insensitively.
	//
	// Defined in RFC 4790.
	ASCIICasemap = "i;ascii-casemap"

	// The "i;unicode-casemap" collation is a simple collation which is
	// case-insensitive in its treatment of characters. It provides equality,
	// substring, and ordering operations. The validity test operation returns "valid"
	// for any input.
	//
	// This collation allows strings in arbitrary (and mixed) character sets,
	// as long as the character set for each string is identified and it is
	// possible to convert the string to Unicode. Strings which have an
	// unidentified character set and/or cannot be converted to Unicode are not
	// rejected, but are treated as binary.
	//
	// Defined in RFC 5051.
	UnicodeCasemap = "i;unicode-casemap"

	// Octet collation is left out intentionally: "Protocols that want to make
	// this collation available have to do so by explicitly allowing it. If not
	// explicitly allowed, it MUST NOT be used."
)
