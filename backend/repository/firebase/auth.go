package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/opentracing/opentracing-go"
)

func (a *admin) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.VerifyIDToken")
	defer span.Finish()

	token, err := a.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("failed to call firebase sdk: %w", err)
	}

	return token, nil
}

func (a *admin) CreateSession(ctx context.Context, idToken string) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.CreateSession")
	defer span.Finish()

	cookie, err := a.authClient.SessionCookie(ctx, idToken, a.sessionLifetime)
	if err != nil {
		return "", fmt.Errorf("failed to call firebase sdk: %w", err)
	}

	return cookie, nil
}

func (a *admin) VerifySessionCookie(ctx context.Context, cookie string) (*auth.Token, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.VerifySessionCookie")
	defer span.Finish()

	token, err := a.authClient.VerifySessionCookie(ctx, cookie)
	if err != nil {
		return nil, fmt.Errorf("failed to call firebase sdk: %w", err)
	}

	return token, nil
}
