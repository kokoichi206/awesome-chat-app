package usecase_test

import (
	"context"
	"errors"
	"testing"

	"firebase.google.com/go/v4/auth"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/kokoichi206/awesome-chat-app/backend/usecase"
)

func TestVerifyIDToken(t *testing.T) {
	t.Parallel()

	type args struct {
		token string
	}

	testCases := map[string]struct {
		args     args
		makeMock func(m *MockFirebase)
		wantErr  string
	}{
		"success": {
			args: args{
				token: "ok-token",
			},
			makeMock: func(m *MockFirebase) {
				m.
					EXPECT().
					VerifyIDToken(gomock.Any(), "ok-token").
					Return(&auth.Token{}, nil)
			},
		},
		"failure: verify token": {
			args: args{
				token: "ng-token",
			},
			makeMock: func(m *MockFirebase) {
				m.
					EXPECT().
					VerifyIDToken(gomock.Any(), "ng-token").
					Return(nil, errors.New("error in test"))
			},
			wantErr: "failed to verify token: error in test",
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

			m := NewMockFirebase(ctrl)
			tc.makeMock(m)

			u := usecase.New(nil, m, nil)

			// Act
			err := u.VerifyIDToken(context.Background(), tc.args.token)

			// Assert
			if tc.wantErr == "" {
				assert.Nil(t, err, "error should be nil")
			} else {
				assert.Equal(t, tc.wantErr, err.Error(), "result does not match")
			}
		})
	}
}
func TestPostLogin(t *testing.T) {
	t.Parallel()

	type args struct {
		token string
	}

	testCases := map[string]struct {
		name     string
		args     args
		makeMock func(md *MockDatabase, mf *MockFirebase)
		want     string
		wantErr  string
	}{
		"success": {
			args: args{
				token: "ok-token",
			},
			makeMock: func(md *MockDatabase, mf *MockFirebase) {
				mf.
					EXPECT().
					VerifyIDToken(gomock.Any(), "ok-token").
					Return(&auth.Token{
						Claims: map[string]interface{}{
							"name":    "kokoichi206",
							"email":   "kokoichi206@test.com",
							"picture": "https://kokoichi206.test.com",
						},
					}, nil)
				md.
					EXPECT().
					UpsertUser(gomock.Any(), "kokoichi206", "kokoichi206@test.com", "https://kokoichi206.test.com", gomock.Any()).
					Return(nil)
				mf.
					EXPECT().
					CreateSession(gomock.Any(), "ok-token").
					Return("session-value", nil)
			},
			want: "session-value",
		},
		"failure: verify id token": {
			args: args{
				token: "ok-token",
			},
			makeMock: func(md *MockDatabase, mf *MockFirebase) {
				mf.
					EXPECT().
					VerifyIDToken(gomock.Any(), "ok-token").
					Return(&auth.Token{}, errors.New("error in test"))
			},
			wantErr: "failed to verify token: error in test",
		},
		"failure: upsert": {
			args: args{
				token: "ok-token",
			},
			makeMock: func(md *MockDatabase, mf *MockFirebase) {
				mf.
					EXPECT().
					VerifyIDToken(gomock.Any(), "ok-token").
					Return(&auth.Token{
						Claims: map[string]interface{}{
							"name":    "kokoichi206",
							"email":   "kokoichi206@test.com",
							"picture": "https://kokoichi206.test.com",
						},
					}, nil)
				md.
					EXPECT().
					UpsertUser(gomock.Any(), "kokoichi206", "kokoichi206@test.com", "https://kokoichi206.test.com", gomock.Any()).
					Return(errors.New("error in test"))
			},
			wantErr: "failed to upsert user: error in test",
		},
		"failure: create session": {
			args: args{
				token: "ok-token",
			},
			makeMock: func(md *MockDatabase, mf *MockFirebase) {
				mf.
					EXPECT().
					VerifyIDToken(gomock.Any(), "ok-token").
					Return(&auth.Token{
						Claims: map[string]interface{}{
							"name":    "kokoichi206",
							"email":   "kokoichi206@test.com",
							"picture": "https://kokoichi206.test.com",
						},
					}, nil)
				md.
					EXPECT().
					UpsertUser(gomock.Any(), "kokoichi206", "kokoichi206@test.com", "https://kokoichi206.test.com", gomock.Any()).
					Return(nil)
				mf.
					EXPECT().
					CreateSession(gomock.Any(), "ok-token").
					Return("", errors.New("error in test"))
			},
			wantErr: "failed to create session: error in test",
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

			md := NewMockDatabase(ctrl)
			mf := NewMockFirebase(ctrl)
			tc.makeMock(md, mf)

			u := usecase.New(md, mf, nil)

			// Act
			got, err := u.PostLogin(context.Background(), tc.args.token)

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

func TestVerifySessionCookie(t *testing.T) {
	t.Parallel()

	type args struct {
		session string
	}

	testCases := map[string]struct {
		name     string
		args     args
		makeMock func(m *MockFirebase)
		want     *auth.Token
		wantErr  string
	}{
		"success": {
			args: args{
				session: "ok-session",
			},
			makeMock: func(m *MockFirebase) {
				m.
					EXPECT().
					VerifySessionCookie(gomock.Any(), "ok-session").
					Return(&auth.Token{
						UID: "ok-uid",
					}, nil)
			},
			want: &auth.Token{
				UID: "ok-uid",
			},
		},
		"failure: verify token": {
			args: args{
				session: "ng-session",
			},
			makeMock: func(m *MockFirebase) {
				m.
					EXPECT().
					VerifySessionCookie(gomock.Any(), "ng-session").
					Return(nil, errors.New("error in test"))
			},
			wantErr: "failed to verify session: error in test",
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

			m := NewMockFirebase(ctrl)
			tc.makeMock(m)

			u := usecase.New(nil, m, nil)

			// Act
			got, err := u.VerifySessionCookie(context.Background(), tc.args.session)

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
