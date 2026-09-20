package ports

import "web-app/internal/domain/thing"

type ThingService interface {
	Create(name string) (thing.Thing, error)
	GetAll() []thing.Thing
}
