package database_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kokoichi206/awesome-chat-app/backend/repository/database"
	"github.com/stretchr/testify/assert"
)

func TestUpsertUser(t *testing.T) {
	t.Parallel()

	type args struct {
		name       string
		email      string
		pictureUrl string
		updateAt   time.Time
	}

	mockTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	testCases := map[string]struct {
		args     args
		makeMock func(m sqlmock.Sqlmock)
		wantErr  string
	}{
		"success": {
			args: args{
				name:       "kokoichi206",
				email:      "kokoichi206@test.com",
				pictureUrl: "https://kokoichi206.test.com",
				updateAt:   mockTime,
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectExec(regexp.QuoteMeta(database.UpsertUserStmt)).
					WithArgs("kokoichi206", "kokoichi206@test.com", "https://kokoichi206.test.com", mockTime).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"failure: exec error": {
			args: args{
				name:       "kokoichi206",
				email:      "kokoichi206@test.com",
				pictureUrl: "https://kokoichi206.test.com",
				updateAt:   mockTime,
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectExec(regexp.QuoteMeta(database.UpsertUserStmt)).
					WithArgs("kokoichi206", "kokoichi206@test.com", "https://kokoichi206.test.com", mockTime).
					WillReturnError(errors.New("error in test"))
			},
			wantErr: "failed to upsert user: error in test",
		},
	}

	for name, tc := range testCases {
		name := name
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tc.makeMock(mock)

			database := database.New(db, nil)

			// Act
			err = database.UpsertUser(context.Background(), tc.args.name, tc.args.email, tc.args.pictureUrl, tc.args.updateAt)

			// Assert
			if tc.wantErr == "" {
				assert.Nil(t, err, "error should be nil")
			} else {
				assert.Equal(t, tc.wantErr, err.Error(), "result does not match")
			}
		})
	}
}
