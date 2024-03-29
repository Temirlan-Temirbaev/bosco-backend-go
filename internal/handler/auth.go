package handler

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/utils"
	"bosco-backend/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) login(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.ValidateAuth(c, input); err != nil {
		return
	}

	user, err := h.services.Authorization.GetUser(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(user)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) registration(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.ValidateAuth(c, input); err != nil {
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
	id, err := GetUserId(c)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Authorization.GetUserById(id)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
