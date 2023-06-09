package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/usecase"
	"github.com/kokoichi206/awesome-chat-app/backend/util/logger"
)

type handler struct {
	logger  logger.Logger
	usecase usecase.Usecase

	Engine *gin.Engine
}

func New(logger logger.Logger, usecase usecase.Usecase) *handler {
	r := gin.Default()

	h := &handler{
		logger:  logger,
		usecase: usecase,
		Engine:  r,
	}
	h.setupRoutes()

	return h
}

func (h *handler) setupRoutes() {
	h.Engine.GET("/health", h.Health)
}

func (h *handler) Health(c *gin.Context) {
	ctx := c.Request.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.Health")
	defer span.Finish()

	c.JSON(http.StatusOK, gin.H{
		"health": "ok",
	})
}
