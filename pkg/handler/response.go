package handler

import "github.com/gin-gonic/gin"

func (h *Handler) sendResponse(c *gin.Context, status int, data interface{}) {
	c.AbortWithStatusJSON(status, data)
}
