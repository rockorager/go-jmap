package mail

import (
	"git.sr.ht/~rockorager/go-jmap"
)

// A Mailbox represents a named set of Emails. This is the primary mechanism
// for organising Emails within an account. It is analogous to a folder or a
// label in other systems. A Mailbox may perform a certain role in the system;
// see below for more details.
//
// For compatibility with IMAP, an Email MUST belong to one or more Mailboxes.
// The Email id does not change if the Email changes Mailboxes.
type Mailbox struct {
	// The id of the Mailbox.
	ID jmap.ID `json:"id,omitempty"`

	// User-visible name for the Mailbox, e.g., “Inbox”. This MUST be a
	// Net-Unicode string [@!RFC5198] of at least 1 character in length,
	// subject to the maximum size given in the capability object. There
	// MUST NOT be two sibling Mailboxes with both the same parent and the
	// same name. Servers MAY reject names that violate server policy
	// (e.g., names containing a slash (/) or control characters).
	Name string `json:"name,omitempty"`

	// The Mailbox id for the parent of this Mailbox, or null if this
	// Mailbox is at the top level. Mailboxes form acyclic graphs (forests)
	// directed by the child-to-parent relationship. There MUST NOT be a
	// loop.
	ParentID jmap.ID `json:"parentId,omitempty"`

	// Identifies Mailboxes that have a particular common purpose (e.g.,
	// the “inbox”), regardless of the name property (which may be
	// localised).
	//
	// This value is shared with IMAP (exposed in IMAP via the SPECIAL-USE
	// extension [@!RFC6154]). However, unlike in IMAP, a Mailbox MUST only
	// have a single role, and there MUST NOT be two Mailboxes in the same
	// account with the same role. Servers providing IMAP access to the
	// same data are encouraged to enforce these extra restrictions in IMAP
	// as well. Otherwise, modifying the IMAP attributes to ensure
	// compliance when exposing the data over JMAP is implementation
	// dependent.
	//
	// The value MUST be one of the Mailbox attribute names listed in the
	// IANA IMAP Mailbox Name Attributes registry, as established in
	// [@!RFC8457], converted to lowercase. New roles may be established
	// here in the future.
	//
	// An account is not required to have Mailboxes with any particular
	// roles.
	Role Role `json:"role,omitempty"`

	// Defines the sort order of Mailboxes when presented in the client’s
	// UI, so it is consistent between devices. The number MUST be an
	// integer in the range 0 <= sortOrder < 2^31.
	//
	// A Mailbox with a lower order should be displayed before a Mailbox
	// with a higher order (that has the same parent) in any Mailbox
	// listing in the client’s UI. Mailboxes with equal order SHOULD be
	// sorted in alphabetical order by name. The sorting should take into
	// account locale-specific character order convention.
	SortOrder jmap.UnsignedInt `json:"sortOrder,omitempty"`

	// The number of Emails in this Mailbox.
	TotalEmails jmap.UnsignedInt `json:"totalEmails,omitempty"`

	// The number of Emails in this Mailbox that have neither the $seen
	// keyword nor the $draft keyword.
	UnreadEmails jmap.UnsignedInt `json:"unreadEmails,omitempty"`

	// The number of Threads where at least one Email in the Thread is in
	// this Mailbox.
	TotalThreads jmap.UnsignedInt `json:"totalThreads,omitempty"`

	// An indication of the number of “unread” Threads in the Mailbox.
	//
	// For compatibility with existing implementations, the way “unread
	// Threads” is determined is not mandated in this document. The
	// simplest solution to implement is simply the number of Threads where
	// at least one Email in the Thread is both in this Mailbox and has
	// neither the $seen nor $draft keywords.
	//
	// However, a quality implementation will return the number of unread
	// items the user would see if they opened that Mailbox. A Thread is
	// shown as unread if it contains any unread Emails that will be
	// displayed when the Thread is opened. Therefore, unreadThreads should
	// be the number of Threads where at least one Email in the Thread has
	// neither the $seen nor the $draft keyword AND at least one Email in
	// the Thread is in this Mailbox. Note that the unread Email does not
	// need to be the one in this Mailbox. In addition, the trash Mailbox
	// (that is, a Mailbox whose role is trash) requires special treatment:
	//
	//     Emails that are only in the trash (and no other Mailbox) are
	//     ignored when calculating the unreadThreads count of other
	//     Mailboxes. Emails that are not in the trash are ignored when
	//     calculating the unreadThreads count for the trash Mailbox.
	//
	// The result of this is that Emails in the trash are treated as though
	// they are in a separate Thread for the purposes of unread counts. It
	// is expected that clients will hide Emails in the trash when viewing
	// a Thread in another Mailbox, and vice versa. This allows you to
	// delete a single Email to the trash out of a Thread.
	//
	// For example, suppose you have an account where the entire contents
	// is a single Thread with 2 Emails: an unread Email in the trash and a
	// read Email in the inbox. The unreadThreads count would be 1 for the
	// trash and 0 for the inbox.
	UnreadThreads jmap.UnsignedInt `json:"unreadThreads,omitempty"`

	// The set of rights (Access Control Lists (ACLs)) the user has in
	// relation to this Mailbox. These are backwards compatible with IMAP
	// ACLs, as defined in [@!RFC4314].
	Rights *MailboxRights `json:"myRights,omitempty"`

	// Has the user indicated they wish to see this Mailbox in their
	// client? This SHOULD default to false for Mailboxes in shared
	// accounts the user has access to and true for any new Mailboxes
	// created by the user themself. This MUST be stored separately per
	// user where multiple users have access to a shared Mailbox.
	//
	// A user may have permission to access a large number of shared
	// accounts, or a shared account with a very large set of Mailboxes,
	// but only be interested in the contents of a few of these. Clients
	// may choose to only display Mailboxes where the isSubscribed property
	// is set to true, and offer a separate UI to allow the user to see and
	// subscribe/unsubscribe from the full set of Mailboxes. However,
	// clients MAY choose to ignore this property, either entirely for ease
	// of implementation or just for an account where isPersonal is true
	// (indicating it is the user’s own rather than a shared account).
	//
	// This property corresponds to IMAP [@?RFC3501] Mailbox subscriptions.
	IsSubscribed bool `json:"isSubscribed,omitempty"`
}

