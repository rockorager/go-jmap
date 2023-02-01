package subscription

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/core"
)

type Set struct {
	// A map of a creation id (a temporary id set by the client) to Foo
	// objects, or null if no objects are to be created.
	//
	// The Foo object type definition may define default values for
	// properties. Any such property may be omitted by the client.
	//
	// The client MUST omit any properties that may only be set by the
	// server (for example, the id property on most object types).
	Create map[jmap.ID]*PushSubscription `json:"create,omitempty"`

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
	Update map[jmap.ID]*jmap.Patch `json:"update,omitempty"`

	// A list of ids for Foo objects to permanently delete, or null if no
	// objects are to be destroyed.
	Destroy []jmap.ID `json:"destroy,omitempty"`
}

func (m *Set) Name() string { return "PushSubscription/set" }

func (m *Set) Requires() []jmap.URI { return []jmap.URI{core.URI} }

type SetResponse struct {
	// A map of the creation id to an object containing any properties of
	// the created Foo object that were not sent by the client. This
	// includes all server-set properties (such as the id in most object
	// types) and any properties that were omitted by the client and thus
	// set to a default by the server.
	//
	// This argument is null if no Foo objects were successfully created.
	Created map[jmap.ID]*PushSubscription `json:"created,omitempty"`

	// The keys in this map are the ids of all Foos that were successfully
	// updated.
	//
	// The value for each id is a Foo object containing any property that
	// changed in a way not explicitly requested by the PatchObject sent to
	// the server, or null if none. This lets the client know of any
	// changes to server-set or computed properties.
	//
	// This argument is null if no Foo objects were successfully updated.
	Updated map[jmap.ID]*PushSubscription `json:"updated,omitempty"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []jmap.ID `json:"destroyed,omitempty"`

	// A map of ID to a SetError for each record that failed to be created
	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`

	// A map of ID to a SetError for each record that failed to be updated
	NotUpdated map[jmap.ID]*jmap.SetError `json:"notUpdated,omitempty"`

	// A map of ID to a SetError for each record that failed to be destroyed
	NotDestroyed map[jmap.ID]*jmap.SetError `json:"notDestroyed,omitempty"`
}

func newSetResponse() jmap.MethodResponse { return &SetResponse{} }
