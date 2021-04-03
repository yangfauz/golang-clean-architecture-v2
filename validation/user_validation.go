package validation

import (
	"api_project/exception"
	"api_project/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Hobby, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
