package services

import (
	"net/http"
	"trello-clone-backend/internal/helpers"
	"trello-clone-backend/internal/repositories"
)

type HealthService interface {
	HealthCheck(w http.ResponseWriter) error
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

func (h *healthService) HealthCheck(w http.ResponseWriter) error {
	err := h.healthRepo.Ping()
	if err != nil {
		return err
	}

	response := helpers.Envelope{
		"message": "Health check is ok",
	}

	helpers.WriteJSON(w, http.StatusOK, response, nil)
	return nil
}
