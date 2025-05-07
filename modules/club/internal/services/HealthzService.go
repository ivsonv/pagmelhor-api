package services

import (
	responses "app/modules/club/domain/dto/responses/healthz"
	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
	"context"
	"log"
	"net/http"
	"time"
)

type HealthzService struct {
	healthzRepository repository.IHealthzRepository
}

func NewHealthzService(healthzRepository repository.IHealthzRepository) services.IHealthzServices {
	return &HealthzService{healthzRepository: healthzRepository}
}

func (s *HealthzService) Get(ctx context.Context) results.Result[responses.GetHealthzResponseDto] {
	healthz, err := s.healthzRepository.Get(ctx)

	if err != nil {
		log.Printf("Error on Get Healthz Service: %s", err)

		errorRs := results.Error{
			StatusCode: http.StatusInternalServerError,
			Tag:        "HEALTHZ_SERVICE_GET",
			Message:    err.Error(),
		}

		return results.Failure[responses.GetHealthzResponseDto](errorRs)
	}

	return results.Success(
		responses.GetHealthzResponseDto{
			UpdatedAt: time.Now().UTC(),
			Database: responses.DatabaseStatus{
				AvailableConnections: healthz.AvailableConnections,
				OpenConnections:      healthz.OpenConnections,
				MaxConnections:       healthz.MaxConnections,
				Version:              healthz.Version,
			},
		})
}
