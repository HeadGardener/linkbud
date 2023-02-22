package handlers

import (
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data to bind user")
		return
	}

	id, err := h.service.Authorization.Create(user)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var userInput models.UserInput
	if err := c.BindJSON(&userInput); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data to bind user")
		return
	}

	token, err := h.service.Authorization.GenerateToken(userInput)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"jwt-token": token,
	})
}
