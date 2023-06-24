package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
)

func (u *usecase) PostMessage(ctx context.Context, roomID, userID, content string, messageType model.MessageType, postedAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.PostMessage")
	defer span.Finish()

	if err := u.database.InsertMessage(ctx, roomID, userID, content, messageType, postedAt); err != nil {
		return fmt.Errorf("failed to insert messages to db: %w", err)
	}

	if err := u.publishMessage(ctx, roomID, userID, content, messageType, postedAt); err != nil {
		return fmt.Errorf("failed to publish messages: %w", err)
	}

	return nil
}

func (u *usecase) SubscribeMessages(ctx context.Context, conn *net.Conn, email string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.SubscribeMessages")
	defer span.Finish()

	user, err := u.database.SelectUser(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	u.msgConns[user.ID.String()] = conn

	pingInterval := 55 * time.Second

	ticker := time.NewTicker(pingInterval)
	for range ticker.C {
		if err := wsutil.WriteServerMessage(*conn, ws.OpPing, nil); err != nil {
			// close された conn は削除。
			if errors.Is(err, syscall.EPIPE) {
				delete(u.msgConns, user.ID.String())

				return nil
			}

			return fmt.Errorf("failed to ping: %w", err)
		}
	}

	return nil
}

func (u *usecase) publishMessage(ctx context.Context, roomID, userID, content string, messageType model.MessageType, postedAt time.Time) error {
	type RMS struct {
		RoomID      string    `json:"room_id"`
		UserID      string    `json:"user_id"`
		MessageType string    `json:"type"`
		Content     string    `json:"content"`
		PostedAt    time.Time `json:"posted_at"`
	}

	body, err := json.Marshal(RMS{
		RoomID:      roomID,
		UserID:      userID,
		MessageType: messageType.String(),
		Content:     content,
		PostedAt:    postedAt,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal to json: %w", err)
	}

	for uid, conn := range u.msgConns {
		if uid == userID || !u.canSubscribe(ctx, uid, roomID) {
			continue
		}

		if err := wsutil.WriteServerMessage(*conn, ws.OpText, body); err != nil {
			// close された conn は削除。
			if errors.Is(err, syscall.EPIPE) {
				delete(u.msgConns, uid)

				continue
			}

			u.logger.Errorf(ctx, "failed to write message: %v", err)
		}
	}

	return nil
}

func (u *usecase) canSubscribe(ctx context.Context, userID, gruopID string) bool {
	ok, err := u.database.IsUserInRoom(ctx, userID, gruopID)
	if err != nil {
		u.logger.Warnf(ctx, "failed to check if user is in room: %v", err)

		return false
	}

	return ok
}
