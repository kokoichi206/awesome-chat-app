package usecase

import (
	"context"
	"fmt"

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

func (u *usecase) PostLogin(ctx context.Context, token string) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.PostLogin")
	defer span.Finish()

	session, err := u.firebase.CreateSession(ctx, token)
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
