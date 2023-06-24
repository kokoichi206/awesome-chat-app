package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/model/response"
)

const selectRoomUsersStmt = `
SELECT
	user_id,
	u.username,
	u.profile_info,
	u.profile_picture_url,
	last_read_at 
FROM room_users
INNER JOIN users AS u
ON room_users.user_id = u.id
WHERE room_id = $1;
`

func (d *database) SelectRoomUsers(ctx context.Context, roomID string) ([]*response.RoomUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.SelectRoomUsers")
	defer span.Finish()

	rows, err := d.db.QueryContext(ctx, selectRoomUsersStmt, roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to select users: %w", err)
	}

	resp := []*response.RoomUser{}

	for rows.Next() {
		var user response.RoomUser
		var pi sql.NullString

		if err := rows.Scan(&user.ID, &user.Name, &pi, &user.PictureURL, &user.LastReadAt); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		if pi.Valid {
			user.Profile = pi.String
		}

		resp = append(resp, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	return resp, nil
}
