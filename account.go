package jmap

import "encoding/json"

// An account is a collection of data. A single account may contain an
// arbitrary set of data types, for example a collection of mail, contacts and
// calendars.
//
// See draft-ietf-jmap-core-17, section 1.6.2 for details.
// The documentation is taked from draft-ietf-jmap-core-17, section 2.
type Account struct {
	// The ID of the account
	ID string `json:"-"`

	// A user-friendly string to show when presenting content from this
	// account, e.g. the email address representing the owner of the account.
	Name string `json:"name"`

	// This is true if the account belongs to the authenticated user, rather
	// than a group account or a personal account of another user that has been
	// shared with them.
	IsPersonal bool `json:"isPersonal"`

	// This is true if the entire account is read-only.
	IsReadOnly bool `json:"isReadOnly"`

	// The set of capability URIs for the methods supported in this account.
	// Each key is a URI for a capability that has methods you can use with
	// this account. The value for each of these keys is an object with further
	// information about the account’s permissions and restrictions with
	// respect to this capability, as defined in the capability’s
	// specification.
	Capabilities map[URI]Capability `json:"-"`

	// The raw JSON of accountCapabilities
	RawCapabilities map[URI]json.RawMessage `json:"accountCapabilities"`
}

type account Account

func (a *Account) UnmarshalJSON(data []byte) error {
	raw := (*account)(a)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	a.Capabilities = make(map[URI]Capability)
	for key, cap := range capabilities {
		rawCap, ok := raw.RawCapabilities[key]
		if !ok {
			continue
		}
		newCap := cap.New()
		err := json.Unmarshal(rawCap, newCap)
		if err != nil {
			return err
		}
		a.Capabilities[key] = newCap
	}

	return nil
}
