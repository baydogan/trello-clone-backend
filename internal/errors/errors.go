package errors

import (
	"fmt"
	"log/slog"
	"net/http"
	"trello-clone-backend/internal/helpers"
)

var logger *slog.Logger

// func InitLogger(l *slog.Logger) {
// 	logger = l
// }

// func logError(r *http.Request, err error) {
// 	var (
// 		method = r.Method
// 		uri    = r.URL.RequestURI()
// 	)

// 	logger.Error(err.Error(), "method", method, "uri", uri)
// }

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := helpers.Envelope{"error": message}

	err := helpers.WriteJSON(w, status, env, nil)
	if err != nil {
		// logError(r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// logError(r, err)

	message := "the server encountered a problem"
	errorResponse(w, r, 500, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested source could not found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported this source", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusUnprocessableEntity, err)
}
