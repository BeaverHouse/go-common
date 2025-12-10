package validation

import (
	"fmt"
	"strings"

	"github.com/BeaverHouse/go-common/errorhandle"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// Global validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("notblank", validators.NotBlank)
}

// ValidateStruct validates a struct using the global validator instance
func ValidateStruct(s any) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		message := formatValidationErrors(validationErrors)
		return errorhandle.ErrValidationFailed(message)
	}
	return errorhandle.ErrValidationFailed("Invalid input data.")
}

// RegisterValidation registers a custom validation function
func RegisterValidation(tag string, fn validator.Func) error {
	return validate.RegisterValidation(tag, fn)
}

// formatValidationErrors formats validator.ValidationErrors into a user-friendly string
func formatValidationErrors(errors validator.ValidationErrors) string {
	var messages []string
	for _, err := range errors {
		message := formatSingleValidationError(err)
		messages = append(messages, message)
	}
	return strings.Join(messages, "; ")
}

// formatSingleValidationError formats a single validation error for API responses
func formatSingleValidationError(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()
	param := err.Param()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, param)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, param)
	case "lte":
		return fmt.Sprintf("%s must be at most %s", field, param)
	case "gte":
		return fmt.Sprintf("%s must be at least %s", field, param)
	case "uuid4":
		return fmt.Sprintf("%s must be a valid UUID", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "notblank":
		return fmt.Sprintf("%s must not be blank", field)
	default:
		return fmt.Sprintf("%s failed validation for tag '%s'", field, tag)
	}
}
