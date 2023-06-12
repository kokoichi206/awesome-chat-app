package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Profile    string
	PictureURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
