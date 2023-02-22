package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const userCtx = "userID"

func (h *Handler) checkUsername(c *gin.Context) {
	username := c.Param("username")
	id, err := h.service.Authorization.CheckUsername(username)
	if err != nil {
		newErrResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, id)
}

func (h *Handler) identifyUser(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrResponse(c, http.StatusBadRequest, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrResponse(c, http.StatusBadRequest, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrResponse(c, http.StatusBadRequest, "jwt token is empty")
	}

	userID, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	oldID, _ := getUserId(c)
	if userID != oldID {
		newErrResponse(c, http.StatusUnauthorized, "you don't have enough rules")
		return
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
