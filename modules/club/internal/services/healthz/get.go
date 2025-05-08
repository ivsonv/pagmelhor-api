package healthz

import (
	responses "app/modules/club/domain/dto/responses/healthz"
	"app/modules/club/domain/results"
	"context"
	"log"
	"time"
)

func (s *HealthzService) Get(ctx context.Context) results.Result[responses.GetHealthzResponseDto] {
	healthz, err := s.healthzRepository.Get(ctx)

	if err != nil {
		log.Printf("Error on Get Healthz Service: %s", err)
		return results.Failure[responses.GetHealthzResponseDto](ErrGetHealthz)
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
