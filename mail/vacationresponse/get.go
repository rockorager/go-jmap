package vacationresponse

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// An Identity/get request
type Get struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId,omitempty"`

	// The IDs of Identity objects to return. Leave blank to return all,
	// subject to the MaxObjectsInGet limit of the server
	IDs []jmap.ID `json:"ids,omitempty"`

	// Only the supplied properties will be returned
	Properties []string `json:"properties,omitempty"`
}

func (m *Get) Name() string { return "VacationResponse/get" }

func (m *Get) Requires() []jmap.URI { return []jmap.URI{mail.URI, URI} }

// This is a standard “/get” method as described in [@!RFC8620], Section 5.1.
type GetResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId,omitempty"`

	// State for all Identity objects on the server for this account
	State string `json:"state,omitempty"`

	// The Identity objects requested
	List []*VacationResponse `json:"list,omitempty"`

	// Slice of objects not found. Only present if specific IDs were
	// requested
	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
