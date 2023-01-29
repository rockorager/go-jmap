package jmap

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sessionBlob = `{
  "capabilities": {
    "urn:ietf:params:jmap:core": {
      "maxSizeUpload": 50000000,
      "maxConcurrentUpload": 8,
      "maxSizeRequest": 10000000,
      "maxConcurrentRequest": 8,
      "maxCallsInRequest": 32,
      "maxObjectsInGet": 256,
      "maxObjectsInSet": 128,
      "collationAlgorithms": [
        "i;ascii-numeric",
        "i;ascii-casemap",
        "i;unicode-casemap"
      ]
    },
    "test:jmap:capability": {
      "testValue": 500
    },
    "urn:ietf:params:jmap:mail": {},
    "urn:ietf:params:jmap:contacts": {},
    "https://example.com/apis/foobar": {
      "maxFoosFinangled": 42
    }
  },
  "accounts": {
    "A13824": {
      "name": "john@example.com",
      "isPersonal": true,
      "isReadOnly": false,
      "accountCapabilities": {
        "urn:ietf:params:jmap:mail": {
          "maxMailboxesPerEmail": null,
          "maxMailboxDepth": 10
        },
        "urn:ietf:params:jmap:contacts": {
        }
      }
    },
    "A97813": {
      "name": "jane@example.com",
      "isPersonal": false,
      "isReadOnly": true,
      "accountCapabilities": {
        "urn:ietf:params:jmap:mail": {
          "maxMailboxesPerEmail": 1,
          "maxMailboxDepth": 10
        }
      }
    }
  },
  "primaryAccounts": {
    "urn:ietf:params:jmap:mail": "A13824",
    "urn:ietf:params:jmap:contacts": "A13824"
  },
  "username": "john@example.com",
  "apiUrl": "https://jmap.example.com/api/",
  "downloadUrl": "https://jmap.example.com/download/{accountId}/{blobId}/{name}?accept={type}",
  "uploadUrl": "https://jmap.example.com/upload/{accountId}/",
  "eventSourceUrl": "https://jmap.example.com/eventsource/?types={types}&closeafter={closeafter}&ping={ping}",
  "state": "75128aab4b1b"
}`

func TestSessionUnmarshal(t *testing.T) {
	RegisterCapability(&testCapability{})
	assert := assert.New(t)
	s := &Session{}
	err := json.Unmarshal([]byte(sessionBlob), s)
	assert.NoError(err)

	testCap := s.Capabilities["test:jmap:capability"].(*testCapability)
	assert.Equal(500, testCap.TestValue)
	assert.Equal("john@example.com", s.Accounts["A13824"].Name)
}

func TestSessionMarshal(t *testing.T) {
	assert := assert.New(t)
	s := Session{}
	err := json.Unmarshal([]byte(sessionBlob), &s)
	assert.NoError(err)

	blob, err := json.MarshalIndent(s, "", "  ")
	assert.NoError(err)

	// We can't just compare []byte because order of fields may be different.
	var original, remarshaled map[string]interface{}
	err = json.Unmarshal([]byte(sessionBlob), &original)
	assert.NoError(err)
	err = json.Unmarshal(blob, &remarshaled)
	assert.NoError(err)

	assert.Equal(original, remarshaled)
}

type testCapability struct {
	TestValue int `json:"testValue"`
}

func (tc *testCapability) URI() string {
	return "test:jmap:capability"
}

func (tc *testCapability) New() Capability {
	return &testCapability{}
}
