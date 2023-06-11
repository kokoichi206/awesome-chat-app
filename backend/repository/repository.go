package repository

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

// interface for database operations.
type Database interface {
}

type Firebase interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
	CreateSession(ctx context.Context, idToken string) (string, error)
	VerifySessionCookie(ctx context.Context, cookie string) (*auth.Token, error)
}
