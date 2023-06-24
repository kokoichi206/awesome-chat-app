package handler

import (
	"encoding/json"
	"net/http"

	auth "firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/httphead"
	"github.com/gobwas/ws"
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

func (h *handler) SubscribeMessages(c *gin.Context) {
	w := c.Writer
	r := c.Request

	ctx := r.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.SubscribeMessages")
	defer span.Finish()

	u := ws.HTTPUpgrader{
		Extension: func(opt httphead.Option) bool {
			return false
		},
	}

	conn, _, _, err := u.Upgrade(r, w)
	if err != nil {
		c.Status(http.StatusInternalServerError)

		return
	}
	defer conn.Close()

	tc, ok := c.Get(tokenContext)
	if !ok {
		c.Status(http.StatusInternalServerError)

		return
	}

	token, ok := (tc).(auth.Token)
	if !ok {
		c.Status(http.StatusInternalServerError)

		return
	}

	// 識別子として email を使う。
	email := token.Claims["email"].(string)

	if err := h.usecase.SubscribeMessages(ctx, &conn, email); err != nil {
		h.logger.Warnf(ctx, "failed to subscribe messages: %v", err)

		return
	}
}
