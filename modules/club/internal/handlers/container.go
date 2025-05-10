package handlers

import (
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/internal/handlers/benefit_items"
	"app/modules/club/internal/handlers/benefits"
	"app/modules/club/internal/handlers/contractors"
	"app/modules/club/internal/handlers/healthz"
	"app/modules/club/internal/handlers/partners"
	"app/modules/club/internal/handlers/users"

	"github.com/labstack/echo/v4"
)

type Container struct {
	HealthzHandler      healthz.Handler
	UsersHandler        users.Handler
	ContractorsHandler  contractors.Handler
	PartnersHandler     partners.Handler
	BenefitsHandler     benefits.Handler
	BenefitItemsHandler benefit_items.Handler
}

func NewContainer(
	userService services.IUserServices,
	healthzService services.IHealthzServices,
	contractorService services.IContractorServices,
	partnerService services.IPartnerServices,
	benefitService services.IBenefitServices,
	benefitItemService services.IBenefitItemServices,
) *Container {
	return &Container{
		UsersHandler:        users.NewHandler(userService),
		HealthzHandler:      healthz.NewHandler(healthzService),
		ContractorsHandler:  contractors.NewHandler(contractorService),
		PartnersHandler:     partners.NewHandler(partnerService),
		BenefitsHandler:     benefits.NewHandler(benefitService),
		BenefitItemsHandler: benefit_items.NewHandler(benefitItemService),
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

	// partners routers
	partners := api.Group("/partners")
	partners.POST("", c.PartnersHandler.Create)

	// benefits routers
	benefits := api.Group("/benefits")
	benefits.POST("", c.BenefitsHandler.Create)

	// benefit items routers
	benefitItems := api.Group("/benefit-items")
	benefitItems.POST("", c.BenefitItemsHandler.Create)
}
