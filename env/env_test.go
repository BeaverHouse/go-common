package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "test_value")
	defer os.Unsetenv("TEST_ENV_VAR")

	tests := []struct {
		name         string
		key          string
		defaultValue string
		expected     string
	}{
		{
			name:         "existing env var",
			key:          "TEST_ENV_VAR",
			defaultValue: "default",
			expected:     "test_value",
		},
		{
			name:         "non-existing env var returns default",
			key:          "NON_EXISTING_VAR",
			defaultValue: "default",
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetEnv(tt.key, tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetEnv_EmptyValue(t *testing.T) {
	os.Setenv("EMPTY_VAR", "")
	defer os.Unsetenv("EMPTY_VAR")

	result := GetEnv("EMPTY_VAR", "default")
	assert.Equal(t, "default", result)
}

func TestGetIntEnv(t *testing.T) {
	os.Setenv("INT_VAR", "42")
	os.Setenv("INVALID_INT", "not_a_number")
	defer os.Unsetenv("INT_VAR")
	defer os.Unsetenv("INVALID_INT")

	tests := []struct {
		name         string
		key          string
		defaultValue int
		expected     int
	}{
		{
			name:         "valid integer",
			key:          "INT_VAR",
			defaultValue: 0,
			expected:     42,
		},
		{
			name:         "non-existing returns default",
			key:          "NON_EXISTING_INT",
			defaultValue: 100,
			expected:     100,
		},
		{
			name:         "invalid integer returns default",
			key:          "INVALID_INT",
			defaultValue: 50,
			expected:     50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetIntEnv(tt.key, tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsGoEnv(t *testing.T) {
	defer os.Unsetenv("GO_ENV")

	t.Run("default is local when not set", func(t *testing.T) {
		os.Unsetenv("GO_ENV")
		assert.True(t, IsGoEnv("local"))
		assert.False(t, IsGoEnv("production"))
	})

	t.Run("matches when set", func(t *testing.T) {
		os.Setenv("GO_ENV", "production")
		assert.True(t, IsGoEnv("production"))
		assert.False(t, IsGoEnv("local"))
	})
}
