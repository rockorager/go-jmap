package mail

import "git.sr.ht/~rockorager/go-jmap"

type Email struct {
	// The id of the Email object. Note that this is the JMAP object id,
	// NOT the Message-ID header field value of the message [@!RFC5322].
	ID jmap.ID `json:"id,omitempty"`

	// The id representing the raw octets of the message [@!RFC5322] for
	// this Email. This may be used to download the raw original message or
	// to attach it directly to another Email, etc.
	BlobID jmap.ID `json:"blobId,omitempty"`

	// The id of the Thread to which this Email belongs.
	ThreadID jmap.ID `json:"threadId,omitempty"`

	// The set of Mailbox ids this Email belongs to. An Email in the mail
	// store MUST belong to one or more Mailboxes at all times (until it
	// is destroyed). The set is represented as an object, with each key
	// being a Mailbox id. The value for each key in the object MUST be
	// true.
	MailboxIDs map[jmap.ID]bool `json:"mailboxIds,omitempty"`

	// A set of keywords that apply to the Email. The set is represented as
	// an object, with the keys being the keywords. The value for each key
	// in the object MUST be true.
	Keywords map[string]bool `json:"keywords,omitempty"`

	// The size, in octets, of the raw data for the message [@!RFC5322] (as
	// referenced by the blobId, i.e., the number of octets in the file the
	// user would download).
	Size jmap.UnsignedInt `json:"size,omitempty"`

	// The date the Email was received by the message store. This is the
	// internal date in IMAP [@?RFC3501].
	ReceivedAt jmap.Date `json:"receivedAt,omitempty"`

	// This is a list of all header fields [@!RFC5322], in the same order
	// they appear in the message.
	Headers []*EmailHeader `json:"headers,omitempty"`

	// The value is identical to the value of
	// header:Message-ID:asMessageIds. For messages conforming to RFC 5322
	// this will be an array with a single entry.
	MessageID []string `json:"messageId,omitempty"`

	// The value is identical to the value of
	// header:In-Reply-To:asMessageIds.
	InReplyTo []string `json:"inReplyTo,omitempty"`

	// The value is identical to the value of
	// header:References:asMessageIds.mailAccount
	References []string `json:"references,omitempty"`

	// The value is identical to the value of header:Sender:asAddresses.
	Sender []*EmailAddress `json:"sender,omitempty"`

	// The value is identical to the value of header:From:asAddresses.
	From []*EmailAddress `json:"from,omitempty"`

	// The value is identical to the value of header:To:asAddresses.
	To []*EmailAddress `json:"to,omitempty"`

	// The value is identical to the value of header:Cc:asAddresses.
	CC []*EmailAddress `json:"cc,omitempty"`

	// The value is identical to the value of header:Bcc:asAddresses.
	BCC []*EmailAddress `json:"bcc,omitempty"`

	// The value is identical to the value of header:Reply-To:asAddresses.
	ReplyTo []*EmailAddress `json:"replyTo,omitempty"`

	// The value is identical to the value of header:Subject:asText.
	Subject string `json:"subject,omitempty"`

	// The value is identical to the value of header:Date:asDate.
	SentAt jmap.Date `json:"sentAt,omitempty"`

	// This is the full MIME structure of the message body, without
	// recursing into message/rfc822 or message/global parts. Note that
	// EmailBodyParts may have subParts if they are of type multipart/*.
	BodyStructure *EmailBodyPart `json:"bodyStructure,omitempty"`

	// A list of text/plain, text/html, image/*, audio/*, and/or video/*
	// parts to display (sequentially) as the message body, with a
	// preference for text/plain when alternative versions are available.
	TextBody []*EmailBodyPart `json:"textBody,omitempty"`

	// A list of text/plain, text/html, image/*, audio/*, and/or video/*
	// parts to display (sequentially) as the message body, with a
	// preference for text/html when alternative versions are available.
	HTMLBody []*EmailBodyPart `json:"htmlBody,omitempty"`

	// A list, traversing depth-first, of all parts in bodyStructure that
	// satisfy either of the following conditions:
	//
	//     not of type multipart/* and not included in textBody or htmlBody
	//
	//     of type image/*, audio/*, or video/* and not in both textBody
	//     and htmlBody
	//
	// None of these parts include subParts, including message/* types.
	// Attached messages may be fetched using the Email/parse method and
	// the blobId.
	//
	// Note that a text/html body part HTML may reference image parts in
	// attachments by using cid: links to reference the Content-Id, as
	// defined in [@!RFC2392], or by referencing the Content-Location.
	Attachments []*EmailBodyPart `json:"attachments,omitempty"`

	// This is true if there are one or more parts in the message that a
	// client UI should offer as downloadable. A server SHOULD set
	// hasAttachment to true if the attachments list contains at least one
	// item that does not have Content-Disposition: inline. The server MAY
	// ignore parts in this list that are processed automatically in some
	// way or are referenced as embedded images in one of the text/html
	// parts of the message.
	//
	// The server MAY set hasAttachment based on implementation-defined or
	// site-configurable heuristics.
	HasAttachment bool `json:"hasAttachment,omitempty"`

	// A plaintext fragment of the message body. This is intended to be
	// shown as a preview line when listing messages in the mail store and
	// may be truncated when shown. The server may choose which part of the
	// message to include in the preview; skipping quoted sections and
	// salutations and collapsing white space can result in a more useful
	// preview.
	//
	// This MUST NOT be more than 256 characters in length.
	//
	// As this is derived from the message content by the server, and the
	// algorithm for doing so could change over time, fetching this for an
	// Email a second time MAY return a different result. However, the
	// previous value is not considered incorrect, and the change SHOULD
	// NOT cause the Email object to be considered as changed by the
	// server.
	Preview string `json:"preview,omitempty"`
}

