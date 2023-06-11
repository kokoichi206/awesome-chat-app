package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/usecase"
	"github.com/kokoichi206/awesome-chat-app/backend/util/logger"
	"github.com/kokoichi206/awesome-chat-app/openapi/gen/go/openapi"
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
	// TODO: openapi に移行する？
	h.setupRoutes(r)

	openapi.RegisterHandlers(r, h)

	return h
}

func (h *handler) setupRoutes(r *gin.Engine) {
	r.GET("/health", h.Health)
}

func (h *handler) Health(c *gin.Context) {
	ctx := c.Request.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.Health")
	defer span.Finish()

	c.JSON(http.StatusOK, gin.H{
		"health": "ok",
	})
}
