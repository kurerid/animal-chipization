package handler

import (
	"animal-chipization/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) getLocationById(c *gin.Context) {
	var input models.LocationGetByIdInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
	}

	output, err := h.services.Location.GetById(input.Id)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) createLocation(c *gin.Context) {
	var input models.LocationNewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	err := h.services.Location.New(input.Latitude, input.Longitude)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) modifyLocation(c *gin.Context) {
	var input models.LocationModifyInput
	id := c.Param("pointId")
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 500, err)
		return
	}
	temp, _ := strconv.Atoi(id)
	input.Id = int64(temp)

	err := h.services.Location.Modify(input.Id, input.Latitude, input.Longitude)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) removeLocation(c *gin.Context) {
	var input models.LocationRemoveInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	err := h.services.Location.Remove(input.Id)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}
