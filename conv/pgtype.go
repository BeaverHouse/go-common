package conv

import (
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Naming: each pgtype has a symmetric pair named by the PG type —
// FromPg<Type> (pgtype -> Go pointer, nil when NULL) and
// ToPg<Type> (Go pointer -> pgtype, invalid when nil).

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
