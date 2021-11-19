package helper

import (
	"github.com/fuadsuleyman/go-couriers/models"
	"github.com/go-playground/validator/v10"

)


type ErrorResponse struct {
    FailedField string
    Tag         string
    Value       string
}

func ValidateStruct(courier models.Courier) []*ErrorResponse {
    var errors []*ErrorResponse
    validate := validator.New()
    err := validate.Struct(courier)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var element ErrorResponse
            element.FailedField = err.StructNamespace()
            element.Tag = err.Tag()
            element.Value = err.Param()
            errors = append(errors, &element)
        }
    }
    return errors
}
