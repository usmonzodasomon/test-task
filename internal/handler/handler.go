package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/test-task/internal/service"
)

type handler struct {
	service *service.Service
	logg    *slog.Logger
}

func NewHandler(service *service.Service, logg *slog.Logger) *handler {
	return &handler{
		service: service,
		logg:    logg,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		api.POST("/person", h.AddPerson)
		api.DELETE("/person/:id", h.DeletePerson)
	}
	return router
}
