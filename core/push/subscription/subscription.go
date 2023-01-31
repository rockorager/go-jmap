package subscription

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

func init() {
	jmap.RegisterMethod("PushSubscription/get", newGetResponse)
	jmap.RegisterMethod("PushSubscription/set", newSetResponse)
}

// A PushSubscription object
type PushSubscription struct {
	// The ID of the push subscription
	//
	// immutable;server-set
	ID jmap.ID `json:"id,omitempty"`

	// An ID that uniquely identifies the client + device the subscription
	// is running on
	//
	// immutable
	DeviceClientID string `json:"deviceClientId,omitempty"`

	// An absolute URL where the JMAP server will POST the data for the push
	// message. This must start with "https://"
	//
	// immutable
	URL string `json:"url,omitempty"`

	// Client-generated encryption keys. If specified, the server will
	// encrypt the push data
	Keys *Key `json:"keys,omitempty"`

	// This must be null or omitted when the subscription is created. The
	// JMAP server will generate a code and send it in a push message. The
	// client must then update this field with that code
	VerificationCode string `json:"verificationCode,omitempty"`

	// The time this subscription expires, if specified. If not specified,
	// the subscription does not expire, however the server may specify a
	// time
	//
	// Must be in UTC
	Expires *time.Time `json:"expires,omitempty"`

	// A list of type changes the client is subscribing to, using the same
	// keys as a TypeState object
	Types []string `json:"types,omitempty"`
}

// A Push Subscription Encryption key. This key must be a P-256 ECDH key
type Key struct {
	// The public key, base64 encoded
	Public string `json:"p256dh"`
	// The authentication secret, base64 encoded
	Auth string `json:"auth"`
}
