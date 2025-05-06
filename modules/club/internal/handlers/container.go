package handlers

import (
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/internal/handlers/users"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type Container struct {
	UsersHandler users.Handler
}

func NewContainer(userService services.IUserServices) *Container {
	return &Container{
		UsersHandler: users.NewHandler(userService),
	}
}

func (c *Container) AddRouters(api *echo.Group) {
	// health check
	api.GET("/health", func(c echo.Context) error {
		now := time.Now().Format("2006-01-02 15:04:05.000000")
		host, _ := os.Hostname()

		result := fmt.Sprintf("Running on %s at %s", host, now)
		return c.JSON(http.StatusOK, result)
	})

	// users routers
	page := api.Group("/users")
	page.GET("", c.UsersHandler.Get)
}
