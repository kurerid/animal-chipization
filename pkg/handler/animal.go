package handler

import (
	"animal-chipization/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) animalGetById(c *gin.Context) {
	var input models.AnimalGetByIdInput

	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	output, err := h.services.Animal.GetById(input.Id)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) searchAnimal(c *gin.Context) {
	var input models.AnimalSearchInput

	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	output, err := h.services.Animal.Search(&input)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) getAnimalLocations(c *gin.Context) {
	var input models.AnimalVisitedLocationListInput

	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	output, err := h.services.Locations(&input)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) createAnimal(c *gin.Context) {
	var input models.AnimalNewInput

	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	output, err := h.services.Animal.New(&input)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) updateAnimal(c *gin.Context) {
	var input models.AnimalUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	output, err := h.services.Animal.Update(&input)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) removeAnimal(c *gin.Context) {
	var input models.AnimalRemoveInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	err := h.services.Animal.Remove(input.Id)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) addTypeToAnimal(c *gin.Context) {
	var input models.AnimalAddTypeInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	err := h.services.Animal.AddType(input.Animal, input.Type)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) changeAnimalType(c *gin.Context) {
	var input models.AnimalChangeTypeInput
	animal := c.Param("animalID")
	animalID, err := strconv.Atoi(animal)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}
	input.Animal = int64(animalID)

	err = h.services.Animal.ChangeType(input.Animal, input.OldType, input.NewType)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) removeTypeFromAnimal(c *gin.Context) {
	var input models.AnimalRemoveTypeInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	err := h.services.Animal.RemoveType(input.Animal, input.Type)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) getVisitedLocations(c *gin.Context) {
	var input models.AnimalVisitedLocationListInput

	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	output, err := h.services.Animal.Locations(&input)
	if err != nil {
		h.sendResponse(c, 500, err)
	}

	c.JSON(201, &output)
}

func (h *Handler) createVisitedLocation(c *gin.Context) {
	var input models.AnimalVisitedLocationNewInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	err := h.services.Animal.AddLocation(input.Animal, input.Location)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) changeVisitedLocation(c *gin.Context) {
	var input models.AnimalVisitedLocationChangeInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	err := h.services.Animal.ChangeLocation(input.Animal, input.VisitedLocation, input.Location)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}

func (h *Handler) removeVisitedLocation(c *gin.Context) {
	var input models.AnimalVisitedLocationRemoveInput
	if err := c.ShouldBindUri(&input); err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	err := h.services.Animal.RemoveLocation(input.Animal, input.VisitedLocation)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}
