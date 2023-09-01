package handler

import (
	"net/http"
	"strconv"

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

	c.JSON(http.StatusOK, listOfSegments{
		Data: lists,
	})
}

func (h *Handler) DeleteSegment(c *gin.Context) {
	var input avitoStartApp.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Segments.DeleteSegment(input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Ok",
	})
}

func (h *Handler) AddUserToSegments(c *gin.Context) {
	var result int
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input listofStrings
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err = h.services.UserSegmentPair.AddUserToSegments(input.Data, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"added rows": result,
	})
}

func (h *Handler) GetUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var lists []avitoStartApp.UserSegmentPair

	lists, err = h.services.UserSegmentPair.GetUserSegments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, listofUsersegmentPairs{
		Data: lists,
	})
}
func (h *Handler) DeleteUserFromSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input listofStrings
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var result int
	result, err = h.services.UserSegmentPair.DeleteUserFromSegments(input.Data, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deleted rows": result,
	})
}

type listOfSegments struct {
	Data []avitoStartApp.Segment `json:"data"`
}
type listofUsersegmentPairs struct {
	Data []avitoStartApp.UserSegmentPair `json:"data"`
}
type listofStrings struct {
	Data []string `json:"data"`
}
