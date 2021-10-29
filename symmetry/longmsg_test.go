package symmetry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrptLongMsg(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		msg    string
		cypher string
	}{
		{
			name:   "case 1",
			key:    "hello",
			msg:    "Hello World",
			cypher: "3M[[a\v?^a^O",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := EncrptLongMsg(tt.msg, tt.key)
			assert.Equal(t, tt.cypher, res)
		})
	}
}

func TestDecryptLongMsg(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		cypher string
		msg    string
	}{
		{
			name:   "case 2",
			key:    "hello",
			msg:    "Hello World",
			cypher: "3M[[a\v?^a^O",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DecryptLongMsg(tt.cypher, tt.key)
			assert.Equal(t, tt.msg, res)
		})
	}
}
