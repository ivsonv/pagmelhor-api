package results

// Error representa um erro da aplicação com código e status HTTP
type Error struct {
	Tag        string            `json:"tag"`
	Message    string            `json:"message"`
	Details    map[string]string `json:"details,omitempty"`
	StatusCode int               `json:"-"`
}

// Error implementa a interface error
func (e Error) Error() string {
	return e.Message
}

// NewError cria um novo erro
func NewError(tag string, message string, statusCode int) Error {
	return Error{
		Tag:        tag,
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewErrorWithDetails(tag, message string, details map[string]string) Error {
	return Error{
		Tag:     tag,
		Message: message,
		Details: details,
	}
}
