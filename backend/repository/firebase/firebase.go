package firebase

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"

	"github.com/kokoichi206/awesome-chat-app/backend/repository"
)

type admin struct {
	app        *firebase.App
	authClient *auth.Client

	sessionLifetime time.Duration
}

func New(ctx context.Context, credentialPath string) (repository.Firebase, error) {
	opt := option.WithCredentialsFile(credentialPath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize app: %w", err)
	}

	ac, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth client : %w", err)
	}

	admin := &admin{
		app:        app,
		authClient: ac,

		// 5 days
		sessionLifetime: 60 * 60 * 24 * 5 * time.Second,
	}

	return admin, nil
}
