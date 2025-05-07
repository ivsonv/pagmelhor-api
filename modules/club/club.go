package club

import (
	"app/configs"
	"app/modules/club/internal/handlers"
	"app/modules/club/internal/repositories"
	"app/modules/club/internal/services"
	"log"

	databases "app/modules/club/libs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(api *echo.Group, cfg *configs.Config) {
	log.Println("preparing api club...")
	initConfigServer(api)

	log.Println("setting api club cors...")
	setConfigCors(api)

	log.Println("setting api club routers...")
	setRouters(api, cfg)
}

func initConfigServer(api *echo.Group) {
	api.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	api.Use(middleware.BodyLimit("1MB"))
	api.Use(middleware.Recover())
	api.Use(middleware.Logger())
}

func setConfigCors(api *echo.Group) {
	AllowMethods := []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS}
	AllowHeaders := []string{"Origin", "Content-Type", "Authorization"}
	AllowOrigins := []string{"*"}

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: AllowOrigins,
		AllowMethods: AllowMethods,
		AllowHeaders: AllowHeaders,
	}))
}

func setRouters(api *echo.Group, cfg *configs.Config) {
	container := getContainer(cfg)
	container.AddRouters(api)
}

// getContainer loads the container dependencies
func getContainer(cfg *configs.Config) *handlers.Container {
	// drive de connection
	db := databases.NewPostgres(cfg)

	// repositories
	repository := repositories.NewRepository(db)
	userRepository := repositories.NewUserRepository(repository)
	healthzRepository := repositories.NewHealthzRepository(repository)

	// services
	userService := services.NewUserService(userRepository)
	healthzService := services.NewHealthzService(healthzRepository)

	// handlers
	return handlers.NewContainer(userService, healthzService)
}
