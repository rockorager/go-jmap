package subscription

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

func init() {
	jmap.RegisterMethod("PushSubscription/get", newGetResponse)
	jmap.RegisterMethod("PushSubscription/set", newSetResponse)
}

// Server side push notification
// https://www.rfc-editor.org/rfc/rfc8620.html#section-7.2
type PushSubscription struct {
	ID jmap.ID `json:"id,omitempty"`

	DeviceClientID string `json:"deviceClientId,omitempty"`

	URL string `json:"url,omitempty"`

	Keys *Key `json:"keys,omitempty"`

	VerificationCode string `json:"verificationCode,omitempty"`

	Expires *time.Time `json:"expires,omitempty"`

	Types []string `json:"types,omitempty"`
}

// A Push Subscription Encryption key. This key must be a P-256 ECDH key
type Key struct {
	// The public key, base64 encoded
	Public string `json:"p256dh"`
	// The authentication secret, base64 encoded
	Auth string `json:"auth"`
}
