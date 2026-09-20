package thing

import (
	"errors"
	"strings"
	"sync"
)

type Service struct {
	mu     sync.RWMutex
	things []Thing
	nextID int64
}

func NewService() *Service {
	return &Service{
		things: make([]Thing, 0),
		nextID: 1,
	}
}

func (s *Service) Create(name string) (Thing, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return Thing{}, errors.New("thing name cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	thing := Thing{
		ID:   s.nextID,
		Name: name,
	}

	s.things = append(s.things, thing)
	s.nextID++

	return thing, nil
}

func (s *Service) GetAll() []Thing {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Thing, len(s.things))
	copy(result, s.things)

	return result
}
