package errorhandle

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrValidationFailed(t *testing.T) {
	err := ErrValidationFailed("Name is required")

	assert.Equal(t, KindInvalid, KindOf(err))
	assert.Equal(t, http.StatusBadRequest, HTTPStatusOf(err))
	assert.Contains(t, err.Error(), "COM_VALIDATION_FAILED")
	assert.Contains(t, err.Error(), "Name is required")
}

func TestErrInternal(t *testing.T) {
	cause := errors.New("db connection failed")
	err := ErrInternal(cause)

	assert.Equal(t, KindInternal, KindOf(err))
	assert.Equal(t, http.StatusInternalServerError, HTTPStatusOf(err))
	assert.ErrorIs(t, err, cause)
	assert.Contains(t, err.Error(), "db connection failed")
}