// The set of rights (Access Control Lists (ACLs)) the user has in relation to
// this Mailbox. These are backwards compatible with IMAP ACLs, as defined in
// [@!RFC4314].
type MailboxRights struct {
	// If true, the user may use this Mailbox as part of a filter in an
	// Email/query call, and the Mailbox may be included in the mailboxIds
	// property of Email objects. Email objects may be fetched if they are
	// in at least one Mailbox with this permission. If a sub-Mailbox is
	// shared but not the parent Mailbox, this may be false. Corresponds to
	// IMAP ACLs lr (if mapping from IMAP, both are required for this to be
	// true).
	MayReadItems bool `json:"mayReadItems,omitempty"`

	// The user may add mail to this Mailbox (by either creating a new
	// Email or moving an existing one). Corresponds to IMAP ACL i.
	MayAddItems bool `json:"mayAddItems,omitempty"`

	// The user may remove mail from this Mailbox (by either changing the
	// Mailboxes of an Email or destroying the Email). Corresponds to IMAP
	// ACLs te (if mapping from IMAP, both are required for this to be
	// true).
	MayRemoveItems bool `json:"mayRemoveItems,omitempty"`

	// The user may add or remove the $seen keyword to/from an Email. If an
	// Email belongs to multiple Mailboxes, the user may only modify $seen
	// if they have this permission for all of the Mailboxes. Corresponds
	// to IMAP ACL s.
	MaySetSeen bool `json:"maySetSeen,omitempty"`

	// The user may add or remove any keyword other than $seen to/from an
	// Email. If an Email belongs to multiple Mailboxes, the user may only
	// modify keywords if they have this permission for all of the
	// Mailboxes. Corresponds to IMAP ACL w.
	MaySetKeywords bool `json:"maySetKeywords,omitempty"`

	// The user may create a Mailbox with this Mailbox as its parent.
	// Corresponds to IMAP ACL k.
	MayCreateChild bool `json:"mayCreateChild,omitempty"`

	// The user may rename the Mailbox or make it a child of another
	// Mailbox. Corresponds to IMAP ACL x (although this covers both rename
	// and delete permissions).
	MayRename bool `json:"mayRename,omitempty"`

	// The user may delete the Mailbox itself. Corresponds to IMAP ACL x
	// (although this covers both rename and delete permissions).
	MayDelete bool `json:"mayDelete,omitempty"`

	// Messages may be submitted directly to this Mailbox. Corresponds to
	// IMAP ACL p.
	MaySubmit bool `json:"maySubmit,omitempty"`
}

// Identifies Mailboxes that have a particular common purpose (e.g., the
// “inbox”), regardless of the name property (which may be localised).
//
// This value is shared with IMAP (exposed in IMAP via the SPECIAL-USE
// extension [@!RFC6154]). However, unlike in IMAP, a Mailbox MUST only have a
// single role, and there MUST NOT be two Mailboxes in the same account with
// the same role. Servers providing IMAP access to the same data are encouraged
// to enforce these extra restrictions in IMAP as well. Otherwise, modifying
// the IMAP attributes to ensure compliance when exposing the data over JMAP is
// implementation dependent.
//
// The value MUST be one of the Mailbox attribute names listed in the IANA IMAP
// Mailbox Name Attributes registry, as established in [@!RFC8457], converted
// to lowercase. New roles may be established here in the future.
//
// An account is not required to have Mailboxes with any particular roles.
//
// Reference:
// https://www.iana.org/assignments/imap-mailbox-name-attributes/imap-mailbox-name-attributes.xhtml
type Role string

