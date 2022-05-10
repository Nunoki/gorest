package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidatorError is a helper for the go-playground/validator package which will loop through its
// validation errors, and return them as a string, joined with a comma
func ValidatorError(err error) string {
	// TODO: Since outputting only the validation rule may not always be helpful on the frontend,
	// some other mechanism should be introduced to offer more detailed validation messages
	errorMessages := make([]string, 0)

	for _, err := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages,
			fmt.Sprintf(
				"Parameter %s violates rule %s",
				err.Field(),
				err.Tag(),
			),
		)
	}

	return strings.Join(errorMessages, ", ")
}