type EmailAddress struct {
	// The display-name of the mailbox [@!RFC5322]. If this is a
	// quoted-string:
	//
	//     The surrounding DQUOTE characters are removed. Any quoted-pair
	//     is decoded. White space is unfolded, and then any leading and
	//     trailing white space is removed.
	//
	// If there is no display-name but there is a comment immediately
	// following the addr-spec, the value of this SHOULD be used instead.
	// Otherwise, this property is null.
	Name string `json:"name,omitempty"`

	// The addr-spec of the mailbox [@!RFC5322].
	Email string `json:"email,omitempty"`
}

type EmailAddressGroup struct {
	// The display-name of the group [@!RFC5322], or null if the addresses
	// are not part of a group. If this is a quoted-string, it is processed
	// the same as the name in the EmailAddress type.
	Name string `json:"name,omitempty"`

	// The mailbox values that belong to this group, represented as
	// EmailAddress objects.
	Addresses []*EmailAddress `json:"addresses,omitempty"`
}

type EmailHeader struct {
	// The header field name as defined in [@!RFC5322], with the same
	// capitalization that it has in the message.
	Name string `json:"name,omitempty"`

	// The header field value as defined in [@!RFC5322], in Raw form.
	Value string `json:"value,omitempty"`
}

// These properties are derived from the message body [@!RFC5322] and its MIME
// entities [@RFC2045].
type EmailBodyPart struct {
	// Identifies this part uniquely within the Email. This is scoped to
	// the emailId and has no meaning outside of the JMAP Email object
	// representation. This is null if, and only if, the part is of type
	// multipart/*.
	PartID string `json:"partId,omitempty"`

	// The id representing the raw octets of the contents of the part,
	// after decoding any known Content-Transfer-Encoding (as defined in
	// [@!RFC2045]), or null if, and only if, the part is of type
	// multipart/*. Note that two parts may be transfer-encoded differently
	// but have the same blob id if their decoded octets are identical and
	// the server is using a secure hash of the data for the blob id. If
	// the transfer encoding is unknown, it is treated as though it had no
	// transfer encoding.
	BlobID jmap.ID `json:"blobId,omitempty"`

	// The size, in octets, of the raw data after content transfer decoding
	// (as referenced by the blobId, i.e., the number of octets in the file
	// the user would download).
	Size jmap.UnsignedInt `json:"size,omitempty"`

	// This is a list of all header fields in the part, in the order they
	// appear in the message. The values are in Raw form.
	Headers []*EmailHeader `json:"headers,omitempty"`

	// This is the decoded filename parameter of the Content-Disposition
	// header field per [@!RFC2231], or (for compatibility with existing
	// systems) if not present, then it’s the decoded name parameter of the
	// Content-Type header field per [@!RFC2047].
	Name string `json:"name,omitempty"`

	// The value of the Content-Type header field of the part, if present;
	// otherwise, the implicit type as per the MIME standard (text/plain or
	// message/rfc822 if inside a multipart/digest). CFWS is removed and
	// any parameters are stripped.
	Type string `json:"type,omitempty"`

	// The value of the charset parameter of the Content-Type header
	// field, if present, or null if the header field is present but not
	// of type text/*. If there is no Content-Type header field, or it
	// exists and is of type text/* but has no charset parameter, this is
	// the implicit charset as per the MIME standard: us-ascii.
	Charset string `json:"charset,omitempty"`

	// The value of the Content-Disposition header field of the part, if
	// present; otherwise, it’s null. CFWS is removed and any parameters
	// are stripped.
	Disposition string `json:"disposition,omitempty"`

	// The value of the Content-Id header field of the part, if present;
	// otherwise it’s null. CFWS and surrounding angle brackets (<>) are
	// removed. This may be used to reference the content from within a
	// text/html body part HTML using the cid: protocol, as defined in
	// [@!RFC2392].
	CID string `json:"cid,omitempty"`

	// The list of language tags, as defined in [@!RFC3282], in the
	// Content-Language header field of the part, if present.
	Language []string `json:"language,omitempty"`

	// The URI, as defined in [@!RFC2557], in the Content-Location header
	// field of the part, if present.
	Location string `json:"location,omitempty"`

	// If the type is multipart/*, this contains the body parts of each
	// child.
	SubParts []*EmailBodyPart `json:"subParts,omitempty"`
}

