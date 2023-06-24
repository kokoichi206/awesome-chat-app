package repository

import (
	"context"
	"time"

	"firebase.google.com/go/v4/auth"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
)

// interface for database operations.
type Database interface {
	SelectUser(ctx context.Context, email string) (*model.User, error)
	UpsertUser(ctx context.Context, name, email, pictureUrl string, updatedAt time.Time) error

	InsertMessage(ctx context.Context, roomID, userID, content string, messageType model.MessageType, postedAt time.Time) error
}

type Firebase interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
	CreateSession(ctx context.Context, idToken string) (string, error)
	VerifySessionCookie(ctx context.Context, cookie string) (*auth.Token, error)
}
