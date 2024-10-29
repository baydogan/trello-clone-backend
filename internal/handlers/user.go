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

	token, err := helpers.GenaerateJwtToken(input.Email)
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	err = h.userService.SendActivationEmail(input.Email, token)
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := helpers.Envelope{"message": "User registered successfully. Please check your email to activate your account"}
	if err := helpers.WriteJSON(w, http.StatusCreated, response, nil); err != nil {
		errors.ServerErrorResponse(w, r, err)
	}
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := helpers.ReadJSON(w, r, &input); err != nil {
		errors.BadRequestResponse(w, r, err)
		return
	}

	user, token, err := h.userService.Login(context.Background(), input.Email, input.Password)
	if err != nil {
		errors.BadRequestResponse(w, r, err)
		return
	}

	response := helpers.Envelope{
		"user":  user,
		"token": token,
	}

	if err := helpers.WriteJSON(w, http.StatusOK, response, nil); err != nil {
		errors.ServerErrorResponse(w, r, err)
	}
}

func (h *Handler) ActivateUserHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	err := h.userService.ActivateUserWithToken(context.Background(), token)
	if err != nil {
		errors.BadRequestResponse(w, r, err)
		return
	}

	response := helpers.Envelope{"message": "User activated successfully"}
	if err := helpers.WriteJSON(w, http.StatusOK, response, nil); err != nil {
		errors.ServerErrorResponse(w, r, err)
	}
}
