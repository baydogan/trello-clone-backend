package handlers

import (
	"context"
	"net/http"
	"trello-clone-backend/internal/errors"
	"trello-clone-backend/internal/helpers"
	"trello-clone-backend/internal/services"
)

type UserHandler struct {
	userService services.UserService
}

type UserHandlerConfig struct {
	UserService services.UserService
}

func NewUserHandler(c *UserHandler) *UserHandler {
	return &UserHandler{
		userService: c.userService,
	}
}

func (h *Handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
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
