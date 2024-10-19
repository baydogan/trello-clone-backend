package errors

import (
	"log/slog"
	"net/http"
	"trello-clone-backend/internal/helpers"
)

type errors struct {
	logger *slog.Logger
}

func (e *errors) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	e.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (e *errors) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := helpers.Envelope{"error": message}

	err := helpers.WriteJSON(w, status, env, nil)
	if err != nil {
		e.logError(r, err)
		w.WriteHeader(500)
	}
}
