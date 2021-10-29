package symmetry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name   string
		msg    string
		key    string
		cypher string
	}{
		{
			name:   "hello",
			msg:    "hello",
			key:    "wldse",
			cypher: "bTSbW",
		},
		{
			name:   "world",
			msg:    "world",
			key:    "xnbag",
			cypher: "r`WPN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Encrypt(tt.msg, tt.key)
			assert.Equal(t, tt.cypher, res)
		})
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
		msg    string
		key    string
	}{
		{
			name:   "hello",
			key:    "wldse",
			cypher: "bTSbW",
			msg:    "hello",
		},
		{
			name:   "world",
			msg:    "world",
			key:    "xnbag",
			cypher: "r`WPN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Decrypt(tt.cypher, tt.key)
			assert.Equal(t, tt.msg, res)
		})
	}
}
