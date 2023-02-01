package jmap

import (
	"encoding/json"
)

type Session struct {
	// An object specifying the capabilities of this server. Each key is a URI
	// for a capability supported by the server. The value for each of these
	// keys is an object with further information about the server’s
	// capabilities in relation to that capability.
	Capabilities map[URI]Capability `json:"-"`

	RawCapabilities map[URI]json.RawMessage `json:"capabilities"`

	// A map of account id to Account object for each account the user has
	// access to.
	Accounts map[ID]Account `json:"accounts"`

	// A map of capability URIs (as found in Capabilities) to the
	// account id to be considered the user’s main or default account for data
	// pertaining to that capability.
	PrimaryAccounts map[URI]ID `json:"primaryAccounts"`

	// The username associated with the given credentials, or the empty string
	// if none.
	Username string `json:"username"`

	// The URL to use for JMAP API requests.
	APIURL string `json:"apiUrl"`

	// The URL endpoint to use when downloading files, in RFC 6570 URI
	// Template (level 1) format.
	DownloadURL string `json:"downloadUrl"`

	// The URL endpoint to use when uploading files, in RFC 6570 URI
	// Template (level 1) format.
	UploadURL string `json:"uploadUrl"`

	// The URL to connect to for push events, as described in section 7.3, in
	// RFC 6570 URI Template (level 1) format.
	EventSourceURL string `json:"eventSourceUrl"`

	// A string representing the state of this object on the server. If the
	// value of any other property on the session object changes, this string
	// will change.
	//
	// The current value is also returned on the API Response object, allowing
	// clients to quickly determine if the session information has changed
	// (e.g. an account has been added or removed) and so they need to refetch
	// the object.
	State string `json:"state"`
}

type session Session

func (s *Session) UnmarshalJSON(data []byte) error {
	raw := (*session)(s)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	s.Capabilities = make(map[URI]Capability)
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
		s.Capabilities[key] = newCap
	}

	return nil
}
