package repository

import (
	"context"
	"time"

	"firebase.google.com/go/v4/auth"
)

// interface for database operations.
type Database interface {
	UpsertUser(ctx context.Context, name, email, pictureUrl string, updatedAt time.Time) error
}

type Firebase interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
	CreateSession(ctx context.Context, idToken string) (string, error)
	VerifySessionCookie(ctx context.Context, cookie string) (*auth.Token, error)
}
