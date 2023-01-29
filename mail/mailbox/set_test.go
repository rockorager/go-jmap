package mailbox

import (
	"encoding/json"
	"testing"

	"git.sr.ht/~rockorager/go-jmap"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	assert := assert.New(t)
	set := &Set{
		Account: "xyz",
		Update: map[jmap.ID]*jmap.Patch{
			"mailbox-id": {
				Path:  "name",
				Value: "New Name",
			},
		},
	}
	data, err := json.Marshal(set)
	assert.NoError(err)
	expected := `{"accountId":"xyz","update":{"mailbox-id":{"name":"New Name"}}}`
	assert.Equal(expected, string(data))

	set = &Set{
		Account: "xyz",
		Update: map[jmap.ID]*jmap.Patch{
			"mailbox-id": {
				Path:  "parentId",
				Value: nil,
			},
		},
	}
	data, err = json.Marshal(set)
	assert.NoError(err)
	expected = `{"accountId":"xyz","update":{"mailbox-id":{"parentId":null}}}`
	assert.Equal(expected, string(data))
}
