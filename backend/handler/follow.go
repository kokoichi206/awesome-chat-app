package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kokoichi206/awesome-chat-app/openapi/gen/go/openapi"
)

func (h *handler) GetFollowers(c *gin.Context) {
}

func (h *handler) GetFollowing(c *gin.Context) {
}

func (h *handler) PatchFollowing(c *gin.Context, userId openapi.UserIdPath) {
}

func (h *handler) PostFollowing(c *gin.Context, userId openapi.UserIdPath) {
}
