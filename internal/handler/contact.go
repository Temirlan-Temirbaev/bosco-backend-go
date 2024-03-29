package handler

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createContact(c *gin.Context) {
	var input model.Contact
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Create(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getContacts(c *gin.Context) {

}

func (h *Handler) editContact(c *gin.Context) {

}

func (h *Handler) deleteContact(c *gin.Context) {

}
