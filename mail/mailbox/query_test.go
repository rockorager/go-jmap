package mailbox

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	assert := assert.New(t)
	query := &Query{
		Account: "xyz",
		Filter: &FilterCondition{
			Name: "Inbox",
		},
		Sort: []*SortComparator{
			{
				Property: "name",
			},
		},
		Limit: 10,
	}
	data, err := json.Marshal(query)
	assert.NoError(err)
	expected := `{"accountId":"xyz","filter":{"name":"Inbox"},"sort":[{"property":"name","isAscending":false}],"limit":10}`
	assert.Equal(expected, string(data))
}
