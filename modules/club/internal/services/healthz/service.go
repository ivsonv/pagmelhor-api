package healthz

import (
	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
	"net/http"
)

type HealthzService struct {
	healthzRepository repository.IHealthzRepository
}

func NewHealthzService(healthzRepository repository.IHealthzRepository) services.IHealthzServices {
	return &HealthzService{healthzRepository: healthzRepository}
}

var (
	ErrGetHealthz = results.NewError("GET_HEALTHZ_ERROR", "error getting healthz", http.StatusInternalServerError)
)
