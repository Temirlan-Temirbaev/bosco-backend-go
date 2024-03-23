package handler

import (
	"bosco-backend/pkg/model"
	"bosco-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) login(c *gin.Context) {

}

func (h *Handler) registration(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) checkLogin(c *gin.Context) {

}
