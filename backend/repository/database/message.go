package database

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
)

const insertMessage = `
INSERT INTO messages (
	room_id,
	user_id,
	type,
	content,
	posted_at
) VALUES (
	$1,
	$2,
	(SELECT id FROM message_types WHERE name = $3),
	$4,
	$5
);
`

func (d *database) InsertMessage(ctx context.Context, roomID, userID, content string, messageType model.MessageType, postedAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.InsertMessage")
	defer span.Finish()

	_, err := d.db.ExecContext(ctx, insertMessage, roomID, userID, messageType.String(), content, postedAt)
	if err != nil {
		return fmt.Errorf("failed to insert messages: %w", err)
	}

	return nil
}

const isUserInRoomQuery = `
SELECT COUNT(*)
FROM room_users
WHERE room_id = '' AND user_id = '';
`

// メッセージの送信先の roomID に対象の userID が含まれているかどうかを返す。
func (d *database) IsUserInRoom(ctx context.Context, userID, roomID string) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.IsUserInRoom")
	defer span.Finish()

	var count int

	if err := d.db.QueryRowContext(ctx, isUserInRoomQuery, roomID, userID).Scan(&count); err != nil {
		return false, fmt.Errorf("failed to query: %w", err)
	}

	return count == 1, nil
}
