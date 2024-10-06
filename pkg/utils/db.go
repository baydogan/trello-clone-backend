package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lpernett/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectToDatabase() (*mongo.Client, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	log.Println(username)

	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")

	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Failed to ping mongodb: %v", err)
	}

	log.Println("Connecting to database...")
	Collection = client.Database("trello").Collection("users")

	return client, nil
}
