package handler

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCategory(c *gin.Context) {
	var input model.Category
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Category.Create(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getCategories(c *gin.Context) {
	categories, err := h.services.Category.GetAll()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": categories,
	})

}

func (h *Handler) getCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	category, err := h.services.Category.GetById(id)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": category,
	})
}

func (h *Handler) editCategory(c *gin.Context) {
	var input model.Category
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.Category.Update(id, input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) deleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.Category.Delete(id)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
