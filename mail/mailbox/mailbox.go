package mailbox

import "git.sr.ht/~rockorager/go-jmap"

const MailCapability = "urn:ietf:params:jmap:mail"

func init() {
	jmap.RegisterMethods(
		&Get{},
		&Changes{},
		&Query{},
		&QueryChanges{},
		&Set{},
	)
}

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
	SortOrder uint64 `json:"sortOrder,omitempty"`

	// The number of Emails in this Mailbox.
	TotalEmails uint64 `json:"totalEmails,omitempty"`

	// The number of Emails in this Mailbox that have neither the $seen
	// keyword nor the $draft keyword.
	UnreadEmails uint64 `json:"unreadEmails,omitempty"`

	// The number of Threads where at least one Email in the Thread is in
	// this Mailbox.
	TotalThreads uint64 `json:"totalThreads,omitempty"`

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
	UnreadThreads uint64 `json:"unreadThreads,omitempty"`

	// The set of rights (Access Control Lists (ACLs)) the user has in
	// relation to this Mailbox. These are backwards compatible with IMAP
	// ACLs, as defined in [@!RFC4314].
	Rights *Rights `json:"myRights,omitempty"`

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
type Rights struct {
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
