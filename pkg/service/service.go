package service

import "bosco-backend/pkg/repository"

type Authorization interface {
}

type Product interface {
}

type Category interface {
}
type Contact interface {
}

type Service struct {
	Product
	Authorization
	Category
	Contact
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
