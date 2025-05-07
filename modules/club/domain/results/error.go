package results

// Error representa um erro da aplicação com código e status HTTP
type Error struct {
	Tag        string `json:"tag"`
	Message    string `json:"message"`
	StatusCode int
}

// Error implementa a interface error
func (e *Error) Error() string {
	return e.Message
}

// NewError cria um novo erro
func NewError(tag string, message string, statusCode int) *Error {
	return &Error{
		Tag:        tag,
		Message:    message,
		StatusCode: statusCode,
	}
}
