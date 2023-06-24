package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/model"
	"github.com/kokoichi206/awesome-chat-app/backend/model/request"
	"github.com/kokoichi206/awesome-chat-app/backend/util"
)

func (h *handler) GetMessages(c *gin.Context) {
}

func (h *handler) PostMessage(c *gin.Context) {
	ctx := c.Request.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.PostMessage")
	defer span.Finish()

	roomID := c.Param("roomID")
	if roomID == "" {
		c.Status(http.StatusBadRequest)

		return
	}

	var body request.PostMessage
	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.Status(http.StatusBadRequest)

		return
	}

	mt, ok := model.MessageTypeStrings[body.Type]
	if !ok {
		c.Status(http.StatusBadRequest)

		return
	}

	postedAt, err := util.FromISO8601(body.PostedAt)
	if err != nil {
		c.Status(http.StatusBadRequest)

		return
	}

	if err := h.usecase.PostMessage(ctx, roomID, body.UserID, body.Content, mt, postedAt); err != nil {
		c.Status(http.StatusInternalServerError)

		return
	}

	c.Status(http.StatusNoContent)
}
