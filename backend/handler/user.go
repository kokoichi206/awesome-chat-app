package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kokoichi206/awesome-chat-app/openapi/gen/go/openapi"
	"github.com/opentracing/opentracing-go"
)

func (h *handler) PostLogin(c *gin.Context) {
	ctx := c.Request.Context()

	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.PostLogin")
	defer span.Finish()

	var lb openapi.LoginBody
	if err := c.ShouldBind(&lb); err != nil {
		c.String(http.StatusBadRequest, "the body should be json")
	}

	session, err := h.usecase.PostLogin(ctx, lb.Token)
	if err != nil {
		// TODO: 真面目にエラーハンドリングする。
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.SetCookie(h.sessionCookieName, session, h.sessionMaxAge, "/", "localhost", false, true)

	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h *handler) GetMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "getme",
	})
}

func (h *handler) PostMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "post me",
	})
}

func (h *handler) GetUserByID(c *gin.Context) {
}