const (
	// All messages
	RoleAll Role = "all"

	// Archived Messages
	RoleArchive Role = "archive"

	// Messages that are working drafts
	RoleDrafts Role = "drafts"

	// Messages with the \Flagged flag
	RoleFlagged Role = "flagged"

	// Has accessible child mailboxes
	//
	// Not used by JMAP
	RoleHasChildren Role = "haschildren"

	// Has no accessible child mailboxes
	//
	// Not used by JMAP
	RoleHasNoChildren Role = "hasnochildren"

	// Messages deemed important to user
	RoleImportant Role = "important"

	// New mail is delivered here by default
	//
	// JMAP only
	RoleInbox Role = "inbox"

	// Messages identified as Spam/Junk
	RoleJunk Role = "junk"

	// Server has marked the mailbox as "interesting"
	//
	// Not used by JMAP
	RoleMarked Role = "marked"

	// No hierarchy under this name
	//
	// Not used by JMAP
	RoleNoInferiors Role = "noinferiors"

	// The mailbox name doesn't actually exist
	//
	// Not used by JMAP
	RoleNonExistent Role = "nonexistent"

	// The mailbox is not selectable
	//
	// Not used by JMAP
	RoleNoSelect Role = "noselect"

	// The mailbox exists on a remote server
	//
	// Not used by JMAP
	RoleRemote Role = "remote"

	// Sent mail
	RoleSent Role = "sent"

	// The mailbox is subscribed to
	RoleSubscribed Role = "subscribed"

	// Messages the user has discarded
	RoleTrash Role = "trash"

	// No new messages since last select
	//
	// Not used by JMAP
	RoleUnmarked Role = "unmarked"
)

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
//
// Objects of type Mailbox are fetched via a call to Mailbox/get The ids
// argument may be null to fetch all at once.
type MailboxGetRequest struct {
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
}

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
type MailboxGetResponse struct {
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
	List []*Mailbox `json:"list,omitempty"`

	// This array contains the ids passed to the method for records that do
	// not exist. The array is empty if all requested ids were found or if
	// the ids argument passed in was either null or an empty array.
	NotFound []jmap.ID `json:"notFound,omitempty"`
}

// This is a standard “/changes” method as described in [@!RFC8620], Section
// 5.2.
//
// When the state of the set of Foo records in an account changes on the server
// (whether due to creation, updates, or deletion), the state property of the
// Foo/get response will change. The Foo/changes method allows a client to
// efficiently update the state of its Foo cache to match the new state on the
// server.
type MailboxChangesRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The current state of the client. This is the string that was
	// returned as the state argument in the Foo/get response. The server
	// will return the changes that have occurred since this state.
	SinceState string `json:"sinceState,omitempty"`

	// The maximum number of ids to return in the response. The server MAY
	// choose to return fewer than this value but MUST NOT return more. If
	// not given by the client, the server may choose how many to return.
	// If supplied by the client, the value MUST be a positive integer
	// greater than 0. If a value outside of this range is given, the
	// server MUST reject the call with an invalidArguments error.
	MaxChanges jmap.UnsignedInt `json:"maxChanges,omitempty"`
}

