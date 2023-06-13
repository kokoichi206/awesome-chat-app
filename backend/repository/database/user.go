package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
)

const selectUserStmt = `
SELECT
	id, username, email, profile_info, profile_picture_url, created_at, updated_at
FROM users
WHERE email = $1;
`

func (d *database) SelectUser(ctx context.Context, email string) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.SelectUser")
	defer span.Finish()

	m := model.User{}
	prof := sql.NullString{}
	row := d.db.QueryRowContext(ctx, selectUserStmt, email)

	err := row.Scan(&m.ID, &m.Name, &m.Email, &prof, &m.PictureURL, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	return &m, nil
}

const upsertUserStmt = `
INSERT INTO users (
	username, email, profile_picture_url, created_at, updated_at
) VALUES ($1, $2, $3, $4, $4)
ON CONFLICT (email)
DO UPDATE SET
	username = $1,
	profile_picture_url = $3,
	updated_at = $4;
`

func (d *database) UpsertUser(ctx context.Context, name, email, pictureUrl string, updatedAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.UpsertUser")
	defer span.Finish()

	_, err := d.db.ExecContext(ctx, upsertUserStmt, name, email, pictureUrl, updatedAt)
	if err != nil {
		return fmt.Errorf("failed to upsert user: %w", err)
	}

	return nil
}
