package validator

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateAuth(c *gin.Context, user model.User) error {
	if len(user.Username) < 5 {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Login must be longer than 5")
		return errors.New("Bad login")
	}
	if len(user.Password) < 5 {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Password must be longer than 5")
		return errors.New("Bad password")
	}
	return nil
}
