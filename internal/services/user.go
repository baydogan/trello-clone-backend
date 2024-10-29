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
	Login(ctx context.Context, username, password string) (*models.User, *string, error)
	SendActivationEmail(email, token string) error
	ActivateUserWithToken(ctx context.Context, token string) error
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

func (s *userService) Login(ctx context.Context, email, password string) (*models.User, *string, error) {
	auth, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	if !helpers.CompareEncrypt(password, auth.Password.Hash) {
		return nil, nil, err
	}

	token, err := helpers.GenaerateJwtToken(email)
	if err != nil {
		return nil, nil, err
	}

	return auth, &token, nil
}

func (s *userService) SendActivationEmail(email, token string) error {
	return nil
}

func (s *userService) ActivateUserWithToken(ctx context.Context, token string) error {
	user, err := s.userRepository.FindUserByToken(ctx, token)
	if err != nil {
		return err
	}

	return s.userRepository.SetUserActive(ctx, user.ActivationToken)

}
