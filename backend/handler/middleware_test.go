package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/kokoichi206/awesome-chat-app/backend/handler"
)

func TestAuthMiddleware(t *testing.T) {
	t.Parallel()

	path := "/test-path"
	testCases := map[string]struct {
		cookie        string
		makeMock      func(m *MockUsecase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		"success": {
			cookie: "ok-session",
			makeMock: func(m *MockUsecase) {
				m.
					EXPECT().
					VerifySessionCookie(gomock.Any(), "ok-session").
					Return(&auth.Token{UID: "ok-user"}, nil).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				t.Helper()
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		"failure: no cookie": {
			makeMock: func(m *MockUsecase) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				t.Helper()
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		"failure: invalid cookie": {
			cookie: "ng-session",
			makeMock: func(m *MockUsecase) {
				m.
					EXPECT().
					VerifySessionCookie(gomock.Any(), "ng-session").
					Return(nil, errors.New("error in test")).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				t.Helper()
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
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

			m := NewMockUsecase(ctrl)
			tc.makeMock(m)

			h := handler.New(nil, m)
			_, r := gin.CreateTestContext(httptest.NewRecorder())

			r.GET(
				path,
				h.SessionCheck(),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, path, nil)

			if tc.cookie != "" {
				req.AddCookie(&http.Cookie{
					Name:  "awesome-chat-app-session",
					Value: tc.cookie,
				})
			}

			// Act
			r.ServeHTTP(recorder, req)

			// Assert
			tc.checkResponse(t, recorder)
		})
	}
}
