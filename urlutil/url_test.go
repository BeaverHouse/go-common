package urlutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "URL without protocol and with whitespace",
			input:    "  google.com  ",
			expected: "https://google.com",
		},
		{
			name:     "URL with https unchanged",
			input:    "https://example.com",
			expected: "https://example.com",
		},
		{
			name:     "URL with port and http protocol",
			input:    "http://localhost:8080",
			expected: "http://localhost:8080",
		},
		{
			name:     "URL with port and no protocol",
			input:    "localhost:8080",
			expected: "https://localhost:8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NormalizeURL(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
