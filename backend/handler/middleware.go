package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var tokenContext = "token-context"

func (h *handler) sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session, err := c.Cookie(h.sessionCookieName)
		if errors.Is(err, http.ErrNoCookie) {
			c.Status(http.StatusUnauthorized)
			c.Abort()

			return
		}

		token, err := h.usecase.VerifySessionCookie(c.Request.Context(), session)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()

			return
		}

		c.Set(tokenContext, *token)

		c.Next()
	}
}
