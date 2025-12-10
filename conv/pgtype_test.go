package conv

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestNullString(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Text
		expected *string
	}{
		{
			name:     "valid string",
			input:    pgtype.Text{String: "hello", Valid: true},
			expected: strPtr("hello"),
		},
		{
			name:     "invalid string",
			input:    pgtype.Text{Valid: false},
			expected: nil,
		},
		{
			name:     "empty valid string",
			input:    pgtype.Text{String: "", Valid: true},
			expected: strPtr(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullString(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestNullTime(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		input    pgtype.Timestamptz
		expected *time.Time
	}{
		{
			name:     "valid time",
			input:    pgtype.Timestamptz{Time: now, Valid: true},
			expected: &now,
		},
		{
			name:     "invalid time",
			input:    pgtype.Timestamptz{Valid: false},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullTime(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestNullInt32(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Int4
		expected *int32
	}{
		{
			name:     "valid int32",
			input:    pgtype.Int4{Int32: 42, Valid: true},
			expected: int32Ptr(42),
		},
		{
			name:     "invalid int32",
			input:    pgtype.Int4{Valid: false},
			expected: nil,
		},
		{
			name:     "zero valid int32",
			input:    pgtype.Int4{Int32: 0, Valid: true},
			expected: int32Ptr(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullInt32(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestNullBool(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Bool
		expected *bool
	}{
		{
			name:     "valid true",
			input:    pgtype.Bool{Bool: true, Valid: true},
			expected: boolPtr(true),
		},
		{
			name:     "valid false",
			input:    pgtype.Bool{Bool: false, Valid: true},
			expected: boolPtr(false),
		},
		{
			name:     "invalid bool",
			input:    pgtype.Bool{Valid: false},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullBool(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestToPgText(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected pgtype.Text
	}{
		{
			name:     "non-nil string",
			input:    strPtr("hello"),
			expected: pgtype.Text{String: "hello", Valid: true},
		},
		{
			name:     "nil string",
			input:    nil,
			expected: pgtype.Text{Valid: false},
		},
		{
			name:     "empty string",
			input:    strPtr(""),
			expected: pgtype.Text{String: "", Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPgText(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToPgTimestamptz(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		input    *time.Time
		expected pgtype.Timestamptz
	}{
		{
			name:     "non-nil time",
			input:    &now,
			expected: pgtype.Timestamptz{Time: now, Valid: true},
		},
		{
			name:     "nil time",
			input:    nil,
			expected: pgtype.Timestamptz{Valid: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPgTimestamptz(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToPgInt4(t *testing.T) {
	tests := []struct {
		name     string
		input    *int32
		expected pgtype.Int4
	}{
		{
			name:     "non-nil int32",
			input:    int32Ptr(42),
			expected: pgtype.Int4{Int32: 42, Valid: true},
		},
		{
			name:     "nil int32",
			input:    nil,
			expected: pgtype.Int4{Valid: false},
		},
		{
			name:     "zero int32",
			input:    int32Ptr(0),
			expected: pgtype.Int4{Int32: 0, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPgInt4(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToPgBool(t *testing.T) {
	tests := []struct {
		name     string
		input    *bool
		expected pgtype.Bool
	}{
		{
			name:     "non-nil true",
			input:    boolPtr(true),
			expected: pgtype.Bool{Bool: true, Valid: true},
		},
		{
			name:     "non-nil false",
			input:    boolPtr(false),
			expected: pgtype.Bool{Bool: false, Valid: true},
		},
		{
			name:     "nil bool",
			input:    nil,
			expected: pgtype.Bool{Valid: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPgBool(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func strPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}
