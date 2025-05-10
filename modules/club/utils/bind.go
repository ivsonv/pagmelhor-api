package utils

import (
	"app/modules/club/domain/validators"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Bind valida e associa os dados da requisição ao objeto passado
// Retorna um map com as mensagens de erro e o erro original
func Bind[T any](c echo.Context, item T) (map[string]string, error) {
	if err := c.Bind(item); err != nil {
		return map[string]string{
			"error": "Erro ao processar a requisição",
		}, err
	}

	validate := validator.New()
	validators.RegisterCustomValidators(validate)

	if err := validate.Struct(item); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = validators.GetValidationErrorMessage(err)
		}
		return errors, err
	}

	return nil, nil
}
