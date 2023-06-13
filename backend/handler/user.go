package handler

import (
	"net/http"

	auth "firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/openapi/gen/go/openapi"
)

func (h *handler) PostLogin(c *gin.Context) {
	ctx := c.Request.Context()

	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.PostLogin")
	defer span.Finish()

	var lb openapi.LoginBody
	if err := c.ShouldBind(&lb); err != nil {
		c.String(http.StatusBadRequest, "the body should be json")

		return
	}

	session, err := h.usecase.PostLogin(ctx, lb.Token)
	if err != nil {
		// TODO: 真面目にエラーハンドリングする。
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.SetCookie(h.sessionCookieName, session, h.sessionMaxAge, "/", "localhost", false, true)

	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h *handler) GetMe(c *gin.Context) {
	ctx := c.Request.Context()

	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.PostLogin")
	defer span.Finish()

	tc, ok := c.Get(tokenContext)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token is not found",
		})

		return
	}

	token, ok := (tc).(auth.Token)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token is invalid",
		})

		return
	}

	email := token.Claims["email"].(string)

	user, err := h.usecase.GetUser(ctx, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user is not found",
		})

		return
	}

	me := openapi.Me{
		Id:       user.ID,
		Email:    email,
		ImgUrl:   &user.PictureURL,
		Username: user.Name,
	}

	c.JSON(http.StatusOK, me)
}

func (h *handler) PostMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "post me",
	})
}

func (h *handler) GetUserByID(c *gin.Context) {
}
