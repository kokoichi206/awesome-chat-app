package usecase

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/kokoichi206/awesome-chat-app/backend/model"
	"github.com/opentracing/opentracing-go"
)

func (u *usecase) GetUser(ctx context.Context, email string) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.GetUser")
	defer span.Finish()

	user, err := u.database.SelectUser(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to select user: %w", err)
	}

	return user, nil
}

func (u *usecase) VerifyIDToken(ctx context.Context, token string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.VerifyIDToken")
	defer span.Finish()

	_, err := u.firebase.VerifyIDToken(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to verify token: %w", err)
	}

	return nil
}

func (u *usecase) PostLogin(ctx context.Context, idToken string) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.PostLogin")
	defer span.Finish()

	token, err := u.firebase.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", fmt.Errorf("failed to verify token: %w", err)
	}

	claims := token.Claims
	if err := u.database.UpsertUser(ctx, claims["name"].(string), claims["email"].(string), claims["picture"].(string), time.Now()); err != nil {
		return "", fmt.Errorf("failed to upsert user: %w", err)
	}

	session, err := u.firebase.CreateSession(ctx, idToken)
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

func (u *usecase) VerifySessionCookie(ctx context.Context, session string) (*auth.Token, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.VerifySessionCookie")
	defer span.Finish()

	token, err := u.firebase.VerifySessionCookie(ctx, session)
	if err != nil {
		return nil, fmt.Errorf("failed to verify session: %w", err)
	}

	return token, nil
}
