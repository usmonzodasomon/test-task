package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteUser(c *gin.Context) {
	h.logg.Info(fmt.Sprintf("Delete user with id: %s", c.Param("id")))
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	h.logg.Info(fmt.Sprintf("User with id %d deleted", id))
}
