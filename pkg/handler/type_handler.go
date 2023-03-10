package handler

import (
	"animal-chipization/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) getAnimalType(c *gin.Context) {
	var input models.AnimalTypeGetByIdInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	output, err := h.services.Type.GetById(input.Id)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) createAnimalType(c *gin.Context) {
	var input models.AnimalTypeNewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	err := h.services.Type.New(input.Type)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) modifyAnimalType(c *gin.Context) {
	var input models.AnimalTypeModifyInput
	id := c.Param("typeId")
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}
	temp, _ := strconv.Atoi(id)
	input.Id = int64(temp)

	err := h.services.Type.Modify(input.Id, input.Type)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) removeAnimalType(c *gin.Context) {
	var input models.AnimalTypeRemoveInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	err := h.services.Type.Remove(input.Id)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}
