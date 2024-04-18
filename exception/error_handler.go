package exception

import (
	"errors"

	"github.com/AhmadIkbalDjaya/go-simple-pos/models/api"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	errorResponse := api.BaseError{
		Code: fiber.StatusInternalServerError,
		Status: "Internal Server Error",
		Message: "Internal Server Error",
		Errors: err.Error(),
	}
	var e *fiber.Error
	if errors.As(err, &e) {
		errorResponse.Code = e.Code
	}
	if errorResponse.Code == 404 {
		errorResponse.Code = fiber.StatusNotFound
		errorResponse.Status = "Not Found"
		errorResponse.Message = "Route Not Found"
		errorResponse.Errors = err.Error()
	}

	switch err := err.(type){
	case *NotFoundError:
		errorResponse.Code = fiber.StatusNotFound
		errorResponse.Status = "Not Found"
		errorResponse.Message = err.Model + " Not Found"
		errorResponse.Errors = err.Error()
	case validator.ValidationErrors:
		errorMessages := GetFieldErrors(err)
		errorResponse.Code = fiber.StatusUnprocessableEntity
		errorResponse.Status = "Unprocessable Content"
		errorResponse.Message = "Data Is Not Valid"
		errorResponse.Errors = errorMessages
	}
	
	return ctx.Status(errorResponse.Code).JSON(errorResponse)
}

func GetFieldErrors(err validator.ValidationErrors) map[string][]string {
	errorMessages := make(map[string][]string)
		for _, err := range err {
			field := err.Field()
			message := err.Tag()

			errorMessages[field] = append(errorMessages[field], message)
		}
		return errorMessages;
}