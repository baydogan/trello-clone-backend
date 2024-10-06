package data

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json: "name"`
	Password  password  `json:"_"`
	Activated bool      `json:"-"`
}

type password struct {
	plaintext *string
	hash      []byte
}

type UserModel struct {
	DB *mongo.Client
}
