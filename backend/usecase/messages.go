package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
)

func (u *usecase) PostMessage(ctx context.Context, roomID, userID, content string, messageType model.MessageType, postedAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.PostMessage")
	defer span.Finish()

	if err := u.database.InsertMessage(ctx, roomID, userID, content, messageType, postedAt); err != nil {
		return fmt.Errorf("failed to insert messages to db: %w", err)
	}

	return nil
}
