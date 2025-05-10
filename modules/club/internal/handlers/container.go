package handlers

import (
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/internal/handlers/contractors"
	"app/modules/club/internal/handlers/healthz"
	"app/modules/club/internal/handlers/users"

	"github.com/labstack/echo/v4"
)

type Container struct {
	HealthzHandler     healthz.Handler
	UsersHandler       users.Handler
	ContractorsHandler contractors.Handler
}

func NewContainer(
	userService services.IUserServices,
	healthzService services.IHealthzServices,
	contractorService services.IContractorServices,
) *Container {
	return &Container{
		UsersHandler:       users.NewHandler(userService),
		HealthzHandler:     healthz.NewHandler(healthzService),
		ContractorsHandler: contractors.NewHandler(contractorService),
	}
}

func (c *Container) AddRouters(api *echo.Group) {
	// health check
	api.GET("/healthz", c.HealthzHandler.Get)

	// users routers
	users := api.Group("/users")
	users.GET("", c.UsersHandler.Get)

	// contractors routers
	contractors := api.Group("/contractors")
	contractors.POST("", c.ContractorsHandler.Create)
}
