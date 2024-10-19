package repositories

import (
	"context"
	"log"
	"time"
	"trello-clone-backend/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Insert(user models.User) (*models.User, error)
}

type userRepository struct {
	db       *mongo.Collection
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

type UserRepoConfig struct {
	Client *mongo.Client
}

func NewUserRepository(u *UserRepoConfig) UserRepository {

	collection := u.Client.Database("trello").Collection("users")

	return &userRepository{
		db: collection,
	}
}

func (u *userRepository) Insert(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.db.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