// This is a standard “/changes” method as described in [@!RFC8620], Section
// 5.2 but with one extra argument to the response: updatedProperties
type MailboxChangesResponse struct {
	// The id of the account used for the call.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// This is the sinceState argument echoed back; it’s the state from
	// which the server is returning changes.
	OldState string `json:"oldState,omitempty"`

	// This is the state the client will be in after applying the set of
	// changes to the old state.
	NewState string `json:"newState,omitempty"`

	// If true, the client may call Foo/changes again with the newState
	// returned to get further updates. If false, newState is the current
	// server state.
	HasMoreChanges bool `json:"hasMoreChanges,omitempty"`

	// An array of ids for records that have been created since the old
	// state.
	Created []jmap.ID `json:"created,omitempty"`

	// An array of ids for records that have been updated since the old
	// state.
	Updated []jmap.ID `json:"updated,omitempty"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []jmap.ID `json:"destroyed,omitempty"`

	// If only the “totalEmails”, “unreadEmails”, “totalThreads”, and/or
	// “unreadThreads” Mailbox properties have changed since the old state,
	// this will be the list of properties that may have changed. If the
	// server is unable to tell if only counts have changed, it MUST just
	// be null.
	UpdatedProperties []string `json:"updatedProperties,omitempty"`
}

// This is a standard “/query” method as described in [@!RFC8620], Section 5.5,
// but with the following additional request argument: sortAsTree, filterAsTree
type MailboxQueryRequest struct {
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
	Sort []*MailboxSortComparator `json:"sort,omitempty"`

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

	// If true, when sorting the query results and comparing Mailboxes A
	// and B:
	//
	// - If A is an ancestor of B, it always comes first regardless of the
	// sort comparators. Similarly, if A is descendant of B, then B always
	// comes first.
	//
	// - Otherwise, if A and B do not share a parentId, find the nearest
	// ancestors of each that do have the same parentId and compare the
	// sort properties on those Mailboxes instead.
	//
	// The result of this is that the Mailboxes are sorted as a tree
	// according to the parentId properties, with each set of children with
	// a common parent sorted according to the standard sort comparators.
	SortAsTree bool `json:"sortAsTree,omitempty"`

	// If true, a Mailbox is only included in the query if all its
	// ancestors are also included in the query according to the filter.
	FilterAsTree bool `json:"filterAsTree,omitempty"`
}

// Determines the set of Mailboxes returned in the results. If null, all
// objects in the account of this type are included in the results.
type MailboxFilter struct {
	// This MUST be one of the following strings: “AND” / “OR” / “NOT”
	Operator jmap.FilterOperator `json:"operator,omitempty"`

	// The conditions to evaluate against each record.
	Conditions []MailboxFilterCondition `json:"conditions,omitempty"`
}

func (f *MailboxFilter) AddCondition(cond MailboxFilterCondition) {
	f.Conditions = append(f.Conditions, cond)
}

// MailboxFilterCondition is an interface that represents FilterCondition
// objects. A filter condition object can be either a named struct, ie
// MailboxFilterConditionName, or a MailboxFilter itself. MailboxFilters can
// be used to create complex filtering ie return mailboxes which are subscribed
// and NOT named Inbox
type MailboxFilterCondition interface{}

type MailboxFilterConditionParentID struct {
	// The Mailbox parentId property must match the given value exactly.
	ParentID jmap.ID `json:"parentId,omitempty"`
}

type MailboxFilterConditionName struct {
	// The Mailbox name property contains the given string.
	Name string `json:"name,omitempty"`
}

type MailboxFilterConditionRole struct {
	// The Mailbox role property must match the given value exactly.
	Role Role `json:"role,omitempty"`
}

type MailboxFilterConditionHasAnyRole struct {
	// If true, a Mailbox matches if it has any non-null value for its role
	// property.
	HasAnyRole bool `json:"hasAnyRole,omitempty"`
}

type MailboxFilterConditionIsSubscribed struct {
	// The isSubscribed property of the Mailbox must be identical to the
	// value given to match the condition.
	IsSubscribed bool `json:"isSubscribed,omitempty"`
}

type MailboxSortComparator struct {
	// The name of the property on the Foo objects to compare. Servers MUST
	// support sorting by the following properties:
	// - sortOrder
	// - name
	Property string `json:"property,omitempty"`

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

type MailboxQueryResponse struct {
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

type MailboxQueryChangesRequest struct {
	// The id of the account to use.
	AccountID jmap.ID `json:"accountId,omitempty"`

	// The filter argument that was used with Foo/query.
	//
	// Each implementation must supply it's own Filter property
	Filter interface{} `json:"filter,omitempty"`

	// The sort argument that was used with Foo/query.
	//
	// Each implementation must supply it's own Sort property
	Sort []*MailboxSortComparator `json:"sort,omitempty"`

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
}

type MailboxQueryChangesResponse struct {
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

// This is a standard “/set” method as described in [@!RFC8620], Section 5.3,
// but with the following additional request argument: onDestroyRemoveEmails
type MailboxSetRequest struct {
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
	Create map[jmap.ID]*Mailbox `json:"create,omitempty"`

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

	// If false, any attempt to destroy a Mailbox that still has Emails in
	// it will be rejected with a mailboxHasEmail SetError. If true, any
	// Emails that were in the Mailbox will be removed from it, and if in
	// no other Mailboxes, they will be destroyed when the Mailbox is
	// destroyed.
	OnDestroyRemoveEmails bool `json:"onDestroyRemoveEmails,omitempty"`
}

type MailboxSetResponse struct {
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
	Created map[jmap.ID]interface{} `json:"created,omitempty"`

	// The keys in this map are the ids of all Foos that were successfully
	// updated.
	//
	// The value for each id is a Foo object containing any property that
	// changed in a way not explicitly requested by the PatchObject sent to
	// the server, or null if none. This lets the client know of any
	// changes to server-set or computed properties.
	//
	// This argument is null if no Foo objects were successfully updated.
	Updated map[jmap.ID]interface{} `json:"updated,omitempty"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []jmap.ID `json:"destroyed,omitempty"`
}