// This is a map of partId to an EmailBodyValue object for none, some, or all
// text/* parts. Which parts are included and whether the value is truncated is
// determined by various arguments to Email/get and Email/parse.
type EmailBodyValue struct {
	// The value of the body part after decoding Content-Transfer-Encoding
	// and the Content-Type charset, if both known to the server, and with
	// any CRLF replaced with a single LF. The server MAY use heuristics to
	// determine the charset to use for decoding if the charset is unknown,
	// no charset is given, or it believes the charset given is incorrect.
	// Decoding is best effort; the server SHOULD insert the unicode
	// replacement character (U+FFFD) and continue when a malformed section
	// is encountered.
	//
	// Note that due to the charset decoding and line ending normalisation,
	// the length of this string will probably not be exactly the same as
	// the size property on the corresponding EmailBodyPart.
	Value string `json:"value,omitempty"`

	// This is true if malformed sections were found while decoding the
	// charset, or the charset was unknown, or the
	// content-transfer-encoding was unknown.
	IsEncodingProblem bool `json:"isEncodingProblem,omitempty"`

	// This is true if the value has been truncated
	Istruncated bool `json:"isTruncated"`
}

// This is a standard get request
//
// If the standard "properties" argument is omitted or null, the following
// default MUST be used instead of "all" properties:
//
// [ "id", "blobId", "threadId", "mailboxIds", "keywords", "size",
// "receivedAt", "messageId", "inReplyTo", "references", "sender", "from",
// "to", "cc", "bcc", "replyTo", "subject", "sentAt", "hasAttachment",
// "preview", "bodyValues", "textBody", "htmlBody", "attachments" ]
type EmailGetRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

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
	MaxBodyValueBytes jmap.UnsignedInt `json:"maxBodyValueBytes,omitempty"`
}

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
type EmailGetResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

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

// This is a standard "/changes" method as described in [RFC8620], Section 5.2.
// If generating intermediate states for a large set of changes, it is
// recommended that newer changes be returned first, as these are generally of
// more interest to users.
type EmailChangesRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId"`

	// The current state of the client. This is the string that was
	// returned as the state argument in the Foo/get response. The server
	// will return the changes that have occurred since this state.
	SinceState string `json:"sinceState"`

	// The maximum number of ids to return in the response. The server MAY
	// choose to return fewer than this value but MUST NOT return more. If
	// not given by the client, the server may choose how many to return.
	// If supplied by the client, the value MUST be a positive integer
	// greater than 0. If a value outside of this range is given, the
	// server MUST reject the call with an invalidArguments error.
	MaxChanges jmap.UnsignedInt `json:"maxChanges"`
}

