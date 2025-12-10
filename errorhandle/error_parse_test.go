package errorhandle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractHTTPStatusFromError(t *testing.T) {
	tests := []struct {
		name         string
		errorMessage string
		modulePrefix string
		expected     int
	}{
		{
			name:         "COM prefix (common module)",
			errorMessage: "COM400-00: validation failed",
			modulePrefix: "",
			expected:     400,
		},
		{
			name:         "AU prefix (auth module)",
			errorMessage: "AU401-01: unauthorized",
			modulePrefix: "",
			expected:     401,
		},
		{
			name:         "custom prefix",
			errorMessage: "LO404-01: not found",
			modulePrefix: "LO",
			expected:     404,
		},
		{
			name:         "wrong custom prefix",
			errorMessage: "DA404-01: not found",
			modulePrefix: "LO",
			expected:     500,
		},
		{
			name:         "no match returns 500",
			errorMessage: "random error",
			modulePrefix: "",
			expected:     500,
		},
		{
			name:         "wrong format returns 500",
			errorMessage: "COM40010: missing dash",
			modulePrefix: "",
			expected:     500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractHTTPStatusFromError(tt.errorMessage, tt.modulePrefix)
			assert.Equal(t, tt.expected, result)
		})
	}
}
