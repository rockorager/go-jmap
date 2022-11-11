package mail

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"git.sr.ht/~rockorager/go-jmap"
	"github.com/stretchr/testify/assert"
)

// Test whether our service creates correct json requests
func TestMailboxRequests(t *testing.T) {
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
	t.Run("Mailbox/get", func(t *testing.T) {
		get := mail.Mailboxes.Get(username)
		get.IDs("abc")
		get.Properties("name")
		get.CallID("call1")
		get.Do()
		io.Copy(buf, c.req.Body)
		expected := `{"using":["urn:ietf:params:jmap:core","urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/get",{"accountId":"A13824","ids":["abc"],"properties":["name"]},"call1"]]}`
		assert.Equal(t, expected, buf.String())
		buf.Reset()
	})

	t.Run("Mailbox/changes", func(t *testing.T) {
		changes := mail.Mailboxes.Changes(username, "123")
		changes.MaxChanges(4)
		changes.CallID("call1")
		changes.Do()
		io.Copy(buf, c.req.Body)
		expected := []byte(`{"using":["urn:ietf:params:jmap:core","urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/changes",{"accountId":"A13824","sinceState":"123","maxChanges":4},"call1"]]}`)
		assert.Equal(t, expected, buf.Bytes())
		buf.Reset()
	})

	t.Run("Mailbox/query", func(t *testing.T) {
		query := mail.Mailboxes.Query(username)
		query.SortAsTree(true)
		query.FilterAsTree(true)
		query.CallID("call1")
		filter := mail.Mailboxes.NewFilter()
		cond1 := &MailboxFilterConditionName{
			Name: "Inbox",
		}
		filter.AddCondition(cond1)
		query.Filter(filter)
		sort := &MailboxSortComparator{
			Property:    "name",
			IsAscending: true,
		}
		query.Sort(sort)
		query.Do()
		io.Copy(buf, c.req.Body)
		expected := []byte(`{"using":["urn:ietf:params:jmap:core","urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/query",{"accountId":"A13824","filter":{"name":"Inbox"},"sort":[{"property":"name","isAscending":true}],"sortAsTree":true,"filterAsTree":true},"call1"]]}`)
		assert.Equal(t, expected, buf.Bytes())
		buf.Reset()
	})

	t.Run("Mailbox/queryChanges", func(t *testing.T) {
		queryChanges := mail.Mailboxes.QueryChanges(username, "123")
		filter := mail.Mailboxes.NewFilter()
		cond1 := &MailboxFilterConditionName{
			Name: "Inbox",
		}
		cond2 := &MailboxFilterConditionIsSubscribed{
			IsSubscribed: true,
		}
		filter.AddCondition(cond1)
		filter.AddCondition(cond2)
		sort := &MailboxSortComparator{
			Property:    "name",
			IsAscending: true,
		}
		filter.Operator = jmap.FilterOperatorAND
		queryChanges.Filter(filter)
		queryChanges.Sort(sort)
		queryChanges.CallID("call1")
		queryChanges.MaxChanges(2)
		queryChanges.UpToID("124")
		queryChanges.Do()
		io.Copy(buf, c.req.Body)
		expected := []byte(`{"using":["urn:ietf:params:jmap:core","urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/queryChanges",{"accountId":"A13824","filter":{"operator":"AND","conditions":[{"name":"Inbox"},{"isSubscribed":true}]},"sort":[{"property":"name","isAscending":true}],"sinceQueryState":"123","maxChanges":2,"upToId":"124"},"call1"]]}`)
		assert.Equal(t, expected, buf.Bytes())
		buf.Reset()
	})

	t.Run("Mailbox/set", func(t *testing.T) {
		set := mail.Mailboxes.Set(username)
		set.CallID("call1")
		mbox := &Mailbox{
			Name: "Inbox",
		}
		set.Create("123", mbox)
		set.OnDestroyRemoveEmails(true)
		set.Do()
		io.Copy(buf, c.req.Body)
		expected := `{"using":["urn:ietf:params:jmap:core","urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/set",{"accountId":"A13824","create":{"123":{"name":"Inbox"}},"onDestroyRemoveEmails":true},"call1"]]}`
		assert.Equal(t, expected, buf.String())
		buf.Reset()
	})
}

// Test whether our service decodes a response properly
func TestMailboxResponses(t *testing.T) {
	c := &TestHTTPClient{
		statusCode: 200,
		resp:       openFile(t, "./test-data/session-response.json"),
	}

	username := "john@example.com"
	mail, err := NewServiceWithClient(c, "", "")
	if err != nil {
		t.Fatalf("error creating test client")
	}

	// buf := bytes.NewBuffer(nil)
	// RequestError
	t.Run("RequestError", func(t *testing.T) {
		r := strings.NewReader(`{
  "type": "urn:ietf:params:jmap:error:limit",
  "limit": "maxSizeRequest",
  "status": 400,
  "detail": "The request is larger than the server is willing to process."
}`)
		c.resp = io.NopCloser(r)
		c.statusCode = 400
		get := mail.Mailboxes.Get(username)
		_, err = get.Do()
		assert.NotNil(t, err)
		assert.IsType(t, jmap.RequestError{}, err)
	})

	// MethodError
	t.Run("MethodError", func(t *testing.T) {
		r := strings.NewReader(`{"sessionState": "123","methodResponses":[["error",{"type":"unknownMethod"},"call1"]]}`)
		c.statusCode = 200
		c.resp = io.NopCloser(r)
		get := mail.Mailboxes.Get(username)
		_, err = get.Do()
		assert.NotNil(t, err)
		assert.IsType(t, jmap.MethodErrorArgs{}, err)
		e := err.(jmap.MethodErrorArgs)
		assert.Equal(t, jmap.ErrorCode("unknownMethod"), e.Type)
	})

	t.Run("Method/get", func(t *testing.T) {
		// TODO Finish response tests for each method
	})
}
