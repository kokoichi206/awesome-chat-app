package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kokoichi206/awesome-chat-app/backend/model/response"
	"github.com/opentracing/opentracing-go"
)

func (h *handler) GetRooms(c *gin.Context) {
}

func (h *handler) PatchRoom(c *gin.Context) {
}

func (h *handler) PostRoom(c *gin.Context) {
}

func (h *handler) GetRoomUsers(c *gin.Context) {
	ctx := c.Request.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.GetRoomUsers")
	defer span.Finish()

	roomID := c.Param("roomID")
	if roomID == "" {
		c.Status(http.StatusBadRequest)

		return
	}

	users, err := h.usecase.GetRoomUsers(ctx, roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.GetRoomUsers{
		Users: users,
	})
}
