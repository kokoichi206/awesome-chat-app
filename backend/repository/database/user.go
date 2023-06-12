package database

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
)

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
