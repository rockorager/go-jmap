package email

import "git.sr.ht/~rockorager/go-jmap"

// This is a standard get request
//
// If the standard "properties" argument is omitted or null, the following
// default MUST be used instead of "all" properties:
//
// [ "id", "blobId", "threadId", "mailboxIds", "keywords", "size",
// "receivedAt", "messageId", "inReplyTo", "references", "sender", "from",
// "to", "cc", "bcc", "replyTo", "subject", "sentAt", "hasAttachment",
// "preview", "bodyValues", "textBody", "htmlBody", "attachments" ]
type Get struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId,omitempty"`

	// The ids of the Foo objects to return. If null, then all records of
	// the data type are returned, if this is supported for that data type
	// and the number of records does not exceed the maxObjectsInGet limit.
	IDs []jmap.ID `json:"ids,omitempty"`

	// If supplied, only the properties listed in the array are returned
	// for each Foo object. If null, all properties of the object are
	// returned. The id property of the object is always returned, even if
	// not explicitly requested. If an invalid property is requested, the
	// call MUST be rejected with an invalidArguments error.
	Properties []string `json:"properties,omitempty"`

	// A list of properties to fetch for each EmailBodyPart returned.  If
	// omitted, this defaults to:
	//
	//    [ "partId", "blobId", "size", "name", "type", "charset",
	//      "disposition", "cid", "language", "location" ]
	BodyProperties []string `json:"bodyProperties,omitempty"`

	// If true, the "bodyValues" property includes any "text/*" part in
	// the "textBody" property.
	FetchTextBodyValues bool `json:"fetchTextBodyValues,omitempty"`

	// If true, the "bodyValues" property includes any "text/*" part in
	// the "htmlBody" property.
	FetchHTMLBodyValues bool `json:"fetchHTMLBodyValues,omitempty"`

	// If true, the "bodyValues" property includes any "text/*" part in
	// the "bodyStructure" property.
	FetchAllBodyValues bool `json:"fetchAllBodyValues,omitempty"`

	// If greater than zero, the "value" property of any EmailBodyValue
	// object returned in "bodyValues" MUST be truncated if necessary so
	// it does not exceed this number of octets in size.  If 0 (the
	// default), no truncation occurs.
	//
	// The server MUST ensure the truncation results in valid UTF-8 and
	// does not occur mid-codepoint.  If the part is of type "text/html",
	// the server SHOULD NOT truncate inside an HTML tag, e.g., in the
	// middle of "<a href="https://example.com">".  There is no
	// requirement for the truncated form to be a balanced tree or valid
	// HTML (indeed, the original source may well be neither of these
	// things).
	MaxBodyValueBytes uint64 `json:"maxBodyValueBytes,omitempty"`
}

func (m *Get) Name() string {
	return "Mailbox/get"
}

func (m *Get) Uses() string {
	return MailCapability
}

func (m *Get) NewResponse() interface{} {
	return &GetResponse{}
}

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
type GetResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// A (preferably short) string representing the state on the server for
	// all the data of this type in the account (not just the objects
	// returned in this call). If the data changes, this string MUST
	// change. If the Foo data is unchanged, servers SHOULD return the same
	// state string on subsequent requests for this data type.
	//
	// When a client receives a response with a different state string to a
	// previous call, it MUST either throw away all currently cached
	// objects for the type or call Foo/changes to get the exact changes.
	State string `json:"state,omitempty"`

	// An array of the Foo objects requested. This is the empty array
	// if no objects were found or if the ids argument passed in was also
	// an empty array. The results MAY be in a different order to the ids
	// in the request arguments. If an identical id is included more than
	// once in the request, the server MUST only include it once in either
	// the list or the notFound argument of the response.
	//
	// Each specification must define it's own List property
	List []*Email `json:"list,omitempty"`

	// This array contains the ids passed to the method for records that do
	// not exist. The array is empty if all requested ids were found or if
	// the ids argument passed in was either null or an empty array.
	NotFound []jmap.ID `json:"notFound,omitempty"`
}
