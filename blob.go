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
