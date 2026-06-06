package errorhandle

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorString(t *testing.T) {
	assert.Equal(t, "USER_NOT_FOUND: user not found", New(KindNotFound, "USER_NOT_FOUND", "user not found").Error())
	assert.Equal(t, "no code message", New(KindInvalid, "", "no code message").Error())

	cause := errors.New("connection refused")
	assert.Equal(t, "DB_DOWN: query failed: connection refused", Wrap(KindUnavailable, "DB_DOWN", "query failed", cause).Error())
}

func TestUnwrapAndIs(t *testing.T) {
	cause := errors.New("root cause")
	err := Wrap(KindInternal, "X", "wrapped", cause)
	assert.ErrorIs(t, err, cause)
}

func TestKindOf(t *testing.T) {
	assert.Equal(t, KindInternal, KindOf(nil))
	assert.Equal(t, KindInternal, KindOf(errors.New("plain")))
	assert.Equal(t, KindNotFound, KindOf(New(KindNotFound, "C", "m")))

	// KindOf must see through additional wrapping.
	wrapped := fmt.Errorf("context: %w", New(KindConflict, "C", "m"))
	assert.Equal(t, KindConflict, KindOf(wrapped))
}

func TestHTTPStatus(t *testing.T) {
	cases := map[Kind]int{
		KindInternal:           http.StatusInternalServerError,
		KindInvalid:            http.StatusBadRequest,
		KindNotFound:           http.StatusNotFound,
		KindConflict:           http.StatusConflict,
		KindUnauthenticated:    http.StatusUnauthorized,
		KindPermissionDenied:   http.StatusForbidden,
		KindFailedPrecondition: http.StatusUnprocessableEntity,
		KindUnavailable:        http.StatusServiceUnavailable,
	}
	for kind, want := range cases {
		assert.Equal(t, want, HTTPStatus(kind))
	}
}

func TestHTTPStatusOf(t *testing.T) {
	assert.Equal(t, http.StatusNotFound, HTTPStatusOf(New(KindNotFound, "C", "m")))
	assert.Equal(t, http.StatusInternalServerError, HTTPStatusOf(errors.New("plain")))
}
