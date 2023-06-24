package usecase

import (
	"context"
	"net"
	"time"

	"firebase.google.com/go/v4/auth"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
	"github.com/kokoichi206/awesome-chat-app/backend/model/response"
	"github.com/kokoichi206/awesome-chat-app/backend/repository"
	"github.com/kokoichi206/awesome-chat-app/backend/util/logger"
)

type Usecase interface {
	GetUser(ctx context.Context, email string) (*model.User, error)
	VerifyIDToken(ctx context.Context, token string) error
	PostLogin(ctx context.Context, token string) (string, error)
	VerifySessionCookie(ctx context.Context, session string) (*auth.Token, error)

	GetRoomUsers(ctx context.Context, roomID string) ([]*response.RoomUser, error)

	GetMessages(ctx context.Context, roomID, userID string, lastReadAt time.Time) ([]*response.Message, error)
	PostMessage(ctx context.Context, roomID, userID, content string, messageType model.MessageType, postedAt time.Time) error
	SubscribeMessages(ctx context.Context, conn *net.Conn, email string) error
}

type usecase struct {
	database repository.Database
	firebase repository.Firebase

	logger logger.Logger

	// userID -> conn
	msgConns map[string]*net.Conn
}

func New(database repository.Database, firebase repository.Firebase,
	logger logger.Logger) Usecase {
	usecase := &usecase{
		database: database,
		firebase: firebase,
		logger:   logger,

		msgConns: map[string]*net.Conn{},
	}

	return usecase
}
