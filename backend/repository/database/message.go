package database

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
	"github.com/kokoichi206/awesome-chat-app/backend/model/response"
)

const selectMessagesStmt = `
SELECT
	id,
	user_id,
	type,
	content,
	posted_at
FROM messages
WHERE
	room_id = $1
AND
	posted_at >= $2;
`

func (d *database) SelectMessages(ctx context.Context, roomID, userID string, lastReadAt time.Time) ([]*response.Message, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.SelectMessages")
	defer span.Finish()

	rows, err := d.db.QueryContext(ctx, selectMessagesStmt, roomID, lastReadAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert messages: %w", err)
	}

	resp := []*response.Message{}

	for rows.Next() {
		var msg response.Message

		if err := rows.Scan(&msg.ID, &msg.UserID, &msg.Type, &msg.Content, &msg.PostedAt); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		resp = append(resp, &msg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return resp, nil
}

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
