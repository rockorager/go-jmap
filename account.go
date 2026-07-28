package jmap

import "encoding/json"

// An account is a collection of data the authenticated user has access to
//
// See RFC 8620 section 1.6.2 for details.
type Account struct {
	// The ID of the account
	ID string `json:"-"`

	// A user-friendly string to show when presenting content from this
	// account, e.g. the email address representing the owner of the account.
	Name string `json:"name"`

	// True if this account belongs to the authenticated user
	IsPersonal bool `json:"isPersonal"`

	IsReadOnly bool `json:"isReadOnly"`

	// The set of capability URIs for the methods supported in this account.
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
