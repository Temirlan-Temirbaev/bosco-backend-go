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
}
type Contact interface {
	Create(contact model.Contact) (int, error)
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
	}
}
