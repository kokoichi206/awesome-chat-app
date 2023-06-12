package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
)

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

func (u *usecase) VerifySessionCookie(ctx context.Context, session string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.VerifySessionCookie")
	defer span.Finish()

	_, err := u.firebase.VerifySessionCookie(ctx, session)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	return nil
}
