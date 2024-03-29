package handler

import (
	"bosco-backend/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AuthorizationHeader = "Authorization"
	userCtx             = "user_id"
)

func (handler *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(AuthorizationHeader)

	if header == "" {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "Not authorized")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		return
	}
	userId, err := handler.services.Authorization.GetIdFromToken(headerParts[1])
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "User id is not found")
		return 0, errors.New("")
	}
	idInt, ok := id.(int)
	if !ok {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "Invalid type of User id")
		return 0, errors.New("User id has invalid type")
	}
	return idInt, nil
}

func (h *Handler) CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		if user.Role != role {
			utils.NewErrorResponse(c, http.StatusMethodNotAllowed, "You have no rights")
			return
		}
	}
}
