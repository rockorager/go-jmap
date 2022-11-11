package mail

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test whether our service creates correct json requests
func TestThreadRequests(t *testing.T) {
	c := &TestHTTPClient{
		statusCode: 200,
		resp:       openFile(t, "./test-data/session-response.json"),
	}

	username := "john@example.com"
	mail, err := NewServiceWithClient(c, "", "")
	if err != nil {
		t.Fatalf("error creating test client")
	}

	buf := bytes.NewBuffer(nil)
	// Mailbox/get
	t.Run("Thread/get", func(t *testing.T) {
		get := mail.Threads.Get(username)
		get.IDs("abc")
		get.Properties("name")
		get.CallID("call1")
		get.Do()
		io.Copy(buf, c.req.Body)
		expected := `{"using":["urn:ietf:params:jmap:core","urn:ietf:params:jmap:mail"],"methodCalls":[["Thread/get",{"accountId":"A13824","ids":["abc"],"properties":["name"]},"call1"]]}`
		assert.Equal(t, expected, buf.String())
		buf.Reset()
	})

	t.Run("Thread/changes", func(t *testing.T) {
		changes := mail.Threads.Changes(username, "123")
		changes.MaxChanges(4)
		changes.CallID("call1")
		changes.Do()
		io.Copy(buf, c.req.Body)
		expected := `{"using":["urn:ietf:params:jmap:core","urn:ietf:params:jmap:mail"],"methodCalls":[["Thread/changes",{"accountId":"A13824","sinceState":"123","maxChanges":4},"call1"]]}`
		assert.Equal(t, expected, buf.String())
		buf.Reset()
	})
}

// Test whether our service decodes a response properly
func TestThreadResponses(t *testing.T) {
}
