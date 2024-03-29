package handler

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createContact(c *gin.Context) {
	var input model.Contact
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Contact.Create(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getContacts(c *gin.Context) {
	contacts, err := h.services.Contact.GetAll()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": contacts,
	})
}

func (h *Handler) editContact(c *gin.Context) {

}

func (h *Handler) deleteContact(c *gin.Context) {
	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Contact.Delete(contactId)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
