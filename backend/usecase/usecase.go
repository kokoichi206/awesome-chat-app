package usecase

import (
	"context"

	"firebase.google.com/go/v4/auth"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
	"github.com/kokoichi206/awesome-chat-app/backend/repository"
	"github.com/kokoichi206/awesome-chat-app/backend/util/logger"
)

type Usecase interface {
	GetUser(ctx context.Context, email string) (*model.User, error)
	VerifyIDToken(ctx context.Context, token string) error
	PostLogin(ctx context.Context, token string) (string, error)
	VerifySessionCookie(ctx context.Context, session string) (*auth.Token, error)
}

type usecase struct {
	database repository.Database
	firebase repository.Firebase

	logger logger.Logger
}

func New(database repository.Database, firebase repository.Firebase,
	logger logger.Logger) Usecase {
	usecase := &usecase{
		database: database,
		firebase: firebase,
		logger:   logger,
	}

	return usecase
}
