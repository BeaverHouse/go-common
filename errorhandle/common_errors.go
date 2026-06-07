package errorhandle

// Closed, cross-service error set: one constructor per situation so every service
// raises it identically. Services build domain errors on New/Wrap with these
// Kinds and never redeclare these. Order: by Kind (kind.go), then alphabetical.

func ErrConfigMissing(key string) error {
	return New(KindInternal, "COM_CONFIG_MISSING", "required configuration not set: "+key)
}

func ErrDBOperation(operation string, err error) error {
	return Wrap(KindInternal, "COM_DB_OPERATION", operation+" failed", err)
}

func ErrInternal(err error) error {
	return Wrap(KindInternal, "COM_INTERNAL", "internal error", err)
}

// ErrInternalMsg reports an internal failure described by a message (no Go cause).
func ErrInternalMsg(message string) error {
	return New(KindInternal, "COM_INTERNAL", message)
}

func ErrValidationFailed(message string) error {
	return New(KindInvalid, "COM_VALIDATION_FAILED", "validation failed: "+message)
}

func ErrNotFound(resource string) error {
	return New(KindNotFound, "COM_NOT_FOUND", resource+" not found")
}

func ErrConflict(resource string) error {
	return New(KindConflict, "COM_CONFLICT", resource+" already exists")
}

func ErrUnauthorized(message string) error {
	return New(KindUnauthenticated, "COM_UNAUTHORIZED", message)
}

func ErrPermissionDenied(message string) error {
	return New(KindPermissionDenied, "COM_PERMISSION_DENIED", message)
}

func ErrPaymentRequired(message string) error {
	return New(KindPaymentRequired, "COM_PAYMENT_REQUIRED", message)
}
