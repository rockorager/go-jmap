package mailbox

import "git.sr.ht/~rockorager/go-jmap"

func init() {
	jmap.RegisterMethod("Mailbox/get", newGetResponse)
	jmap.RegisterMethod("Mailbox/changes", newChangesResponse)
	jmap.RegisterMethod("Mailbox/query", newQueryResponse)
	jmap.RegisterMethod("Mailbox/queryChanges", newQueryChangesResponse)
	jmap.RegisterMethod("Mailbox/set", newSetResponse)
}

// Named set of Email objects. Can be viewed as a folder or a label.
// An email must be part of at least one Mailbox.
// https://www.rfc-editor.org/rfc/rfc8621.html#section-2
type Mailbox struct {
	ID jmap.ID `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ParentID jmap.ID `json:"parentId,omitempty"`

	Role Role `json:"role,omitempty"`

	SortOrder uint64 `json:"sortOrder,omitempty"`

	TotalEmails uint64 `json:"totalEmails,omitempty"`

	UnreadEmails uint64 `json:"unreadEmails,omitempty"`

	TotalThreads uint64 `json:"totalThreads,omitempty"`

	UnreadThreads uint64 `json:"unreadThreads,omitempty"`

	Rights *Rights `json:"myRights,omitempty"`

	IsSubscribed bool `json:"isSubscribed,omitempty"`
}

// Access Control Lists (ACLs)
type Rights struct {
	MayReadItems bool `json:"mayReadItems,omitempty"`

	MayAddItems bool `json:"mayAddItems,omitempty"`

	MayRemoveItems bool `json:"mayRemoveItems,omitempty"`

	MaySetSeen bool `json:"maySetSeen,omitempty"`

	MaySetKeywords bool `json:"maySetKeywords,omitempty"`

	MayCreateChild bool `json:"mayCreateChild,omitempty"`

	MayRename bool `json:"mayRename,omitempty"`

	MayDelete bool `json:"mayDelete,omitempty"`

	MaySubmit bool `json:"maySubmit,omitempty"`
}

// Identifies Mailboxes that have a particular common purpose (e.g., the
// “inbox”), regardless of the name property (which may be localised).
type Role string

const (
	RoleAll Role = "all"

	RoleArchive Role = "archive"

	RoleDrafts Role = "drafts"

	RoleFlagged Role = "flagged"

	RoleHasChildren Role = "haschildren"

	RoleHasNoChildren Role = "hasnochildren"

	RoleImportant Role = "important"

	RoleInbox Role = "inbox"

	RoleJunk Role = "junk"

	RoleMarked Role = "marked"

	RoleNoInferiors Role = "noinferiors"

	RoleNonExistent Role = "nonexistent"

	RoleNoSelect Role = "noselect"

	RoleRemote Role = "remote"

	RoleSent Role = "sent"

	RoleSubscribed Role = "subscribed"

	RoleTrash Role = "trash"

	RoleUnmarked Role = "unmarked"
)
