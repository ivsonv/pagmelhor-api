package users

import (
	"app/modules/club/domain/interfaces/services"
)

type Handler struct {
	UserService services.IUserServices
}

func NewHandler(userService services.IUserServices) Handler {
	return Handler{UserService: userService}
}
