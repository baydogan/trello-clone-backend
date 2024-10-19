package services

import (
	"net/http"
	"trello-clone-backend/internal/helpers"
	"trello-clone-backend/internal/repositories"
)

type HealthService interface {
	HealthCheck(w http.ResponseWriter)
}

type healthService struct {
	healthRepo repositories.HealthRepository
}

type HealthServiceConfig struct {
	HealthRepository repositories.HealthRepository
}

func NewHealthService(c *HealthServiceConfig) HealthService {
	return &healthService{
		healthRepo: c.HealthRepository,
	}
}

func (h *healthService) HealthCheck(w http.ResponseWriter) {
	err := h.healthRepo.Ping()
	if err != nil {
		response := helpers.BuildErrorResponse(500, "Database is down")
		helpers.WriteJSON(w, 500, response, nil)
		return
	}

	resp := helpers.BuildSuccessResponse(200, "Health check is ok", "")
	helpers.WriteJSON(w, 200, resp, nil)
}
