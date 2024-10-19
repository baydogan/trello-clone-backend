package handlers

import (
	"net/http"
	"trello-clone-backend/internal/errors"
	"trello-clone-backend/internal/services"
)

type HealthHandler struct {
	healthService services.HealthService
}

type HealthHandlerConfig struct {
	HealthService services.HealthService
}

func NewHealthHandler(c *HealthHandlerConfig) *HealthHandler {
	return &HealthHandler{
		healthService: c.HealthService,
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	err := h.healthService.HealthCheck(w)

	if err != nil {
		// Use error handler to wrap in Envelope
		errors.ServerErrorResponse(w, r, err)
		return
	}
}
