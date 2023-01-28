package searchsnippet

type Get struct {
	// The id of the account to use.
	AccountID string `json:"accountId,omitempty"`

	// Determines the set of Foos returned in the results. If null, all
	// objects in the account of this type are included in the results.
	//
	// Each implementation must implement it's own Filter
	Filter interface{} `json:"filter,omitempty"`

	// The ids of the Emails to fetch snippets for.
	EmailIDs []string `json:"emailIds,omitempty"`
}

func (m *Get) Name() string {
	return "Mailbox/get"
}

func (m *Get) Uses() string {
	return MailCapability
}

func (m *Get) NewResponse() interface{} {
	return &GetResponse{}
}

type GetResponse struct {
	// The id of the account used for the call
	AccountID string `json:"accountId,omitempty"`

	// An array of SearchSnippet objects for the requested Email ids.
	// This may not be in the same order as the ids that were in the
	// request.
	List []*SearchSnippet `json:"list,omitempty"`

	// An array of Email ids requested that could not be found, or null
	// if all ids were found.
	NotFound []string `json:"notFound,omitempty"`
}
