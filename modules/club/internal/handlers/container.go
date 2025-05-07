package handlers

import (
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/internal/handlers/healthz"
	"app/modules/club/internal/handlers/users"

	"github.com/labstack/echo/v4"
)

type Container struct {
	HealthzHandler healthz.Handler
	UsersHandler   users.Handler
}

func NewContainer(userService services.IUserServices, healthzService services.IHealthzServices) *Container {
	return &Container{
		UsersHandler:   users.NewHandler(userService),
		HealthzHandler: healthz.NewHandler(healthzService),
	}
}

func (c *Container) AddRouters(api *echo.Group) {
	// health check
	api.GET("/healthz", c.HealthzHandler.Get)

	// users routers
	page := api.Group("/users")
	page.GET("", c.UsersHandler.Get)
}
