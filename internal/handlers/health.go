package handlers

import (
	"net/http"
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
	h.healthService.HealthCheck(w)
}
