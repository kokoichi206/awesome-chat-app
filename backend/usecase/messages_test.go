package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	model "github.com/kokoichi206/awesome-chat-app/backend/model"
	"github.com/kokoichi206/awesome-chat-app/backend/model/response"
	"github.com/kokoichi206/awesome-chat-app/backend/usecase"
)

func TestGetMessages(t *testing.T) {
	t.Parallel()

	type args struct {
		roomID     string
		lastReadAt time.Time
	}

	testCases := map[string]struct {
		args     args
		makeMock func(m *MockDatabase)
		want     []*response.Message
		wantErr  string
	}{
		"success": {
			args: args{
				roomID:     "c36d313b-bb54-464c-a1b6-fb8ffc728e6a",
				lastReadAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			makeMock: func(m *MockDatabase) {
				m.
					EXPECT().
					SelectMessages(gomock.Any(), "c36d313b-bb54-464c-a1b6-fb8ffc728e6a", time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)).
					Return([]*model.Message{
						{
							ID:       "553b32e0-3c06-4b13-a5cc-9bf8f534eacd",
							UserID:   "b4619b69-1629-42de-98cd-4de5a8afbcb7",
							Type:     "text",
							Content:  "Hello",
							PostedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
						},
					}, nil)
			},
			want: []*response.Message{
				{
					ID:       "553b32e0-3c06-4b13-a5cc-9bf8f534eacd",
					UserID:   "b4619b69-1629-42de-98cd-4de5a8afbcb7",
					Type:     "text",
					Content:  "Hello",
					PostedAt: "2021-03-01T00:00:00Z",
				},
			},
		},
		"failure: select messages": {
			args: args{
				roomID:     "c36d313b-bb54-464c-a1b6-fb8ffc728e6a",
				lastReadAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			makeMock: func(m *MockDatabase) {
				m.
					EXPECT().
					SelectMessages(gomock.Any(), "c36d313b-bb54-464c-a1b6-fb8ffc728e6a", time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)).
					Return(nil, errors.New("test error"))
			},
			wantErr: "failed to select messages: test error",
		},
	}

	for name, tc := range testCases {
		name := name
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := NewMockDatabase(ctrl)
			tc.makeMock(m)

			u := usecase.New(m, nil, nil)

			// Act
			got, err := u.GetMessages(context.Background(), tc.args.roomID, tc.args.lastReadAt)

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

func TestPostMessage(t *testing.T) {
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
		makeMock func(m *MockDatabase)
		wantErr  string
	}{
		"success": {
			args: args{
				roomID:      "c36d313b-bb54-464c-a1b6-fb8ffc728e6a",
				userID:      "b4619b69-1629-42de-98cd-4de5a8afbcb7",
				content:     "Hello",
				messageType: model.MessageTypeText,
				postedAt:    time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			makeMock: func(m *MockDatabase) {
				m.
					EXPECT().
					InsertMessage(gomock.Any(), "c36d313b-bb54-464c-a1b6-fb8ffc728e6a", "b4619b69-1629-42de-98cd-4de5a8afbcb7", "Hello", gomock.Any(), time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)).
					Return(nil)
			},
		},
		"failure: insert message": {
			args: args{
				roomID:      "c36d313b-bb54-464c-a1b6-fb8ffc728e6a",
				userID:      "b4619b69-1629-42de-98cd-4de5a8afbcb7",
				content:     "Hello",
				messageType: model.MessageTypeText,
				postedAt:    time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			makeMock: func(m *MockDatabase) {
				m.
					EXPECT().
					InsertMessage(gomock.Any(), "c36d313b-bb54-464c-a1b6-fb8ffc728e6a", "b4619b69-1629-42de-98cd-4de5a8afbcb7", "Hello", gomock.Any(), time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)).
					Return(errors.New("test error"))
			},
			wantErr: "failed to insert messages to db: test error",
		},
	}

	for name, tc := range testCases {
		name := name
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := NewMockDatabase(ctrl)
			tc.makeMock(m)

			u := usecase.New(m, nil, nil)

			// Act
			got := u.PostMessage(context.Background(), tc.args.roomID, tc.args.userID, tc.args.content, tc.args.messageType, tc.args.postedAt)

			// Assert
			if tc.wantErr == "" {
				assert.Nil(t, got, "error should be nil")
			} else {
				assert.Equal(t, tc.wantErr, got.Error(), "result does not match")
			}
		})
	}
}
