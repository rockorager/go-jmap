package email

// This method allows you to parse blobs as messages [RFC5322] to get
// Email objects.  The server MUST support messages with EAI headers
// [RFC6532].  This can be used to parse and display attached messages
// without having to import them as top-level Email objects in the mail
// store in their own right.
//
// The following metadata properties on the Email objects will be null
// if requested:
//
// o  id
//
// o  mailboxIds
//
// o  keywords
//
// o  receivedAt
//
// The "threadId" property of the Email MAY be present if the server can
// calculate which Thread the Email would be assigned to were it to be
// imported.  Otherwise, this too is null if fetched.
type Parse struct {
	// The id of the account to use.
	AccountID string `json:"accountId,omitempty"`

	// The ids of the blobs to parse.
	BlobIDs []string `json:"blobIds,omitempty"`

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

func (m *Parse) Name() string {
	return "Email/parse"
}

func (m *Parse) Uses() string {
	return MailCapability
}

func (m *Parse) NewResponse() interface{} {
	return &ParseResponse{}
}

type ParseResponse struct {
	// The id of the account used for the call
	AccountID string `json:"accountId,omitempty"`

	// A map of blob id to parsed Email representation for each
	// successfully parsed blob, or null if none.
	Parsed map[string]*Email `json:"parsed,omitempty"`

	// A list of ids given that corresponded to blobs that could not be
	// parsed as Emails, or null if none.
	NotParsable []string `json:"notParsable,omitempty"`

	// A list of blob ids given that could not be found, or null if none.
	NotFound []string `json:"notFound,omitempty"`
}
