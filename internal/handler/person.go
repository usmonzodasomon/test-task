package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/test-task/internal/models"
	"github.com/usmonzodasomon/test-task/internal/service"
)

func (h *handler) GetPerson(c *gin.Context) {
	var params models.GetPersonRequest
	var err error

	params.Limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid limit query")
		return
	}

	params.Offset, err = strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid offset query")
		return
	}

	params.Age, err = strconv.Atoi(c.DefaultQuery("age", "-1"))
	if err != nil {
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid age query")
		return
	}
	params.Gender = c.Query("gender")
	params.Nationality = c.Query("nationality")

	people, err := h.service.GetPerson(params)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, people)
	h.logg.Info(fmt.Sprintf("Get users with params: %+#v", params))
}

func (h *handler) AddPerson(c *gin.Context) {
	var input models.AddPersonInput

	if err := c.BindJSON(&input); err != nil {
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.AddPerson(input)
	if err != nil {
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
	h.logg.Info(fmt.Sprintf("User with id %d created", id))
}

func (h *handler) ChangePerson(c *gin.Context) {
	h.logg.Info(fmt.Sprintf("Change user with id: %s", c.Param("id")))
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.Person

	if err := c.BindJSON(&input); err != nil {
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.ChangePerson(id, input); err != nil {
		if errors.Is(err, service.ErrRecordNotFound) {
			h.logg.Error(err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "person not found")
			return
		}
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	h.logg.Info(fmt.Sprintf("User with id %d changed", id))
}

func (h *handler) DeletePerson(c *gin.Context) {
	h.logg.Info(fmt.Sprintf("Delete user with id: %s", c.Param("id")))
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.service.DeletePerson(id); err != nil {
		if errors.Is(err, service.ErrRecordNotFound) {
			h.logg.Error(err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "person not found")
			return
		}
		h.logg.Error(err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	h.logg.Info(fmt.Sprintf("User with id %d deleted", id))
}
