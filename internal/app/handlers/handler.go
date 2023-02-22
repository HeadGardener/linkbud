package handlers

import (
	"github.com/HeadGardener/linkbud/internal/app/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/:username", h.checkUsername)
	{
		linklist := api.Group("/", h.identifyUser)
		{
			linklist.POST("/", h.createList)
			linklist.GET("/", h.getAllLists)
			linklist.GET("/:title", h.getListByTitle)
			linklist.PUT("/:title", h.updateList)
			linklist.DELETE("/:title", h.deleteList)

			links := linklist.Group(":title/")
			{
				links.POST("/", h.createLink)
				// links.GET("/", h.getAllLinks)
				// links.GET("/:link_title", h.getLinkByTitle)
				links.PUT("/:link_id", h.updateLink)
				links.DELETE("/:link_id", h.deleteLink)
			}
		}
	}

	api.GET("/:title/", h.getAllLinks)
	api.GET("/:title/:link_title", h.getLinkByTitle)
	return router
}
