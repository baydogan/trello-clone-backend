package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"trello-clone-backend/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	FindUserByToken(ctx context.Context, token string) (*models.User, error)
	SetUserActive(ctx context.Context, token string) error
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

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {

	filter := bson.M{"email": email}
	result := u.db.FindOne(ctx, filter)

	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindUserByToken(ctx context.Context, token string) (*models.User, error) {
	filter := bson.M{"activation_token": token}
	result := u.db.FindOne(ctx, filter)

	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) SetUserActive(ctx context.Context, email string) error {
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"activated": true}}

	_, err := u.db.UpdateOne(ctx, filter, update)
	return err
}