// This is a standard "/changes" method as described in [RFC8620], Section 5.2.
// If generating intermediate states for a large set of changes, it is
// recommended that newer changes be returned first, as these are generally of
// more interest to users.
type EmailChangesResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId"`

	// This is the sinceState argument echoed back; it’s the state from
	// which the server is returning changes.
	OldState string `json:"oldState"`

	// This is the state the client will be in after applying the set of
	// changes to the old state.
	NewState string `json:"newState"`

	// If true, the client may call Foo/changes again with the newState
	// returned to get further updates. If false, newState is the current
	// server state.
	HasMoreChanges bool `json:"hasMoreChanges"`

	// An array of ids for records that have been created since the old
	// state.
	Created []jmap.ID `json:"created"`

	// An array of ids for records that have been updated since the old
	// state.
	Updated []jmap.ID `json:"updated"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []jmap.ID `json:"destroyed"`
}

// This is a standard "/query" method as described in [RFC8620], Section 5.5
// but with the following additional request arguments:
type EmailQueryRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// Determines the set of Foos returned in the results. If null, all
	// objects in the account of this type are included in the results.
	//
	// Each implementation must implement it's own Filter
	Filter interface{} `json:"filter,omitempty"`

	// Lists the names of properties to compare between two Foo records,
	// and how to compare them, to determine which comes first in the sort.
	// If two Foo records have an identical value for the first comparator,
	// the next comparator will be considered, and so on. If all
	// comparators are the same (this includes the case where an empty
	// array or null is given as the sort argument), the sort order is
	// server dependent, but it MUST be stable between calls to Foo/query.
	//
	// Each implementation must define it's own Sort property. The
	// SortComparator object can be used as a basis
	Sort []*EmailSortComparator `json:"sort,omitempty"`

	// The zero-based index of the first id in the full list of results to
	// return.
	//
	// If a negative value is given, it is an offset from the end of the
	// list. Specifically, the negative value MUST be added to the total
	// number of results given the filter, and if still negative, it’s
	// clamped to 0. This is now the zero-based index of the first id to
	// return.
	//
	// If the index is greater than or equal to the total number of objects
	// in the results list, then the ids array in the response will be
	// empty, but this is not an error.
	Position jmap.Int `json:"position,omitempty"`

	// A Foo id. If supplied, the position argument is ignored. The index
	// of this id in the results will be used in combination with the
	// anchorOffset argument to determine the index of the first result to
	// return (see below for more details).
	//
	// If an anchor argument is given, the anchor is looked for in the
	// results after filtering and sorting. If found, the anchorOffset is
	// then added to its index. If the resulting index is now negative, it
	// is clamped to 0. This index is now used exactly as though it were
	// supplied as the position argument. If the anchor is not found, the
	// call is rejected with an anchorNotFound error.
	//
	// If an anchor is specified, any position argument supplied by the
	// client MUST be ignored. If no anchor is supplied, any anchorOffset
	// argument MUST be ignored.
	//
	// A client can use anchor instead of position to find the index of an
	// id within a large set of results.
	Anchor jmap.ID `json:"anchor,omitempty"`

	// The index of the first result to return relative to the index of the
	// anchor, if an anchor is given. This MAY be negative. For example, -1
	// means the Foo immediately preceding the anchor is the first result
	// in the list returned (see below for more details).
	AnchorOffset jmap.Int `json:"anchorOffset,omitempty"`

	// The maximum number of results to return. If null, no limit presumed.
	// The server MAY choose to enforce a maximum limit argument. In this
	// case, if a greater value is given (or if it is null), the limit is
	// clamped to the maximum; the new limit is returned with the response
	// so the client is aware. If a negative value is given, the call MUST
	// be rejected with an invalidArguments error.
	Limit jmap.UnsignedInt `json:"limit,omitempty"`

	// Does the client wish to know the total number of results in the
	// query? This may be slow and expensive for servers to calculate,
	// particularly with complex filters, so clients should take care to
	// only request the total when needed.
	CalculateTotal bool `json:"calculateTotal,omitempty"`

	// If true, Emails in the same Thread as a previous Email in the list
	// (given the filter and sort order) will be removed from the list.
	// This means only one Email at most will be included in the list for
	// any given Thread.
	CollapseThreads bool `json:"collapseThreads,omitempty"`
}

// Determines the set of Emails returned in the results. If null, all objects
// in the account of this type are included in the results.
type EmailFilter struct {
	// This MUST be one of the following strings: “AND” / “OR” / “NOT”
	Operator jmap.FilterOperator `json:"operator,omitempty"`

	// The conditions to evaluate against each record.
	Conditions []EmailFilterCondition `json:"conditions,omitempty"`
}

func (f *EmailFilter) AddCondition(cond EmailFilterCondition) {
	f.Conditions = append(f.Conditions, cond)
}

// EmailFilterCondition is an interface that represents FilterCondition
// objects. A filter condition object can be either a named struct, ie
// EmailFilterConditionName, or an EmailFilter itself. EmailFilters can
// be used to create complex filtering
type EmailFilterCondition interface{}

// A Mailbox id.  An Email must be in this Mailbox to match the condition.
type EmailFilterConditionIMailbox struct {
	InMailbox jmap.ID `json:"inMailbox,omitempty"`
}

// A list of Mailbox ids.  An Email must be in at least one Mailbox not in this
// list to match the condition.  This is to allow messages solely in trash/spam
// to be easily excluded from a search.
type EmailFilterConditionInMailboxOtherThan struct {
	InMailboxOtherThan []jmap.ID `json:"inMailboxOtherThan,omitempty"`
}

// The "receivedAt" date-time of the Email must be before this date- time to
// match the condition.
type EmailFilterConditionBefore struct {
	Before jmap.Date `json:"before,omitempty"`
}

// The "receivedAt" date-time of the Email must be the same or after this
// date-time to match the condition.
type EmailFilterConditionAfter struct {
	After jmap.Date `json:"after,omitempty"`
}

// The "size" property of the Email must be equal to or greater than this
// number to match the condition.
type EmailFilterConditionMinSize struct {
	MinSize jmap.UnsignedInt `json:"minSize,omitempty"`
}

// The "size" property of the Email must be less than this number to match the
// condition.
type EmailFilterConditionMaxSize struct {
	MaxSize jmap.UnsignedInt `json:"maxSize,omitempty"`
}

// All Emails (including this one) in the same Thread as this Email must have
// the given keyword to match the condition.
type EmailFilterConditionAllInThreadHaveKeyword struct {
	AllInThreadHaveKeyword string `json:"allInThreadHaveKeyword,omitempty"`
}

// At least one Email (possibly this one) in the same Thread as this Email must
// have the given keyword to match the condition.
type EmailFilterConditionSomeInThreadHaveKeyword struct {
	SomeInThreadHaveKeyword string `json:"someInThreadHaveKeyword,omitempty"`
}

// All Emails (including this one) in the same Thread as this Email must *not*
// have the given keyword to match the condition.
type EmailFilterConditionNoneInThreadHaveKeyword struct {
	NoneInThreadHaveKeyword string `json:"noneInThreadHaveKeyword,omitempty"`
}

// This Email must have the given keyword to match the condition.
type EmailFilterConditionHasKeyword struct {
	HasKeyword string `json:"hasKeyword,omitempty"`
}

// This Email must not have the given keyword to match the condition.
type EmailFilterConditionNotKeyword struct {
	NotKeyword string `json:"notKeyword,omitempty"`
}

// The "hasAttachment" property of the Email must be identical to the value
// given to match the condition.
type EmailFilterConditionHasAttachment struct {
	HasAttachment bool `json:"hasAttachment,omitempty"`
}

// Looks for the text in Emails.  The server MUST look up text in the From, To,
// Cc, Bcc, and Subject header fields of the message and SHOULD look inside any
// "text/*" or other body parts that may be converted to text by the server.
// The server MAY extend the search to any additional textual property.
type EmailFilterConditionText struct {
	Text string `json:"text,omitempty"`
}

// Looks for the text in the From header field of the message.
type EmailFilterConditionFrom struct {
	From string `json:"from,omitempty"`
}

// Looks for the text in the To header field of the message.
type EmailFilterConditionTo struct {
	To string `json:"to,omitempty"`
}

// Looks for the text in the Cc header field of the message.
type EmailFilterConditionCc struct {
	Cc string `json:"cc,omitempty"`
}

// Looks for the text in the Bcc header field of the message.
type EmailFilterConditionBcc struct {
	Bcc string `json:"bcc,omitempty"`
}

// Looks for the text in the Subject header field of the message.
type EmailFilterConditionSubject struct {
	Subject string `json:"subject,omitempty"`
}

// Looks for the text in one of the body parts of the message.  The server MAY
// exclude MIME body parts with content media types other than "text/*" and
// "message/*" from consideration in search matching.  Care should be taken to
// match based on the text content actually presented to an end user by viewers
// for that media type or otherwise identified as appropriate for search
// indexing. Matching document metadata uninteresting to an end user (e.g.,
// markup tag and attribute names) is undesirable.
type EmailFilterConditionBody struct {
	Body string `json:"body,omitempty"`
}

// The array MUST contain either one or two elements.  The first element is the
// name of the header field to match against.  The second (optional) element is
// the text to look for in the header field value.  If not supplied, the
// message matches simply if it has a header field of the given name.
type EmailFilterConditionHeader struct {
	Header []string `json:"header,omitempty"`
}

type EmailSortComparator struct {
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

type EmailQueryResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// A string encoding the current state of the query on the server. This
	// string MUST change if the results of the query (i.e., the matching
	// ids and their sort order) have changed. The queryState string MAY
	// change if something has changed on the server, which means the
	// results may have changed but the server doesn’t know for sure.
	//
	// The queryState string only represents the ordered list of ids that
	// match the particular query (including its sort/filter). There is no
	// requirement for it to change if a property on an object matching the
	// query changes but the query results are unaffected (indeed, it is
	// more efficient if the queryState string does not change in this
	// case). The queryState string only has meaning when compared to
	// future responses to a query with the same type/sort/filter or when
	// used with /queryChanges to fetch changes.
	//
	// Should a client receive back a response with a different queryState
	// string to a previous call, it MUST either throw away the currently
	// cached query and fetch it again (note, this does not require
	// fetching the records again, just the list of ids) or call
	// Foo/queryChanges to get the difference.
	QueryState string `json:"queryState,omitempty"`

	// This is true if the server supports calling Foo/queryChanges with
	// these filter/sort parameters. Note, this does not guarantee that the
	// Foo/queryChanges call will succeed, as it may only be possible for a
	// limited time afterwards due to server internal implementation
	// details.
	CanCalculateChanges bool `json:"canCalculateChanges,omitempty"`

	// The zero-based index of the first result in the ids array within the
	// complete list of query results.
	Position jmap.UnsignedInt `json:"position,omitempty"`

	// The list of ids for each Foo in the query results, starting at the
	// index given by the position argument of this response and continuing
	// until it hits the end of the results or reaches the limit number of
	// ids. If position is >= total, this MUST be the empty list.
	IDs []jmap.ID `json:"ids,omitempty"`

	// The total number of Foos in the results (given the filter). This
	// argument MUST be omitted if the calculateTotal request argument is
	// not true.
	Total jmap.UnsignedInt `json:"total,omitempty"`

	// The limit enforced by the server on the maximum number of results to
	// return. This is only returned if the server set a limit or used a
	// different limit than that given in the request.
	Limit jmap.UnsignedInt `json:"limit,omitempty"`
}

// This is a standard "/queryChanges" method as described in [RFC8620], Section
// 5.6 with the following additional request argument: collapseThreads
type EmailQueryChangesRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The filter argument that was used with Foo/query.
	//
	// Each implementation must supply it's own Filter property
	Filter interface{} `json:"filter,omitempty"`

	// The sort argument that was used with Foo/query.
	//
	// Each implementation must supply it's own Sort property
	Sort []*EmailSortComparator `json:"sort,omitempty"`

	// The current state of the query in the client. This is the string
	// that was returned as the queryState argument in the Foo/query
	// response with the same sort/filter. The server will return the
	// changes made to the query since this state.
	SinceQueryState string `json:"sinceQueryState,omitempty"`

	// The maximum number of changes to return in the response. See error
	// descriptions below for more details.
	MaxChanges jmap.UnsignedInt `json:"maxChanges,omitempty"`

	// The last (highest-index) id the client currently has cached from the
	// query results. When there are a large number of results, in a common
	// case, the client may have only downloaded and cached a small subset
	// from the beginning of the results. If the sort and filter are both
	// only on immutable properties, this allows the server to omit changes
	// after this point in the results, which can significantly increase
	// efficiency. If they are not immutable, this argument is ignored.
	UpToID jmap.ID `json:"upToId,omitempty"`

	// Does the client wish to know the total number of results now in the
	// query? This may be slow and expensive for servers to calculate,
	// particularly with complex filters, so clients should take care to
	// only request the total when needed.
	CalculateTotal bool `json:"calculateTotal,omitempty"`

	// If true, Emails in the same Thread as a previous Email in the list
	// (given the filter and sort order) will be removed from the list.
	// This means only one Email at most will be included in the list for
	// any given Thread.
	CollapseThreads bool `json:"collapseThreads,omitempty"`
}

// This is a standard "/queryChanges" method as described in [RFC8620], Section
// 5.6
type EmailQueryChangesResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// This is the sinceQueryState argument echoed back; that is, the state
	// from which the server is returning changes.
	OldQueryState string `json:"oldQueryState,omitempty"`

	// This is the state the query will be in after applying the set of
	// changes to the old state.
	NewQueryState string `json:"newQueryState,omitempty"`

	// The id for every Foo that was in the query results in the old state
	// and that is not in the results in the new state.
	//
	// If the server cannot calculate this exactly, the server MAY return
	// the ids of extra Foos in addition that may have been in the old
	// results but are not in the new results.
	//
	// If the sort and filter are both only on immutable properties and an
	// upToId is supplied and exists in the results, any ids that were
	// removed but have a higher index than upToId SHOULD be omitted.
	//
	// If the filter or sort includes a mutable property, the server MUST
	// include all Foos in the current results for which this property may
	// have changed. The position of these may have moved in the results,
	// so must be reinserted by the client to ensure its query cache is
	// correct.
	Removed []jmap.ID `json:"removed,omitempty"`

	// The id and index in the query results (in the new state) for every
	// Foo that has been added to the results since the old state AND every
	// Foo in the current results that was included in the removed array
	// (due to a filter or sort based upon a mutable property).
	//
	// If the sort and filter are both only on immutable properties and an
	// upToId is supplied and exists in the results, any ids that were
	// added but have a higher index than upToId SHOULD be omitted.
	//
	// The array MUST be sorted in order of index, with the lowest index
	// first.
	Added []jmap.AddedItem `json:"added,omitempty"`
}

// This is a standard "/set" method as described in [RFC8620], Section 5.3.
// The "Email/set" method encompasses:
//
// o  Creating a draft
//
// o  Changing the keywords of an Email (e.g., unread/flagged status)
//
// o  Adding/removing an Email to/from Mailboxes (moving a message)
//
// o  Deleting Emails
//
// The format of the "keywords"/"mailboxIds" properties means that when
// updating an Email, you can either replace the entire set of keywords/
// Mailboxes (by setting the full value of the property) or add/remove
// individual ones using the JMAP patch syntax (see [RFC8620],
// Section 5.3 for the specification and Section 5.7 for an example).
//
// Due to the format of the Email object, when creating an Email, there
// are a number of ways to specify the same information.  To ensure that
// the message [RFC5322] to create is unambiguous, the following
// constraints apply to Email objects submitted for creation:
//
// o  The "headers" property MUST NOT be given on either the top-level
//    Email or an EmailBodyPart -- the client must set each header field
//    as an individual property.
//
// o  There MUST NOT be two properties that represent the same header
//    field (e.g., "header:from" and "from") within the Email or
//    particular EmailBodyPart.
//
// o  Header fields MUST NOT be specified in parsed forms that are
//    forbidden for that particular field.
//
// o  Header fields beginning with "Content-" MUST NOT be specified on
//    the Email object, only on EmailBodyPart objects.
//
// o  If a "bodyStructure" property is given, there MUST NOT be
//    "textBody", "htmlBody", or "attachments" properties.
//
// o  If given, the "bodyStructure" EmailBodyPart MUST NOT contain a
//    property representing a header field that is already defined on
//    the top-level Email object.
//
// o  If given, textBody MUST contain exactly one body part and it MUST
//    be of type "text/plain".
//
// o  If given, htmlBody MUST contain exactly one body part and it MUST
//    be of type "text/html".
//
// o  Within an EmailBodyPart:
//
//    *  The client may specify a partId OR a blobId, but not both.  If
//       a partId is given, this partId MUST be present in the
//       "bodyValues" property.
//
//    *  The "charset" property MUST be omitted if a partId is given
//       (the part's content is included in bodyValues, and the server
//       may choose any appropriate encoding).
//
//    *  The "size" property MUST be omitted if a partId is given.  If a
//       blobId is given, it may be included but is ignored by the
//       server (the size is actually calculated from the blob content
//       itself).
//
//    *  A Content-Transfer-Encoding header field MUST NOT be given.
type EmailSetRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// This is a state string as returned by the Foo/get method
	// (representing the state of all objects of this type in the account).
	// If supplied, the string must match the current state; otherwise, the
	// method will be aborted and a stateMismatch error returned. If null,
	// any changes will be applied to the current state.
	IfInState string `json:"ifInState,omitempty"`

	// A map of a creation id (a temporary id set by the client) to Foo
	// objects, or null if no objects are to be created.
	//
	// The Foo object type definition may define default values for
	// properties. Any such property may be omitted by the client.
	//
	// The client MUST omit any properties that may only be set by the
	// server (for example, the id property on most object types).
	Create map[jmap.ID]*Email `json:"create,omitempty"`

	// A map of an id to a Patch object to apply to the current Foo object
	// with that id, or null if no objects are to be updated.
	//
	// A PatchObject is of type String[*] and represents an unordered set
	// of patches. The keys are a path in JSON Pointer Format [@!RFC6901],
	// with an implicit leading “/” (i.e., prefix each key with “/” before
	// applying the JSON Pointer evaluation algorithm).
	//
	// All paths MUST also conform to the following restrictions; if there
	// is any violation, the update MUST be rejected with an invalidPatch
	// error:
	//
	//     The pointer MUST NOT reference inside an array (i.e., you MUST
	//     NOT insert/delete from an array; the array MUST be replaced in
	//     its entirety instead). All parts prior to the last (i.e., the
	//     value after the final slash) MUST already exist on the object
	//     being patched. There MUST NOT be two patches in the PatchObject
	//     where the pointer of one is the prefix of the pointer of the
	//     other, e.g., “alerts/1/offset” and “alerts”.
	//
	// The value associated with each pointer determines how to apply that
	// patch:
	//
	//     If null, set to the default value if specified for this
	//     property; otherwise, remove the property from the patched
	//     object. If the key is not present in the parent, this a no-op.
	//     Anything else: The value to set for this property (this may be a
	//     replacement or addition to the object being patched).
	//
	// Any server-set properties MAY be included in the patch if their
	// value is identical to the current server value (before applying the
	// patches to the object). Otherwise, the update MUST be rejected with
	// an invalidProperties SetError.
	//
	// This patch definition is designed such that an entire Foo object is
	// also a valid PatchObject. The client may choose to optimise network
	// usage by just sending the diff or may send the whole object; the
	// server processes it the same either way.
	Update map[jmap.ID]map[string]interface{} `json:"update,omitempty"`

	// A list of ids for Foo objects to permanently delete, or null if no
	// objects are to be destroyed.
	Destroy []jmap.ID `json:"destroy,omitempty"`
}

type EmailSetResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The state string that would have been returned by Foo/get before
	// making the requested changes, or null if the server doesn’t know
	// what the previous state string was.
	OldState string `json:"oldState,omitempty"`

	// The state string that will now be returned by Foo/get.
	NewState string `json:"newState,omitempty"`

	// A map of the creation id to an object containing any properties of
	// the created Foo object that were not sent by the client. This
	// includes all server-set properties (such as the id in most object
	// types) and any properties that were omitted by the client and thus
	// set to a default by the server.
	//
	// This argument is null if no Foo objects were successfully created.
	Created map[jmap.ID]*Email `json:"created,omitempty"`

	// The keys in this map are the ids of all Foos that were successfully
	// updated.
	//
	// The value for each id is a Foo object containing any property that
	// changed in a way not explicitly requested by the PatchObject sent to
	// the server, or null if none. This lets the client know of any
	// changes to server-set or computed properties.
	//
	// This argument is null if no Foo objects were successfully updated.
	Updated map[jmap.ID]*Email `json:"updated,omitempty"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []jmap.ID `json:"destroyed,omitempty"`
}

