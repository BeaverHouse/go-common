package conv

import "encoding/json"

// MarshalJSONOrDefault marshals v to JSON, returning the fallback bytes when v is
// nil. Keeps "create" paths terse where a JSONB column needs a fixed default
// (e.g. "[]") instead of NULL.
func MarshalJSONOrDefault(v any, fallback string) ([]byte, error) {
	if v == nil {
		return []byte(fallback), nil
	}
	return json.Marshal(v)
}

// MarshalJSONOrNil marshals v to JSON, returning (nil, nil) when v is nil so the
// JSONB column receives SQL NULL.
func MarshalJSONOrNil(v any) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	return json.Marshal(v)
}
