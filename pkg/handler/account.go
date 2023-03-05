package handler

import (
	"animal-chipization/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) signUp(c *gin.Context) {
	if c.GetHeader("Authorization") != "" {
		h.sendResponse(c, 403, "Already authorised user")
		return
	}
	var input models.SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}
	ok, err := h.services.Account.Exist(input.Email)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	} else if ok == true {
		h.sendResponse(c, 409, gin.H{})
		return
	}
	output, err := h.services.Authorization.SignUp(&input)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}
	c.JSON(201, &output)
}

func (h *Handler) accountGetById(c *gin.Context) {
	var input models.AccountGetByIdInput

	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}
	output, err := h.services.Account.GetById(input.Id)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) searchAccount(c *gin.Context) {
	var input models.AccountSearchInput

	if err := c.ShouldBindQuery(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}
	output, err := h.services.Account.Search(&input)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) updateAccount(c *gin.Context) {
	var input models.AccountUpdateInput
	var err error

	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	input.Id, err = strconv.Atoi(c.Param("accountId"))
	if err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	output, err := h.services.Account.Update(&input)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.JSON(201, &output)
}

func (h *Handler) removeAccount(c *gin.Context) {
	var input models.AccountRemoveInput
	var err error

	input.Id, err = strconv.Atoi(c.Param("accountId"))
	if err != nil {
		h.sendResponse(c, 400, err)
		return
	}

	err = h.services.Account.Remove(input.Id)
	if err != nil {
		h.sendResponse(c, 500, err)
		return
	}

	c.Status(201)
}
