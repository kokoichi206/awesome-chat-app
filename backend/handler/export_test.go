package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) SessionCheck() gin.HandlerFunc {
	return h.sessionCheck()
}
