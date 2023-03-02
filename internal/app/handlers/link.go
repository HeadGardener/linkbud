package handlers

import (
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createLink(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	title := c.Param("title")

	var link models.Link
	if err := c.BindJSON(&link); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data to bind link")
		return
	}

	listID, err := h.getListID(userID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	linkID, err := h.service.LinkInterface.Create(userID, link, listID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"linkID": linkID,
	})
}

func (h *Handler) getAllLinks(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	title := c.Param("title")

	listID, err := h.getListID(userID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.service.ListInterface.GetList(userID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	links, err := h.service.LinkInterface.GetAll(userID, listID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.ListsLinks{
		List:  list,
		Links: links,
	})
}

func (h *Handler) getLinkByTitle(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	title := c.Param("title")
	linkTitle := c.Param("link_title")

	listID, err := h.getListID(userID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.service.ListInterface.GetList(userID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	link, err := h.service.LinkInterface.GetByID(userID, listID, title, linkTitle)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var links []models.Link
	links = append(links, link)
	c.JSON(http.StatusCreated, models.ListsLinks{
		List:  list,
		Links: links,
	})
}

func (h *Handler) updateLink(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	title := c.Param("title")

	listID, err := h.getListID(userID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	linkID, err := strconv.Atoi(c.Param("link_id"))
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "invalid id type")
		return
	}

	var linkInput models.LinkInput
	if err := c.BindJSON(&linkInput); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data to bind link")
		return
	}

	linkInput.ID = linkID

	err = h.service.LinkInterface.Update(userID, listID, linkInput, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "updated",
	})
}

func (h *Handler) deleteLink(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid userID header")
		return
	}

	title := c.Param("title")

	linkID, err := strconv.Atoi(c.Param("link_id"))
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "invalid id type")
		return
	}

	err = h.service.LinkInterface.Delete(userID, linkID, title)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
	})
}

func (h *Handler) getListID(userID int, title string) (int, error) {
	list, err := h.service.ListInterface.GetList(userID, title)
	if err != nil {
		return 0, err
	}

	return list.ID, nil
}
