package services

import (
	"context"
	"time"
	"trello-clone-backend/internal/helpers"
	"trello-clone-backend/internal/models"
	r "trello-clone-backend/internal/repositories"
)

type UserService interface {
	RegisterUser(ctx context.Context, name, email, password string) error
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

func (s *userService) RegisterUser(ctx context.Context, name, email, password string) error {
	var user models.User

	user.Name = name
	user.Email = email
	user.CreatedAt = time.Now()
	user.Activated = false
	user.Version = 1

	hashedPassword, err := helpers.Encrypt(password)
	if err != nil {
		return err
	}

	user.Password = models.Password{
		Hash: hashedPassword,
	}

	err = s.userRepository.InsertUser(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