// This is a standard "/copy" method as described in [RFC8620], Section 5.4,
// except only the "mailboxIds", "keywords", and "receivedAt" properties may be
// set during the copy.  This method cannot modify the message represented by
// the Email.
type EmailCopyRequest struct {
	// The id of the account to copy records from.
	FromAccountID jmap.ID `json:"fromAccountId,omitempty"`

	// This is a state string as returned by the Foo/get method. If
	// supplied, the string must match the current state of the account
	// referenced by the fromAccountId when reading the data to be copied;
	// otherwise, the method will be aborted and a stateMismatch error
	// returned. If null, the data will be read from the current state.
	IfFromInState string `json:"ifFromInState,omitempty"`

	// The id of the account to copy records to. This MUST be different to
	// the fromAccountId.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// This is a state string as returned by the Foo/get method. If
	// supplied, the string must match the current state of the account
	// referenced by the accountId; otherwise, the method will be aborted
	// and a stateMismatch error returned. If null, any changes will be
	// applied to the current state.
	IfInState string `json:"ifInState,omitempty"`

	// A map of the creation id to a Foo object. The Foo object MUST
	// contain an id property, which is the id (in the fromAccount) of the
	// record to be copied. When creating the copy, any other properties
	// included are used instead of the current value for that property on
	// the original.
	Create map[jmap.ID]*Email `json:"create,omitempty"`

	// If true, an attempt will be made to destroy the original records
	// that were successfully copied: after emitting the Foo/copy response,
	// but before processing the next method, the server MUST make a single
	// call to Foo/set to destroy the original of each successfully copied
	// record; the output of this is added to the responses as normal, to
	// be returned to the client.
	OnSuccessDestroyOriginal bool `json:"onSuccessDestroyOriginal,omitempty"`

	// This argument is passed on as the ifInState argument to the implicit
	// Foo/set call, if made at the end of this request to destroy the
	// originals that were successfully copied.
	DestroyFromIfInState string `json:"destroyFromIfInState,omitempty"`
}

