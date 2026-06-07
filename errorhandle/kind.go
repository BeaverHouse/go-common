package errorhandle

import (
	"errors"
	"fmt"
	"net/http"
)

// Kind is a protocol-agnostic error category, named after Rob Pike's Upspin
// errors.Kind (gRPC calls it codes.Code; Kubernetes, StatusReason). Core logic
// raises errors by Kind; each adapter (HTTP, MCP, CLI) maps the Kind to its own
// representation (HTTP status, error text, exit code) at the boundary, so business
// logic never depends on a transport.
//
// This is a CLOSED, generic taxonomy of standard status categories — not a registry
// of domain errors. go-common owns the categories; services own their error
// instances (errorhandle.New/Wrap with an existing Kind) and must NOT add
// business-specific Kinds here. New values are added only to fill a missing
// standard status category, which is rare.
type Kind int

const (
	// KindInternal is an unexpected, unclassified failure. It is the zero value,
	// so an unrecognized error is treated as a server fault by default.
	KindInternal Kind = iota
	// KindInvalid is a malformed or semantically invalid request.
	KindInvalid
	// KindNotFound is a missing resource.
	KindNotFound
	// KindConflict is a state conflict, such as a duplicate or version mismatch.
	KindConflict
	// KindUnauthenticated is a missing or invalid credential.
	KindUnauthenticated
	// KindPermissionDenied is an authenticated caller lacking permission.
	KindPermissionDenied
	// KindFailedPrecondition is a request rejected by the current system state.
	KindFailedPrecondition
	// KindUnavailable is a transient failure that the caller may retry.
	KindUnavailable
	// KindPaymentRequired means payment or quota (e.g. a subscription tier) is
	// required before the request can proceed.
	KindPaymentRequired
)

// Error is a structured, protocol-agnostic domain error. Code is a stable,
// service-defined identifier (e.g. "USER_NOT_FOUND") for catalogs, metrics, and
// i18n; it must not embed a transport status such as an HTTP code. Err, when
// set, is the wrapped cause and is reachable via errors.Is / errors.As. Details,
// when set via WithDetails, carries structured context an adapter may surface
// alongside the message (e.g. an HTTP body's data field).
type Error struct {
	Kind    Kind
	Code    string
	Message string
	Err     error
	Details any
}

// Error implements the error interface.
func (e *Error) Error() string {
	switch {
	case e.Code == "" && e.Err != nil:
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	case e.Code == "":
		return e.Message
	case e.Err != nil:
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	default:
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	}
}

// Unwrap returns the wrapped cause so errors.Is and errors.As traverse it.
func (e *Error) Unwrap() error { return e.Err }

// New creates an Error with the given kind, stable code, and message.
func New(kind Kind, code, message string) *Error {
	return &Error{Kind: kind, Code: code, Message: message}
}

// Wrap creates an Error that wraps cause with the given kind, code, and message.
func Wrap(kind Kind, code, message string, cause error) *Error {
	return &Error{Kind: kind, Code: code, Message: message, Err: cause}
}

// WithDetails attaches structured context to the error and returns the receiver,
// for the rare case an adapter must surface a data payload alongside the message
// (e.g. an HTTP 409 listing the conflicting resources).
func (e *Error) WithDetails(details any) *Error {
	e.Details = details
	return e
}

// DetailsOf returns the structured details attached to err via WithDetails,
// unwrapping as needed; nil when none.
func DetailsOf(err error) any {
	if e, ok := errors.AsType[*Error](err); ok {
		return e.Details
	}
	return nil
}

// KindOf reports the Kind of err, unwrapping as needed. A nil error or an error
// that is not an *Error is reported as KindInternal.
func KindOf(err error) Kind {
	if e, ok := errors.AsType[*Error](err); ok {
		return e.Kind
	}
	return KindInternal
}

// HTTPStatus maps a Kind to an HTTP status code. Use this in HTTP adapters only;
// other adapters map Kind to their own representation.
func HTTPStatus(kind Kind) int {
	switch kind {
	case KindInvalid:
		return http.StatusBadRequest
	case KindNotFound:
		return http.StatusNotFound
	case KindConflict:
		return http.StatusConflict
	case KindUnauthenticated:
		return http.StatusUnauthorized
	case KindPermissionDenied:
		return http.StatusForbidden
	case KindFailedPrecondition:
		return http.StatusUnprocessableEntity
	case KindUnavailable:
		return http.StatusServiceUnavailable
	case KindPaymentRequired:
		return http.StatusPaymentRequired
	default:
		return http.StatusInternalServerError
	}
}

// HTTPStatusOf maps an error to an HTTP status code via its Kind. It is a
// convenience for HTTP adapters over HTTPStatus(KindOf(err)).
func HTTPStatusOf(err error) int {
	return HTTPStatus(KindOf(err))
}
