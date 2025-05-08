package users

import (
	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
	"net/http"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) services.IUserServices {
	return &UserService{userRepository: userRepository}
}

var (
	ErrGetUsers  = results.NewError("GET_USERS_ERROR", "error getting users", http.StatusInternalServerError)
	ErrNoContent = results.NewError("NO_CONTENT_ERROR", "user is empty", http.StatusNoContent)
)