type EmailCopyResponse struct {
	// The id of the account records were copied from.
	FromAccountID jmap.ID `json:"fromAccountId,omitempty"`

	// The id of the account records were copied to.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The state string that would have been returned by Foo/get on the
	// account records that were copied to before making the requested
	// changes, or null if the server doesn’t know what the previous state
	// string was.
	OldState string `json:"oldState,omitempty"`

	// The state string that will now be returned by Foo/get on the account
	// records were copied to.
	NewState string `json:"newState,omitempty"`

	// A map of the creation id to an object containing any properties of
	// the copied Foo object that are set by the server (such as the id in
	// most object types; note, the id is likely to be different to the id
	// of the object in the account it was copied from).
	//
	// This argument is null if no Foo objects were successfully copied.
	Created map[jmap.ID]*Email `json:"created,omitempty"`

	// A map of the creation id to a SetError object for each record that
	// failed to be copied, or null if none.
	NotCreated map[jmap.ID]*jmap.MethodErrorArgs `json:"notCreated,omitempty"`
}

// The "Email/import" method adds messages [RFC5322] to the set of Emails in an
// account.  The server MUST support messages with Email Address
// Internationalization (EAI) headers [RFC6532].  The messages must first be
// uploaded as blobs using the standard upload mechanism.
type EmailImportRequest struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// This is a state string as returned by the Foo/get method
	// (representing the state of all objects of this type in the account).
	// If supplied, the string must match the current state; otherwise, the
	// method will be aborted and a stateMismatch error returned. If null,
	// any changes will be applied to the current state.
	IfInState string `json:"ifInState,omitempty"`

	// A map of creation id (client specified) to EmailImport objects.
	Emails map[jmap.ID]*EmailImport `json:"emails,omitempty"`
}

