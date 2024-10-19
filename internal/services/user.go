package services

import (
	r "trello-clone-backend/internal/repositories"
)

type UserService interface {
}

type userService struct {
	userRepository r.UserRepository
}

type UserServiceConfig struct {
	UserRepository r.UserRepository
}

func NewUserService(u *UserServiceConfig) UserService {
	return &userService{
		userRepository: u.UserRepository,
	}
}
