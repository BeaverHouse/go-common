package validation

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=0,lte=150"`
}

func TestValidateStruct_Valid(t *testing.T) {
	s := TestStruct{
		Name:  "John",
		Email: "john@example.com",
		Age:   30,
	}

	err := ValidateStruct(s)
	assert.NoError(t, err)
}

func TestValidateStruct_SingleError(t *testing.T) {
	s := TestStruct{
		Email: "john@example.com",
		Age:   30,
	}

	err := ValidateStruct(s)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Name is required")
	assert.Contains(t, err.Error(), "COM400-00")
}

func TestValidateStruct_MultipleErrors(t *testing.T) {
	s := TestStruct{
		Email: "invalid-email",
		Age:   -1,
	}

	err := ValidateStruct(s)
	assert.Error(t, err)
	errStr := err.Error()
	assert.True(t, strings.Contains(errStr, ";"), "multiple errors should be joined with semicolon")
	assert.True(t, strings.Contains(errStr, "Name is required"))
}

func TestValidateStruct_NotBlank(t *testing.T) {
	type TestStructNotBlank struct {
		Name string `validate:"notblank"`
	}

	s := TestStructNotBlank{
		Name: "   ",
	}

	err := ValidateStruct(s)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Name must not be blank")
}

func TestRegisterValidation(t *testing.T) {
	err := RegisterValidation("customtag", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "custom"
	})
	assert.NoError(t, err)

	type CustomStruct struct {
		Value string `validate:"customtag"`
	}

	t.Run("valid", func(t *testing.T) {
		err = ValidateStruct(CustomStruct{Value: "custom"})
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err = ValidateStruct(CustomStruct{Value: "other"})
		assert.Error(t, err)
	})
}
