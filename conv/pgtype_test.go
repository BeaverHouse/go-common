package conv

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestFromPgText(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Text
		expected *string
	}{
		{"valid string", pgtype.Text{String: "hello", Valid: true}, strPtr("hello")},
		{"invalid string", pgtype.Text{Valid: false}, nil},
		{"empty valid string", pgtype.Text{String: "", Valid: true}, strPtr("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromPgText(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestFromPgTimestamptz(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		input    pgtype.Timestamptz
		expected *time.Time
	}{
		{"valid time", pgtype.Timestamptz{Time: now, Valid: true}, &now},
		{"invalid time", pgtype.Timestamptz{Valid: false}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromPgTimestamptz(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestFromPgDate(t *testing.T) {
	day := time.Date(2026, 6, 5, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		input    pgtype.Date
		expected *time.Time
	}{
		{"valid date", pgtype.Date{Time: day, Valid: true}, &day},
		{"invalid date", pgtype.Date{Valid: false}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromPgDate(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestFromPgInt2(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Int2
		expected *int16
	}{
		{"valid int16", pgtype.Int2{Int16: 7, Valid: true}, int16Ptr(7)},
		{"invalid int16", pgtype.Int2{Valid: false}, nil},
		{"zero valid int16", pgtype.Int2{Int16: 0, Valid: true}, int16Ptr(0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromPgInt2(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestFromPgInt4(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Int4
		expected *int32
	}{
		{"valid int32", pgtype.Int4{Int32: 42, Valid: true}, int32Ptr(42)},
		{"invalid int32", pgtype.Int4{Valid: false}, nil},
		{"zero valid int32", pgtype.Int4{Int32: 0, Valid: true}, int32Ptr(0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromPgInt4(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestFromPgBool(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Bool
		expected *bool
	}{
		{"valid true", pgtype.Bool{Bool: true, Valid: true}, boolPtr(true)},
		{"valid false", pgtype.Bool{Bool: false, Valid: true}, boolPtr(false)},
		{"invalid bool", pgtype.Bool{Valid: false}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromPgBool(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestFromPgNumeric(t *testing.T) {
	t.Run("invalid numeric", func(t *testing.T) {
		assert.Nil(t, FromPgNumeric(pgtype.Numeric{Valid: false}))
	})
	t.Run("valid numeric round-trips", func(t *testing.T) {
		result := FromPgNumeric(ToPgNumeric(float64Ptr(3.14)))
		assert.NotNil(t, result)
		assert.InDelta(t, 3.14, *result, 1e-9)
	})
}

func TestToPgText(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected pgtype.Text
	}{
		{"non-nil string", strPtr("hello"), pgtype.Text{String: "hello", Valid: true}},
		{"nil string", nil, pgtype.Text{Valid: false}},
		{"empty string", strPtr(""), pgtype.Text{String: "", Valid: true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ToPgText(tt.input))
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
		{"non-nil time", &now, pgtype.Timestamptz{Time: now, Valid: true}},
		{"nil time", nil, pgtype.Timestamptz{Valid: false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ToPgTimestamptz(tt.input))
		})
	}
}

func TestToPgDate(t *testing.T) {
	day := time.Date(2026, 6, 5, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		input    *time.Time
		expected pgtype.Date
	}{
		{"non-nil date", &day, pgtype.Date{Time: day, Valid: true}},
		{"nil date", nil, pgtype.Date{Valid: false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ToPgDate(tt.input))
		})
	}
}

func TestToPgInt2(t *testing.T) {
	tests := []struct {
		name     string
		input    *int16
		expected pgtype.Int2
	}{
		{"non-nil int16", int16Ptr(7), pgtype.Int2{Int16: 7, Valid: true}},
		{"nil int16", nil, pgtype.Int2{Valid: false}},
		{"zero int16", int16Ptr(0), pgtype.Int2{Int16: 0, Valid: true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ToPgInt2(tt.input))
		})
	}
}

func TestToPgInt4(t *testing.T) {
	tests := []struct {
		name     string
		input    *int32
		expected pgtype.Int4
	}{
		{"non-nil int32", int32Ptr(42), pgtype.Int4{Int32: 42, Valid: true}},
		{"nil int32", nil, pgtype.Int4{Valid: false}},
		{"zero int32", int32Ptr(0), pgtype.Int4{Int32: 0, Valid: true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ToPgInt4(tt.input))
		})
	}
}

func TestToPgBool(t *testing.T) {
	tests := []struct {
		name     string
		input    *bool
		expected pgtype.Bool
	}{
		{"non-nil true", boolPtr(true), pgtype.Bool{Bool: true, Valid: true}},
		{"non-nil false", boolPtr(false), pgtype.Bool{Bool: false, Valid: true}},
		{"nil bool", nil, pgtype.Bool{Valid: false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ToPgBool(tt.input))
		})
	}
}

func TestToPgNumeric(t *testing.T) {
	t.Run("nil numeric", func(t *testing.T) {
		assert.Equal(t, pgtype.Numeric{Valid: false}, ToPgNumeric(nil))
	})
	t.Run("valid numeric", func(t *testing.T) {
		n := ToPgNumeric(float64Ptr(12.5))
		assert.True(t, n.Valid)
		f, err := n.Float64Value()
		assert.NoError(t, err)
		assert.InDelta(t, 12.5, f.Float64, 1e-9)
	})
}

func strPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}

func int16Ptr(i int16) *int16 {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}

func boolPtr(b bool) *bool {
	return &b
}
