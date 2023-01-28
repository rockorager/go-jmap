package jmap

import (
	"encoding/json"
	"time"
)

// Date is a time.Time that is serialized to JSON in RFC 3339 format (without
// the fractional part).
type Date time.Time

func (d Date) MarshalText() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339))
	b = time.Time(d).AppendFormat(b, time.RFC3339)
	return b, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*(*time.Time)(d), err = time.Parse(time.RFC3339, s)
	if err != nil {
		*(*time.Time)(d) = time.Time{}
	}
	return nil
}

// Date is a time.Time that is serialized to JSON in RFC 3339 format (without
// the fractional part) in UTC timezone.
//
// If UTCDate value is not in UTC, it will be converted to UTC during
// serialization.
type UTCDate time.Time

func (d UTCDate) MarshalText() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = time.Time(d).UTC().AppendFormat(b, time.RFC3339)
	return b, nil
}

func (d *UTCDate) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*(*time.Time)(d), err = time.ParseInLocation(time.RFC3339, s, time.UTC)
	return err
}
