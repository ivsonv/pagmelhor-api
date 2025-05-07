package healthz

import (
	"app/modules/club/domain/interfaces/services"
)

type Handler struct {
	HealthzService services.IHealthzServices
}

func NewHandler(healthzService services.IHealthzServices) Handler {
	return Handler{HealthzService: healthzService}
}
