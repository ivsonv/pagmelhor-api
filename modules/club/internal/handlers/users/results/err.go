package results

import (
	"net/http"
)

// Error representa um erro da aplicação com código e status HTTP
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Error implementa a interface error
func (e *Error) Error() string {
	return e.Message
}

// NewError cria um novo erro
func NewError(code string, message string, status int) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// UserErrors contém todos os erros relacionados a usuários
const (
	InternalServerError = "INTERNAL_SERVER_ERROR"
	EmailExists         = "EMAIL_ALREADY_EXISTS"
	UserNotFound        = "USER_NOT_FOUND"
)

var (
	ErrInternalServerError = NewError(InternalServerError, "Internal server error", http.StatusInternalServerError)
	ErrEmailExists         = NewError(EmailExists, "Email already registered", http.StatusConflict)
	ErrUserNotFound        = NewError(UserNotFound, "User not found", http.StatusNotFound)
)
