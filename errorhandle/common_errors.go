package errorhandle

// ErrValidationFailed returns a KindInvalid error describing failed input
// validation. It is what validation.ValidateStruct returns.
func ErrValidationFailed(message string) error {
	return New(KindInvalid, "COM_VALIDATION_FAILED", "validation failed: "+message)
}

// ErrInternal wraps an unexpected cause as a KindInternal error.
func ErrInternal(err error) error {
	return Wrap(KindInternal, "COM_INTERNAL", "internal error", err)
}
