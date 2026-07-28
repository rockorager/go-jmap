package jmap

import (
	"encoding/json"
)

type Session struct {
	// Capabilities specifies the capabililities the server has.
	Capabilities map[URI]Capability `json:"-"`

	RawCapabilities map[URI]json.RawMessage `json:"capabilities"`

	Accounts map[ID]Account `json:"accounts"`

	// PrimaryAccounts maps a Capability to the primary account associated
	// with it
	PrimaryAccounts map[URI]ID `json:"primaryAccounts"`

	// The username associated with the given credentials
	Username string `json:"username"`

	// The URL to use for JMAP API requests.
	APIURL string `json:"apiUrl"`

	// The URL endpoint to use when downloading files
	DownloadURL string `json:"downloadUrl"`

	// The URL endpoint to use when uploading files
	UploadURL string `json:"uploadUrl"`

	// The URL to connect to for push events
	EventSourceURL string `json:"eventSourceUrl"`

	// A string representing the state of this object on the server
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
