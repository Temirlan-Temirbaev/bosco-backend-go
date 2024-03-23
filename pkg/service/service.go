package service

import (
	"bosco-backend/pkg/model"
	"bosco-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
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
