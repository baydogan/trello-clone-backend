package repositories

import (
	"context"
	"log"
	"trello-clone-backend/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
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

func (u *userRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := u.db.InsertOne(ctx, user)
	return err
}
