package handler

import (
	"animal-chipization/models"
	"github.com/gin-gonic/gin"
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

func (h *Handler) signIn(c *gin.Context) {

}

func (h *Handler) accountGetById(c *gin.Context) {

}

func (h *Handler) searchAccount(c *gin.Context) {

}

func (h *Handler) updateAccount(c *gin.Context) {

}

func (h *Handler) removeAccount(c *gin.Context) {

}
