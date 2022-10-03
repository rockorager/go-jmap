package mail

import "git.sr.ht/~rockorager/go-jmap"

// When doing a search on a "String" property, the client may wish to
// show the relevant section of the body that matches the search as a
// preview and to highlight any matching terms in both this and the
// subject of the Email.  Search snippets represent this data.
type SearchSnippet struct {
	// The Email id the snippet applies to.
	EmailID jmap.ID `json:"emailId,omitempty"`

	// If text from the filter matches the subject, this is the subject
	// of the Email with the following transformations:
	//
	// 1.  Any instance of the following three characters MUST be
	//     replaced by an appropriate HTML entity: & (ampersand), <
	//     (less-than sign), and > (greater-than sign) [HTML].  Other
	//     characters MAY also be replaced with an HTML entity form.
	//
	// 2.  The matching words/phrases from the filter are wrapped in HTML
	//     "<mark></mark>" tags.
	//
	// If the subject does not match text from the filter, this property
	// is null.
	Subject string `json:"subject,omitempty"`

	// If text from the filter matches the plaintext or HTML body, this is
	// the relevant section of the body (converted to plaintext if
	// originally HTML), with the same transformations as the "subject"
	// property.  It MUST NOT be bigger than 255 octets in size.  If the
	// body does not contain a match for the text from the filter, this
	// property is null.
	Preview string `json:"preview,omitempty"`
}

type SearchSnippetGetRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// Determines the set of Foos returned in the results. If null, all
	// objects in the account of this type are included in the results.
	//
	// Each implementation must implement it's own Filter
	Filter interface{} `json:"filter,omitempty"`

	// The ids of the Emails to fetch snippets for.
	EmailIDs []jmap.ID `json:"emailIds,omitempty"`
}

type SearchSnippetGetResponse struct {
	// The id of the account used for the call
	AccountID jmap.ID `json:"accountId,omitempty"`

	// An array of SearchSnippet objects for the requested Email ids.
	// This may not be in the same order as the ids that were in the
	// request.
	List []*SearchSnippet `json:"list,omitempty"`

	// An array of Email ids requested that could not be found, or null
	// if all ids were found.
	NotFound []jmap.ID `json:"notFound,omitempty"`
}
