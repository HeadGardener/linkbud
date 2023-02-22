package handlers

import (
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	var list models.LinkList
	if err := c.BindJSON(&list); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data to bind list")
		return
	}

	listID, err := h.service.ListInterface.Create(userID, list)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"listID": listID,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListByTitle(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
