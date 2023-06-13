package database_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/kokoichi206/awesome-chat-app/backend/model"
	"github.com/kokoichi206/awesome-chat-app/backend/repository/database"
	"github.com/stretchr/testify/assert"
)

func TestSelectUser(t *testing.T) {
	t.Parallel()

	type args struct {
		email string
	}

	testCases := map[string]struct {
		args     args
		makeMock func(m sqlmock.Sqlmock)
		want     *model.User
		wantErr  string
	}{
		"success": {
			args: args{
				email: "kokoichi206@test.com",
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.SelectUserStmt)).
					WithArgs("kokoichi206@test.com").
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "username", "email", "profile_info", "profile_picture_url", "created_at", "updated_at"}).
							AddRow("6cfd4322-6a2f-4e69-ae21-db6420488816", "kokoichi206", "kokoichi206@test.com", "", "https://kokoichi206.test.com", time.Time{}, time.Time{}),
					)
			},
			want: &model.User{
				ID:         uuid.MustParse("6cfd4322-6a2f-4e69-ae21-db6420488816"),
				Name:       "kokoichi206",
				Email:      "kokoichi206@test.com",
				PictureURL: "https://kokoichi206.test.com",
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
			},
		},
		"failure: query error": {
			args: args{
				email: "kokoichi206@test.com",
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.SelectUserStmt)).
					WithArgs("kokoichi206@test.com").
					WillReturnError(errors.New("error in test"))
			},
			wantErr: "failed to scan user: error in test",
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
			got, err := database.SelectUser(context.Background(), tc.args.email)

			// Assert
			assert.Equal(t, tc.want, got, "result does not match")
			if tc.wantErr == "" {
				assert.Nil(t, err, "error should be nil")
			} else {
				assert.Equal(t, tc.wantErr, err.Error(), "result does not match")
			}
		})
	}
}

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
