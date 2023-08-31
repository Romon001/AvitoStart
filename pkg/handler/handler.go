package handler

import (
	"github.com/Romon001/AvitoStart-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/GetSegments", h.GetAll)
		api.POST("/CreateSegment", h.CreateSegment)
		api.DELETE("/DeleteSegment/:name", h.DeleteSegment)
		api.GET("/GetUserSegments/:user_id", h.GetUserSegments)

	}

	return router
}
