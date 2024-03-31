package service

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(user model.User) (model.User, error)
	GetUserById(id int) (model.User, error)
	GenerateToken(user model.User) (string, error)
	GetIdFromToken(accessToken string) (int, error)
	GetRoleFromToken(accessToken string) (string, error)
}

type Product interface {
}

type Category interface {
	Create(category model.Category) (int, error)
	Update(id int, category model.Category) error
	Delete(id int) error
	GetAll() ([]model.Category, error)
	GetById(id int) (model.Category, error)
}
type Contact interface {
	Create(contact model.Contact) (int, error)
	GetAll() ([]model.Contact, error)
	Delete(id int) error
	Update(id int, contact model.Contact) error
}

type Service struct {
	Authorization
	Product
	Category
	Contact
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
		Contact:       NewContactService(r.Contact),
		Category:      NewCategoryService(r.Category),
	}
}
