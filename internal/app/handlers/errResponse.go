package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"Message"`
}

func newErrResponse(c *gin.Context, statusCode int, msg string) {
	logrus.Errorf(msg)
	c.AbortWithStatusJSON(statusCode, Error{
		Message: msg,
	})
}
