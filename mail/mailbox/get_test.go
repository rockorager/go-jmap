package mailbox

import (
	"encoding/json"
	"testing"

	"git.sr.ht/~rockorager/go-jmap"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	m := &Get{
		AccountID: "account-id",
	}
	req := &jmap.Request{}

	id := req.Invoke(m)
	assert.Equal(t, "0", id)

	data, err := json.Marshal(req)
	assert.NoError(t, err)
	expected := `{"using":["urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/get",{"accountId":"account-id"},"0"]]}`
	assert.Equal(t, expected, string(data))

	t.Run("manual", func(t *testing.T) {
		m := &Get{
			AccountID: "account-id",
		}
		req = &jmap.Request{
			Using: []string{MailCapability},
			Calls: []*jmap.Invocation{
				{
					Name:   "Mailbox/get",
					Args:   m,
					CallID: "manual",
				},
			},
		}
		data, err := json.Marshal(req)
		assert.NoError(t, err)
		expected := `{"using":["urn:ietf:params:jmap:mail"],"methodCalls":[["Mailbox/get",{"accountId":"account-id"},"manual"]]}`
		assert.Equal(t, expected, string(data))
	})
}
