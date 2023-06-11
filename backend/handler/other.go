package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func (h *handler) Health(c *gin.Context) {
	ctx := c.Request.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.Health")
	defer span.Finish()

	c.JSON(http.StatusOK, gin.H{
		"health": "ok",
	})
}
