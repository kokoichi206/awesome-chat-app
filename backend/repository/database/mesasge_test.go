package database_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
	"github.com/kokoichi206/awesome-chat-app/backend/repository/database"
)

func TestSelectMessages(t *testing.T) {
	t.Parallel()

	type args struct {
		roomID     string
		lastReadAt time.Time
	}

	testCases := map[string]struct {
		args     args
		makeMock func(m sqlmock.Sqlmock)
		want     []*model.Message
		wantErr  string
	}{
		"success": {
			args: args{
				roomID:     "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
				lastReadAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.SelectMessagesStmt)).
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21", time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "user_id", "type", "content", "posted_at"}).
							AddRow("6cfd4322-6a2f-4e69-ae21-db6420488816", "4de5a8af-1629-42de-98cd-4de5a8afbcb7", "text", "pien", time.Date(2023, 06, 23, 11, 00, 34, 512627, time.UTC)),
					)
			},
			want: []*model.Message{
				{
					ID:       "6cfd4322-6a2f-4e69-ae21-db6420488816",
					UserID:   "4de5a8af-1629-42de-98cd-4de5a8afbcb7",
					Type:     "text",
					Content:  "pien",
					PostedAt: time.Date(2023, 06, 23, 11, 00, 34, 512627, time.UTC),
				},
			},
		},
		"failure: query error": {
			args: args{
				roomID:     "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
				lastReadAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.SelectMessagesStmt)).
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21", time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)).
					WillReturnError(errors.New("error in test"))
			},
			wantErr: "failed to select messages: error in test",
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
			got, err := database.SelectMessages(context.Background(), tc.args.roomID, tc.args.lastReadAt)

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

func TestInsertMessage(t *testing.T) {
	t.Parallel()

	type args struct {
		roomID      string
		userID      string
		content     string
		messageType model.MessageType
		postedAt    time.Time
	}

	testCases := map[string]struct {
		args     args
		makeMock func(m sqlmock.Sqlmock)
		wantErr  string
	}{
		"success": {
			args: args{
				roomID:      "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
				userID:      "b4619b69-1629-42de-98cd-4de5a8afbcb7",
				content:     "paon",
				messageType: model.MessageTypeText,
				postedAt:    time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectExec(regexp.QuoteMeta(database.InsertMessage)).
					// MessageTypeText は text で保存される。
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21", "b4619b69-1629-42de-98cd-4de5a8afbcb7", "text", "paon", time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"failure: query error": {
			args: args{
				roomID:      "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
				userID:      "b4619b69-1629-42de-98cd-4de5a8afbcb7",
				content:     "paon",
				messageType: model.MessageTypeText,
				postedAt:    time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectExec(regexp.QuoteMeta(database.InsertMessage)).
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21", "b4619b69-1629-42de-98cd-4de5a8afbcb7", "text", "paon", time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)).
					WillReturnError(errors.New("error in test"))
			},
			wantErr: "failed to insert message: error in test",
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
			got := database.InsertMessage(context.Background(), tc.args.roomID, tc.args.userID, tc.args.content, tc.args.messageType, tc.args.postedAt)

			// Assert
			if tc.wantErr == "" {
				assert.Nil(t, got, "error should be nil")
			} else {
				assert.Equal(t, tc.wantErr, got.Error(), "result does not match")
			}
		})
	}
}

func TestIsUserInRoom(t *testing.T) {
	t.Parallel()

	type args struct {
		roomID string
		userID string
	}

	testCases := map[string]struct {
		args     args
		makeMock func(m sqlmock.Sqlmock)
		want     bool
		wantErr  string
	}{
		"success": {
			args: args{
				roomID: "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
				userID: "b4619b69-1629-42de-98cd-4de5a8afbcb7",
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.IsUserInRoomQuery)).
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21", "b4619b69-1629-42de-98cd-4de5a8afbcb7").
					WillReturnRows(
						sqlmock.NewRows([]string{"count"}).
							AddRow(1))

			},
			want: true,
		},
		"success: haven't joined": {
			args: args{
				roomID: "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
				userID: "b4619b69-1629-42de-98cd-4de5a8afbcb7",
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.IsUserInRoomQuery)).
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21", "b4619b69-1629-42de-98cd-4de5a8afbcb7").
					WillReturnRows(
						sqlmock.NewRows([]string{"count"}).
							AddRow(0))

			},
			want: false,
		},
		"failure: query error": {
			args: args{
				roomID: "9a1eaad6-0572-46be-95fc-6c3edb3bdf21",
				userID: "b4619b69-1629-42de-98cd-4de5a8afbcb7",
			},
			makeMock: func(m sqlmock.Sqlmock) {
				m.
					ExpectQuery(regexp.QuoteMeta(database.IsUserInRoomQuery)).
					WithArgs("9a1eaad6-0572-46be-95fc-6c3edb3bdf21", "b4619b69-1629-42de-98cd-4de5a8afbcb7").
					WillReturnError(errors.New("error in test"))
			},
			want:    false,
			wantErr: "failed to query: error in test",
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
			got, err := database.IsUserInRoom(context.Background(), tc.args.userID, tc.args.roomID)

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
