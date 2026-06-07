package conv

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// Naming: each pgtype has a symmetric pair named by the PG type —
// FromPg<Type> (pgtype -> Go pointer, nil when NULL) and
// ToPg<Type> (Go pointer -> pgtype, invalid when nil). A few variants at the end of
// each section cover non-pointer or string representations.

// =============================================================================
// pgtype -> Go
// =============================================================================

// FromPgText converts pgtype.Text to *string
func FromPgText(v pgtype.Text) *string {
	if !v.Valid {
		return nil
	}
	return &v.String
}

// FromPgTimestamptz converts pgtype.Timestamptz to *time.Time
func FromPgTimestamptz(v pgtype.Timestamptz) *time.Time {
	if !v.Valid {
		return nil
	}
	return &v.Time
}

// FromPgDate converts pgtype.Date to *time.Time
func FromPgDate(v pgtype.Date) *time.Time {
	if !v.Valid {
		return nil
	}
	return &v.Time
}

// FromPgInt2 converts pgtype.Int2 to *int16
func FromPgInt2(v pgtype.Int2) *int16 {
	if !v.Valid {
		return nil
	}
	return &v.Int16
}

// FromPgInt4 converts pgtype.Int4 to *int32
func FromPgInt4(v pgtype.Int4) *int32 {
	if !v.Valid {
		return nil
	}
	return &v.Int32
}

// FromPgBool converts pgtype.Bool to *bool
func FromPgBool(v pgtype.Bool) *bool {
	if !v.Valid {
		return nil
	}
	return &v.Bool
}

// FromPgBoolOrFalse converts pgtype.Bool to a bool value, returning false when
// NULL/invalid. Use it when a false fallback is fine and you don't need to tell
// NULL apart; use FromPgBool (returns *bool) when you do.
func FromPgBoolOrFalse(v pgtype.Bool) bool {
	return v.Valid && v.Bool
}

// FromPgNumeric converts pgtype.Numeric to *float64. Returns nil when the value
// is NULL or cannot be represented as a float64 (NaN/Inf).
func FromPgNumeric(v pgtype.Numeric) *float64 {
	if !v.Valid {
		return nil
	}
	f, err := v.Float64Value()
	if err != nil || !f.Valid {
		return nil
	}
	return &f.Float64
}

// FromPgFloat8 converts pgtype.Float8 to *float64
func FromPgFloat8(v pgtype.Float8) *float64 {
	if !v.Valid {
		return nil
	}
	return &v.Float64
}

// FromPgUUID converts pgtype.UUID to *uuid.UUID (nil when NULL).
func FromPgUUID(v pgtype.UUID) *uuid.UUID {
	if !v.Valid {
		return nil
	}
	u := uuid.UUID(v.Bytes)
	return &u
}

// FromPgDateString formats pgtype.Date as a "YYYY-MM-DD" string (nil when NULL). Use
// when an API represents a date-only value as a string rather than a time.Time.
func FromPgDateString(v pgtype.Date) *string {
	if !v.Valid {
		return nil
	}
	s := v.Time.Format("2006-01-02")
	return &s
}

// FromPgTimestamptzOrZero converts pgtype.Timestamptz to a time.Time value,
// returning the zero time when NULL/invalid. Use it when a zero fallback is fine
// and you don't need to tell NULL apart; use FromPgTimestamptz (returns *time.Time)
// when you do.
func FromPgTimestamptzOrZero(v pgtype.Timestamptz) time.Time {
	if !v.Valid {
		return time.Time{}
	}
	return v.Time
}

// =============================================================================
// Go -> pgtype
// =============================================================================

// ToPgText converts *string to pgtype.Text
func ToPgText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

// ToPgTimestamptz converts *time.Time to pgtype.Timestamptz
func ToPgTimestamptz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{Time: *t, Valid: true}
}

// ToPgDate converts *time.Time to pgtype.Date
func ToPgDate(t *time.Time) pgtype.Date {
	if t == nil {
		return pgtype.Date{Valid: false}
	}
	return pgtype.Date{Time: *t, Valid: true}
}

// ToPgInt2 converts *int16 to pgtype.Int2
func ToPgInt2(i *int16) pgtype.Int2 {
	if i == nil {
		return pgtype.Int2{Valid: false}
	}
	return pgtype.Int2{Int16: *i, Valid: true}
}

// ToPgInt4 converts *int32 to pgtype.Int4
func ToPgInt4(i *int32) pgtype.Int4 {
	if i == nil {
		return pgtype.Int4{Valid: false}
	}
	return pgtype.Int4{Int32: *i, Valid: true}
}

// ToPgBool converts *bool to pgtype.Bool
func ToPgBool(b *bool) pgtype.Bool {
	if b == nil {
		return pgtype.Bool{Valid: false}
	}
	return pgtype.Bool{Bool: *b, Valid: true}
}

// ToPgNumeric converts *float64 to pgtype.Numeric. Returns an invalid Numeric
// when the input is nil or not a finite number.
func ToPgNumeric(f *float64) pgtype.Numeric {
	if f == nil {
		return pgtype.Numeric{Valid: false}
	}
	var n pgtype.Numeric
	if err := n.Scan(strconv.FormatFloat(*f, 'f', -1, 64)); err != nil {
		return pgtype.Numeric{Valid: false}
	}
	return n
}

// ToPgFloat8 converts *float64 to pgtype.Float8
func ToPgFloat8(f *float64) pgtype.Float8 {
	if f == nil {
		return pgtype.Float8{Valid: false}
	}
	return pgtype.Float8{Float64: *f, Valid: true}
}

// ToPgUUID converts uuid.UUID to pgtype.UUID. The nil UUID maps to invalid (NULL).
func ToPgUUID(u uuid.UUID) pgtype.UUID {
	if u == uuid.Nil {
		return pgtype.UUID{Valid: false}
	}
	return pgtype.UUID{Bytes: u, Valid: true}
}

// ToPgDateString parses a "YYYY-MM-DD" string into pgtype.Date (invalid when nil,
// empty, or unparseable).
func ToPgDateString(s *string) pgtype.Date {
	if s == nil || *s == "" {
		return pgtype.Date{Valid: false}
	}
	t, err := time.Parse("2006-01-02", *s)
	if err != nil {
		return pgtype.Date{Valid: false}
	}
	return pgtype.Date{Time: t, Valid: true}
}
