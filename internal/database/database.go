package database

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	GetClient() *mongo.Client
}

type service struct {
	db *mongo.Client
}

var (
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
)

func New() Service {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://root:password1234@%s:%s", host, port)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection established")
	return &service{
		db: client,
	}
}

func (s *service) GetClient() *mongo.Client {
	return s.db
}
