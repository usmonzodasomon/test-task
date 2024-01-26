package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/test-task/internal/models"
)

func (h *handler) AddPerson(c *gin.Context) {
	var input models.AddPersonInput

	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.AddPerson(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	h.logg.Info(fmt.Sprintf("User with id %d created", id))
}

func (h *handler) DeletePerson(c *gin.Context) {
	h.logg.Info(fmt.Sprintf("Delete user with id: %s", c.Param("id")))
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.service.DeletePerson(id); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	h.logg.Info(fmt.Sprintf("User with id %d deleted", id))
}
