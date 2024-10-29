package models

import (
	"time"
)

type User struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        Password  `json:"-"`
	Activated       bool      `json:"activated"`
	Version         int       `json:"-"`
	ActivationToken string    `json:"activation_token,omitempty"` // Yeni alan

}

type Password struct {
	Hash string
}
