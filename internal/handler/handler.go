package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/test-task/internal/service"
)

type handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *handler {
	return &handler{service: service}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}
	return router
}
