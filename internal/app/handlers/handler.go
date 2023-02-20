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

	/*auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	*/
	api := router.Group(":username" /*, h.identifyUser*/)
	{
		linklist := api.Group("/linklist")
		{
			linklist.POST("/", h.createList)
			linklist.GET("/", h.getAllLists)
			linklist.GET("/:title", h.getListByID)
			linklist.PUT("/:title", h.updateList)
			linklist.DELETE("/:title", h.deleteList)

			links := linklist.Group(":title/links")
			{
				links.POST("/", h.createLink)
				links.GET("/", h.getAllLinks)
				links.GET("/:link_name", h.getLinkByID)
				links.PUT("/:link_id", h.updateLink)
				links.DELETE("/:link_id", h.deleteLink)
			}
		}
	}

	return router
}
