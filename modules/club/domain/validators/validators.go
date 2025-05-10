package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	cpfRegex   = regexp.MustCompile(`^\d{11}$`)
	cnpjRegex  = regexp.MustCompile(`^\d{14}$`)
)

// RegisterCustomValidators registra os validadores customizados
func RegisterCustomValidators(v *validator.Validate) {
	// Validador de email
	v.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		email = strings.ReplaceAll(email, " ", "")
		if emailRegex.MatchString(email) {
			fl.Field().SetString(email)
			return true
		}
		return false
	})

	// Validador de CPF/CNPJ
	v.RegisterValidation("cpf_cnpj", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		value = strings.ReplaceAll(value, ".", "")
		value = strings.ReplaceAll(value, "-", "")
		value = strings.ReplaceAll(value, "/", "")
		value = strings.ReplaceAll(value, " ", "")

		if len(value) == 11 {
			if cpfRegex.MatchString(value) {
				fl.Field().SetString(value)
				return true
			}
			return false
		}
		if len(value) == 14 {
			if cnpjRegex.MatchString(value) {
				fl.Field().SetString(value)
				return true
			}
			return false
		}
		return false
	})

	// Validador de telefone
	v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		phone = strings.ReplaceAll(phone, "(", "")
		phone = strings.ReplaceAll(phone, ")", "")
		phone = strings.ReplaceAll(phone, "-", "")
		phone = strings.ReplaceAll(phone, " ", "")
		if len(phone) >= 10 && len(phone) <= 11 {
			fl.Field().SetString(phone)
			return true
		}
		return false
	})

	// Validador de slug
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		slug := fl.Field().String()
		slug = strings.ReplaceAll(slug, " ", "-")
		if len(slug) >= 3 && len(slug) <= 100 {
			fl.Field().SetString(slug)
			return true
		}
		return false
	})

	// Validador de cep
	v.RegisterValidation("zip_code", func(fl validator.FieldLevel) bool {
		zipCode := fl.Field().String()
		zipCode = strings.ReplaceAll(zipCode, "-", "")
		zipCode = strings.ReplaceAll(zipCode, " ", "")
		if len(zipCode) == 8 {
			fl.Field().SetString(zipCode)
			return true
		}
		return false
	})

	// Validador de senha
	v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return len(password) >= 6 && len(password) <= 30
	})
}

// GetValidationErrorMessage retorna mensagens de erro customizadas
func GetValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("O campo %s é obrigatório", err.Field())
	case "email":
		return fmt.Sprintf("O campo %s deve ser um email válido", err.Field())
	case "cpf_cnpj":
		return fmt.Sprintf("O campo %s deve ser um CPF ou CNPJ válido", err.Field())
	case "phone":
		return fmt.Sprintf("O campo %s deve ser um telefone válido", err.Field())
	case "min":
		return fmt.Sprintf("O campo %s deve ter no mínimo %s caracteres", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("O campo %s deve ter no máximo %s caracteres", err.Field(), err.Param())
	default:
		return fmt.Sprintf("O campo %s é inválido", err.Field())
	}
}
