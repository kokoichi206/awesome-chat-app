package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session, err := c.Cookie(h.sessionCookieName)
		if errors.Is(err, http.ErrNoCookie) {
			c.Status(http.StatusUnauthorized)
			c.Abort()

			return
		}

		if err := h.usecase.VerifySessionCookie(c.Request.Context(), session); err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()

			return
		}

		c.Next()
	}
}
