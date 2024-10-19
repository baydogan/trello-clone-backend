package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type HealthRepository interface {
	Ping() error
}

type healthRepository struct {
	client *mongo.Client
}

type HealthRepoConfig struct {
	Client *mongo.Client
}

func NewHealthRepository(c *HealthRepoConfig) HealthRepository {
	return &healthRepository{
		client: c.Client,
	}
}

func (h *healthRepository) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	return h.client.Ping(ctx, nil)
}
