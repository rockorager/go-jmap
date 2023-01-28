package mailbox

import (
	"encoding/json"
	"testing"

	"git.sr.ht/~rockorager/go-jmap"
	"github.com/stretchr/testify/assert"
)

func TestChanges(t *testing.T) {
	m := &Changes{
		AccountID: "account-id",
		SinceState: "1234",
	}
	req := &jmap.Request{}

	id := req.Invoke(m)
	assert.Equal(t, "0", id)

	data, err := json.Marshal(req)
	assert.NoError(t, err)
	expected := `{"using":["urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/changes",{"accountId":"account-id","sinceState":"1234"},"0"]]}`
	assert.Equal(t, expected, string(data))

	t.Run("manual", func(t *testing.T) {
		m := &Changes{
			AccountID: "account-id",
			SinceState: "1234",
		}
		req = &jmap.Request{
			Using: []string{MailCapability},
			Calls: []*jmap.Invocation{
				{
					Name:   "Mailbox/changes",
					Args:   m,
					CallID: "manual",
				},
			},
		}
		data, err := json.Marshal(req)
		assert.NoError(t, err)
		expected := `{"using":["urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/changes",{"accountId":"account-id","sinceState":"1234"},"manual"]]}`
		assert.Equal(t, expected, string(data))
	})
}
