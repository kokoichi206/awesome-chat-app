package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/kokoichi206/awesome-chat-app/backend/model/response"
	"github.com/kokoichi206/awesome-chat-app/backend/usecase"
)

func TestGetRoomUsers(t *testing.T) {
	t.Parallel()

	type args struct {
		roomID string
	}

	testCases := map[string]struct {
		args     args
		makeMock func(m *MockDatabase)
		want     []*response.RoomUser
		wantErr  string
	}{
		"success": {
			args: args{
				roomID: "c36d313b-bb54-464c-a1b6-fb8ffc728e6a",
			},
			makeMock: func(m *MockDatabase) {
				m.
					EXPECT().
					SelectRoomUsers(gomock.Any(), "c36d313b-bb54-464c-a1b6-fb8ffc728e6a").
					Return([]*response.RoomUser{
						{
							ID:         "553b32e0-3c06-4b13-a5cc-9bf8f534eacd",
							Name:       "kokoichi206",
							PictureURL: "https://lh3.goo",
							LastReadAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
						},
					}, nil)
			},
			want: []*response.RoomUser{
				{
					ID:         "553b32e0-3c06-4b13-a5cc-9bf8f534eacd",
					Name:       "kokoichi206",
					PictureURL: "https://lh3.goo",
					LastReadAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		"failure: select users": {
			args: args{
				roomID: "c36d313b-bb54-464c-a1b6-fb8ffc728e6a",
			},
			makeMock: func(m *MockDatabase) {
				m.
					EXPECT().
					SelectRoomUsers(gomock.Any(), "c36d313b-bb54-464c-a1b6-fb8ffc728e6a").
					Return(nil, errors.New("test error"))
			},
			wantErr: "failed to select users: test error",
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
			got, err := u.GetRoomUsers(context.Background(), tc.args.roomID)

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
