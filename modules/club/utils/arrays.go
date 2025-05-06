package utils

func IsEmpty[T any](arr []T) bool {
	return len(arr) == 0
}

func IsNotEmpty[T any](arr []T) bool {
	return len(arr) > 0
}

// Map aplica uma funcao a cada elemento do slice e retorna um novo slice com os resultados
func Map[T any, U any](elems []T, fn func(i int) U) []U {
	result := make([]U, len(elems))
	for i := range elems {
		result[i] = fn(i)
	}
	return result
}

// Reduce agrega os elementos do slice em um unico valor de acordo uma função
func Reduce[T any, U any](elems []T, fn func(acc U, i int) U, initialValue U) U {
	result := initialValue
	for i := range elems {
		result = fn(result, i)
	}
	return result
}

// Some verifica se algum elemento satisfaz a condicao
func Some[T any](elems []T, fn func(i int) bool) bool {
	for i := range elems {
		if fn(i) {
			return true
		}
	}
	return false
}

// Every verifica se todos elementos satisfazem a condicao
func Every[T any](elems []T, fn func(i int) bool) bool {
	for i := range elems {
		if !fn(i) {
			return false
		}
	}
	return true
}

// ChunkBy divide um slice em sub-slices de tamanho especificado
func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

// Unique retorna um slice de itens unicos (descartando as duplicatas)
func Unique[T comparable](s []T) []T {
	inResult := make(map[T]struct{})
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = struct{}{}
			result = append(result, str)
		}
	}
	return result
}

// Contains verifica se um slice contem um determinado valor
func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// Filter filtra um slice de acordo com uma funcao de filtro que deve retornar true caso o item deva ser mantido no slice. Similar a funções "retain" em outras linguagens
func Filter[T any](items []T, fn func(i int) bool) []T {
	filteredItems := []T{}
	for i, value := range items {
		if fn(i) {
			filteredItems = append(filteredItems, value)
		}
	}
	return filteredItems
}
