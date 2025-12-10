package conv

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// NullString converts pgtype.Text to *string
func NullString(v pgtype.Text) *string {
	if !v.Valid {
		return nil
	}
	return &v.String
}

// NullTime converts pgtype.Timestamptz to *time.Time
func NullTime(v pgtype.Timestamptz) *time.Time {
	if !v.Valid {
		return nil
	}
	return &v.Time
}

// NullInt32 converts pgtype.Int4 to *int32
func NullInt32(v pgtype.Int4) *int32 {
	if !v.Valid {
		return nil
	}
	return &v.Int32
}

// NullBool converts pgtype.Bool to *bool
func NullBool(v pgtype.Bool) *bool {
	if !v.Valid {
		return nil
	}
	return &v.Bool
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
