package handler

import (
	"bosco-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/registration", h.registration)
		auth.GET("/check-login", h.UserIdentity, h.checkLogin)
	}
	api := router.Group("/api")
	{
		contact := api.Group("/contact")
		{
			contact.POST("/", h.UserIdentity, h.CheckRole("admin"), h.createContact)
			contact.PUT("/:id", h.UserIdentity, h.CheckRole("admin"), h.editContact)
			contact.GET("/", h.getContacts)
			contact.DELETE("/:id", h.UserIdentity, h.CheckRole("admin"), h.deleteContact)
		}
		product := api.Group("/product")
		{
			product.POST("/", h.UserIdentity, h.createProduct)
			product.GET("/", h.getProducts)
			product.DELETE("/:id", h.UserIdentity, h.deleteProduct)
			product.PUT("/:id", h.UserIdentity, h.editProduct)
		}
		category := api.Group("/category")
		{
			category.POST("/", h.UserIdentity, h.createCategory)
			category.GET("/", h.getCategories)
			category.GET("/:id", h.getCategoryById)
			category.PUT("/:id", h.UserIdentity, h.editCategory)
			category.DELETE("/:id", h.UserIdentity, h.deleteCategory)
		}
	}
	return router
}
