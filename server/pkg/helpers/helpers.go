package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetValidationMessage(fe validator.FieldError) (key string, msg string) {
	switch fe.Tag() {
	case "required":
		return fe.Field(), "is required"
	case "email":
		return fe.Field(), "Invalid email"
	case "max":
		return fe.Field(), fmt.Sprintf("maximum length (%v) exited", fe.Param())
	case "min":
		return fe.Field(), fmt.Sprintf("too short minimum length (%v)", fe.Param())
	case "alpha":
		return fe.Field(), "invalid value only alpha value accepted"
	default:
		return "error", fe.Error() // default error
	}
}
