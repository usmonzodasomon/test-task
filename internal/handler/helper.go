package handler

import "github.com/gin-gonic/gin"

func (h *handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	h.logg.Error(message)
	c.AbortWithStatusJSON(statusCode, gin.H{
		"status": "error",
		"error":  message,
	})
}
