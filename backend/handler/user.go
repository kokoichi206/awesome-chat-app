package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kokoichi206/awesome-chat-app/openapi/gen/go/openapi"
)

func (h *handler) PostLogin(c *gin.Context) {
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

func (h *handler) GetUserByID(c *gin.Context, userId openapi.UserIdPath) {
}
