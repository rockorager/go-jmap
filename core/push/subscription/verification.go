package subscription

// A PushVerification object is sent by the server to the created Subscriptions'
// URL. This object contains the ID of the subscription and the Verification
// code required. The Client must update the PushSubscription using a Set method
// with the code.
type Verification struct {
	// The MUST be "PushVerification"
	Type string `json:"@type,omitempty"`

	// The ID of the Push Subscription that was created
	SubscriptionID string `json:"pushSubscriptionId,omitempty"`

	// The verification code to add to the subscription
	Code string `json:"verificationCode,omitempty"`
}
