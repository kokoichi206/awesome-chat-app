package database_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/kokoichi206/awesome-chat-app/backend/model/response"
	"github.com/kokoichi206/awesome-chat-app/backend/repository/database"
)

func TestSelectRoomUsers(t *testing.T) {
	t.Parallel()

	type args struct {
		roomID string
	}

	testCases := map[string]struct {
		args     args
		makeMock func(m sqlmock.Sqlmock)
		want     []*response.RoomUser
		wantErr  string
	}{
		"success": {
			args: args{
				roomID: "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.SelectRoomUsersStmt)).
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21").
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "username", "profile_info", "profile_picture_url", "last_read_at"}).
							AddRow("6cfd4322-6a2f-4e69-ae21-db6420488816", "kokoichi206", "", "https://kokoichi206.test.com", time.Date(2023, 06, 23, 11, 00, 34, 512627, time.UTC)),
					)
			},
			want: []*response.RoomUser{
				{
					ID:         "6cfd4322-6a2f-4e69-ae21-db6420488816",
					Name:       "kokoichi206",
					PictureURL: "https://kokoichi206.test.com",
					LastReadAt: time.Date(2023, 06, 23, 11, 00, 34, 512627, time.UTC),
				},
			},
		},
		"failure: query error": {
			args: args{
				roomID: "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.SelectRoomUsersStmt)).
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21").
					WillReturnError(errors.New("error in test"))
			},
			wantErr: "failed to select users: error in test",
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
			got, err := database.SelectRoomUsers(context.Background(), tc.args.roomID)

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
