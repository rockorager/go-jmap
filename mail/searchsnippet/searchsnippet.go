package searchsnippet

import "git.sr.ht/~rockorager/go-jmap"

func init() {
	jmap.RegisterMethod("SearchSnippet/get", newGetResponse)
}

// Search preview snippet
// https://www.rfc-editor.org/rfc/rfc8621.html#section-5
type SearchSnippet struct {
	Email jmap.ID `json:"emailId,omitempty"`

	Subject string `json:"subject,omitempty"`

	Preview string `json:"preview,omitempty"`
}
