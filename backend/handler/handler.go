package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/kokoichi206/awesome-chat-app/backend/usecase"
	"github.com/kokoichi206/awesome-chat-app/backend/util/logger"
)

const (
	sessionCookieName = "awesome-chat-app-session"
	sessionMaxAge     = 60 * 60 * 24 * 7
)

type handler struct {
	logger  logger.Logger
	usecase usecase.Usecase

	Engine *gin.Engine

	sessionCookieName string
	sessionMaxAge     int
}

func New(logger logger.Logger, usecase usecase.Usecase) *handler {
	r := gin.Default()

	h := &handler{
		logger:  logger,
		usecase: usecase,
		Engine:  r,

		sessionCookieName: sessionCookieName,
		sessionMaxAge:     sessionMaxAge,
	}

	h.setupRoutes()

	return h
}

func (h *handler) setupRoutes() {
	r := h.Engine.Group("/api")

	r.GET("/health", h.Health)
	r.POST("/login", h.PostLogin)

	r.Use(h.sessionCheck())

	r.GET("/rooms", h.GetRooms)
	r.PATCH("/rooms", h.PatchRoom)
	r.POST("/rooms", h.PostRoom)
	r.GET("/rooms/:roomID/messages", h.GetMessages)
	r.POST("/rooms/:roomID/messages", h.PostMessage)
	r.GET("/rooms/:roomID/users", h.GetRoomUsers)

	r.GET("/users/followers", h.GetFollowers)
	r.GET("/users/following", h.GetFollowing)
	r.PATCH("/users/following/:user_id", h.PatchFollowing)
	r.POST("/users/following/:user_id", h.PostFollowing)
	r.GET("/users/me", h.GetMe)
	r.PATCH("/users/me", h.PostMe)
	r.GET("/users/:user_id", h.GetUserByID)

	r.GET("/users/:user_id/messages", h.SubscribeMessages)
}
