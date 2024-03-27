package service

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(user model.User) (model.User, error)
	GenerateToken(user model.User) (string, error)
}

type Product interface {
}

type Category interface {
}
type Contact interface {
}

type Service struct {
	Authorization
	Product
	Category
	Contact
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r),
	}
}
