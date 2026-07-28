package subscription

// A PushVerification object is sent by the server to the created Subscriptions'
// URL.
type Verification struct {
	// The MUST be "PushVerification"
	Type string `json:"@type,omitempty"`

	// The ID of the Push Subscription that was created
	SubscriptionID string `json:"pushSubscriptionId,omitempty"`

	// The verification code to add to the subscription
	Code string `json:"verificationCode,omitempty"`
}
