package domain

import (
	"errors"
	"time"
)

type User struct {
	UserID     int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	RegisterAt time.Time `json:"registered_at"`
}

var (
	ErrCredentilalNotFound = errors.New("credentil not found")
)
