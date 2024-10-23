package handlers

import (
	"trello-clone-backend/internal/database"
	"trello-clone-backend/internal/repositories"
	r "trello-clone-backend/internal/repositories"
	"trello-clone-backend/internal/services"
	s "trello-clone-backend/internal/services"
)

type Handler struct {
	userService   s.UserService
	healthService s.HealthService
}

type HandlerConfig struct {
	UserService   s.UserService
	HealthService s.HealthService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		userService:   c.UserService,
		healthService: c.HealthService,
	}
}

func InitAllHandlers() *Handler {
	client := database.New().GetClient()

	userRepo := r.NewUserRepository((&r.UserRepoConfig{Client: client}))
	healthRepo := repositories.NewHealthRepository((&r.HealthRepoConfig{Client: client}))

	userService := s.NewUserService(&s.UserServiceConfig{UserRepository: userRepo})
	healthService := services.NewHealthService(&s.HealthServiceConfig{HealthRepository: healthRepo})

	h := New(&HandlerConfig{
		UserService:   userService,
		HealthService: healthService,
	})

	return h
}
