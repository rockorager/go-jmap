package vacationresponse

import (
	"encoding/json"
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

const URI jmap.URI = "urn:ietf:params:jmap:vacationresponse"

func init() {
	jmap.RegisterCapability(&Capability{})
	jmap.RegisterMethod("VacationResponse/get", newGetResponse)
	jmap.RegisterMethod("VacationResponse/set", newSetResponse)
}

// The VacationResponse capability is an empty object
type Capability struct{}

func (m *Capability) URI() jmap.URI { return URI }

func (m *Capability) New() jmap.Capability { return &Capability{} }

// Automatic reply when a message is delivered to the mail store
// https://www.rfc-editor.org/rfc/rfc8621.html#section-8
type VacationResponse struct {
	ID string `json:"id,omitempty"`

	IsEnabled bool `json:"isEnabled,omitempty"`

	FromDate *time.Time `json:"fromDate,omitempty"`

	ToDate *time.Time `json:"toDate,omitempty"`

	Subject *string `json:"subject,omitempty"`

	TextBody *string `json:"textBody,omitempty"`

	HTMLBody *string `json:"htmlBody,omitempty"`
}

func (v *VacationResponse) MarshalJson() ([]byte, error) {
	if v.FromDate != nil && v.FromDate.Location() != time.UTC {
		utc := v.FromDate.UTC()
		v.FromDate = &utc
	}
	if v.ToDate != nil && v.ToDate.Location() != time.UTC {
		utc := v.ToDate.UTC()
		v.ToDate = &utc
	}
	type Alias VacationResponse
	return json.Marshal((*Alias)(v))
}
