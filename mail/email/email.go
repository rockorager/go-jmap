package email

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

func init() {
	jmap.RegisterCapability(&smimeVerify{})
	jmap.RegisterMethod("Email/get", newGetResponse)
	jmap.RegisterMethod("Email/changes", newChangesResponse)
	jmap.RegisterMethod("Email/query", newQueryResponse)
	jmap.RegisterMethod("Email/queryChanges", newQueryChangesResponse)
	jmap.RegisterMethod("Email/set", newSetResponse)
	jmap.RegisterMethod("Email/copy", newCopyResponse)
	jmap.RegisterMethod("Email/import", newImportResponse)
	jmap.RegisterMethod("Email/parse", newParseResponse)
}

// Representation of an RFC5322 message
// https://www.rfc-editor.org/rfc/rfc8621.html#section-4
type Email struct {
	ID jmap.ID `json:"id,omitempty"`

	BlobID jmap.ID `json:"blobId,omitempty"`

	ThreadID jmap.ID `json:"threadId,omitempty"`

	MailboxIDs map[jmap.ID]bool `json:"mailboxIds,omitempty"`

	Keywords map[string]bool `json:"keywords,omitempty"`

	Size uint64 `json:"size,omitempty"`

	ReceivedAt *time.Time `json:"receivedAt,omitempty"`

	Headers []*Header `json:"headers,omitempty"`

	MessageID []string `json:"messageId,omitempty"`

	InReplyTo []string `json:"inReplyTo,omitempty"`

	References []string `json:"references,omitempty"`

	Sender []*mail.Address `json:"sender,omitempty"`

	From []*mail.Address `json:"from,omitempty"`

	To []*mail.Address `json:"to,omitempty"`

	CC []*mail.Address `json:"cc,omitempty"`

	BCC []*mail.Address `json:"bcc,omitempty"`

	ReplyTo []*mail.Address `json:"replyTo,omitempty"`

	Subject string `json:"subject,omitempty"`

	SentAt *time.Time `json:"sentAt,omitempty"`

	BodyStructure *BodyPart `json:"bodyStructure,omitempty"`

	BodyValues map[string]*BodyValue `json:"bodyValues,omitempty"`

	TextBody []*BodyPart `json:"textBody,omitempty"`

	HTMLBody []*BodyPart `json:"htmlBody,omitempty"`

	Attachments []*BodyPart `json:"attachments,omitempty"`

	HasAttachment bool `json:"hasAttachment,omitempty"`

	Preview string `json:"preview,omitempty"`

	SMIMEStatus string `json:"smimeStatus,omitempty"`

	SMIMEStatusAtDelivery string `json:"smimeStatusAtDelivery,omitempty"`

	SMIMEErrors []string `json:"smimeErrors,omitempty"`

	SMIMEVerifiedAt *time.Time `json:"smimeVerifiedAt,omitempty"`
}

type AddressGroup struct {
	Name string `json:"name,omitempty"`

	Addresses []*mail.Address `json:"addresses,omitempty"`
}

type Header struct {
	Name string `json:"name,omitempty"`

	Value string `json:"value,omitempty"`
}

type BodyPart struct {
	PartID string `json:"partId,omitempty"`

	BlobID jmap.ID `json:"blobId,omitempty"`

	Size uint64 `json:"size,omitempty"`

	Headers []*Header `json:"headers,omitempty"`

	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`

	Charset string `json:"charset,omitempty"`

	Disposition string `json:"disposition,omitempty"`

	CID string `json:"cid,omitempty"`

	Language []string `json:"language,omitempty"`

	Location string `json:"location,omitempty"`

	SubParts []*BodyPart `json:"subParts,omitempty"`
}

type BodyValue struct {
	Value string `json:"value,omitempty"`

	IsEncodingProblem bool `json:"isEncodingProblem,omitempty"`

	IsTruncated bool `json:"isTruncated"`
}
