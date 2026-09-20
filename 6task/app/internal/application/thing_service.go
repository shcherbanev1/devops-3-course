package application

import (
	"log/slog"

	"web-app/internal/domain/thing"
	"web-app/internal/ports"
)

type ThingService struct {
	domainService *thing.Service
	metrics       ports.Metrics
}

func NewThingService(
	domainService *thing.Service,
	metrics ports.Metrics,
) *ThingService {
	return &ThingService{
		domainService: domainService,
		metrics:       metrics,
	}
}

func (s *ThingService) Create(name string) (thing.Thing, error) {
	result, err := s.domainService.Create(name)
	if err != nil {
		slog.Error(
			"failed to create thing",
			"thing_name", name,
			"error", err,
		)

		return thing.Thing{}, err
	}

	s.metrics.ThingCreated()

	slog.Info(
		"thing created",
		"thing_id", result.ID,
		"thing_name", result.Name,
	)

	return result, nil
}

func (s *ThingService) GetAll() []thing.Thing {
	things := s.domainService.GetAll()

	slog.Info(
		"things requested",
		"count", len(things),
	)

	return things
}