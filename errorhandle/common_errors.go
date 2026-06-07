package errorhandle

// Closed set of cross-service error constructors: the same situation reads
// identically everywhere. Services build domain errors on New/Wrap with these
// Kinds and must not re-declare these. Order: by Kind (kind.go), then alphabetical.

// ErrConfigMissing reports a required configuration value (e.g. an env var) not set.
func ErrConfigMissing(key string) error {
	return New(KindInternal, "COM_CONFIG_MISSING", "required configuration not set: "+key)
}

// ErrDBOperation wraps a failed database operation, named by operation.
func ErrDBOperation(operation string, err error) error {
	return Wrap(KindInternal, "COM_DB_OPERATION", operation+" failed", err)
}

// ErrInternal wraps an unexpected cause as an internal error.
func ErrInternal(err error) error {
	return Wrap(KindInternal, "COM_INTERNAL", "internal error", err)
}

// ErrInternalMsg reports an internal failure described by a message (no Go cause).
func ErrInternalMsg(message string) error {
	return New(KindInternal, "COM_INTERNAL", message)
}

// ErrValidationFailed reports failed input validation. validation.ValidateStruct returns this.
func ErrValidationFailed(message string) error {
	return New(KindInvalid, "COM_VALIDATION_FAILED", "validation failed: "+message)
}

// ErrNotFound reports a missing resource, named by resource (e.g. "student").
func ErrNotFound(resource string) error {
	return New(KindNotFound, "COM_NOT_FOUND", resource+" not found")
}

// ErrConflict reports a resource that already exists or a state conflict.
func ErrConflict(resource string) error {
	return New(KindConflict, "COM_CONFLICT", resource+" already exists")
}

// ErrUnauthorized reports a missing or invalid credential (HTTP 401).
func ErrUnauthorized(message string) error {
	return New(KindUnauthenticated, "COM_UNAUTHORIZED", message)
}

// ErrPermissionDenied reports an authenticated caller lacking permission (HTTP 403).
func ErrPermissionDenied(message string) error {
	return New(KindPermissionDenied, "COM_PERMISSION_DENIED", message)
}

// ErrPaymentRequired reports that a subscription tier or quota is required (HTTP 402).
func ErrPaymentRequired(message string) error {
	return New(KindPaymentRequired, "COM_PAYMENT_REQUIRED", message)
}
