package email

import "git.sr.ht/~rockorager/go-jmap"

type SortComparator struct {
	// The name of the property on the Foo objects to compare. Servers MUST
	// support sorting by the following properties:
	// - receivedAt
	//
	// Additional supported properties are reported in the mail capability
	// object
	Property string `json:"property,omitempty"`

	// When specifying a "hasKeyword", "allInThreadHaveKeyword", or
	// "someInThreadHaveKeyword" sort, the Comparator object MUST also have
	// a "keyword" property.
	Keyword string `json:"keyword,omitempty"`

	// If true, sort in ascending order. If false, reverse the comparator’s
	// results to sort in descending order.
	IsAscending bool `json:"isAscending,omitempty"`

	// The identifier, as registered in the collation registry defined in
	// [@!RFC4790], for the algorithm to use when comparing the order of
	// strings. The algorithms the server supports are advertised in the
	// capabilities object returned with the Session object (see Section
	// 2).
	//
	// If omitted, the default algorithm is server-dependent, but:
	//
	//     It MUST be unicode-aware. It MAY be selected based on an
	//     Accept-Language header in the request (as defined in
	//     [@!RFC7231], Section 5.3.5), or out-of-band information about
	//     the user’s language/locale. It SHOULD be case insensitive where
	//     such a concept makes sense for a language/locale. Where the
	//     user’s language is unknown, it is RECOMMENDED to follow the
	//     advice in Section 5.2.3 of [@!RFC8264].
	//
	// The “i;unicode-casemap” collation [@!RFC5051] and the Unicode
	// Collation Algorithm (http://www.unicode.org/reports/tr10/) are two
	// examples that fulfil these criterion and provide reasonable
	// behaviour for a large number of languages.
	//
	// When the property being compared is not a string, the collation
	// property is ignored, and the following comparison rules apply based
	// on the type. In ascending order:
	//
	//     Boolean: false comes before true. Number: A lower number comes
	//     before a higher number. Date/UTCDate: The earlier date comes
	//     first.
	Collation jmap.CollationAlgo `json:"collation,omitempty"`
}
