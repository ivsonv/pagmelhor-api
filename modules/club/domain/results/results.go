package results

// Result representa uma resposta padrão da API
type Result[T any] struct {
	Value     T     `json:"value,omitempty"`
	Error     Error `json:"error,omitempty"`
	IsSuccess bool
}

// Success cria um novo resultado de sucesso
func Success[T any](value T) Result[T] {
	return Result[T]{
		IsSuccess: true,
		Value:     value,
	}
}

// Failure cria um novo resultado de erro
func Failure[T any](error Error) Result[T] {
	return Result[T]{
		IsSuccess: false,
		Error:     error,
	}
}

// Paginate representa uma resposta paginada
type Paginate struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}
