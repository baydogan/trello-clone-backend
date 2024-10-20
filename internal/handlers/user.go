package handlers

import (
	"context"
	"net/http"
	"trello-clone-backend/internal/errors"
	"trello-clone-backend/internal/helpers"
)

func (h *Handler) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := helpers.ReadJSON(w, r, &input)
	if err != nil {
		errors.BadRequestResponse(w, r, err)
		return
	}

	err = h.userService.RegisterUser(context.Background(), input.Name, input.Email, input.Password)
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
