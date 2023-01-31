package vacationresponse

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

func init() {
	jmap.RegisterMethod("VacationResponse/get", newGetResponse)
	jmap.RegisterMethod("VacationResponse/set", newSetResponse)
}

type VacationResponse struct {
	// The ID of the object. There is only ever one VacationResponse object,
	// and it's ID is constant: "singleton"
	//
	// immutable;server-set;constant
	ID string `json:"id,omitempty"`

	// If the response is enabled
	IsEnabled bool `json:"isEnabled,omitempty"`

	// If IsEnabled is true, the response is active for messages received
	// after this time. Must be UTC
	FromDate *time.Time `json:"fromDate,omitempty"`

	// If IsEnabled is true, the response is active for messages received
	// before this time. Must be UTC
	ToDate *time.Time `json:"toDate,omitempty"`

	// The subject for the response. If null, the server MAY set a suitable
	// subject
	Subject *string `json:"subject,omitempty"`

	// The plaintext body to send in the response
	TextBody *string `json:"textBody,omitempty"`

	// The HTML body to send in the response
	HTMLBody *string `json:"htmlBody,omitempty"`
}
