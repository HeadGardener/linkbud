package handlers

import (
	"fmt"
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

	var list models.ListInput
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
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	lists, err := h.service.ListInterface.GetAll(userID)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, lists)
}

func (h *Handler) getListByTitle(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	title := c.Param("title")

	list, err := h.service.ListInterface.GetList(userID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	title := c.Param("title")

	var list models.ListInput
	if err := c.BindJSON(&list); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data to bind list")
		return
	}

	listID, err := h.service.ListInterface.Update(userID, title, list)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status": fmt.Sprintf("list with id %d updated", listID),
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	title := c.Param("title")

	listID, err := h.service.ListInterface.Delete(userID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status": fmt.Sprintf("list with id %d deleted", listID),
	})
}
