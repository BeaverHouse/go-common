package errorhandle

import "fmt"

// Common validation errors (COM[http status code]-[error code])

// ErrValidationFailed returns a validation failed error
//
//   - Code: COM400-00
//   - Message: Validation failed: {message}
func ErrValidationFailed(message string) error {
	return fmt.Errorf("COM400-00: Validation failed: %s", message)
}

// ErrInternal returns an internal error in the common module
//
//   - Code: COM500-00
//   - Message: Internal error: {error}
func ErrInternal(err error) error {
	return fmt.Errorf("COM500-00: Internal error: %v", err)
}
