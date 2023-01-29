package jmap

// BlobInfo is the object returned in response to blob upload.
type BlobInfo struct {
	// The id of the account used for the call.
	AccountID string `json:"accountId"`

	// The id representing the binary data uploaded. The data for this id is
	// immutable. The id only refers to the binary data, not any metadata.
	BlobID string `json:"blobId"`

	// The media type of the file (as specified in RFC 6838, section 4.2) as
	// set in the Content-Type header of the upload HTTP request.
	Type string `json:"type"`

	// The size of the file in octets.
	Size uint64 `json:"size"`
}

// Binary data may be copied between two different accounts using the Blob/copy
// method rather than having to download and then reupload on the client.
type BlobCopy struct {
	// The ID of the account to copy blobs from
	FromAccountID string `json:"fromAccountId,omitempty"`

	// The ID of the account to copy blobs to
	AccountID string `json:"accountId,omitempty"`

	// A list of IDs of blobs to copy
	BlobIDs []string `json:"blobIds,omitempty"`
}

func (m *BlobCopy) Name() string { return "Blob/copy" }

func (m *BlobCopy) Uses() string { return "" }

func (m *BlobCopy) NewResponse() interface{} { return &BlobCopy{} }

type BlobCopyResponse struct {
	// The ID of the account blobs were copied from
	FromAccountID string `json:"fromAccountId,omitempty"`

	// The ID of the account blobs were copied to
	AccountID string `json:"accountId,omitempty"`

	// A map of the blobId in the fromAccount to the ID of the blob in the
	// account it was copied to. Map is null if no blobs were copied
	Copied map[string]string `json:"blobIds,omitempty"`

	// A map of blobId to a SetError object for each blob that failed to be
	// copied, or null if none.
	NotCopied map[id]SetError `json:"notCopied,omitempty"`
}
