package helper

import "github.com/go-playground/validator/v10"

func GetFieldErrors(err validator.ValidationErrors) map[string][]string {
	errorMessages := make(map[string][]string)
		for _, err := range err {
			field := err.Field()
			message := err.Tag()

			errorMessages[field] = append(errorMessages[field], message)
		}
		return errorMessages;
}