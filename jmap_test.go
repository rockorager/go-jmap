package jmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdLength(t *testing.T) {
	cases := []struct {
		id    string
		valid bool
	}{
		{
			id:    "",
			valid: false,
		},
		{
			// Length 256
			id:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV",
			valid: false,
		},
	}

	for _, c := range cases {
		id := ID(c.id)

		ok, err := id.Valid()
		assert.Equal(t, ok, c.valid, "%v", err)
	}
}
