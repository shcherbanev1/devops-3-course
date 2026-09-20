package http

import (
	"encoding/json"
	"net/http"
	"log/slog"

	"web-app/internal/adapters/http/dto"
	"web-app/internal/ports"
)

type Handler struct {
	service ports.ThingService
}

func NewHandler(service ports.ThingService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateThing(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateThingRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		slog.Error(
			"failed to decode request",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)

		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	thing, err := h.service.Create(request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := dto.ThingResponse{
		ID:   thing.ID,
		Name: thing.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	things := h.service.GetAll()

	response := make([]dto.ThingResponse, 0, len(things))

	for _, item := range things {
		response = append(response, dto.ThingResponse{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}
