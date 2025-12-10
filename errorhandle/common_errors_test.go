package errorhandle

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrValidationFailed(t *testing.T) {
	err := ErrValidationFailed("Name is required")

	assert.Contains(t, err.Error(), "COM400-00")
	assert.Contains(t, err.Error(), "Name is required")
}

func TestErrInternal(t *testing.T) {
	originalErr := errors.New("db connection failed")
	err := ErrInternal(originalErr)

	assert.Contains(t, err.Error(), "COM500-00")
	assert.Contains(t, err.Error(), "db connection failed")
}
