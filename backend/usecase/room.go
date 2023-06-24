package usecase

import (
	"context"
	"fmt"

	"github.com/kokoichi206/awesome-chat-app/backend/model/response"
	"github.com/opentracing/opentracing-go"
)

func (u *usecase) GetRoomUsers(ctx context.Context, roomID string) ([]*response.RoomUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.GetRoomUsers")
	defer span.Finish()

	msgs, err := u.database.SelectRoomUsers(ctx, roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to select users: %w", err)
	}

	return msgs, nil
}
