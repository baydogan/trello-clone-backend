package routes

import (
	"net/http"
	"trello-clone-backend/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

func Routes(h *handlers.Handler) *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/health", h.HealthCheck)

	return router
}