type EmailImport struct {
	// The id of the blob containing the raw message [RFC5322].
	BlobID jmap.ID `json:"blobId,omitempty"`

	// The ids of the Mailboxes to assign this Email to.  At least one
	// Mailbox MUST be given.
	MailboxIDs map[jmap.ID]bool `json:"mailboxIds,omitempty"`

	// The keywords to apply to the Email.
	Keywords map[string]bool `json:"keywords,omitempty"`

	// The "receivedAt" date to set on the Email.
	ReceivedAt jmap.Date `json:"receivedAt,omitempty"`
}

type EmailImportResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The state string that would have been returned by Foo/get before
	// making the requested changes, or null if the server doesn’t know
	// what the previous state string was.
	OldState string `json:"oldState,omitempty"`

	// The state string that will now be returned by Foo/get.
	NewState string `json:"newState,omitempty"`

	// A map of the creation id to an object containing any properties of
	// the created Foo object that were not sent by the client. This
	// includes all server-set properties (such as the id in most object
	// types) and any properties that were omitted by the client and thus
	// set to a default by the server.
	//
	// This argument is null if no Foo objects were successfully created.
	Created map[jmap.ID]*Email `json:"created,omitempty"`

	// A map of the creation id to a SetError object for each Email that
	// failed to be created, or null if all successful.  The possible
	// errors are defined above.
	NotCreated map[jmap.ID]*jmap.MethodErrorArgs `json:"notCreated,omitempty"`
}

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
type EmailParseRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The ids of the blobs to parse.
	BlobIDs []jmap.ID `json:"blobIds,omitempty"`

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
	MaxBodyValueBytes jmap.UnsignedInt `json:"maxBodyValueBytes,omitempty"`
}

type EmailParseResponse struct {
	// The id of the account used for the call
	AccountID jmap.ID `json:"accountId,omitempty"`

	// A map of blob id to parsed Email representation for each
	// successfully parsed blob, or null if none.
	Parsed map[jmap.ID]*Email `json:"parsed,omitempty"`

	// A list of ids given that corresponded to blobs that could not be
	// parsed as Emails, or null if none.
	NotParsable []jmap.ID `json:"notParsable,omitempty"`

	// A list of blob ids given that could not be found, or null if none.
	NotFound []jmap.ID `json:"notFound,omitempty"`
}
