package handler

import (
	"net/http"

	avitoStartApp "github.com/Romon001/AvitoStart-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSegment(c *gin.Context) {
	var input avitoStartApp.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Segments.Create(input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	var lists []avitoStartApp.Segment
	lists, err := h.services.Segments.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) DeleteSegment(c *gin.Context) {

}

func (h *Handler) AddUserToSegments(c *gin.Context) {

}

func (h *Handler) GetUserSegments(c *gin.Context) {

}
func (h *Handler) DeleteUserFromSegments(c *gin.Context) {

}

type getAllListsResponse struct {
	Data []avitoStartApp.Segment `json:"data"`
}
